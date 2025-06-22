#!/usr/bin/env python3
# AWS オブザーバビリティ eコマースアプリのアーキテクチャ図生成スクリプト（SVG専用、画像埋め込み）

from diagrams import Diagram, Cluster, Edge
from diagrams.aws.network import CloudFront, APIGateway
from diagrams.aws.security import Cognito
from diagrams.aws.storage import S3
from diagrams.aws.compute import Lambda
from diagrams.aws.integration import SQS
from diagrams.aws.database import RDS
from diagrams.aws.management import Cloudwatch
from diagrams.aws.devtools import XRay
from diagrams.aws.compute import Fargate
from diagrams.onprem.client import User as CustomerUser
from diagrams.onprem.client import Client as AdminClient
import os
import base64
from pathlib import Path

# 図の設定
graph_attr = {
    "fontname": "sans-serif",
    "fontsize": "30",
    "bgcolor": "transparent",
    "layout": "dot",
    "margin": "0.3,0.3",
    "pad": "1.0",
    "nodesep": "0.8",
    "ranksep": "1.0",
    "splines": "spline"
}

# ノードの設定
node_attr = {
    "fontname": "sans-serif",
    "fontsize": "12",
}

# エッジの設定
edge_attr = {
    "fontname": "sans-serif",
    "fontsize": "10",
}

# 生成されたSVGファイルを修正する関数
def fix_svg_image_paths(svg_file):
    # SVGファイルを読み込む
    with open(svg_file, 'r', encoding='utf-8') as f:
        svg_content = f.read()

    # 画像パスを修正する正規表現パターン
    import re
    pattern = r'xlink:href="(/[^"]+)"'

    # パスを抽出してBase64エンコードに置き換える
    def replace_with_embedded_image(match):
        path = match.group(1)
        try:
            # 画像が存在する場合は埋め込む
            if os.path.exists(path):
                with open(path, 'rb') as img_file:
                    img_data = img_file.read()
                    img_base64 = base64.b64encode(img_data).decode('utf-8')
                    img_ext = os.path.splitext(path)[1].lstrip('.')
                    if img_ext.lower() in ['jpg', 'jpeg']:
                        mime_type = 'image/jpeg'
                    elif img_ext.lower() == 'png':
                        mime_type = 'image/png'
                    else:
                        mime_type = f'image/{img_ext}'
                    return f'xlink:href="data:{mime_type};base64,{img_base64}"'
            # 画像が存在しない場合はプレースホルダーを使用
            return 'xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYAAABw4pVUAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH4QoTCxUHCzJQMQAAAB1pVFh0Q29tbWVudAAAAAAAQ3JlYXRlZCB3aXRoIEdJTVBkLmUHAAACIklEQVR42u3dsU7CQBiA0Y/ExGDi6KiDk5OrgyM+kg/B4AtJHB0dXRzd3AwOJsbqYLUJkUHv/5IupUf5coVLSdxut7vLMAxPwzg8l2V53drv8Xi8Ho/Hm1arVWVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0A1Dlw9Bi6IwAyGEYRieD4fDRVmWNz88z9fr9WYYhkVRFEVK6Xkchpv9fr9YLpePKaWnyWRy1e/3q5TSMsa42O12i9ls9pRSep5Op1c/PzUY/LiV8sX7iqI4uxlFUVRlWeb57OM4Lsufy/I7W8jXJ+SnwyfhuwN58P5ABEFCZFkIsgcRIgQJkWUhyB5EiBAkRJaFIHsQIUKQEFkWguxBhAhBQmRZCLIHESIECZFlIcgeRIgQJESWhSB7ECHfPaj7ib/neX5XH+DRarVez+fzxwPY3W63Q/t9cbdaLper1erBV+S7t37rul6V9f9lCSFCkBBZFoLsQYQIQUJkWQiyBxEiBAn55yhgaFu5EOvHMebDv4y86nQ6dYzxrtPp1DHGZZ7nt/9/P91ut845L/M8v+v1enW32627/f6vfKXEBf3mYZwQWRaC7EGECEFCZFkIsgcRIgQJkWUhyB5EiBAkRJaFIHsQIUKQEFkWguxBhAhBQmRZCLIHESIECZFlIcgeRIgQJESWhSB7ECFC0Lvea1q56Kfz1DUAAAAASUVORK5CYII="'
        except Exception as e:
            print(f"Error processing image {path}: {e}")
            # エラーの場合もプレースホルダーを使用
            return 'xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGQAAABkCAYAAABw4pVUAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH4QoTCxUHCzJQMQAAAB1pVFh0Q29tbWVudAAAAAAAQ3JlYXRlZCB3aXRoIEdJTVBkLmUHAAACIklEQVR42u3dsU7CQBiA0Y/ExGDi6KiDk5OrgyM+kg/B4AtJHB0dXRzd3AwOJsbqYLUJkUHv/5IupUf5coVLSdxut7vLMAxPwzg8l2V53drv8Xi8Ho/Hm1arVWVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0GVZ5tBlWebQZVnm0A1Dlw9Bi6IwAyGEYRieD4fDRVmWNz88z9fr9WYYhkVRFEVK6Xkchpv9fr9YLpePKaWnyWRy1e/3q5TSMsa42O12i9ls9pRSep5Op1c/PzUY/LiV8sX7iqI4uxlFUVRlWeb57OM4Lsufy/I7W8jXJ+SnwyfhuwN58P5ABEFCZFkIsgcRIgQJkWUhyB5EiBAkRJaFIHsQIUKQEFkWguxBhAhBQmRZCLIHESIECZFlIcgeRIgQJESWhSB7ECHfPaj7ib/neX5XH+DRarVez+fzxwPY3W63Q/t9cbdaLper1erBV+S7t37rul6V9f9lCSFCkBBZFoLsQYQIQUJkWQiyBxEiBAn55yhgaFu5EOvHMebDv4y86nQ6dYzxrtPp1DHGZZ7nt/9/P91ut845L/M8v+v1enW32627/f6vfKXEBf3mYZwQWRaC7EGECEFCZFkIsgcRIgQJkWUhyB5EiBAkRJaFIHsQIUKQEFkWguxBhAhBQmRZCLIHESIECZFlIcgeRIgQJESWhSB7ECFC0Lvea1q56Kfz1DUAAAAASUVORK5CYII="'

    fixed_svg_content = re.sub(pattern, replace_with_embedded_image, svg_content)

    # 修正したSVGを保存
    with open(svg_file, 'w', encoding='utf-8') as f:
        f.write(fixed_svg_content)

    print(f"SVGファイルを修正しました: {svg_file}")

