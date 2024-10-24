package locale

import (
	"github.com/snivilised/li18ngo"
)

func Use(options ...li18ngo.UseOptionFn) error {
	return li18ngo.Use(options...)
}
