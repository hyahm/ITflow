package gaconfig

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//const middle = "========="
const SEP = "=" // key 和 value 分隔符

var ConfigPath string // 配置文件路径，保存后方便重新加载配置文件
var ConfigKeyValue map[string]string
var NOTE = "#[" // #和[开头的为注释

// 读取配置文件

func InitConf(configpath string) {
	fp := flag.String("conf", configpath, "specify configfile path")

	flag.Parse()
	//filepath.Dir(*fp)
	ConfigPath = filepath.Join(filepath.Dir(*fp), filepath.Base(*fp))

	fmt.Println("configfile:", ConfigPath)
	LoopKey()
	// 检查日志目录
	logpath := ReadString("logpath")
	err := os.MkdirAll(logpath, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// 读取配置文件到全局变量，并检查重复项
func LoopKey() {
	var err error
	//获取文件字节
	cf, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	line := 0
	var bs []string
	if runtime.GOOS == "windows" {
		bs = strings.Split(string(cf), "\r\n")
	} else {
		bs = strings.Split(string(cf), "\n")
	}

	//换行符切割字符串
	ConfigKeyValue = make(map[string]string, 0)
	for i := 0; i < len(bs); i++ {
		//fmt.Println()

		line++

		//去掉2边的空格
		sbs := strings.Trim(bs[i], " ")
		//  #开头是注释， [ 开头是模块 , 空行

		if sbs == "" || strings.ContainsAny(sbs[0:1], NOTE) {
			continue
		}
		index := strings.Index(sbs, SEP)
		key := strings.Trim(sbs[:index], " ")

		if _, ok := ConfigKeyValue[key]; ok {
			log.Fatal(fmt.Sprintf("key:%s duplicate，line:%d \n", key, line))
		} else {
			fmt.Printf("Key:%s --- Value: %s \n", key, strings.Trim(sbs[index+1:], " "))
			ConfigKeyValue[key] = strings.Trim(sbs[index+1:], " ")
		}

	}
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