# SVGのみ出力するように設定
with Diagram(
    "AWS Observability eCommerce Architecture",
    show=False,
    filename="aws_observability_ecommerce_architecture",
    outformat="svg",
    graph_attr=graph_attr,
    node_attr=node_attr,
    edge_attr=edge_attr
):

    # ユーザー
    with Cluster("Users"):
        customer = CustomerUser("Customer Browser")
        admin = AdminClient("Admin Browser")

    with Cluster("AWS Cloud"):
        # フロントエンド - 顧客
        with Cluster("Frontend - Customer"):
            cf_customer = CloudFront("CloudFront:\nCustomer Site")
            s3_customer = S3("S3 Bucket:\nCustomer Assets")
            cf_customer >> s3_customer

        # フロントエンド - 管理者
        with Cluster("Frontend - Admin"):
            cf_admin = CloudFront("CloudFront:\nAdmin Panel")
            s3_admin = S3("S3 Bucket:\nAdmin Assets")
            cf_admin >> s3_admin

        # API & 認証
        with Cluster("API & Auth"):
            api_gw = APIGateway("API Gateway:\nREST API")
            cognito = Cognito("Amazon Cognito:\nUser Pools")
            api_gw - Edge(label="Auth") >> cognito

        # バックエンドサービス
        with Cluster("Backend Services"):
            lambda_order = Lambda("Lambda:\nOrder API (Go)")
            lambda_product = Lambda("Lambda:\nProduct API (Go)")
            fargate_stock = Fargate("Fargate:\nStock Service (Go)")
            sqs = SQS("SQS:\nOrder Queue")

            lambda_order >> Edge(label="Stock Check") >> fargate_stock
            lambda_order >> Edge(label="Send Msg") >> sqs
            fargate_stock << Edge(label="Consume Msg") << sqs

        # データストア
        with Cluster("Data Stores"):
            rds = RDS("RDS:\nMySQL/PostgreSQL")
            s3_images = S3("S3 Bucket:\nProduct Images")

            lambda_product >> rds
            lambda_order >> rds
            fargate_stock >> rds
            lambda_product >> Edge(label="Read Images") >> s3_images

        # オブザーバビリティツール
        with Cluster("Observability Tools"):
            cloudwatch = Cloudwatch("CloudWatch:\nLogs, Metrics,\nAlarms, Dashboards")
            xray = XRay("AWS X-Ray")

        # APIへの接続
        api_gw << Edge(label="Public & Future APIs") << cf_customer
        api_gw << Edge(label="Admin APIs") << cf_admin

        # APIからバックエンドへの接続
        api_gw >> Edge(label="/products") >> lambda_product
        api_gw >> Edge(label="/orders") >> lambda_order
        api_gw >> Edge(label="/admin/*") >> lambda_order
        api_gw >> Edge(label="/admin/*") >> lambda_product
        api_gw >> Edge(label="/admin/*") >> fargate_stock

        # オブザーバビリティへの接続
        cf_customer >> Edge(label="Logs/Metrics") >> cloudwatch
        cf_admin >> Edge(label="Logs/Metrics") >> cloudwatch
        api_gw >> Edge(label="Logs/Metrics/Traces") >> cloudwatch
        api_gw >> Edge(label="Traces") >> xray
        lambda_order >> Edge(label="Logs/Metrics/Traces") >> cloudwatch
        lambda_order >> Edge(label="Traces") >> xray
        lambda_product >> Edge(label="Logs/Metrics/Traces") >> cloudwatch
        lambda_product >> Edge(label="Traces") >> xray
        fargate_stock >> Edge(label="Logs/Metrics/Traces") >> cloudwatch
        fargate_stock >> Edge(label="Traces") >> xray
        sqs >> Edge(label="Metrics") >> cloudwatch
        rds >> Edge(label="Logs/Metrics") >> cloudwatch
        cognito >> Edge(label="Logs") >> cloudwatch

    # ユーザーとフロントエンドの接続
    customer >> cf_customer
    admin >> cf_admin

# SVGファイル内の画像パスを修正
svg_file = "aws_observability_ecommerce_architecture.svg"
if os.path.exists(svg_file):
    fix_svg_image_paths(svg_file)
    print(f"SVGファイル '{svg_file}' を生成し、画像を埋め込みました。")
else:
    print(f"エラー: SVGファイル '{svg_file}' が見つかりません。")
