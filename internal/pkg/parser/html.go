package parser

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/pkg/errors"

	"github.com/Banh-Canh/salaryman/internal/models"
	"github.com/Banh-Canh/salaryman/internal/pkg/template"
	"github.com/Banh-Canh/salaryman/internal/utils/fs"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

type HTMLParser struct {
	OutputDir      string
	OutputHtmlFile string
	TmplManager    *template.Manager
}

func NewHTMLParser(outputDir, outputHtmlFile string, templateMgr *template.Manager) *HTMLParser {
	return &HTMLParser{
		OutputDir:      outputDir,
		OutputHtmlFile: outputHtmlFile,
		TmplManager:    templateMgr,
	}
}

func (p *HTMLParser) ParseToHtml(resumeData models.Resume) (string, error) {
	startedAt := time.Now()

	fs.EnsureDir(p.OutputDir)

	htmlOut, err := fs.CreateFile(p.OutputHtmlFile)
	if err != nil {
		return "", err
	}
	defer htmlOut.Close()

	if err := p.updateResumeLabels(&resumeData); err != nil {
		return "", err
	}

	t, err := p.TmplManager.GetTemplate(resumeData.Meta.Template)
	if err != nil {
		return "", err
	}

	err = t.Execute(htmlOut, resumeData)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("ParseToHtml %s - Execute", resumeData.Meta.Template))
	}
	logger.Logger.Info("HTML generated", slog.Float64("duration_seconds", time.Since(startedAt).Seconds()))

	return htmlOut.Name(), nil
}

func (p *HTMLParser) updateResumeLabels(resumeData *models.Resume) error {
	resumeData.Labels.Education = resumeData.GetEducationLabel()
	resumeData.Labels.Experiences = resumeData.GetExperiencesLabel()
	resumeData.Labels.Projects = resumeData.GetProjectsLabel()
	resumeData.Labels.Publications = resumeData.GetPublicationsLabel()
	resumeData.Labels.Skills = resumeData.GetSkillsLabel()
	resumeData.Labels.SoftSkills = resumeData.GetSoftSkillsLabel()
	resumeData.Labels.Languages = resumeData.GetLanguagesLabel()
	resumeData.Labels.Interests = resumeData.GetInterestsLabel()
	resumeData.Labels.Profile = resumeData.GetProfileLabel()
	resumeData.Labels.Since = resumeData.GetSinceLabel()
	resumeData.Labels.Certificates = resumeData.GetCertificatesLabel()
	resumeData.Labels.Socials = resumeData.GetSocialsLabel()
	if resumeData.Meta.Template == "" {
		resumeData.Meta.Template = models.DefaultTemplate
	}

	return nil
}
