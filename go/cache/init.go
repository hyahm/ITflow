package cache

import (
	"os"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

// 生成初始化本地缓存

var (
	ImgDir      string
	PrivateKey  string
	ShowBaseUrl string
	Salt        string
	Expirontion time.Duration
	ShareDir    string
)

type UG struct {
	Ugid int64
	Name string
	Uids string
}

// cached
var (
	CacheRidRole       map[int64]string
	CacheRoleRid       map[string]int64
	DefaultCreateSid   int64
	DefaultCompleteSid int64

	CacheEmail *Email

	DEADLINE time.Duration
	SUPERID  int64
)

func LoadConfig() {
	SUPERID = goconfig.ReadInt64("adminid", 1)
	ImgDir = goconfig.ReadString("imgdir", "/data/bugimg/")
	err := os.MkdirAll(ImgDir, 0755)
	if err != nil {
		panic(err)
	}
	ShowBaseUrl = goconfig.ReadWithEndSlash("showbaseurl", " http://127.0.0.1:10001/showimg/")
	Salt = goconfig.ReadString("salt", "hjkkaksjdhfryuooweqzmbvc")
	ShareDir = goconfig.ReadString("sharedir", "/share/")
	// 创建共享文件夹
	err = os.MkdirAll(ShareDir, 0755)
	if err != nil {
		panic(err)
	}

	CacheRidRole = make(map[int64]string, 0)
	CacheRoleRid = make(map[string]int64, 0)

	Expirontion = goconfig.ReadDuration("expiration", 120*time.Minute)
	golog.Info("cookie过期时间为：", Expirontion)
	initCache()
	// 添加一个admin 用户的权限，默认全是1
	cacheemail()

}
