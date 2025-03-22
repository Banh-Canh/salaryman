package cmd

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/Banh-Canh/salaryman/internal/pkg/parser"
	"github.com/Banh-Canh/salaryman/internal/pkg/pdf"
	"github.com/Banh-Canh/salaryman/internal/pkg/template"
	"github.com/Banh-Canh/salaryman/internal/services"
	"github.com/Banh-Canh/salaryman/internal/utils/fs"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

var (
	resumeDataFile string
	templateName   string
	outputFile     string
	htmlFile       bool
)

func init() {
     RootCmd.AddCommand(localCmd)
	localCmd.Flags().StringVarP(&resumeDataFile, "input", "f", "", "input json file")
	localCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output html|pdf file")
	localCmd.Flags().StringVarP(&templateName, "template", "t", "", "override template")
	localCmd.Flags().BoolVarP(&htmlFile, "html", "", false, "output an html file")
	err := localCmd.MarkFlagRequired("input")
	logger.InitializeLogger(slog.LevelInfo)
	if err != nil {
		logger.Logger.Error("Failed to mark 'input' flag as required", slog.Any("error", err))
	}
	err = localCmd.MarkFlagRequired("output")
	if err != nil {
		logger.Logger.Error("Failed to mark 'output' flag as required", slog.Any("error", err))
	}
}

var localCmd = &cobra.Command{
	Use:     "local",
	Short:   "Generates output locally",
	PreRunE: preRunLocalCommand,
	RunE:    runLocalCommand,
}

func preRunLocalCommand(cmd *cobra.Command, args []string) error {
	filePath := cmd.Flag("input").Value.String()
	err := fs.EnsureNonEmptyFile(filePath)
	if err != nil {
		return err
	}

	return nil
}

func runLocalCommand(cmd *cobra.Command, args []string) error {
	logger.Logger.Info("Generating output")
	tmpDir, err := os.MkdirTemp("", "salaryman-*")
	if err != nil {
		logger.Logger.Error("Failed to create temporary directory.")
	}
	templateManager := template.NewTemplateManager("ui")
	parser := parser.NewHTMLParser(tmpDir, tmpDir+"/resume.html", templateManager)
	pdfGenerator := pdf.NewPDFGenerator()
	resumeService := services.NewResumeService(parser, pdfGenerator)

	fileData, err := fs.ReadFile(resumeDataFile)
	if err != nil {
		return err
	}
	resumeData, err := resumeService.UnmarshalResume(fileData)
	if err != nil {
		return err
	}
	if templateName != "" {
		resumeData.Meta.Template = templateName
	}
	pdfData, err := resumeService.GeneratePDF(resumeData, outputFile)
	if err != nil {
		return err
	}
	if htmlFile {
		htmlFilePath := tmpDir + "/resume.html"
		if err := os.Link(htmlFilePath, outputFile); err != nil {
			return err
		}
	} else {
		if err := fs.WriteFile(outputFile, pdfData); err != nil {
			return err
		}
	}

	return nil
}

func Cmd() *cobra.Command {
	return localCmd
}
