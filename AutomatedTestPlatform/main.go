package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/AutomatedTestPlatform/internal/repository"
	"gotest/AutomatedTestPlatform/internal/repository/dao"
	"gotest/AutomatedTestPlatform/internal/service"
	"gotest/AutomatedTestPlatform/internal/web"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := gin.Default()
	//初始化识别的接口
	initSea(db, server)
	//初始化路由
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders: []string{"x-jwt-token"},
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "10.178")
		},
		MaxAge: 12 * time.Hour,
	}))
	server.Run(":18080")
}

func initSea(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewInterfaceDao(db)
	// 创建公共接口服务和路由
	commHttp := web.NewCommon(service.NewCommServiceInterface(repository.NewCommRepInterface(ud)))
	commHttp.CommonRoutes(server)
	// 创建识别接口服务和路由
	searchHttp := web.NewSearchHttpInterface(service.NewSearchServiceInterface(repository.NewRepSearchInterface(ud)))
	searchHttp.RegisterRoutes(server)
	//聚类接口路由
	//clusterHttp := web.NewWebHttpCluster(service.ClusterServiceInterface{}(repository.NewRepSearchInterface(ud)))
	//clusterHttp.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:Cloudwalk#galaxy@tcp(10.178.16.5:33091)/XqpTest"))
	if err != nil {
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
