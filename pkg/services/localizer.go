package services

import (
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Localizer struct {
	// bundle contains the messages that can be returned by the Localizer.
	bundle *i18n.Bundle
}

func NewLocalizer(bundle *i18n.Bundle) *Localizer {
	return &Localizer{
		bundle: bundle,
	}
}

func (l *Localizer) GetLocalizedMessage(ctx echo.Context) string {

	locale := ctx.Get("locale").(string)
	// Create a Localizer for the specified language
	localizer := i18n.NewLocalizer(l.bundle, locale, "en-US")

	// Localize the "hello_world" message
	localizedMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "HelloWorld",
	})
	return localizedMessage
}
