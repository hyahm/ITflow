package cache

// var CacheLidLevel map[LevelId]Level
// var CacheLevelLid map[Level]LevelId

// type Level string
// type LevelId int64

// type LevelList []Level
// type StoreLevelId string

// var DefaultLevelId LevelId = 0

// func (li LevelId) ToString() string {
// 	return strconv.FormatInt(int64(li), 10)
// }

// func (li LevelId) ToInt64() int64 {
// 	return int64(li)
// }

// // 状态名组转状态id
// func (li LevelId) Name() Level {
// 	if v, ok := CacheLidLevel[li]; ok {
// 		return v
// 	} else {
// 		return ""
// 	}
// }

// func (l Level) Id() LevelId {
// 	if v, ok := CacheLevelLid[l]; ok {
// 		return v
// 	} else {
// 		return DefaultLevelId
// 	}
// }

// func (l Level) Trim() Level {
// 	return Level(strings.Trim(string(l), " "))
// }

// func (l Level) ToString() string {
// 	return string(l)
// }

// // func ToStore(ll LevelList) StoreLevelId {
// // 	tmp := make([]string, 0)
// // 	for _, v := range ll {
// // 		tmp = append(tmp, v.Id().ToString())
// // 	}
// // 	return StoreLevelId(strings.Join(tmp, ","))
// // }

// // func (sli StoreLevelId) ToShow() LevelList {
// // 	sl := make([]Level, 0)
// // 	for _, v := range strings.Split(string(sli), ",") {
// // 		i64, err := strconv.ParseInt(v, 10, 64)
// // 		if err != nil {
// // 			continue
// // 		}
// // 		sl = append(sl, LevelId(i64).Name())
// // 	}
// // 	return sl
// // }
