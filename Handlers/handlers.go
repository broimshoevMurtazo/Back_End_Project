package handlers

import (
	"murtazo/app/controlers"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	router := gin.Default()

	router.POST("/signup",controlers.Register)
	router.POST("/login",controlers.Login)
	router.POST("/student",controlers.Addstudent)
	router.POST("/project",controlers.Addproject)
	router.POST("/join",controlers.JoinStudent)
	router.GET("/joinlist",controlers.JoinesList)
	router.POST("/delete",controlers.DeleteStudent)
	router.POST("/deleteproject",controlers.DeleteProject)
	router.POST("/project_list",controlers.Projectlist)


	router.Run(":2020")
}