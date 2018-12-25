package go_common_logger

import (
	lumberjack "github.com/natefinch/lumberjack"
)

//Объект настройки логгера
type LoggerConfig struct {
	Filename   string `json:"filename"`   //имя файла, в который будут писаться логи
	MaxSize    int    `json:"maxsize"`    //(МБ) максимальный размер файла, в который будут писаться логи
	MaxAge     int    `json:"maxage"`     //Кол-во дней, которое хранятся логи
	MaxBackups int    `json:"maxbackups"` //Кол-во фалов логов, которые хранятся
	LocalTime  bool   `json:"localtime"`  //Используется ли локальное время для формирования названия файлов
	Compress   bool   `json:"compress"`   //Архивация старых логов
}

//Объект логгера
type Logger struct {
	logger lumberjack.Logger
	Config LoggerConfig
	isInit bool
}
