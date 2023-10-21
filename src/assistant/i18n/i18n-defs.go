package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const CobrassSourceID = "github.com/snivilised/cobrass"

// These definitions are in support of extendio's Localisable
// interface and other i18n related definitions.

type CobrassTemplData struct{}

func (td CobrassTemplData) SourceID() string {
	return CobrassSourceID
}

type Message = i18n.Message
