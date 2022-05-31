package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/martadrozsa/curso-golang-web/gogin/controller"
)

func HandlerRequests() {
	//servidor com as configs default
	r := gin.Default()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.Run()
}
