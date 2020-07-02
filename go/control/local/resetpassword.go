package local

import (
	"itflow/cache"
	"itflow/encrypt"
	"itflow/model"
)

func ResetAdminPassword(password string) error {
	user := &model.User{}

	// 检查是否存在admin账号
	if err := user.CheckHaveAdminUser(); err != nil {
		return err
	}

	// 加密输入的密码
	password = encrypt.PwdEncrypt(password, cache.Salt)

	// 数据库更新加密的密码
	return user.UpdateAdminPassword(password)

}
