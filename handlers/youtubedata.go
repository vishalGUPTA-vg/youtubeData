package handlers

import (
	"net/http"
	"strconv"
	"youtubedata/services"

	"github.com/gin-gonic/gin"
)

func setYoutubeDataRoute(r *gin.Engine) {
	r.GET("/youtube/videos", Getlist)
	r.GET("/youtube/search/videos", GetlistSearch)
}

func Getlist(c *gin.Context) {
	y := services.NewYouttubeServices()
	pg, _ := strconv.Atoi(c.Query("pg"))
	if pg == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "enter valid pg number ",
		})
		return
	}
	res, err := y.GetList(pg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)

}



func GetlistSearch(c *gin.Context) {
	y := services.NewYouttubeServices()

	q:=c.Query("q")
	if q==""{
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "enter valid q ",
		})
	}
	res, err := y.GetSearchList(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

