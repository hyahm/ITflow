package cache

//

// type Namer interface {
// 	Id() IDer
// 	Trim()
// 	ToString() string
// }

// type IDer interface {
// 	ToString() string
// 	ToInt64() int64
// 	Name() Namer
// }

// type DbIds interface {
// }

// func ToStore(names []Namer) string {
// 	tmp := make([]string, 0)
// 	for _, v := range names {
// 		tmp = append(tmp, v.Id().ToString())
// 	}
// 	return strings.Join(tmp, ",")
// }

// func ToShow(s string) []string {
// 	sl := make([]string, 0)
// 	for _, v := range strings.Split(string(ssi), ",") {
// 		i64, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			continue
// 		}
// 		sl = append(sl, IDer(i64).Name())
// 	}
// 	return sl
// }
