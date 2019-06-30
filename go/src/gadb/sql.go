package gadb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hyahm/goconfig"
	"log"
	"strconv"
	"strings"
)

const SESSION = "session"

var CONNECTDBERROR = errors.New("can't connect db")

type Sconfig struct {
	UserName string
	Password string
	Host     string
	Port     int
	DbName   string
	Driver   string
}

func NewSqlConfig() *Sconfig {
	return &Sconfig{
		UserName: goconfig.ReadString("mysqluser"),
		Password: goconfig.ReadString("mysqlpwd"),
		Host:     goconfig.ReadString("mysqlhost"),
		Port:     goconfig.ReadInt("mysqlport"),
		DbName:   goconfig.ReadString("mysqldb"),
		Driver:   goconfig.ReadString("sqldriver"),
	}
}

type Db struct {
	Db *sql.DB
	Tx *sql.Tx
}

func (md *Sconfig) ConnDB() (*Db, error) {
	var connstring string

	connstring = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		md.UserName, md.Password, md.Host, md.Port, md.DbName,
	)

	db, err := sql.Open(md.Driver, connstring)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Db{
		db,
		nil,
	}, nil
}

func (d *Db) GetConnections() int {
	return d.Db.Stats().OpenConnections
}

func (d *Db) Update(cmd string, args ...interface{}) (int64, error) {

	result, err := d.Db.ExecContext(context.Background(), cmd, args...)
	if err != nil {
		return -1, err
	}
	lastid, lasterr := result.RowsAffected()
	if lasterr != nil {
		return -1, lasterr
	}

	return lastid, nil

}

func (d *Db) InsertWithID(cmd string, args ...interface{}) (int64, error) {
	result, err := d.Db.ExecContext(context.Background(), cmd, args...)

	if err != nil {
		return 0, err
	}

	l, e := result.LastInsertId()
	if e != nil {
		log.Fatal(e)
	}

	if e != nil {
		return 0, nil
	}

	return l, nil
}

func (d *Db) Insert(cmd string, args ...interface{}) (int64, error) {

	result, err := d.Db.ExecContext(context.Background(), cmd, args...)

	if err != nil {
		return -1, err
	}

	l, e := result.LastInsertId()

	if e != nil {
		log.Fatal(e)
	}

	if e != nil {
		return 0, nil
	}

	return l, nil
}

func (d *Db) InsertMany(cmd string, args [][]interface{}) (int, error) {

	l := len(args)
	if l > 1 {
		index := strings.Index(cmd, "values(")
		addcmd := "," + cmd[index+6:]
		for i := 0; i < l-1; i++ {
			cmd = cmd + addcmd
		}
		newargs := make([]interface{}, 0)

		for _, v := range args {
			newargs = append(newargs, v...)
		}

		_, err := d.Insert(cmd, newargs...)
		if err != nil {
			return 0, err
		}
	} else if l == 1 {
		d.Insert(cmd, args[0]...)
	} else {
		return 0, errors.New("format error")
	}

	return l, nil
}

func (d *Db) GetRows(cmd string, args ...interface{}) (*sql.Rows, error) {
	row, err1 := d.Db.QueryContext(context.Background(), cmd, args...)
	if err1 != nil {
		return nil, err1
	}
	return row, nil
}

func (d *Db) GetOne(cmd string, args ...interface{}) *sql.Row {
	return d.Db.QueryRowContext(context.Background(), cmd, args...)
}

