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

	relatorioService := services.NewRelatorioService(dbconnection)
	diarioUseCase := usecases.NewDiarioUsecase(diarioService, relatorioService, obraService, pessoaService)
	diarioController := controller.NewDiarioController(diarioUseCase)

	fornecedorService := services.NewFornecedorService(dbconnection)
	fornecedorUseCase := usecases.NewFornecedorUsecase(fornecedorService)
	fornecedorController := controller.NewFornecedorController(fornecedorUseCase)

	despesaService := services.NewDespesaService(dbconnection)
	despesaUseCase := usecases.NewDespesaUsecase(despesaService)
	despesaController := controller.NewDespesaController(despesaUseCase)

	receitaService := services.NewReceitaService(dbconnection)
	receitaUseCase := usecases.NewReceitaUseCase(receitaService, obraService)
	receitaController := controller.NewReceitaController(receitaUseCase)

	relatorioUseCase := usecases.NewRelatorioUseCase(relatorioService, obraService)
	relatorioController := controller.NewRelatorioController(relatorioUseCase)

	loginService := services.NewLoginService(dbconnection)
	loginUseCase := usecases.NewLoginUsecase(loginService)
	loginController := controller.NewLoginController(loginUseCase)

	// Equipe do Diário
	equipeDiarioService := services.NewEquipeDiarioService(dbconnection)
	equipeDiarioUseCase := usecases.NewEquipeDiarioUseCase(equipeDiarioService)
	equipeDiarioController := controller.NewEquipeDiarioController(equipeDiarioUseCase)

	// Equipamento do Diário
	equipamentoDiarioService := services.NewEquipamentoDiarioService(dbconnection)
	equipamentoDiarioUseCase := usecases.NewEquipamentoDiarioUseCase(equipamentoDiarioService)
	equipamentoDiarioController := controller.NewEquipamentoDiarioController(equipamentoDiarioUseCase)

	// Material do Diário
	materialDiarioService := services.NewMaterialDiarioService(dbconnection)
	materialDiarioUseCase := usecases.NewMaterialDiarioUseCase(materialDiarioService)
	materialDiarioController := controller.NewMaterialDiarioController(materialDiarioUseCase)

	// Atividade Diária (Nova Arquitetura)
	atividadeDiariaService := services.NewAtividadeDiariaService(dbconnection)
	atividadeDiariaUseCase := usecases.NewAtividadeDiariaUseCase(atividadeDiariaService, obraService)
	atividadeDiariaController := controller.NewAtividadeDiariaController(atividadeDiariaUseCase)

	// Ocorrência Diária (Nova Arquitetura)
	ocorrenciaDiariaService := services.NewOcorrenciaDiariaService(dbconnection)
	ocorrenciaDiariaUseCase := usecases.NewOcorrenciaDiariaUseCase(ocorrenciaDiariaService, obraService)
	ocorrenciaDiariaController := controller.NewOcorrenciaDiariaController(ocorrenciaDiariaUseCase)

	// Diário Consolidado (Nova Arquitetura)
	diarioConsolidadoService := services.NewDiarioConsolidadoService(dbconnection)
	diarioMetadadosService := services.NewDiarioMetadadosService(dbconnection)
	diarioConsolidadoUseCase := usecases.NewDiarioConsolidadoUseCase(diarioConsolidadoService, diarioMetadadosService, obraService)
	diarioConsolidadoController := controller.NewDiarioConsolidadoController(diarioConsolidadoUseCase)

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
		protected.POST("/receitas", receitaController.CreateReceita)

		// READ (GET)
		protected.GET("/usuarios", usuarioController.GetUsuarios)
		protected.GET("/usuarios/:id", usuarioController.GetUsuarioById)

		protected.GET("/pessoas", pessoaController.GetPessoas)
		protected.GET("/pessoas/:id", pessoaController.GetPessoaById)

		protected.GET("/obras", obraController.GetObras)
		protected.GET("/obras/:id", obraController.GetObraById)

		protected.GET("/diarios", diarioController.GetDiarios)
		protected.GET("/diarios/relatorio-formatado/:obra_id", diarioController.GetDiarioRelatorioFormatado)
		protected.GET("/diarios/obra/:id", diarioController.GetDiariosByObraId)
		protected.GET("/diarios/:id/relatorio-completo", diarioController.GetRelatorioDiarioCompleto)
		protected.GET("/diarios/:id", diarioController.GetDiarioById)

		protected.GET("/fornecedores", fornecedorController.GetFornecedores)
		protected.GET("/fornecedores/:id", fornecedorController.GetFornecedorById)

		protected.GET("/despesas", despesaController.GetDespesas)
		protected.GET("/despesas/:id", despesaController.GetDespesaById)
		protected.GET("/despesas/relatorio/:obra_id", despesaController.GetRelatorioPorObra)

		protected.GET("/receitas", receitaController.GetReceitas)
		protected.GET("/receitas/:id", receitaController.GetReceitaById)
		protected.GET("/receitas/obra/:obra_id", receitaController.GetReceitasByObra)

		// RELATÓRIOS
		protected.GET("/relatorios/obra/:obra_id", relatorioController.GetRelatorioObra)
		protected.GET("/relatorios/despesas/:obra_id", relatorioController.GetRelatorioDespesasPorCategoria)
		protected.GET("/relatorios/pagamentos/:obra_id", relatorioController.GetRelatorioPagamentos) // ?status=PENDENTE opcional
		protected.GET("/relatorios/materiais/:obra_id", relatorioController.GetRelatorioMateriais)
		protected.GET("/relatorios/profissionais/:obra_id", relatorioController.GetRelatorioProfissionais)

		// EQUIPE DIARIO
		protected.POST("/equipe-diario", equipeDiarioController.Create)
		protected.GET("/equipe-diario/obra/:obra_id", equipeDiarioController.GetByObraId)                 // Nova rota: todas as datas
		protected.GET("/equipe-diario/obra/:obra_id/data/:data", equipeDiarioController.GetByObraAndData) // Nova rota: data específica
		protected.GET("/equipe-diario/diario/:diario_id", equipeDiarioController.GetByDiarioId)           // Legado: deprecated
		protected.PUT("/equipe-diario/:id", equipeDiarioController.Update)
		protected.DELETE("/equipe-diario/:id", equipeDiarioController.Delete)

		// EQUIPAMENTO DIARIO
		protected.POST("/equipamento-diario", equipamentoDiarioController.Create)
		protected.GET("/equipamento-diario/obra/:obra_id", equipamentoDiarioController.GetByObraId)                 // Todas as datas
		protected.GET("/equipamento-diario/obra/:obra_id/data/:data", equipamentoDiarioController.GetByObraAndData) // Data específica
		protected.PUT("/equipamento-diario/:id", equipamentoDiarioController.Update)
		protected.DELETE("/equipamento-diario/:id", equipamentoDiarioController.Delete)

		// MATERIAL DIARIO
		protected.POST("/material-diario", materialDiarioController.Create)
		protected.GET("/material-diario/obra/:obra_id", materialDiarioController.GetByObraId)                 // Todas as datas
		protected.GET("/material-diario/obra/:obra_id/data/:data", materialDiarioController.GetByObraAndData) // Data específica
		protected.PUT("/material-diario/:id", materialDiarioController.Update)
		protected.DELETE("/material-diario/:id", materialDiarioController.Delete)

		// ========== NOVA ARQUITETURA: DIÁRIO NORMALIZADO ==========

		// TAREFAS REALIZADAS
		protected.POST("/tarefas", atividadeDiariaController.CreateAtividade)
		protected.GET("/tarefas", atividadeDiariaController.GetAtividades)
		protected.GET("/tarefas/obra/:obra_id", atividadeDiariaController.GetAtividadesByObra)                // Nova rota: todas as datas
		protected.GET("/tarefas/obra/:obra_id/data/:data", atividadeDiariaController.GetAtividadesByObraData) // Rota específica com data
		protected.PUT("/tarefas/:id", atividadeDiariaController.UpdateAtividade)
		protected.DELETE("/tarefas/:id", atividadeDiariaController.DeleteAtividade)

		// OCORRÊNCIAS
		protected.POST("/ocorrencias", ocorrenciaDiariaController.CreateOcorrencia)
		protected.GET("/ocorrencias", ocorrenciaDiariaController.GetOcorrencias)
		protected.GET("/ocorrencias/obra/:obra_id", ocorrenciaDiariaController.GetOcorrenciasByObra)                // Nova rota: todas as datas
		protected.GET("/ocorrencias/obra/:obra_id/data/:data", ocorrenciaDiariaController.GetOcorrenciasByObraData) // Rota específica com data
		protected.GET("/ocorrencias/gravidade/:gravidade", ocorrenciaDiariaController.GetOcorrenciasByGravidade)
		protected.PUT("/ocorrencias/:id", ocorrenciaDiariaController.UpdateOcorrencia)
		protected.DELETE("/ocorrencias/:id", ocorrenciaDiariaController.DeleteOcorrencia)

		// DIÁRIO CONSOLIDADO (View Agregada)
		protected.GET("/diarios-consolidado", diarioConsolidadoController.GetDiarioConsolidado)
		protected.GET("/diarios-consolidado/obra/:obra_id", diarioConsolidadoController.GetDiarioConsolidadoByObra)
		protected.GET("/diarios-consolidado/data/:data", diarioConsolidadoController.GetDiarioConsolidadoByData)
		protected.POST("/diarios-consolidado/metadados", diarioConsolidadoController.CreateOrUpdateMetadados)

		// ==========================================================

		// UPDATE (PUT)
		protected.PUT("/usuarios/:id", usuarioController.PutUsuarioById)
		protected.PUT("/pessoas/:id", pessoaController.PutPessoaById)
		protected.PUT("/obras/:id", obraController.PutObraById)
		protected.PUT("/diarios/:id", diarioController.PutDiarioById)
		protected.PUT("/fornecedores/:id", fornecedorController.PutFornecedorById)
		protected.PUT("/despesas/:id", despesaController.PutDespesaById)
		protected.PUT("/receitas/:id", receitaController.PutReceitaById)

		// DELETE
		protected.DELETE("/usuarios/:id", usuarioController.DeleteUsuarioById)
		protected.DELETE("/pessoas/:id", pessoaController.DeletePessoaById)
		protected.DELETE("/obras/:id", obraController.DeleteObraById)
		protected.DELETE("/diarios/:id", diarioController.DeleteDiariosById)
		protected.DELETE("/fornecedores/:id", fornecedorController.DeleteFornecedorById)
		protected.DELETE("/despesas/:id", despesaController.DeleteDespesaById)
		protected.DELETE("/receitas/:id", receitaController.DeleteReceitaById)
	}

	// ✅ Inicia servidor
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Porta padrão
	}

	log.Printf("Servidor iniciado na porta %s", port)
	server.Run(":" + port)
}
