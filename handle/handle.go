package handle

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/keepondream/RBAC_service/service"
	"github.com/keepondream/RBAC_service/store"
	"github.com/keepondream/RBAC_service/utils"
	"github.com/sirupsen/logrus"
)

type Handle struct {
	Config   *utils.Config
	Enforcer *casbin.Enforcer
	Logger   *logrus.Logger
	Store    *store.Store
	Service  *service.Service
}

func NewHandle(config *utils.Config, e *casbin.Enforcer, logger *logrus.Logger, store *store.Store, service *service.Service) *Handle {
	return &Handle{
		Config:   config,
		Enforcer: e,
		Logger:   logger,
		Store:    store,
		Service:  service,
	}
}

type Table = string

// 基础数据
const AllPermission Table = "all::permission"           // 所有权限 (同时代表: 超级管理策略)
const AllPermissionGroup Table = "all::permissionGroup" // 所有权限组
const AllRole Table = "all::role"                       // 所有角色
const AllUserGroup Table = "all::userGroup"             // 所有用户组

type PolicyType = string

// 策略标识
const PRole PolicyType = "role::"                       // 角色策略
const PUser PolicyType = "user::"                       // 用户策略
const PPermissionGroup PolicyType = "permissionGroup::" // 权限组策略
const PUserGroup PolicyType = "userGroup::"             // 用户组策略

// 构造策略标识
func MakePolicySign(policyType PolicyType, id string) string {
	return fmt.Sprintf("%s%s", policyType, id)
}
