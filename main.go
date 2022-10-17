package main

import (
	"code_hash/component"
	"code_hash/middleware"
	"code_hash/module/decrypt/decryptTransport"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	if err := runService(); err != nil {
		log.Fatalln(err)
	}

}

func runService() error {
	r := gin.Default()
	appCtx := component.NewAppContext(nil, "MetaAccess_General_@#113412", "opswatmetaaccess")
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	decryptLog := v1.Group("/decrypt")
	{
		decryptLog.POST("/string", decryptTransport.DecryptString(appCtx))
		decryptLog.POST("/file", decryptTransport.DecryptFile(appCtx))
	}

	return r.Run()
}
