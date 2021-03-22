package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

var jwtSecret string

func init() {

	val, ex := os.LookupEnv("SECRET")
	if !ex {
		time.Sleep(time.Duration(5) * time.Second)
		panic("must set env: SECRET")
	}
	jwtSecret = val

	var dbAddr, dbUser, dbPassword string
	dbAddr, ex = os.LookupEnv("DBADDR")
	if !ex {
		time.Sleep(time.Duration(3) * time.Second)
		panic("must set env: DBADDR")
	}

	dbUser, ex = os.LookupEnv("DBUSER")
	if !ex {
		time.Sleep(time.Duration(3) * time.Second)
		panic("must set env: DBUSER")
	}

	dbPassword, ex = os.LookupEnv("DBPASSWORD")
	if !ex {
		time.Sleep(time.Duration(3) * time.Second)
		panic("must set env: DBPASSWORD")
	}

	dbe = new(DBEngine)
	dbe.setDsn("org_feyon", dbAddr, dbUser, dbPassword)
	if dbe.connect() != nil {
		panic("failed to connect database-root")
	}
	if dbe.initTables() != nil {
		panic("failed to init tables")
	}
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	defer glog.Flush()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowMethods = append(config.AllowMethods, "DELETE")
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.StaticFile("/favicon.ico", "/root/favicon.ico")

	v1 := r.Group("/v1")
	v1.GET("/token", basicAuth)
	r.Run("0.0.0.0:8080")
}
