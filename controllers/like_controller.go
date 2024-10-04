package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"

	"exchangeapp/global"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")
	LikeKey := "article:" + articleID + ":likes"

	if err := global.RedisDb.Incr(LikeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")
	LikeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDb.Get(LikeKey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
