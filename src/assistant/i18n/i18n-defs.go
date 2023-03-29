package i18n

const SOURCE_ID = "github.com/snivilised/cobrass"

// These definitions are in support of extendio's Localisable
// interface and other i18n related definitions.

type CobrassTemplData struct{}

func (td CobrassTemplData) SourceId() string {
	return SOURCE_ID
}
