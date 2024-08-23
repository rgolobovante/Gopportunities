package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRouters(router)
	
	//Run the Server
	router.Run() // listen and serve on 0.0.0.0:8080

}
