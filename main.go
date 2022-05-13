package main

import (
	"gin-go-wp/common"
	"gin-go-wp/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//docs.nestjs.cn/8/openapi
	//https://github.com/panjiachen/vue-admin-template  基础模板
	//https://github.com/panjiachen/vue-element-admin 集成模板
	//https://github.com/xiaoxian521/vue-pure-admin
	//https://github.com/haydenzhourepo/
	db := common.InitDB()
	defer db.Close()
	// https://gin-gonic.com/docs/quickstart/  官网
	r := gin.Default()
	r = router.CollectRoute(r)
	panic(r.Run(":1016")) // listen and serve on 0.0.0.0:8080
}
