package router

type (
	IAccountRouter interface {
		SetAccountUserRouter()
	}
)

func (r router) SetAccountUserRouter() {
	r.routerSetting.POST("/account", r.controller.CreateAccountUser)
}
