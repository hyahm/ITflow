package cache

import (
	"os"
	"runtime"
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
	Expirontion int
	ShareDir    string
	CLASSIFY    = []string{"login", "user", "bug", "version", "project", "env", "statusgroup", "role", "status", "restfulproject", "api", "type", "usergroup", "header", "important", "level", "position"}
)

type Email struct {
	CreateUser bool   `json:"createuser"`
	CreateBug  bool   `json:"createbug"`
	PassBug    bool   `json:"passbug"`
	Id         int64  `json:"id"`
	Port       int    `json:"port"`
	EmailAddr  string `json:"emailaddr"`
	Password   string `json:"password"`
	To         string `json:"to"`
	Code       int    `json:"code"`
}

// cached
var (
	CacheRidRole        map[int64]string
	CacheRidInfo        map[int64]string
	CacheRoleRid        map[string]int64
	CacheEidName        map[int64]string
	CacheEnvNameEid     map[string]int64
	CacheUidRealName    map[int64]string
	CacheVidName        map[int64]string
	CacheVersionNameVid map[string]int64
	CacheUidSgid        map[int64]int64
	CacheUidNickName    map[int64]string
	CacheRidGroup       map[int64]string
	CacheUidFilter      map[int64]StoreStatusId
	DefaultCreateSid    StatusId
	DefaultCompleteSid  StatusId

	CacheUidRid  map[int64]int64
	CacheUidJid  map[int64]int64
	CacheTidName map[int64]string
	CacheNameTid map[string]int64

	CacheSgidGroup map[int64]string

	CacheJidJobname map[int64]string
	CacheJobnameJid map[string]int64
	CacheHidHeader  map[int64]string
	CacheHeaderHid  map[string]int64

	CacheNickNameUid map[string]int64
	CacheRealNameUid map[string]int64
	CacheUidEmail    map[int64]string
	CacheEmail       *Email

	CacheGidGroup map[int64]string
	DEADLINE      time.Duration
	SUPERID       int64
)

func LoadConfig() {
	SUPERID = goconfig.ReadInt64("adminid", 1)
	ImgDir = goconfig.ReadString("imgdir", "/data/bugimg/")
	err := os.MkdirAll(ImgDir, 0755)
	if err != nil {
		panic(err)
	}
	ShowBaseUrl = goconfig.ReadString("showbaseurl", " http://127.0.0.1:10001/showimg")
	Salt = goconfig.ReadString("salt", "hjkkaksjdhfryuooweqzmbvc")
	ShareDir = goconfig.ReadString("sharedir", "/share/")
	// 创建共享文件夹
	err = os.MkdirAll(ShareDir, 0755)
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		if ShareDir[len(ShareDir)-1:] == "\\" {
			ShareDir = ShareDir[:len(ShareDir)-1]
		}
	} else {
		if ShareDir[len(ShareDir)-1:] == "/" {
			ShareDir = ShareDir[:len(ShareDir)-1]
		}
	}
	CacheSidStatus = make(map[StatusId]Status, 0)
	CacheRidGroup = make(map[int64]string, 0)
	CacheRidInfo = make(map[int64]string, 0)
	CacheRidRole = make(map[int64]string, 0)
	CacheIidImportant = make(map[ImportantId]Important, 0)
	CacheRoleRid = make(map[string]int64, 0)
	CacheImportantIid = make(map[Important]ImportantId, 0)
	CachePidProject = make(map[ProjectId]Project, 0)
	CacheUidSgid = make(map[int64]int64, 0)
	CacheUidRid = make(map[int64]int64, 0)
	CacheEidName = make(map[int64]string, 0)
	CacheUidRealName = make(map[int64]string, 0)
	CacheUidNickName = make(map[int64]string, 0)
	CacheVidName = make(map[int64]string, 0)
	CacheGidGroup = make(map[int64]string, 0)
	CacheNickNameUid = make(map[string]int64, 0)
	CacheStatusSid = make(map[Status]StatusId, 0)
	CacheRealNameUid = make(map[string]int64, 0)
	CacheProjectPid = make(map[Project]ProjectId, 0)
	CacheVersionNameVid = make(map[string]int64, 0)
	CacheEnvNameEid = make(map[string]int64, 0)
	CacheLidLevel = make(map[LevelId]Level, 0)
	CacheJidJobname = make(map[int64]string, 0)
	CacheJobnameJid = make(map[string]int64, 0)
	CacheLevelLid = make(map[Level]LevelId, 0)
	CacheSgidGroup = make(map[int64]string, 0)
	CacheUidFilter = make(map[int64]StoreStatusId, 0)
	CacheUidJid = make(map[int64]int64, 0)
	CacheUidEmail = make(map[int64]string, 0)

	CacheHidHeader = make(map[int64]string, 0)
	CacheTidName = make(map[int64]string, 0)
	CacheHeaderHid = make(map[string]int64, 0)
	CacheNameTid = make(map[string]int64, 0)
	CacheEmail = &Email{}

	golog.Info("cookie过期时间为：", goconfig.ReadDuration("expiration", 120*time.Minute))
	initCache()
	// 添加一个admin 用户的权限，默认全是1
	cacheemail()

}

func addXieGang(path string) string {
	l := len(path)
	// 如果是windows
	if runtime.GOOS == "windows" {
		if path[l-1:] != "\\" {
			return path + "\\"

		}

	} else {
		if path[l-1:] != "/" {
			return path + "/"

		}
	}
	return path
}
