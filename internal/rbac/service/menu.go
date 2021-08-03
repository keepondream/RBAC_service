package service

import "context"

type MenuService interface {
	List(ctx context.Context, offset, limit int, sort []string, order string, conditions map[string][]string) (*MenuListResponse, error)
}

type MenuQuery struct {
}

type MenuListResponse struct {
	// 数据集
	Items []MenuQuery `json:"items"`

	// 数量
	TotalCount string `json:"total_count"`
}
