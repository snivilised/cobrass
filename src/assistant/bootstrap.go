package assistant

import (
	"github.com/cubiest/jibberjabber"
	"golang.org/x/text/language"
)

type Bootstrap struct {
	Detector LocaleDetector
}

type LocaleDetector interface {
	Scan() language.Tag
}

// Jabber is a LocaleDetector implemented using jibberjabber.
//
type Jabber struct {
}

// Scan returns the detected language tag.
//
func (j *Jabber) Scan() language.Tag {

	lang, _ := jibberjabber.DetectIETF()
	return language.MustParse(lang)
}

// Execute runs the bootstrap.
//
func (b *Bootstrap) Execute(initialise func(LocaleDetector)) {
	if b.Detector == nil {
		b.Detector = &Jabber{}
	}

	initialise(b.Detector)
}
