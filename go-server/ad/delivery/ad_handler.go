package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/wj461/ad-service/domain"
	"github.com/wj461/ad-service/swagger"
)

type adHandler struct {
	adUsecase domain.AdUsecase
}

func NewAdHandler(e *gin.Engine, adUsecase domain.AdUsecase) {
	handler := &adHandler{
		adUsecase: adUsecase,
	}

	v1 := e.Group("api/v1")

	v1.POST("/ad", handler.CreateAd)
	v1.PUT("/ad", handler.CreateAd)
	v1.GET("/ad", handler.CreateAd)
}

func (ah *adHandler) CreateAd(c *gin.Context) {
	ad := &swagger.Ad{}
	if err := c.ShouldBindJSON(ad); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ah.adUsecase.CreateAd(c, ad); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ad)
}
