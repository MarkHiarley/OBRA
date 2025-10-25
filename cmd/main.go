package main

import (
	"codxis-obras/internal/auth"
	controller "codxis-obras/internal/controllers"
	"codxis-obras/internal/services"
	"codxis-obras/internal/usecases"
	"codxis-obras/pkg/postgres"
	"log"
	"os"

	"github.com/gin-contrib/cors"
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

	fornecedorService := services.NewFornecedorService(dbconnection)
	fornecedorUseCase := usecases.NewFornecedorUsecase(fornecedorService)
	fornecedorController := controller.NewFornecedorController(fornecedorUseCase)

	despesaService := services.NewDespesaService(dbconnection)
	despesaUseCase := usecases.NewDespesaUsecase(despesaService)
	despesaController := controller.NewDespesaController(despesaUseCase)

	loginService := services.NewLoginService(dbconnection)
	loginUseCase := usecases.NewLoginUsecase(loginService)
	loginController := controller.NewLoginController(loginUseCase)

	server := gin.Default()

	// ✅ Configurar CORS para permitir todas as origens
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todas as origens
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,     // Deve ser false quando AllowOrigins é "*"
		MaxAge:           12 * 3600, // Cache de 12 horas
	}))

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
		protected.POST("/fornecedores", fornecedorController.CreateFornecedor)
		protected.POST("/despesas", despesaController.CreateDespesa)

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

		protected.GET("/fornecedores", fornecedorController.GetFornecedores)
		protected.GET("/fornecedores/:id", fornecedorController.GetFornecedorById)

		protected.GET("/despesas", despesaController.GetDespesas)
		protected.GET("/despesas/:id", despesaController.GetDespesaById)
		protected.GET("/despesas/relatorio/:obra_id", despesaController.GetRelatorioPorObra)

		// UPDATE (PUT)
		protected.PUT("/usuarios/:id", usuarioController.PutUsuarioById)
		protected.PUT("/pessoas/:id", pessoaController.PutPessoaById)
		protected.PUT("/obras/:id", obraController.PutObraById)
		protected.PUT("/diarios/:id", diarioController.PutDiarioById)
		protected.PUT("/fornecedores/:id", fornecedorController.PutFornecedorById)
		protected.PUT("/despesas/:id", despesaController.PutDespesaById)

		// DELETE
		protected.DELETE("/usuarios/:id", usuarioController.DeleteUsuarioById)
		protected.DELETE("/pessoas/:id", pessoaController.DeletePessoaById)
		protected.DELETE("/obras/:id", obraController.DeleteObraById)
		protected.DELETE("/diarios/:id", diarioController.DeleteDiariosById)
		protected.DELETE("/fornecedores/:id", fornecedorController.DeleteFornecedorById)
		protected.DELETE("/despesas/:id", despesaController.DeleteDespesaById)
	}

	// ✅ Inicia servidor
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Porta padrão
	}

	log.Printf("Servidor iniciado na porta %s", port)
	server.Run(":" + port)
}
