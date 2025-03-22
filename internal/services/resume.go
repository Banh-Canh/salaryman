package services

import (
	"github.com/pkg/errors"

	"github.com/Banh-Canh/salaryman/internal/models"
	"github.com/Banh-Canh/salaryman/internal/pkg/parser"
	"github.com/Banh-Canh/salaryman/internal/pkg/pdf"
	"github.com/Banh-Canh/salaryman/internal/utils/json"
)

type ResumeService struct {
	Parser *parser.HTMLParser
	Pdf    *pdf.Generator
}

func NewResumeService(parser *parser.HTMLParser, pdf *pdf.Generator) *ResumeService {
	return &ResumeService{
		Parser: parser,
		Pdf:    pdf,
	}
}

func (s *ResumeService) GeneratePDF(resumeData models.Resume, outputDir string) ([]byte, error) {
	htmlFile, err := s.Parser.ParseToHtml(resumeData)
	if err != nil {
		return nil, err
	}

	pdfData, err := s.Pdf.GenerateFromHTML(htmlFile)
	if err != nil {
		return nil, err
	}

	return pdfData, err
}

func (s *ResumeService) UnmarshalResume(data []byte) (models.Resume, error) {
	var resumeData models.Resume
	err := json.Unmarshal(data, &resumeData)
	if err != nil {
		return models.Resume{}, errors.Wrap(err, "failed to unmarshal JSON")
	}

	return resumeData, nil
}
