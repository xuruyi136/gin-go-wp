package main

import (
	"gin-go-wp/common"
	"gin-go-wp/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	// https://gin-gonic.com/docs/quickstart/  官网
	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
