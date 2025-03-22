package lang

import (
	"log/slog"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"github.com/Banh-Canh/salaryman/internal/utils/json"
	"github.com/Banh-Canh/salaryman/internal/utils/logger"
)

const localesDir = "internal/utils/lang/locales"

var bundle *i18n.Bundle

var supportedLanguages = []string{"en_US", "fr_FR"}

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, lang := range supportedLanguages {
		translation, err := loadTranslation(lang)
		if err != nil {
			logger.Logger.Error("Failed to load translation", slog.String("language", lang), slog.Any("error", err))
			continue
		}
		err = bundle.AddMessages(language.Make(lang), translation.Messages...)
		if err != nil {
			logger.Logger.Error("Failed to add messages to bundle", slog.String("language", lang), slog.Any("error", err))
			continue
		}
	}
}

func loadTranslation(lang string) (*i18n.MessageFile, error) {
	translationFile := filepath.Join(localesDir, lang+".json")
	return bundle.LoadMessageFile(translationFile)
}

func Translate(lang string, messageID string) string {
	return i18n.NewLocalizer(bundle, lang).
		MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
		})
}
