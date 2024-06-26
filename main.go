package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net-chat/config"
	"net-chat/global"
	"net-chat/initialize"
	"net-chat/pkg"
	"net/http"
	"runtime"
	"time"
)

var (
	configFile = flag.String("f", "config/config.yaml", "the config file")
)

func init() {
	flag.Parse()

	global.Config = new(config.Config)
	loadEngine := config.NewLoad()
	loadEngine.LoadCfg(*configFile, global.Config)

	_, err := initialize.InitJaeger(global.Config.Runtime)
	if err != nil {
		fmt.Printf("初始化jager出错:%s", err.Error())
	}

	global.Log = initialize.InitLogger(global.Config.Runtime.LogLevel)
	global.DB = initialize.GormMysql(global.Config.MainMySQL, global.Config.Runtime.Mode)

	initialize.InitWsHub()
	initialize.InitConsumeMsg()

	pkg.PwdJwt.LoadConfig(global.Config.Runtime.JwtPrivateKey, global.Config.Runtime.JwtPublicKey)

}

func main() {

	engin := initialize.Engin()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.Config.Runtime.HttpPort),
		Handler:        engin,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf(`
   _____ _           _   
  / ____| |         | |  
 | |    | |__   __ _| |_ 
 | |    | '_ \ / _' | __|
 | |____| | | | (_| | |_ 
  \_____|_| |_|\__,_|\__|

 %s Start
 ------------------------------------------------
 Config Path : %s
 Go Version  : %s
 Gin Version : %s
 Run Mode    : %s
 HttpPort    : %d
 ------------------------------------------------
`,
		global.Config.Runtime.ServerName,
		*configFile,
		runtime.Version(),
		gin.Version,
		gin.Mode(),
		global.Config.Runtime.HttpPort)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("err:%+v\n", err)
	}
}
