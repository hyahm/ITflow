package cache

import (
	"os"
	"strconv"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

// 生成初始化本地缓存

var (
	// CacheRoleRid map[string]int64
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
	CacheRoleID map[int64]PageInfo
	DEADLINE    time.Duration
	SUPERID     int64
)

func LoadConfig() {
	var err error
	SUPERID, err = strconv.ParseInt(goconfig.ReadEnv("ADMINID"), 10, 64)
	if err != nil {
		SUPERID = goconfig.ReadInt64("adminid", 1)
	}

	ImgDir = goconfig.ReadString("imgdir", "/data/bugimg/")
	err = os.MkdirAll(ImgDir, 0755)
	if err != nil {
		panic(err)
	}
	ShowBaseUrl = goconfig.ReadEnv("SHOW_URL",
		goconfig.ReadWithEndSlash("showbaseurl", " http://127.0.0.1:10001/showimg/"))
	Salt = goconfig.ReadEnv("SALT", goconfig.ReadString("salt", "hjkkaksjdhfryuooweqzmbvc"))
	ShareDir = goconfig.ReadString("sharedir", "/share/")
	// 创建共享文件夹
	err = os.MkdirAll(ShareDir, 0755)
	if err != nil {
		panic(err)
	}

	// CacheRoleRid = make(map[string]int64, 0)

	Expirontion = goconfig.ReadDuration("expiration", 120*time.Minute)
	golog.Info("cookie过期时间为：", Expirontion)

	// 添加一个admin 用户的权限，默认全是1

}
