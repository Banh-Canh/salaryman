package api

import (
	"cmp"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/Banh-Canh/salaryman/configs"
	"github.com/Banh-Canh/salaryman/internal/api/router"
	"github.com/Banh-Canh/salaryman/internal/pkg/parser"
	"github.com/Banh-Canh/salaryman/internal/pkg/pdf"
	"github.com/Banh-Canh/salaryman/internal/pkg/template"
	"github.com/Banh-Canh/salaryman/internal/services"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

type Api struct {
	config configs.ApiConfig
	router *gin.Engine
}

func New() *Api {
	api := &Api{}
	api.setup()
	return api
}

func (api *Api) setup() {
	outputDir := cmp.Or(os.Getenv("SALARYMAN_OUTPUTDIR"), "./output")
	absOutputPath, err := filepath.Abs(outputDir)
	if err != nil {
		logger.Logger.Error("Error resolving absolute path", "error", err)
		return
	}
	logger.Logger.Info("Output directory set", "directory", absOutputPath)
	gin.SetMode(gin.ReleaseMode)
	templateManager := template.NewTemplateManager("ui")
	parser := parser.NewHTMLParser(outputDir, outputDir+"/resume.html", templateManager)
	pdfGenerator := pdf.NewPDFGenerator()
	resumeService := services.NewResumeService(parser, pdfGenerator)
	api.config = configs.LoadApiConfig()
	api.router = router.Init(resumeService)
}

func (api *Api) Run() error {
	logger.Logger.Info(fmt.Sprintf("%s API running on Port %d", api.config.AppName, api.config.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.config.Port), api.router)
	if err != nil {
		return err
	}
	return nil
}
