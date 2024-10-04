package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"net/http"
	"time"

	"exchangeapp/global"
	"exchangeapp/models"
)

var cashKey = "articles"

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {
	cashData, err := global.RedisDb.Get(cashKey).Result()
	if err == redis.Nil {

		var articles []models.Article

		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error1": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			}
			return
		}
		articleJson, err := json.Marshal(articles)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err})
		}
		if err := global.RedisDb.Set(cashKey, articleJson, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error3": err.Error()})
	} else {
		var articles []models.Article
		if err := json.Unmarshal([]byte(cashData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error4": err})
			return
		}
		if err := global.RedisDb.Del(cashKey).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error5": err})
			return
		}
		ctx.JSON(http.StatusOK, articles)
	}

}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error1": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)
}
