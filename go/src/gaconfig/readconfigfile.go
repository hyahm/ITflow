package gaconfig

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFloat(key string) float64 {
	f64, err := strconv.ParseFloat(ConfigKeyValue[key], 64)
	if err != nil {
		return 0
	}
	return f64
}

func ReadFile(key string) string {
	bs, err := ioutil.ReadFile(ConfigKeyValue[key])
	if err != nil {
		return ""
	}
	return string(bs)
}

func ReadString(key string) string {
	return ConfigKeyValue[key]
}

// 返回int
func ReadInt(key string) int {
	i, err := strconv.Atoi(ConfigKeyValue[key])
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return i
}

func ReadInt16(key string) int16 {
	i := ReadInt(key)
	// 如果大于取值区间，返回0
	if i > ((1<<16)/2)-1 || i < -((1<<16)/2) {
		fmt.Println("out of range")
		return 0
	}
	return int16(ReadInt(key))
}

// 2边需要用到引号
func ReadPassword(key string) string {
	v := ConfigKeyValue[key]
	// 如果头尾不是"
	l := len(v)
	if string(v[0]) != "\"" || string(v[l-1:]) != "\"" {
		fmt.Println("not found quote")
		return ""
	}
	return v[1 : l-1]
}

func ReadBool(key string) bool {
	if ConfigKeyValue[key] == "true" {
		return true
	}
	return false
}

func ReadInt64(key string) int64 {
	i, err := strconv.ParseInt(ConfigKeyValue[key], 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return i
}

func ReadMap(key string) map[string]interface{} {
	// value only three format
	x := ConfigKeyValue[key]
	//x := `{"asdf":"ajsdfkl","type":6,"has":true}`

	l := len(x)
	kv := make(map[string]interface{}, 0)
	if string(x[0]) == "{" && string(x[l-1]) == "}" {

		// 去头和尾和空格
		newstr := strings.Trim(x[1:l-1], " ")

		// 逗号分开组
		sl := strings.Split(newstr, ",")
		for _, v := range sl {
			// 去头尾空格
			var k string

			mstr := strings.Trim(v, " ")
			fmt.Println(mstr)
			// 冒号分开,避免values存在:,以第一个冒号分割
			index := strings.Index(mstr, ":")
			//key是：,要去掉头尾空格
			keyquote := strings.Trim(mstr[:index], " ")
			//还要去掉2边的引号,如果没有冒号，格式不正确
			kl := len(keyquote)
			if string(keyquote[0]) == "\"" && string(keyquote[kl-1]) == "\"" && kl > 2 {
				// key去掉2边的空格， 如果是空的，key值不能为空
				k = strings.Trim(keyquote[1:kl-1], " ")

				if k == "" {
					fmt.Println("key 值不能为空")
					fmt.Println("key:%s is not an right format", "key")
				}
			} else {
				fmt.Println("key 缺少双引号或者没有值")
				fmt.Println("key:%s is not an right format", "key")
				return nil
			}

			// value是：去头尾空格
			valuequote := strings.Trim(mstr[index+1:], " ")
			//查看左右2边是否存在双引号
			vl := len(valuequote)
			if string(valuequote[0]) == "\"" && string(valuequote[vl-1]) == "\"" {
				// 存在双引号，那就是字符串,获取里面的值，不去掉2边的空格
				value := valuequote[1 : vl-1]
				kv[k] = value
				continue

			} else {
				// 否则是数字或者布尔值，先判断布尔值
				if valuequote == "true" {
					kv[k] = true
					continue
				} else if valuequote == "false" {
					kv[k] = false
					continue
				} else if v, err := strconv.ParseInt(valuequote, 10, 64); err == nil {
					//判断数字int64
					kv[k] = v
					continue
				} else if v, err := strconv.ParseFloat(valuequote, 64); err == nil {
					//判断数字float64
					kv[k] = v
					continue
				} else {

					fmt.Println("value 只支持string,int64,float64,bool")
					return nil
				}

			}
		}

	} else {
		fmt.Println("头尾缺少大括号")
		fmt.Println("key:%s is not an right format", "key")
		return nil
	}

	return kv
}

func ReadIntArray(key string) []int {
	il := make([]int, 0)
	vl := ConfigKeyValue[key]
	vlength := len(vl)
	if vlength == 0 {
		return il
	}
	if vl[0:1] == "[" && vl[vlength-1:vlength] == "]" {
		vlist := strings.Split(vl[1:vlength-1], ",")
		//如果没值就返回空数组
		if len(vlist) == 0 {
			return il
		}
		for _, v := range vlist {
			//去掉2边的空格
			i, err := strconv.Atoi(strings.Trim(v, " "))
			if err != nil {
				log.Fatal(fmt.Sprintf("key:%s,%v", key, err))
			}
			il = append(il, i)
		}
		return il
	} else {
		log.Fatal("key:%s,not an int array format \n", key)
	}
	return il
}

func ReadStringArray(key string) []string {
	sl := make([]string, 0)
	vl := ConfigKeyValue[key]
	vlength := len(vl)
	if vlength == 0 {
		return sl
	}
	if vl[0:1] == "[" && vl[vlength-1:vlength] == "]" {
		vlist := strings.Split(vl[1:vlength-1], ",")
		//如果没值就返回空数组
		if len(vlist) == 0 {
			return sl
		}
		for _, v := range vlist {
			//去掉2边的空格
			stringquote := strings.Trim(v, " ")
			// 检查2边是否有双引号
			ql := len(stringquote)
			if ql == 0 {
				return sl
			}
			if stringquote[0:1] == "\"" && stringquote[ql-1:ql] == "\"" {
				stringlist := stringquote[1 : ql-1]
				if stringlist == "" {
					continue
				}
				sl = append(sl, stringlist)
			} else {
				log.Fatal(fmt.Sprintf("key:%s,value must be has quote \n", key))
				return sl
			}

		}
		return sl
	} else {
		log.Fatal("key:%s,not an int array format \n", key)
	}
	return sl
}
