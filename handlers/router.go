package handlers


import "github.com/gin-gonic/gin"


func GetRouter() *gin.Engine {
   router := gin.Default()
   // v1 := router.Group("/v1")


   return router


}





