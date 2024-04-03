package handlers


import "github.com/gin-gonic/gin"


func GetRouter() *gin.Engine {
   router := gin.Default()
   setYoutubeDataRoute(router)


   return router


}





