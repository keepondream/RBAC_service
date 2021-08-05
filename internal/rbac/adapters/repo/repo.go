package repo

import "github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"

type Repo struct {
	EntClient *ent.Client
}

func NewRepo(entClient *ent.Client) *Repo {

	return &Repo{
		EntClient: entClient,
	}
}
