// AWS X-Ray SDK for Node.js
const AWSXRay = require("aws-xray-sdk-core");
// AWS SDK
const AWS = AWSXRay.captureAWS(require("aws-sdk"));

// DynamoDB DocumentClient
const dynamoDB = new AWS.DynamoDB.DocumentClient();

// Projectプレフィックス（環境変数から取得、Terraform変数と一致させる）
const PROJECT_PREFIX = process.env.PROJECT_PREFIX || "ecommerce-observability";

// テーブル名
const PRODUCTS_TABLE = `${PROJECT_PREFIX}-products`;
const CATEGORIES_TABLE = `${PROJECT_PREFIX}-categories`;

// ログ出力関数
const log = (level, message, data = {}) => {
  const timestamp = new Date().toISOString();
  const requestId = process.env._X_AMZN_TRACE_ID || "unknown";

  // 構造化ログ
  const logEntry = {
    timestamp,
    level,
    message,
    requestId,
    service: "catalog-service",
    ...data,
  };

  console.log(JSON.stringify(logEntry));
};

// リクエスト処理の基本関数
const handleRequest = async (event, context, handler) => {
  // イベントからパスとHTTPメソッドを取得
  const path = getPath(event);
  const httpMethod = getHttpMethod(event);

  // リクエスト開始ログ
  log("INFO", "Request received", {
    path,
    httpMethod,
    queryStringParameters: event.queryStringParameters,
  });

  const startTime = Date.now();

  try {
    // 実際の処理を実行
    const result = await handler(event, context);

    // レスポンスタイム計測
    const responseTime = Date.now() - startTime;

    // 成功ログ
    log("INFO", "Request completed successfully", {
      responseTime,
      statusCode: result.statusCode,
    });

    return result;
  } catch (error) {
    // レスポンスタイム計測
    const responseTime = Date.now() - startTime;

    // エラーログ
    log("ERROR", "Error processing request", {
      responseTime,
      errorMessage: error.message,
      stackTrace: error.stack,
    });

    // エラーレスポンス
    return {
      statusCode: 500,
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*", // CORS対応
      },
      body: JSON.stringify({
        message: "Internal server error",
      }),
    };
  }
};

// イベントからパスを取得する関数
const getPath = (event) => {
  if (event.path) {
    return event.path;
  }

  if (event.requestContext) {
    if (event.requestContext.http) {
      return event.requestContext.http.path; // HTTP API (v2)
    }
    if (event.requestContext.path) {
      return event.requestContext.path; // REST API (v1)
    }
    if (event.requestContext.resourcePath) {
      return event.requestContext.resourcePath; // 別のREST API形式
    }
  }

  if (event.resource) {
    return event.resource; // バックアップとしてリソースパスを使用
  }

  log("WARN", "Could not determine path from event", { event });
  return "/"; // デフォルトパス
};

// イベントからHTTPメソッドを取得する関数
const getHttpMethod = (event) => {
  if (event.httpMethod) {
    return event.httpMethod;
  }

  if (event.requestContext && event.requestContext.http) {
    return event.requestContext.http.method; // HTTP API (v2)
  }

  if (event.requestContext && event.requestContext.httpMethod) {
    return event.requestContext.httpMethod; // REST API (v1)
  }

  log("WARN", "Could not determine HTTP method from event", { event });
  return "GET"; // デフォルトメソッド
};

// パスパラメータを取得する関数
const getPathParameter = (event, index) => {
  const path = getPath(event);
  const segments = path.split("/").filter((segment) => segment.length > 0);

  if (segments.length > index) {
    return segments[index];
  }

  // API Gateway v2の場合はpathParametersから取得
  if (event.pathParameters) {
    const keys = Object.keys(event.pathParameters);
    if (keys.length > 0) {
      return event.pathParameters[keys[0]];
    }
  }

  return null;
};

// 全商品を取得
const getProducts = async () => {
  // DynamoDBから商品一覧を取得
  const params = {
    TableName: PRODUCTS_TABLE,
  };

  // X-Rayサブセグメント作成
  const subsegment = AWSXRay.getSegment().addNewSubsegment(
    "DynamoDB-getProducts"
  );

  try {
    subsegment.addAnnotation("table", PRODUCTS_TABLE);
    subsegment.addMetadata("params", params);

    const result = await dynamoDB.scan(params).promise();

    subsegment.addMetadata("itemCount", result.Items.length);
    subsegment.close();

    return result.Items;
  } catch (error) {
    subsegment.addError(error);
    subsegment.close();
    throw error;
  }
};

// 特定のカテゴリの商品を取得
const getProductsByCategory = async (categoryId) => {
  // DynamoDBからカテゴリに属する商品を取得
  const params = {
    TableName: PRODUCTS_TABLE,
    IndexName: "CategoryIndex",
    KeyConditionExpression: "category_id = :cid",
    ExpressionAttributeValues: {
      ":cid": categoryId,
    },
  };

  // X-Rayサブセグメント作成
  const subsegment = AWSXRay.getSegment().addNewSubsegment(
    "DynamoDB-getProductsByCategory"
  );

  try {
    subsegment.addAnnotation("table", PRODUCTS_TABLE);
    subsegment.addAnnotation("categoryId", categoryId);
    subsegment.addMetadata("params", params);

    const result = await dynamoDB.query(params).promise();

    subsegment.addMetadata("itemCount", result.Items.length);
    subsegment.close();

    return result.Items;
  } catch (error) {
    subsegment.addError(error);
    subsegment.close();
    throw error;
  }
};

