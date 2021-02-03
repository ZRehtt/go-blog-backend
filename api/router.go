package api

import "github.com/gin-gonic/gin"

//NewRouter ...
func NewRouter() *gin.Engine {
	router := gin.Default()

	return router
}
