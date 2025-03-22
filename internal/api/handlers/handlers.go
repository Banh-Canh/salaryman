package handlers

import (
	"cmp"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Banh-Canh/salaryman/internal/models"
	"github.com/Banh-Canh/salaryman/internal/services"
	"github.com/Banh-Canh/salaryman/internal/utils/fs"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}

func GetPdf(service *services.ResumeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		outputDir := cmp.Or(os.Getenv("SALARYMAN_OUTPUTDIR"), "./output")
		var resumeData models.Resume

		if err := c.ShouldBindJSON(&resumeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "ShouldBindJSON : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		htmlFile, err := service.Parser.ParseToHtml(resumeData)
		if err != nil {
			logger.Logger.Error("error", slog.Any("error", err))
		}

		pdfData, err := service.Pdf.GenerateFromHTML(htmlFile)
		if err != nil {
			logger.Logger.Error("error", slog.Any("error", err))
		}

		if err := fs.WriteFile(outputDir+"/resume.pdf", pdfData); err != nil {
			logger.Logger.Error("error", slog.Any("error", err))
		}

		c.Writer.Header().Set("Content-type", "application/pdf")
		c.File(outputDir)
		// Return a JSON response confirming the PDF was successfully created
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "PDF generated successfully",
		})
	}
}
