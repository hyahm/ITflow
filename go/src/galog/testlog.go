package galog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func control(name string, format string, args ...interface{}) {

	if _, ok := LogName[name]; !ok {
		path := filepath.Join(Logpath, name+".log")

		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}

		LogName[name] = &file{
			Mu:       &sync.Mutex{},
			Filebyte: f,
		}

	}
	LogName[name].Mu.Lock()
	defer LogName[name].Mu.Unlock()
	getLine(name, format, args...)
}

type file struct {
	Filebyte *os.File    // 文件
	Mu       *sync.Mutex // 文件锁
}

func Close(name string) {
	LogName[name].Filebyte.Close()
}

func getLine(name string, format string, args ...interface{}) {
	line := fmt.Sprintf(format, args...)
	now := time.Now().Format("2006-01-02 15:04:05") + "\t" + line + "\n"

	write(name, now)

}

func write(name string, message string) {
	//fmt.Println("write")

	LogName[name].Filebyte.WriteString(message)
	info, err := LogName[name].Filebyte.Stat()

	if err != nil {
		fmt.Printf("not found %v \n", name)
	}
	localtime := time.Now()
	//每天一次来分割
	if EveryDay && localtime.Day() != info.ModTime().Day() {
		prefix := fmt.Sprintf("%d-%d-%d", localtime.Year(), localtime.Month(), localtime.Day())

		os.Rename(info.Name(), prefix+info.Name())
		LogName[name].Filebyte.Close()
		delete(LogName, name)

	} else if FileSize > 0 && info.Size() >= FileSize {
		// 根据文件大小来分割
		prefix := fmt.Sprintf("%d", time.Now().UnixNano())
		os.Rename(info.Name(), prefix+info.Name())
		LogName[name].Filebyte.Close()
		delete(LogName, name)
	}

}

func printfileline() string {

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	return fmt.Sprintln(file, line)
}
