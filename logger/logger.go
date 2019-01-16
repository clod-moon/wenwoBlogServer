package logger

import (
	"os"
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"sync"
)

var Log = NewLogger()

type Logger struct {
	log *logs.BeeLogger
	day string
	m sync.Mutex
}

func NewLogger() *Logger{
	return &Logger{log:logs.NewLogger()}
}

func init(){

	//day := fmt.Sprintf("%s%d%s")
	Log.day = time.Now().Format("20060102")
	path := fmt.Sprintf("./logs/%s",Log.day)
	CreateLogDir(path)
}

func WriteLog(file string ,format string, v ...interface{}){
	config := fmt.Sprintf(`{"filename":"./logs/%s/%s.log"}`,Log.day,file)
	fmt.Println(config)
	Log.m.Lock()
	Log.log.SetLogger("file",config)
	Log.log.Error( format, v...)
	Log.log.Flush()
	Log.m.Unlock()
}

func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateLogDir(path string) {

	b, err := PathExists(path)
	if err != nil {
		beego.SetLogger("file", `{"filename":"./logs/init.log"}`)
		beego.SetLogFuncCall(true)
		beego.Error(err.Error())
		return
	}
	if !b {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			beego.Error("mkdir failed![%v]\n", err)
		} else {
			beego.Info("mkdir success!\n")
		}
	}
}

func GetLogPath(fileName string) string{
	return fmt.Sprintf("./logs/%s/%s.log",time.Now().Day(),fileName)
}