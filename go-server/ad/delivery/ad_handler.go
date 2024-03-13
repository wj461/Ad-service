package delivery

import (
	"net/http"

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
	{
		v1.POST("/ad", handler.CreateAd)
		v1.GET("/ad", handler.SearchAd)
	}
}

func (ah *adHandler) CreateAd(c *gin.Context) {
	ad := &swagger.Ad{}
	ad.Conditions = &swagger.Conditions{}

	if err := c.ShouldBindJSON(ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ah.adUsecase.CreateAd(c, ad); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (ah *adHandler) SearchAd(c *gin.Context) {
	// adQuery := &swagger.AdQuery{}
	var adQuery swagger.AdQuery

	err := c.BindQuery(&adQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adResponse, err := ah.adUsecase.SearchAd(c, &adQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adResponse)
}
