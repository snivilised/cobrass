package i18n

const CobrassSourceID = "github.com/snivilised/cobrass"

// These definitions are in support of extendio's Localisable
// interface and other i18n related definitions.

type CobrassTemplData struct{}

func (td CobrassTemplData) SourceID() string {
	return CobrassSourceID
}
