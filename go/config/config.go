package main

import (
	"fmt"
	"os"

	"github.com/hyahm/goconfig"
)

// 初始化配置文件, 会清空原有文件内容
func main() {
	conf := "../bug.ini"
	if len(os.Args) > 1 {
		conf = os.Args[1]
	}
	fmt.Println(conf)
	goconfig.InitWriteConf(conf, goconfig.INI)
	goconfig.WriteBool("httpproxy", true, "是否使用了代码，为了获取ip，可能不起作用")
	goconfig.WriteString("listenaddr", ":10001", "监听地址")
	goconfig.WriteString("imgdir", "/data/bugimg/", "存放图片的目录")
	goconfig.WriteString("showbaseurl", "http://127.0.0.1:10001/showimg", "图片显示的地址(用接口的地址)")
	goconfig.WriteString("salt", "hjkkaksjdhfryuooweqzmbvc", "盐值，建议修改，然后用curl http://127.0.0.1:10001/admin/reset?password=123 来修改admin密码")
	goconfig.WriteString("sharedir", "/share/", "共享文件夹根目录")

	// goconfig.WriteString("redis.pwd", "")
	// goconfig.WriteString("redis.host", "127.0.0.1:6379")
	// goconfig.WriteInt("redis.db", 0)
	goconfig.WriteString("expiration", "120m", "token 过期时间")

	goconfig.WriteNotesForModule("ssl", "ssl, 使用ssl")
	goconfig.WriteBool("ssl.on", false)
	goconfig.WriteString("ssl.cert", "")
	goconfig.WriteString("ssl.key", "")

	goconfig.WriteString("log.path", "", "日志目录, 不设置就控制台输出")
	goconfig.WriteInt64("log.size", 0, "日志大小备份一次， 0为不切割大小")
	goconfig.WriteBool("log.everyday", false, "每天备份一次 大小也存在的话，此项优先 ，false为不每天备份一次")

	goconfig.WriteString("mysql.user", "root")
	goconfig.WriteString("mysql.pwd", "")
	goconfig.WriteString("mysql.host", "127.0.0.1")
	goconfig.WriteInt("mysql.port", 3306)
	goconfig.WriteString("mysql.db", "project")
	goconfig.FlushWrite()
}
