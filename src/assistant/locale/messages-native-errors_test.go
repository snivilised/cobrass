package locale_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/cobrass/src/assistant/locale"
	"github.com/snivilised/cobrass/src/internal/lab"
)

type nativeEntry struct {
	Name string
	Args []any
	Fn   any
}

var _ = Describe("MessagesNativeErrors", func() {

	DescribeTable("Native Errors",
		func(entry nativeEntry) {
			err := lab.CallE(entry.Fn, entry.Args)
			GinkgoWriter.Printf("⚠️ NATIVE-ERROR-RESULT: %v", err)
			Expect(err).Error().NotTo(BeNil())
		},
		func(entry nativeEntry) string {
			return fmt.Sprintf("🧪 --> 🐞 given: native error function: '%v'", entry.Name)
		},

		Entry(nil, nativeEntry{
			Name: "NewEnumValueValueAlreadyExistsNativeError",
			Fn:   locale.NewEnumValueValueAlreadyExistsNativeError,
			Args: []any{"foo-bar", 2},
		}),

		Entry(nil, nativeEntry{
			Name: "NewIsNotValidEnumValueNativeError",
			Fn:   locale.NewIsNotValidEnumValueNativeError,
			Args: []any{"foo-bar"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewFailedToAddValidatorAlreadyExistsNativeError",
			Fn:   locale.NewFailedToAddValidatorAlreadyExistsNativeError,
			Args: []any{"foo-flag"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewCommandAlreadyRegisteredNativeError",
			Fn:   locale.NewCommandAlreadyRegisteredNativeError,
			Args: []any{"foo-name"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParentCommandNotRegisteredNativeError",
			Fn:   locale.NewParentCommandNotRegisteredNativeError,
			Args: []any{"foo-parent"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetAlreadyRegisteredNativeError",
			Fn:   locale.NewParamSetAlreadyRegisteredNativeError,
			Args: []any{"foo-name"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetObjectMustBeStructNativeError",
			Fn:   locale.NewParamSetObjectMustBeStructNativeError,
			Args: []any{"foo-name", "foo-typ"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetObjectMustBePointerNativeError",
			Fn:   locale.NewParamSetObjectMustBePointerNativeError,
			Args: []any{"foo-pointer-name", "foo-pointer-typ"},
		}),

		Entry(nil, nativeEntry{
			Name: "NewParamSetNotFoundNativeError",
			Fn:   locale.NewParamSetNotFoundNativeError,
			Args: []any{"foo-name"},
		}),
	)
})
