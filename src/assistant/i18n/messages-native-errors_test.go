package i18n_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/cobrass/src/assistant/i18n"
	"github.com/snivilised/cobrass/src/internal/helpers"
)

type nativeEntry struct {
	Name string
	Args []any
	Fn   any
}

var _ = Describe("MessagesNativeErrors", func() {

	DescribeTable("Native Errors",
		func(entry nativeEntry) {
			err := helpers.CallE(entry.Fn, entry.Args)
			GinkgoWriter.Printf("⚠️ NATIVE-ERROR-RESULT: %v", err)
			Expect(err).Error().NotTo(BeNil())
		},
		func(entry nativeEntry) string {
			return fmt.Sprintf("🧪 --> 🐞 given: native error function: '%v'", entry.Name)
		},

		Entry(nil, nativeEntry{
			Name: "NewEnumValueValueAlreadyExistsNativeError",
			Fn:   i18n.NewEnumValueValueAlreadyExistsNativeError,
			Args: []any{"foo-bar", 2},
		}),

		Entry(nil, nativeEntry{
			Name: "NewIsNotValidEnumValueNativeError",
			Fn:   i18n.NewIsNotValidEnumValueNativeError,
			Args: []any{"foo-bar"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewFailedToAddValidatorAlreadyExistsNativeError",
			Fn:   i18n.NewFailedToAddValidatorAlreadyExistsNativeError,
			Args: []any{"foo-flag"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewCommandAlreadyRegisteredNativeError",
			Fn:   i18n.NewCommandAlreadyRegisteredNativeError,
			Args: []any{"foo-name"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParentCommandNotRegisteredNativeError",
			Fn:   i18n.NewParentCommandNotRegisteredNativeError,
			Args: []any{"foo-parent"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetAlreadyRegisteredNativeError",
			Fn:   i18n.NewParamSetAlreadyRegisteredNativeError,
			Args: []any{"foo-name"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetObjectMustBeStructNativeError",
			Fn:   i18n.NewParamSetObjectMustBeStructNativeError,
			Args: []any{"foo-name", "foo-typ"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetObjectMustBePointerNativeError",
			Fn:   i18n.NewParamSetObjectMustBePointerNativeError,
			Args: []any{"foo-pointer-name", "foo-pointer-typ"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetNotFoundNativeError",
			Fn:   i18n.NewParamSetNotFoundNativeError,
			Args: []any{"foo-name"},
		}),
	)
})
