package setting

import (
    "github.com/Grey-12/gin-blog/pkg/logging"
    "time"

    "github.com/go-ini/ini"
)

var (
    Cfg *ini.File
    RunMode string
    HTTPPort int
    ReadTimeout time.Duration
    WriteTImeOut time.Duration
    PageSize int
    JwtSecret string
)

func init() {
    var err error
    Cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        logging.SugarLogger.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }
    LoadBase()
    LoadServer()
    LoadApp()
}

func LoadBase() {
    RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
    sec, err := Cfg.GetSection("server")
    if err != nil {
        logging.SugarLogger.Fatalf("Fail to get section 'server': %v", err)
    }
    RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
    HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
    ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    WriteTImeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
    sec, err := Cfg.GetSection("app")
    if err != nil {
        logging.SugarLogger.Fatalf("Fail to get section 'app': %v", err)
    }
    JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
    PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}