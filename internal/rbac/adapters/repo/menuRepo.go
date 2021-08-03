package repo

import (
	"context"

	"entgo.io/ent/entc/integration/idtype/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/service"
)

type MenuRepo struct {
	client *ent.Client
}

func NewMenuRepo(c *ent.Client) *MenuRepo {
	return &MenuRepo{
		client: c,
	}
}

func (m *MenuRepo) List(ctx context.Context, offset int, limit int, sort []string, order string, conditions map[string][]string) (*service.MenuListResponse, error) {
	data := service.MenuListResponse{
		Items:      []service.MenuQuery{},
		TotalCount: "0",
	}

	return &data, nil
}
