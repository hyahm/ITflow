package cache

// var CacheVidVersion map[VersionId]Version
// var CacheVersionVid map[Version]VersionId

// type Version string
// type VersionId int64

// type StoreVersionIds string

// func (si StoreVersionIds) ToIds() []VersionId {
// 	its := make([]VersionId, 0)
// 	for _, v := range strings.Split(string(si), ",") {
// 		i64, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		its = append(its, VersionId(i64))
// 	}
// 	return its
// }

// func (si StoreVersionIds) ToNames() []Version {
// 	its := make([]Version, 0)
// 	for _, v := range strings.Split(string(si), ",") {
// 		i64, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		its = append(its, VersionId(i64).Name())
// 	}
// 	return its
// }

// var DefaultVersionId VersionId = 0

// func (si VersionId) ToString() string {
// 	return strconv.FormatInt(int64(si), 10)
// }

// func (si VersionId) ToInt64() int64 {
// 	return int64(si)
// }

// // 状态名组转状态id
// func (si VersionId) Name() Version {
// 	if v, ok := CacheVidVersion[si]; ok {
// 		return v
// 	} else {
// 		return ""
// 	}
// }

// func (s Version) Id() VersionId {
// 	s = s.Trim()
// 	if v, ok := CacheVersionVid[s]; ok {
// 		return v
// 	} else {
// 		return DefaultVersionId
// 	}
// }

// func (s Version) ToString() string {
// 	return string(s)
// }
// func (s Version) Trim() Version {
// 	return Version(strings.Trim(string(s), " "))
// }
