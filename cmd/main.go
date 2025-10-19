package main

import (
	"codxis-obras/internal/auth"
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

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not find or load .env file. Using system environment variables.")
	}

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

	loginService := services.NewLoginService(dbconnection)
	loginUseCase := usecases.NewLoginUsecase(loginService)
	loginController := controller.NewLoginController(loginUseCase)

	server := gin.Default()

	server.POST("/login", loginController.CreateLogin)
	server.POST("/refresh", loginController.RefreshToken)
	server.POST("/usuarios", usuarioController.CreateUsuario)

	protected := server.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		// CREATE (POST)
		protected.POST("/pessoas", pessoaController.CreatePessoa)

		protected.POST("/obras", obraController.CreateObra)
		protected.POST("/diarios", diarioController.CreateDiario)

		// READ (GET)
		protected.GET("/usuarios", usuarioController.GetUsuarios)
		protected.GET("/usuarios/:id", usuarioController.GetUsuarioById)

		protected.GET("/pessoas", pessoaController.GetPessoas)
		protected.GET("/pessoas/:id", pessoaController.GetPessoaById)

		protected.GET("/obras", obraController.GetObras)
		protected.GET("/obras/:id", obraController.GetObraById)

		protected.GET("/diarios", diarioController.GetDiarios)
		protected.GET("/diarios/:id", diarioController.GetDiarioById)
		protected.GET("/diarios/obra/:id", diarioController.GetDiariosByObraId)

		// UPDATE (PUT)
		protected.PUT("/usuarios/:id", usuarioController.PutUsuarioById)
		protected.PUT("/pessoas/:id", pessoaController.PutPessoaById)
		protected.PUT("/obras/:id", obraController.PutObraById)
		protected.PUT("/diarios/:id", diarioController.PutDiarioById)

		// DELETE
		protected.DELETE("/usuarios/:id", usuarioController.DeleteUsuarioById)
		protected.DELETE("/pessoas/:id", pessoaController.DeletePessoaById)
		protected.DELETE("/obras/:id", obraController.DeleteObraById)
		protected.DELETE("/diarios/:id", diarioController.DeleteDiariosById)
	}

	// ✅ Inicia servidor
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Porta padrão
	}

	log.Printf("Servidor iniciado na porta %s", port)
	server.Run(":" + port)
}
