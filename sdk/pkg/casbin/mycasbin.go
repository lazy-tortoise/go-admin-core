package mycasbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	"gorm.io/gorm"

	gormAdapter "github.com/go-admin-team/gorm-adapter/v3"
)

// Initialize the model from a string.
//var text = `
//[request_definition]
//r = sub, obj, act
//
//[policy_definition]
//p = sub, obj, act
//
//[policy_effect]
//e = some(where (p.eft == allow))
//
//[matchers]
//m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
//`

var text = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch2(r.obj, p.obj) && r.act == p.act`

//setup
func Setup(db *gorm.DB, _ string) *casbin.SyncedEnforcer {
	Apter, err := gormAdapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	log.SetLogger(&Logger{})
	e.EnableLog(true)
	return e
}
