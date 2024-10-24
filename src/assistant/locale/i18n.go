package locale

import (
	xi18n "github.com/snivilised/extendio/i18n"
)

func Use(options ...xi18n.UseOptionFn) error {
	return xi18n.Use(options...)
}
