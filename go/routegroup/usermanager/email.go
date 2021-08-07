package usermanager

import (
	"itflow/handle/usercontroller/email"

	"github.com/hyahm/xmux"
)

var UpdateEmailPage *xmux.GroupRoute

func init() {
	UpdateEmailPage = xmux.NewGroupRoute()
	// 获取自己的邮箱
	UpdateEmailPage.Post("/my/email", email.Get)
	// 修改邮箱
	UpdateEmailPage.Post("/email/update", email.Set).ApiDescribe("修改邮箱")
}
