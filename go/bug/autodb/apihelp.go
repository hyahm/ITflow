package autodb

import (
	"database/sql"
	"fmt"
	"github.com/hyahm/goconfig"
	"itflow/bug/bugconfig"
	"log"
	"strconv"
	"strings"
)

func createapi() {
	var pid int64
	// 判断是否存在项目名
	err := bugconfig.Bug_Mysql.GetOne("select id from apiproject where name=?", goconfig.ReadString("apiname")).Scan(&pid)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	// 不存在就创建
	if pid == 0 {
		pid, err = bugconfig.Bug_Mysql.Insert("insert into apiproject(name,ownerid,auth) values(?,1,0)", goconfig.ReadString("apiname"))
		if err != nil {
			log.Fatal(err)
		}
	} else if goconfig.ReadBool("apicover") {

		_, err = bugconfig.Bug_Mysql.Update("delete from apilist where pid=?", pid)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return
	}

	opts := [][]string{{"'username',2,'必须'", "'password',2,'必须'"}}
	apis :=
		[]string{"'登陆',?,'/login/login',?,'POST',1,'json','',0,'{&#34;username&#34;,&#34;admin@qq.com&#34;,&#34;password&#34;:&#34;admin&#34;}','{&#34;username&#34;: &#34;admin&#34;, &#34;password&#34;: &#34;&#34;, &#34;token&#34;: &#34;7ccfb389542777936a64851c71a8297c06b7d6ae&#34;, &#34;statuscode&#34;: 0, &#34;avatar&#34;: &#34;&#34;}'"}
	for k, va := range apis {
		oid := make([]string, 0)
		for _, v := range opts[k] {
			x, err := bugconfig.Bug_Mysql.Insert(fmt.Sprintf("insert into options(name,tid,need) values(%s)", v))
			if err != nil {
				log.Fatal(err)
			}
			oid = append(oid, strconv.FormatInt(x, 10))
		}
		_, err := bugconfig.Bug_Mysql.Insert(
			fmt.Sprintf("insert into apilist(name,pid,url,opts,methods,uid,calltype,information,hid,resp,result) values(%s)", va),
			pid, strings.Join(oid, ","))
		if err != nil {
			log.Fatal(err)
		}
	}

}
