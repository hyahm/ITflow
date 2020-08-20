package assist

import (
	"fmt"
	"itflow/db"
	"strings"
)

type Names []string
type Uid string

func (names Names) RealNameToUsersId() Uid {
	// nickname 转 id的字符串拼接
	ul := make([]string, 0)
	realnames := strings.Join(names, ",")
	rows, err := db.Mconn.GetRows(fmt.Sprintf("select id from user where realname in ('%s')", realnames))
	if err != nil {
		return ""
	}
	realname := new(string)
	for rows.Next() {
		err = rows.Scan(realname)
		ul = append(ul, *realname)
	}

	return Uid(strings.Join(ul, ","))
}

func (uids Uid) UsersIdToRealName() Names {
	// nickname 转 id的字符串拼接
	ul := make([]string, 0)
	rows, err := db.Mconn.GetRows(fmt.Sprintf("select realname from user where id in ('%s')", uids))
	if err != nil {
		return ul
	}
	realname := new(string)
	for rows.Next() {
		err = rows.Scan(realname)
		ul = append(ul, *realname)
	}

	return ul
}

func FormatUserlistToShow(userlist string) Names {
	// 用户id转 真实姓名
	al := make([]string, 0)
	rows, err := db.Mconn.GetRows(fmt.Sprintf("select realname from user where id in ('%s')", userlist))
	if err != nil {
		return al
	}
	realname := new(string)
	for rows.Next() {
		err = rows.Scan(realname)
		al = append(al, *realname)
	}
	return al
}
