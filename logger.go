//библиотека для логирования
//обертка для github.com/natefinch/lumberjack
package go_common_logger

import (
	"fmt"
	config "github.com/B2BFamily/go_common_config"
	lumberjack "github.com/natefinch/lumberjack"
	"time"
)

//Создание экземпляра логгера на основе конфигурации
//
//Пример корфигурации для автоматического создания подключения:
//	{
//	  "example": {
//	    "filename":"test.txt",
//	    "maxsize":2,
//	    "maxage":22,
//	    "maxbackups":5,
//	    "localtime":true,
//	    "compress":true,
//	  }
//	}
//Пример вызова
//	logg := Create("example")
func Create(configPath string) *Logger {
	logg := new(Logger)
	config.GetConfigPath(configPath, &logg.Config)
	logg.init()
	return logg
}

func (base *Logger) Info(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	base.write("INFO", text)
}
func (base *Logger) Trace(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	base.write("TRACE", text)
}
func (base *Logger) Debug(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	base.write("DEBUG", text)
}
func (base *Logger) Error(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	base.write("ERROR", text)
}
func (base *Logger) Panic(format string, a ...interface{}) {
	text := fmt.Sprintf(format, a...)
	base.write("PANIC", text)
}

//инициализация логгера, если задание параметров проводилось в ручную
func (base *Logger) init() {
	base.logger = lumberjack.Logger{
		Filename:   base.Config.Filename,
		MaxSize:    base.Config.MaxSize,
		MaxBackups: base.Config.MaxBackups,
		MaxAge:     base.Config.MaxAge,
		Compress:   base.Config.Compress,
	}
	base.isInit = true
}

//Запись в лог по определенному стандарту
func (base *Logger) write(mode string, text string) {
	if !base.isInit {
		base.init()
	}
	base.logger.Write([]byte(fmt.Sprintf("%v %6v: %v\n", time.Now().Format("2006-01-02T15:04:05.999999-07:00"), mode, text)))
}
