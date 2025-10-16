package main

import (
	controller "codxis-obras/internal/controllers"
	"codxis-obras/internal/services"
	"codxis-obras/internal/usecases"
	"codxis-obras/pkg/postgres"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()
	if err != nil {

		log.Println("Warning: Could not find or load .env file. Using system environment variables.")
	}

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
	server.POST("/pessoas", pessoaController.CreatePessoa)
	server.POST("/usuarios", usuarioController.CreateUsuario)
	server.POST("/obras", obraController.CreateObra)
	server.POST("/diarios", diarioController.CreateDiario)

	//gest
	server.GET("/usuarios", usuarioController.GetUsuarios)
	server.GET("/pessoas", pessoaController.GetPessoas)
	server.GET("/obras", obraController.GetObras)
	server.GET("/diarios", diarioController.GetDiarios)

	server.GET("/usuarios/:id", usuarioController.GetUsuarioById)
	server.GET("/pessoas/:id", pessoaController.GetPessoaById)
	server.GET("/obras/:id", obraController.GetObraById)
	server.GET("/diarios/:id", diarioController.GetDiarioById)
	server.GET("/diarios/:id/obra", diarioController.GetDiariosByObraId)

	//patch

	server.PUT("/usuarios/:id", usuarioController.PutUsuarioById)
	server.PUT("/pessoas/:id", pessoaController.PutPessoaById)
	server.PUT("/obras/:id", obraController.PutObraById)
	server.PUT("/diarios/:id", diarioController.PutDiarioById)

	//delete

	server.DELETE("/usuarios/:id", usuarioController.DeleteUsuarioById)
	port := os.Getenv("API_PORT")
	server.Run(":" + port)
}
