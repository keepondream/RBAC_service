package store

import (
	"database/sql"

	"github.com/keepondream/RBAC_service/store/rbacStore"
)

type Store struct {
	*rbacStore.Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: rbacStore.New(db),
	}
}