// 特定の商品を取得
const getProduct = async (productId) => {
  // DynamoDBから商品を取得
  const params = {
    TableName: PRODUCTS_TABLE,
    Key: {
      id: productId,
    },
  };

  // X-Rayサブセグメント作成
  const subsegment = AWSXRay.getSegment().addNewSubsegment(
    "DynamoDB-getProduct"
  );

  try {
    subsegment.addAnnotation("table", PRODUCTS_TABLE);
    subsegment.addAnnotation("productId", productId);
    subsegment.addMetadata("params", params);

    const result = await dynamoDB.get(params).promise();

    subsegment.close();

    if (!result.Item) {
      throw new Error("Product not found");
    }

    return result.Item;
  } catch (error) {
    subsegment.addError(error);
    subsegment.close();
    throw error;
  }
};

// 全カテゴリを取得
const getCategories = async () => {
  // DynamoDBからカテゴリ一覧を取得
  const params = {
    TableName: CATEGORIES_TABLE,
  };

  // X-Rayサブセグメント作成
  const subsegment = AWSXRay.getSegment().addNewSubsegment(
    "DynamoDB-getCategories"
  );

  try {
    subsegment.addAnnotation("table", CATEGORIES_TABLE);
    subsegment.addMetadata("params", params);

    const result = await dynamoDB.scan(params).promise();

    subsegment.addMetadata("itemCount", result.Items.length);
    subsegment.close();

    return result.Items;
  } catch (error) {
    subsegment.addError(error);
    subsegment.close();
    throw error;
  }
};

// 特定のカテゴリを取得
const getCategory = async (categoryId) => {
  // DynamoDBからカテゴリを取得
  const params = {
    TableName: CATEGORIES_TABLE,
    Key: {
      id: categoryId,
    },
  };

  // X-Rayサブセグメント作成
  const subsegment = AWSXRay.getSegment().addNewSubsegment(
    "DynamoDB-getCategory"
  );

  try {
    subsegment.addAnnotation("table", CATEGORIES_TABLE);
    subsegment.addAnnotation("categoryId", categoryId);
    subsegment.addMetadata("params", params);

    const result = await dynamoDB.get(params).promise();

    subsegment.close();

    if (!result.Item) {
      throw new Error("Category not found");
    }

    return result.Item;
  } catch (error) {
    subsegment.addError(error);
    subsegment.close();
    throw error;
  }
};

// Lambda関数のルーティング処理
exports.handler = async (event, context) => {
  // イベントの全体構造をログに出力（デバッグ用）
  log("DEBUG", "Event received", { event });

  // パスとHTTPメソッドを取得
  const path = getPath(event);
  const httpMethod = getHttpMethod(event);

  // ヘルスチェックエンドポイント
  if (path === "/health") {
    return {
      statusCode: 200,
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
      },
      body: JSON.stringify({
        status: "healthy",
        timestamp: new Date().toISOString(),
        service: "catalog-service",
      }),
    };
  }

  // 商品一覧の取得
  if (path === "/products" && httpMethod === "GET") {
    return handleRequest(event, context, async () => {
      const products = await getProducts();

      return {
        statusCode: 200,
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
        },
        body: JSON.stringify(products),
      };
    });
  }

  // 特定の商品の取得
  if (path.match(/^\/products\/[a-zA-Z0-9-]+$/) && httpMethod === "GET") {
    return handleRequest(event, context, async () => {
      const productId = path.split("/")[2];

      try {
        const product = await getProduct(productId);

        return {
          statusCode: 200,
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(product),
        };
      } catch (error) {
        if (error.message === "Product not found") {
          return {
            statusCode: 404,
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
            },
            body: JSON.stringify({ message: "Product not found" }),
          };
        }
        throw error;
      }
    });
  }

  // カテゴリ一覧の取得
  if (path === "/categories" && httpMethod === "GET") {
    return handleRequest(event, context, async () => {
      const categories = await getCategories();

      return {
        statusCode: 200,
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
        },
        body: JSON.stringify(categories),
      };
    });
  }

  // 特定のカテゴリの取得
  if (path.match(/^\/categories\/[a-zA-Z0-9-]+$/) && httpMethod === "GET") {
    return handleRequest(event, context, async () => {
      const categoryId = path.split("/")[2];

      try {
        const category = await getCategory(categoryId);

        return {
          statusCode: 200,
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(category),
        };
      } catch (error) {
        if (error.message === "Category not found") {
          return {
            statusCode: 404,
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
            },
            body: JSON.stringify({ message: "Category not found" }),
          };
        }
        throw error;
      }
    });
  }

  // カテゴリに属する商品一覧の取得
  if (
    path.match(/^\/categories\/[a-zA-Z0-9-]+\/products$/) &&
    httpMethod === "GET"
  ) {
    return handleRequest(event, context, async () => {
      const categoryId = path.split("/")[2];

      try {
        // カテゴリが存在するか確認
        await getCategory(categoryId);

        // カテゴリに属する商品を取得
        const products = await getProductsByCategory(categoryId);

        return {
          statusCode: 200,
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
          body: JSON.stringify(products),
        };
      } catch (error) {
        if (error.message === "Category not found") {
          return {
            statusCode: 404,
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
            },
            body: JSON.stringify({ message: "Category not found" }),
          };
        }
        throw error;
      }
    });
  }

  // 存在しないパスの場合は404を返す
  return {
    statusCode: 404,
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
    },
    body: JSON.stringify({
      message: "Not found",
      path: path,
      method: httpMethod,
    }),
  };
};
