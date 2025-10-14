package main

import (
	controller "codxis-obras/internal/controllers"
	"codxis-obras/internal/services"
	"codxis-obras/internal/usecases"
	"codxis-obras/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// api := "/v1"

	dbconnection, err := postgres.ConnectDB()
	if err != nil {
		panic(err)

	}

	obraService := services.NewObraService(dbconnection)
	obraUseCase := usecases.NewObraUsecase(obraService)
	obraController := controller.NewObraController(obraUseCase)

	pessoaService := services.NewPessoaService(dbconnection)
	pessoaUseCase := usecases.NewPessoaUsecase(pessoaService)
	pessoaController := controller.NewPessoaController(pessoaUseCase)

	usuarioService := services.NewUsuarioService(dbconnection)
	usuarioUseCase := usecases.NewUsuarioUsecase(usuarioService)
	usuarioController := controller.NewUsuarioController(usuarioUseCase)

	diarioService := services.NewDiarioService(dbconnection)
	diarioUseCase := usecases.NewDiarioUsecase(diarioService)
	diarioController := controller.NewDiarioController(diarioUseCase)

	//posts
	server.POST("/pessoa", pessoaController.CreatePessoa)
	server.POST("/usuario", usuarioController.CreateUsuario)
	server.POST("/obra", obraController.CreateObra)
	server.POST("/diario", diarioController.CreateDiario)

	//gest
	server.GET("/usuarios", usuarioController.GetUsuarios)
	server.GET("/pessoas", pessoaController.GetPessoas)
	server.GET("/obras", obraController.GetObras)
	server.GET("/diarios", diarioController.GetDiarios)

	server.GET("/usuario/:id", usuarioController.GetUsuarioById)
	server.GET("/pessoa/:id", pessoaController.GetPessoaById)
	server.GET("/obra/:id", obraController.GetObraById)
	server.GET("/diario/:id", diarioController.GetDiarioById)
	server.GET("/diarios/:id/obra", diarioController.GetDiariosByObraId)

	//patch

	server.Run(":3000")
}
