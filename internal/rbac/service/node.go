package service

import (
	"context"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type Node struct {
	*Service
}

func NewNode(s *Service) *Node {
	return &Node{
		Service: s,
	}
}

type Noder interface {
	Create(ctx context.Context, params ports.PostNodesJSONBody) (*ports.NodeInfoResponse, error)
	IsUnique(ctx context.Context, tenant, name, node_type string) error
	GetById(ctx context.Context, id string) (*ports.NodeInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	List(ctx context.Context, params ports.GetNodesParams) (*ports.NodeListResponse, error)
	Update(ctx context.Context, params ports.PatchNodesIdJSONBody, id string) (*ports.NodeInfoResponse, error)
}

func (s *Node) Create(ctx context.Context, params ports.PostNodesJSONBody) (*ports.NodeInfoResponse, error) {
	r := repo.NewNode(s.Repo)
	res, err := r.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	// 同步节点对应的权限
	s.SyncCasbinForNode(ctx, res.Id, string(res.Tenant))

	return res, nil
}

func (s *Node) IsUnique(ctx context.Context, tenant string, name string, node_type string) error {
	r := repo.NewNode(s.Repo)
	return r.IsUnique(ctx, tenant, name, node_type)
}

func (s *Node) GetById(ctx context.Context, id string) (*ports.NodeInfoResponse, error) {
	r := repo.NewNode(s.Repo)
	return r.GetById(ctx, id)
}

func (s *Node) DeleteById(ctx context.Context, id string) error {
	r := repo.NewNode(s.Repo)

	res, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}

	err = r.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	// 同步节点对应的权限
	s.SyncCasbinForNode(ctx, res.Id, string(res.Tenant))

	return nil
}

func (s *Node) List(ctx context.Context, params ports.GetNodesParams) (*ports.NodeListResponse, error) {
	r := repo.NewNode(s.Repo)
	return r.List(ctx, params)
}

func (s *Node) Update(ctx context.Context, params ports.PatchNodesIdJSONBody, id string) (*ports.NodeInfoResponse, error) {
	r := repo.NewNode(s.Repo)
	res, err := r.Update(ctx, params, id)
	if err != nil {
		return nil, err
	}

	// 同步节点对应的权限
	s.SyncCasbinForNode(ctx, res.Id, string(res.Tenant))

	return res, nil
}
