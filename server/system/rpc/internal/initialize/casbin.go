package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"sync"
)

var (
	syncedEnforcer  *casbin.SyncedEnforcer
	initCasbinModel sync.Once
)

func InitCasbin(db *gorm.DB) *casbin.SyncedEnforcer {
	initCasbinModel.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(db)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			logx.Error("InitCasbin: import model fail!", err)
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
