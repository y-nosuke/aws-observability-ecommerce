package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// ProductHandler は商品関連のハンドラーを表す構造体
type ProductHandler struct {
	// 後々実際のデータソースに置き換えられる
	products []Product
}

// Product は商品を表す構造体
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	CategoryID  int     `json:"category_id"`
}

// Category はカテゴリーを表す構造体
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PaginatedResponse はページネーション付きの応答を表す構造体
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// NewProductHandler は新しい商品ハンドラーを作成します
func NewProductHandler() *ProductHandler {
	// モックデータを初期化
	return &ProductHandler{
		products: []Product{
			{
				ID:          1,
				Name:        "スマートフォン",
				Description: "高性能なスマートフォン、最新モデル",
				Price:       89900,
				ImageURL:    "https://placehold.co/400x300/4F46E5/FFFFFF?text=スマートフォン",
				CategoryID:  1,
			},
			{
				ID:          2,
				Name:        "ノートパソコン",
				Description: "軽量で高性能なノートパソコン",
				Price:       129900,
				ImageURL:    "https://placehold.co/800x600/4F46E5/FFFFFF?text=ノートパソコン", // 大きいサイズ
				CategoryID:  1,
			},
			{
				ID:          3,
				Name:        "ワイヤレスイヤホン",
				Description: "ノイズキャンセリング機能付きワイヤレスイヤホン",
				Price:       24900,
				ImageURL:    "https://placehold.co/400x300/4F46E5/FFFFFF?delay=2000&text=イヤホン", // 2秒遅延
				CategoryID:  2,
			},
			{
				ID:          4,
				Name:        "スマートウォッチ",
				Description: "健康管理機能付きスマートウォッチ",
				Price:       29900,
				ImageURL:    "https://non-existent-domain.example/smartwatch.jpg", // 404エラー用
				CategoryID:  2,
			},
			{
				ID:          5,
				Name:        "ゲーミングマウス",
				Description: "高精度センサー搭載ゲーミングマウス",
				Price:       8900,
				ImageURL:    "https://placehold.co/400x300/4F46E5/FFFFFF?text=ゲーミングマウス",
				CategoryID:  3,
			},
		},
	}
}

// HandleGetProducts は商品一覧を取得するハンドラー関数
func (h *ProductHandler) HandleGetProducts(c echo.Context) error {
	// リクエストの処理開始をログに記録
	log.Printf("Get products request received from %s", c.RealIP())

	// クエリパラメータからページネーション情報を取得
	page := 1
	pageStr := c.QueryParam("page")
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	pageSize := 10 // デフォルトのページサイズ
	pageSizeStr := c.QueryParam("page_size")
	if pageSizeStr != "" {
		parsedPageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil && parsedPageSize > 0 && parsedPageSize <= 50 {
			pageSize = parsedPageSize
		}
	}

	// カテゴリーIDによるフィルタリング
	categoryID, err := strconv.Atoi(c.QueryParam("category_id"))
	var filteredProducts []Product
	if err == nil && categoryID > 0 {
		// カテゴリーでフィルタリング
		for _, p := range h.products {
			if p.CategoryID == categoryID {
				filteredProducts = append(filteredProducts, p)
			}
		}
	} else {
		// フィルタリングなし
		filteredProducts = h.products
	}

	// 製品の総数
	totalItems := len(filteredProducts)

	// 総ページ数を計算
	totalPages := (totalItems + pageSize - 1) / pageSize

	// 現在のページの開始と終了インデックスを計算
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > totalItems {
		endIndex = totalItems
	}

	// ページに表示する製品を取得
	var pageProducts []Product
	if startIndex < totalItems {
		pageProducts = filteredProducts[startIndex:endIndex]
	} else {
		pageProducts = []Product{}
	}

	// レスポンスを構築
	response := PaginatedResponse{
		Items:      pageProducts,
		TotalItems: totalItems,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	// レスポンスの送信をログに記録
	log.Printf("Get products request completed. Total items: %d, Page: %d, Items returned: %d",
		totalItems, page, len(pageProducts))

	return c.JSON(http.StatusOK, response)
}

// GetCategories はカテゴリー一覧を返すモックデータ
func (h *ProductHandler) GetCategories() []Category {
	return []Category{
		{
			ID:          1,
			Name:        "コンピュータ・タブレット",
			Description: "パソコン、タブレットなどの製品",
		},
		{
			ID:          2,
			Name:        "オーディオ・アクセサリー",
			Description: "イヤホン、スマートウォッチなどのアクセサリー",
		},
		{
			ID:          3,
			Name:        "周辺機器",
			Description: "マウス、キーボードなどの周辺機器",
		},
	}
}

// HandleGetCategories はカテゴリー一覧を取得するハンドラー関数
func (h *ProductHandler) HandleGetCategories(c echo.Context) error {
	// リクエストの処理開始をログに記録
	log.Printf("Get categories request received from %s", c.RealIP())

	categories := h.GetCategories()

	// レスポンスの送信をログに記録
	log.Printf("Get categories request completed. Categories count: %d", len(categories))

	return c.JSON(http.StatusOK, categories)
}
