package main

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
	"news/controller"
	"news/db"
	"news/helper"
	"news/repository"
	"news/service"
	"time"
)

func setupRouter() *gin.Engine {
	var (
		redisStore = persistence.NewRedisCache(helper.GetEnv("REDIS_HOST", "127.0.0.1:6379"), "", time.Second)
		mongo      = db.InitMongo()

		//setup repository
		newsRepo = repository.NewNewsRepo(mongo)
		tagRepo  = repository.NewTagRepo(mongo)
		//setup service
		tagServ  = service.NewTagService(tagRepo)
		newsServ = service.NewNewsService(newsRepo, tagRepo)
		//setup controller
		newsCont = controller.NewNewsCont(newsServ)
		tagCont  = controller.NewTagCont(tagServ)
	)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/tag", tagCont.Create)
		v1.GET("/tags", cache.CachePage(redisStore, time.Second, tagCont.Gets))
		v1.GET("/tag/:slug", cache.CachePage(redisStore, time.Second, tagCont.FindTagBySlug))
		v1.PUT("/tag/:slug", tagCont.UpdateBySlug)
		v1.DELETE("/tag/:slug", tagCont.DeleteBySlug)

		//news?topic=&status=publish
		v1.GET("/news", cache.CachePage(redisStore, time.Second, newsCont.Gets))
		v1.POST("/news", newsCont.Create)
		v1.PUT("/news/:id", newsCont.UpdateByID)
		v1.GET("/news/:id", cache.CachePage(redisStore, time.Second, newsCont.Get))
		v1.DELETE("/news/:id", newsCont.DeleteByID)
	}

	return r
}