//func (d *Db) SelectSlice_Slice(cmd string, args ...interface{}) (record [][]string, r int, err error) {
//	logger := golog.NewLog()
//	starttime := time.Now()
//	selectsql := cmdtostring(cmd, args...)
//	row, err1 := d.Db.QueryContext(context.Background(), cmd, args...)
//
//	if err1 != nil {
//		logger.FailSqlLog(selectsql)
//		logger.ErrorLog(err1.Error())
//		return nil, -1, err1
//	}
//
//	columns, _ := row.Columns()
//	scanArgs := make([]interface{}, len(columns))
//	values := make([]interface{}, len(columns))
//
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//	cols := len(values)
//
//	rows := 0
//
//	record = make([][]string, 0)
//
//	for row.Next() {
//		scanerr := row.Scan(scanArgs...)
//		if scanerr != nil {
//			logger.FailSqlLog(selectsql)
//			logger.ErrorLog(scanerr.Error())
//			return nil, -1, scanerr
//		}
//
//		tmprecord := make([]string, cols)
//
//		for i, colv := range values {
//
//			if colv != nil {
//
//				if _, ok := colv.(int64); ok {
//					tmprecord[i] = strconv.FormatInt(colv.(int64), 10)
//				} else {
//					tmprecord[i] = string(colv.([]byte))
//				}
//			}
//		}
//		record = append(record, tmprecord)
//		rows++
//	}
//	if logger.Selectfile != "" {
//		exectime := time.Since(starttime).Nanoseconds()
//		logger.Selectlog(selectsql, exectime)
//	}
//	return record, rows, nil
//}
//
////[]map[string]string
//func (d *Db) SelectJson_String(cmd string, args ...interface{}) (vuestr string, err error) {
//	logger := golog.NewLog()
//	starttime := time.Now()
//	row, err1 := d.Db.QueryContext(context.Background(), cmd, args...)
//	if err1 != nil {
//		return "", err1
//	}
//
//	columns, err2 := row.Columns()
//	if err2 != nil {
//		return "", err2
//	}
//
//	values := make([]sql.RawBytes, len(columns))
//	scanArgs := make([]interface{}, len(values))
//
//	ret := make([]map[string]string, 0)
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//
//	for row.Next() {
//		err = row.Scan(scanArgs...)
//		if err != nil {
//			return "", err
//		}
//		var value string
//		vmap := make(map[string]string, len(scanArgs))
//		for i, col := range values {
//			if col == nil {
//				value = "NULL"
//			} else {
//				value = string(col)
//			}
//			vmap[columns[i]] = value
//		}
//		ret = append(ret, vmap)
//
//	}
//	if logger.Selectfile != "" {
//		exectime := time.Since(starttime).Nanoseconds()
//		selectsql := cmdtostring(cmd, args...)
//		logger.Selectlog(selectsql, exectime)
//	}
//	data, err := json.Marshal(ret)
//	vuestr = string(data)
//	return vuestr, nil
//}
//
//func (d *Db) SelectJson_StringWithIndex(cmd string, args ...interface{}) (j string, err error) {
//	logger := golog.NewLog()
//	starttime := time.Now()
//
//	row, err1 := d.Db.QueryContext(context.Background(), cmd, args...)
//	if err1 != nil {
//		return "", err1
//	}
//
//	columns, err2 := row.Columns()
//	if err2 != nil {
//		return "", err2
//	}
//	ret := make(map[string]map[string]string, 0)
//	values := make([]sql.RawBytes, len(columns))
//	scanArgs := make([]interface{}, len(values))
//
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//	r := 0
//	for row.Next() {
//		err = row.Scan(scanArgs...)
//		if err != nil {
//			return "", err
//		}
//		var value string
//		vmap := make(map[string]string, 0)
//		for i, col := range values {
//			if col == nil {
//				value = "NULL"
//			} else {
//				value = string(col)
//			}
//			vmap[columns[i]] = value
//		}
//		ret["r"+strconv.Itoa(r)] = vmap
//		r++
//	}
//	if logger.Selectfile != "" {
//		exectime := time.Since(starttime).Nanoseconds()
//		selectsql := cmdtostring(cmd, args...)
//		logger.Selectlog(selectsql, exectime)
//	}
//	index := make(map[string]string, 0)
//	index["index"] = strconv.Itoa(r)
//	ret["max"] = index
//
//	data, err := json.Marshal(ret)
//	j = string(data)
//	return j, nil
//}
//
//func (d *Db) SelectJson_String_Add(addkey string, addvalue string, cmd string, args ...interface{}) (j string, err error) {
//	logger := golog.NewLog()
//	starttime := time.Now()
//
//	row, err1 := d.Db.QueryContext(context.Background(), cmd, args...)
//	if err1 != nil {
//		return "", err1
//	}
//
//	columns, err2 := row.Columns()
//	if err2 != nil {
//		return "", err2
//	}
//	ret := make(map[string]map[string]string, 0)
//	values := make([]sql.RawBytes, len(columns))
//	scanArgs := make([]interface{}, len(values))
//
//	for i := range values {
//		scanArgs[i] = &values[i]
//	}
//	r := 0
//	for row.Next() {
//		err = row.Scan(scanArgs...)
//		if err != nil {
//			return "", err
//		}
//		var value string
//		vmap := make(map[string]string, 0)
//		for i, col := range values {
//			if col == nil {
//				value = "NULL"
//			} else {
//				value = string(col)
//			}
//			vmap[columns[i]] = value
//		}
//		ret["r"+strconv.Itoa(r)] = vmap
//		r++
//	}
//	// 如果没有写配置，就不写入日志
//	if logger.Selectfile != "" {
//		exectime := time.Since(starttime).Nanoseconds()
//		selectsql := cmdtostring(cmd, args...)
//		logger.Selectlog(selectsql, exectime)
//	}
//
//	index := make(map[string]string, 0)
//	add := make(map[string]string, 0)
//	add[addkey] = addvalue
//	index["index"] = strconv.Itoa(r)
//	ret["max"] = index
//	ret[addkey] = add
//
//	data, err := json.Marshal(ret)
//	j = string(data)
//	return j, nil
//}

// 还原sql
func cmdtostring(cmd string, args ...interface{}) string {

	var logstr string

	for _, v := range args {
		switch v.(type) {
		case int64:
			logstr = "'" + strconv.FormatInt(v.(int64), 10) + "'"
		case int:
			logstr = "'" + strconv.Itoa(v.(int)) + "'"
		case string:
			logstr = "'" + v.(string) + "'"
		default:
			logstr = "'" + v.(string) + "'"
			//return
		}
		cmd = strings.Replace(cmd, "?", "%s", 1)
		cmd = fmt.Sprintf(cmd, logstr)
		// public.Printlog(logstr + strconv.Itoa(i))

	}
	return cmd
}
