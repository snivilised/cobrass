# 🐲 ___Powershell Code Generation___
<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->
<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->
<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<p align="right">
  <a href="https://go.dev"><img src="resources/images/PowerShellLogo.png" width="50" /></a>
</p>

Powershell is being used to generate go code and this article documents how this works.

___Ideally, this would have been implemented using the 'go generate' tool chain but as I am new to Go development this is yet another barrier I was not willing to put up with, preferring instead to exploit existing skills, so that the job could just get done (at some point in the future, it is intended to replace the PowerShell script with the 'go generate' equivalent).___

⚠️ __Implementation of code generation in Go has started but not yet complete. See the Go section at the end of this document.__

There are 6 main functions that perform code generation, 3 that generate source code and the remaining generating ginkgo/gomega based test suites. The reason why code generation was needed mainly stems from the `Cobra` api. Since, they are type based functions, in order to fully integrate with it, `Cobrass` must also provide a type based api. During the initial design of cobrass, functionality was built manually, with a lot of copy and pasting. When it was discovered how laborious it was just to build the initial binder functions, it was decided that another more automated approach would be desirable. This way, any future (non-breaking!) changes can be implemented relatively swiftly, without one having to tear ones hair out due to the monotony of code modifications and it's inherent vulnerability to copy and paste errors.

The 6 generator functions (aliases in brackets) are:

- ___Build-Validators(gen-ov)___: `option-validator-auto.go` generates the core validator api types
- ___Build-ParamSet(gen-ps)___: `param-set-auto.go` generates binder functions
- ___Build-PsTestEntry(gen-ps-t)___: `param-set-auto_test.go` generates unit tests for binder functions
- ___Build-TestEntry(gen-ov-t)___: `option-validator-auto_test.go` generates unit tests for validated binder functions
- ___Build-BinderHelpers(gen-help)___: `param-set-binder-helpers-auto.go` generates validation helper functions
- ___Build-BinderHelperTests(gen-help-t)___: `param-set-binder-helpers-auto_test.go` generates unit tests for validation helper functions

The order of the above list is significant, as it reflects the order in which commands should be run (there are exceptions to this depending on the nature of the change being made).

There are an additional 2 functions that have been built to protect against accidental api breaking changes, after all _with great power comes great responsibility_, that is to say, with code generation in play, it could be easy to release a breaking change and the possibility of this has to be prohibited as much as it can be.

These 2 commands are:

- ___Checkpoint-ParamSetSignatures(check-sig)___: invokes all 3 source code generators and creates a hash from the generated api surface
- ___Show-ParamSetSignatures___(show-sig): displays the result of the calculated hash from check-sig. Also compares with the previous hash to provide a quick and easy mechanism to determine if any api changes have occurred

## 🤖 Using the Code Generators

### 👉 Invoking

To help insertion of generated code, the scripts make use of the system clipboard. Once a command is run, the contents are copied to the clipboard. The user then has to insert that content into the source file indicated. There are markers inside the code files which indicates where the new code should be inserted.

When running for the first time, it is recommended to set the current hash inside the powershell session. This is done via an environment variable (___$env:COBRASS_API_HASH___) and this should be set inside the `$profile` script (see [$profile](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.2) for more details). The initial hash can be retrieved from file [`signature.PS-HASH.txt`](generators/signature.PS-HASH.txt). With this hash in place, `source` the code generation script:

> . \<path-to-script\>/generate-option-validators.ps1

Check that the hash is in place:

> $env:COBRASS_API_HASH

should show a sha256 hash.

Then invoke, show-sig:

> show-sig

which should show something like:

```
λ show-sig
---> 🍄 [gen-ps] Signature Counts - 🍅functions: '52', 🥦types: '0'
---> 🍄 [gen-ov] Signature Counts - 🍅functions: '24', 🥦types: '50'
---> 🍄 [gen-help] Signature Counts - 🍅functions: '116', 🥦types: '0'
---> 🍄 Total Counts - 🍅functions: '192', 🥦types: '50'

===> [🤖]  THIS-HASH: 'CC1F622BE3931613C9E629231E1BEE9237D0A06533049B7C47947515927ADEF4'
===> [🛡️] STATUS: '✔️ Hashes are equal' (COBRASS_API_HASH)
```

This output is what is stored in file [`signature.OUTPUT.txt`](generators/signature.OUTPUT.txt)

This is to help keep track of of evolving api changes. In keeping with Go coding guidelines, no breaking changes will be made after version 1.0 is released. But the api may change by the addition of new api functions. These changes are ok and will result in changes of the hash value.

Every time a change occurs, both files [`signature.OUTPUT.txt`](generators/signature.OUTPUT.txt) and [`signature.PS-HASH.txt`](generators/signature.PS-HASH.txt) should be kept up to date and the `$env:COBRASS_API_HASH` should also be updated in the powershell profile.

When a hash change occurs, the output will show both __THIS-HASH__ the new hash and the __OLD-HASH__, eg:

```
λ show-sig
---> 🍄 [gen-ov] Signature Counts - 🍅functions: '24', 🥦types: '50'
---> 🍄 [gen-ps] Signature Counts - 🍅functions: '52', 🥦types: '0'
---> 🍄 [gen-help] Signature Counts - 🍅functions: '116', 🥦types: '0'
---> 🍄 Total Counts - 🍅functions: '192', 🥦types: '50'

===> [🤖]  THIS-HASH: 'CC1F622BE3931613C9E629231E1BEE9237D0A06533049B7C47947515927ADEF4'
===> [👾]   OLD-HASH: '0EDD9221FA1F16EAEF5E9BFC1D8F3DC66356D734CDA10286D6647F8DF4A1B16C'
===> [🛡️] STATUS: '💥 Api changes detected' (COBRASS_API_HASH)
```

## 🧱 General Structure

At the highest level of abstraction, 2 collections of entities have been defined. 1 represents _Types_ and the other represents _Operations_. Conceptually, these are combined to form a matrix. However, due to some exceptions, not every element of this matrix is a valid code generation point, giving rise to exceptions where custom functionality is employed. An example of this is the _Bool_ type, which does not need a validator, because clearly there is little or practically nothing that can be validated for a true or false value. Another example is the pseudo _Enum_ type which performs its validation in the string domain therefore validating in a type different to its own.

### ⚜️ The Types

Inserted into a map object keyed by a logical type name. This logical type name is not the underlying Go type, rather it is a a name similar to the go type but reflects the name used in the `Cobra` api. Eg on the ___FlagSet___, there is a type based api ___BindInt64___, so there is a corresponding entry under ___Int64___.

#### ☑️ Fields of note

- ___TypeName___: same as the type key just described above
- ___GoType___: the Go type, usually in lowercase. For those types that are explicitly imported, includes the package name (eg ___time.Duration___)
- ___UnderlyingTypeName___: only applicable to `Enum` and defines the foreign type in which validation actually occurs
- ___FlagName___: defines which member of the test ___WidgetParameterSet___ the option value is bound to
- ___Short___: single letter short code of the flag
- ___Def___: the default value passed to the binder function for that flag
- ___Validatable___: defines whether a validator (in `option-validator-auto.go`) is to be defined for that type. The ___Bool___ type does not have a validator, so its ___Validatable___ is $false
- ___ForeignValidatorFn___: ___Enum___ is the only type with this set to true and it indicates that validation occurs in a different type (`string`) than its own
- ___SliceFlagName___: defines which slice member of the test ___WidgetParameterSet___ the option value is bound to
- ___SliceShort___: single letter short code of the slice flag
- ___SliceValue___: the value used on the command line to represent slice option values, in tests
- ___OptionValue___: the value used as the expected value in tests
- ___PsTcEntry___: is a sub object that is also used as a switch in ___Build-PsTestEntry___ to activate generation of a table based sub test suite for this type
- ___BindDoc___: used to allow the generation of additional custom documentation in ___Build-ParamSet___: for the BindXXXX function
- ___BindValidatedDoc___: used to allow the generation of additional custom documentation in ___Build-ParamSet___: for the BindValidatedXXXX function
- ___Comparable___: used by ___Build-BinderHelperTests___ to create comparable tests. Comparable operations are those which work with a threshold
- ___BhTests___: is a sub object that defines test data cross referenced by operation, used in ___Build-BinderHelperTests___. Some types inherit this property via its parent as defined by ___BhParent___. The reason why we can shared this test data is that data values are not significant. It's the types that are more important, so test values have been chosen that can be shared as much as possible. So you'll notice that only positive whole numbers have been used so that the same values can be shared amongst signed/unsigned/floats

- ___BhParent___: allows the sharing of test data used in ___Build-BinderHelperTests___ to cut down code duplication
- ___CastLiteralsAs___: used to perform go based type casting of test arguments

### 🔱 The Operators

Defined as a an array objects representing operations to be generated for the __Types__

#### ✅ Fields of note

- ___Name___: directly relates to the operation name used is function names, eg the `GreaterThan` of ___BindValidatedDurationGreaterThan___
- ___Documentation___: used by ___Build-BinderHelpers___ to generate operation specific documentation for the ___BindValidatedXXXX___ functions
- ___Dual___: Used to distinguish operations that require 2 values. This is currently enabled for the `Within` operator
- ___Args___: argument(s) required by the operator
- ___Condition___: most of the operators behave in a similar manner. The part that distinguishes them is the condition, which is the core of the operation
- ___Negate___: defines whether the operation is `notable`. The relational operators are not `notable`, because the opposite operation can be defined using opposing functions
- ___ExcludeTypes___: some types are not compatible with some operations. This value defines which types are not compatible with this operation

## ✨ The Generators

### 💎 Generic Concepts

Describes cross cutting concepts spanning different generators.

#### 🪁 Casting

`CastLiteralsAs` is defined on `Types` and controls how values are cast in Go. It is usually the same as the Go type, but this is not always the case, hence it is a different setting. An example of where this difference arises is with `time.Duration`. For duration literals, a conversion has to be performed to translate from its literal form to an actual duration instance. This can only be done using a function call ___time.ParseDuration___, but because it also returns an error, this has been wrapped into a helper function (duration(), see next section).

#### 💪 Helper functions

Some types are less straight forward to work with, eg `IPNet` and `IPMask`. To simplify working with these types, helper functions have been defined (in `param-set-data_test.go`). This means the gory details are hidden away and don't have to be present in the generation process:

- ipmask(v string) net.IPMask
- ipnet(v string) net.IPNet

As can be seen from the above signatures, both of these functions take a string argument, which is just a logical identity that maps to a real instance that is returned.

For `time.Duration` a ___duration(d string) time.Duration___ function has been defined. Since performing a Go type cast is the same syntactically as invoking a function, the duration function can be used to cast duration literals into duration instances.

### ✈️ Build-Validators(gen-ov)

🎯 Generates content in the form (for all `Types` that are `Validatable`):

```
type XXXXValidatorFn
type XXXXOptionValidator
func (validator XXXXOptionValidator) Validate()
```

### ✈️ Build-ParamSet(gen-ps)

🎯 Generates content in the form (for all `Types`):

```
// Bind$($spec.TypeName) binds $($displayType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//$($spec.BindDoc)
func (params *ParamSet[N]) Bind$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType)) *ParamSet[N] {...
```

🎯 Generates content in the form (and for all `Types` that are `Validatable`):

```
// BindValidated$($spec.TypeName) binds $($displayType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of $($displayType) type.
//$($spec.BindValidatedDoc)
func (params *ParamSet[N]) BindValidated$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType), validator $($validatorFn)) OptionValidator {...
```

🎯 Generates content in the form (and for all `Types` that for which slice definitions should be generated; `GenerateSlice`):

```
// Bind$($sliceTypeName) binds $($sliceType) slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) Bind$($sliceTypeName)(info *FlagInfo, to *$($sliceType)) *ParamSet[N] {...
```

🎯 Generates content in the form (and for all `Types` that are `Validatable`):

```
// BindValidated$($sliceTypeName) binds $($sliceType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of $sliceType type.
//
func (params *ParamSet[N]) BindValidated$($sliceTypeName)(info *FlagInfo, to *$($sliceType), validator $($sliceValidatorFn)) OptionValidator {...
```

### 🚀 Build-PsTestEntry(gen-ps-t)

🎭 `Sides` is designed to create tests for flags defined with and without a short code.

🎯 Generates content in the form (for all `Types` that have `PsTcEntry` and an additional similar content for all `Types` that for which slice definitions should be generated; `GenerateSlice`):

```
Entry(nil, TcEntry{
  Message: "$($displayType) type$($side.MessageAdornments) (auto)",
  Binder: func() {
    paramSet.Bind$($spec.TypeName)(
      $($constructFlagInfo),
      $($bindTo),
    )
  },
  CommandLine: "$commandLine",
  Assert: $($assertion),
}),
```

### 🚀 Build-TestEntry(gen-ov-t)

🎯 Generates content in the form (for all `Types` that are `Validatable` and an additional similar content for all `Types` that for which slice definitions should be generated; `GenerateSlice`):

```
Entry(nil, OvEntry{
  Message: "$($spec.GoType) type (auto)",
  Setup: func() {
    $($setup)
  },
  Validator: func() assistant.OptionValidator {
    $($spec.Assign)
    return paramSet.BindValidated$($spec.TypeName)(
      assistant.NewFlagInfo("$($lowerFlagName)", "$($spec.Short)", $default),
      $bindTo,
      func(value $($spec.GoType)) error {
        $($assert)
        return nil
      },
    )
  },
}),
```

### ✈️ Build-BinderHelpers(gen-help)

🎯 Generates content in the form (for all `Type`/`Operator` combinations):

```
// BindValidated$($methodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method $($op.Documentation).
// 
func (params *ParamSet[N]) BindValidated$($methodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {...
```

🎯 Generates content in the form (for all `Operators` that are `Relatable`):

```
// BindValidated$($notMethodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidated$($methodSubStmt)'.
//
func (params *ParamSet[N]) BindValidated$($notMethodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {...
```

### 🚀 Build-BinderHelperTests(gen-help-t)

🎭 `Sides` is designed to create tests for positive and negated sides of operator, eg `Contains`/`NotContains`

🎯 Generates content in the form (for all `Type`/`Dual` `Operator` combinations)

```
DescribeTable("BindValidated$($side.Method)",
  func(given, should string, value $($spec.GoType), expectNil bool, low, high $($spec.GoType)) {
    validator := paramSet.BindValidated$($side.Method)(
      assistant.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), low, high,
    )
    paramSet.Native.$($spec.FlagName) = value

```

🎯 With test case entry content in the form (for all `Type`/`Dual` `Operator` combinations and repeated for `Not` scenario):

```
Entry(nil, "value is below range", "return error", $($belowEntryArgs)),
Entry(nil, "value is equal to low end of range", "return error", $($equalLoEntryArgs)),
Entry(nil, "value is inside range", "return error", $($insideEntryArgs)),
Entry(nil, "value is equal to high end of range", "return error", $($equalHiEntryArgs)),
Entry(nil, "value is above range", "NOT return error", $($aboveEntryArgs)),
```

🎯 With test case entry content in the form (for all `Type`/`Container` `Operator` combinations and repeated for `Not` scenario):

```
Entry(nil, "collection contains member", "return error", $($doesContainArgs)),
Entry(nil, "collection does not contain member", "return error", $($doesNotContainArgs)),
```

🎯 With test case entry content in the form (for `IsMatch` `Operator` which applied only to `String` repeated for `Not` scenario):

```
Entry(nil, "value matches pattern", "return error", $($doesMatchArgs)),
Entry(nil, "value does not match pattern", "return error", $($doesNotMatchArgs)),
```

🎯 With test case entry content in the form (for all `Type`/`Comparable` `Operator` combinations):

```
Entry(nil, "value is below threshold", "return error", $($belowThresholdArgs)),
Entry(nil, "value is equal threshold", "return error", $($equalThresholdArgs)),
Entry(nil, "value is above threshold", "NOT return error", $($aboveThresholdArgs)),
```

`BhTests` (binder helper tests) is a property that provides the test data for validation helper tests. It turns out that this test data can be shared among different types as long as values are chosen that work with all types. So there is currently a single instance of `BhTests` that just so happens to have been defined in `Int`, and a separate one defined for `enum` but all other compatible types reference `Int` instance via their ___BhParent___ property.

`BhEntry` is a dictionary object mapping the `Operator` name into a test data object (let's call this the test data source). All operations take either 1 or 2 validation arguments and these are represented by ___First___ and ___Second___. For those operations that require just a single argument, the second is considered a dummy and is not referenced, but provided to make code generation logic simpler.

Test cases are generated into a `Gomega` test table structure. The ___Entry___ member of the test data source is another object that defines which entries are inserted into the test table structure. The values inside these entries are the values to be validated according the the validation arguments as they appear in ___First___ and ___Second___. The other value in each entry describes whether nil is expected as a result of the validation operation. The `Not` test cases reverse this logic using the inverse of ___expectNil___, ie ___!expectNil___.

<p align="right">
  <a href="https://go.dev"><img src="resources/images/go-sticker.jpg" width="100" /></a>
</p>

## 🔩 ___Golang code generation___

It was always intended to replace the PowerShell implementation of the Cobra based option validation code with an equivalent Go version, using language specific features such as __go:generate__ and __go:embed__. An issue was raised and now it been partly implemented. The following description will give a hint as to why it remains only partly implemented and includes discussions on the following topics:

- go templating engine
- virtual file system including an in memory solution provided by [avfs](https://github.com/avfs/avfs)

### 🛠️ Go Templating

The following resources were consulted to gain a foothold on how to use go templates

- [how to use templates in go](https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go)
- [go template syntax](https://developer.hashicorp.com/nomad/tutorials/templates/go-template-syntax)
- [using go templates](https://blog.logrocket.com/using-golang-templates/)
- [Go Templates - Simple and Powerful by Donald Feury(YouTube)](https://www.youtube.com/watch?v=dWchPTi9Vc0&t=5s)

For code generation ...

- [std stringer](https://cs.opensource.google/go/x/tools/+/refs/tags/v0.13.0:cmd/stringer/stringer.go)

📍 ___A note about Go's templating___: I considered implementing a basic templating solution, merely to get round Go's infuriating re-use of a C-like format specifiers, which are a real pain; not implementing a name place holder model like PowerShell, makes using Printf and its format specifiers a non starter due to the sheer size of the content that needed to be rendered. (I did consider implementing a scheme, where by a string can contain the placeholders much like go templating, but instead of executing, we implement a custom function, that uses the ___strings___ package functions, eg ___Replace(s, old, new string, n int) string___), to gain equivalent ___Printf___ functionality, without having to use these illegible, non descriptive '%v' placeholders, but part of the rationale of working on this feature was to learn another features of Go, so Go's templating feature won out in the end, even though the resulting code is now hugely more verbose than the the original PowerShell generator.

📍 ___Design dilemmas___

- functions: Which model do use? Do we favour implementing function on the input data types (namely type-spec), or do we use the Funcs model on on the template.

📍 ___Problems/Gotchyas___

- ___go generate___ can't evaluate variables.

📍 ___Template hint and tips___

- how to add a sub tree to a template and what this necessitates (ie using template.Tree and adding that to the parent tree) A __template.Template__ is implemented as a map, of the template name to the underlying template. If you just have a single template, then it can be executed with __template.Execute__. However, if you have multiple templates, then you need to identify which one to execute. This is done with __template.ExecuteTemplate__. The thing to note however is that the template names can either be defined explicitly or they can be derived; this depends on how you create the templates. Initially, I implemented this using an explicit mechanism, where by I created a template file for each identified code section. These were specifically loaded into a global variable via a __go:embed__ statement. This led to having many global variables whose contents were bound by the embed. Incidentally, using this approach to defining templates, necessitated building up a template map manually, with explicit template names. This would require identifying which was the base template (the first template to be created using __template.New()__ and then adding the children to it using __baseTemplate.AddParseTree("child-template-name", childTempl.Tree)__) This was becoming unscalable. I since discovered that there was a much better approach, use __ParseGlob__. With this approach in place, we could do away with the global variables with their embed statements. When the Parse template methods are used, the names of the templates are derived from the name of the file as opposed to needing to be specified explicitly.

### 🧰 Implementation

#### ⚙️ Templates folder

The templates folder at `./generators/gola/templates` contains a directory for each source code file to be generated. Each directory contains file specific template definitions, but notice how the name of the parent directory (eg ___option-validator-auto___) is repeated in the name of the child templates (eg ___option-validator-auto-header.go.tmpl___). I initially wanted the child template names to be as concise as possible so defined the children with names like `header.go.tmpl`, `footer.go.tmpl` etc. But I then learnt that the names have to be unique, otherwise a template name that has already been encountered will be overwritten and therefore not addressable. This is why the names of the parent directory are repeated in the names of the child templates.

#### Implementation Issues

- the verbosity of the code using a templates, splitting the template code over separate files

- attempting to generate code (see hyper-gen) and the mind melting delirium it sends you into as you're trying a create code that generates code that generates code; stairway to insanity ensues!

It was found that the Go implementation of the code generator is hugely work intensive and it makes me wonder if I should continued with it as a result. Since we already have a working solution and the churn in functionality required by the option validators is very low, it feels like the return on investment is very low. However, the return on investment can not be seen solely this way. There are other benefits including learning other parts of the Go eco system, which will hopefully come in handy for future projects. It took more time to implement just the generation of the first file: ___option-validator-auto___. The other files are typically more complicated and will require significantly more work, probably more than is justified for the time being. There are parts of the PowerShell generation code that were seen as a bit of a hack at the time, which I always said should be implemented more cleanly, but due to the amount of work required, I decided to simply mimic the PowerShell implementation in the Go version.

- the taskfile

#### ⚙️ Functions

For ___option-validator-auto___, a funcs map has been defined in ___source-code-container.go___. But in retrospect, these functions are not really necessary. The reason is those functions only implement very basic text substitution functionality, eg:

```go
  "getValidatorFn": func(typeName string) string {
    return typeName + "ValidatorFn"
  },
```

There is very little value of this (all it does is append the string `ValidatorFn` to the typeName passed in); see how it is currently used in ___option-validator-auto/option-validator-auto-body.go.tmpl___:

```
{{- $validatorFn := getValidatorFn .Spec.TypeName -}}
...
// {{ $validatorStruct }} defines the struct that wraps the client defined validator function
// {{ $validatorFn }} for {{ getDisplayType .Spec }} type. This is the instance that is returned by
// validated binder function BindValidated{{ .Spec.TypeName }}.
```

A local variable (__$validatorFn__) is created then referenced. But we can achieve this much more simply by inlining the addition of `ValidatorFn`:

```
// {{ .Spec.TypeName }}ValidatorFn defines the struct that wraps the client defined validator function
// {{ $validatorFn }} for {{ getDisplayType .Spec }} type. This is the instance that is returned by
// validated binder function BindValidated{{ .Spec.TypeName }}.
```

This is much more efficient and should be applied the next time work is done on this feature.

#### 🔐 Signature Check

Checking signature of the source code using the Go version is much more efficient for the development process. The PowerShell version is much more disjoint requiring handling each file individually and manually pasting content into the appropriate files. The Go version handles this much more gracefully.

This has been implemented in the ___cobrass-gen___ tool. A signature can be obtained either by performing a new generation or by requesting the signature of the existing source code. The PowerShell version depended on the user storing the existing hash value (sha256) in the powershell environment under variable ___$env:COBRASS_API_HASH___. But this is quite cumbersome as it requires some initial setup.

In the Go version, the existing hash has been embed into an internal package level variable ___gola.RegisteredHash___. This means that when the unit tests run, if a change in the API has been detected, one or more tests will fail, immediately informing the developer of either the need to check what changes they have made, or if the change is valid, then they can update the ___RegisteredHash___ to the new value.

##### With Generation

When ___cobrass-gen___ is invoked to generate code, it will return a ___SignatureResult___ which contains the newly calculated hash.

##### Inline

The user can invoke ___cobrass-gen___ to calculate the hash of existing code only, without needing to re-generate the code. This is achieved using the ___-sign___ flag. By default, the signature returned is based off the repository source code. When the ___-test___ is specified, the signature will be based off the test  code that has been generated into the test area: `./generators/gola/out/assistant`. Since the current Go generator can only generate code for ___option-validator-auto___, the hash for it, will never match the signature for the repository code, which is also based off ___param-set-auto___ and ___param-set-binder-helpers-auto___. (📌 Actually, what this means is that when we come back to completing the rest of this work, perhaps we should concentrate first on completing these then at leat we can get matching hashes).

### 📁 Virtual Filesystem

I have wondered for a long while ever since taking the decision to use Go as a language of choice for various projects, of the best way to work with the file system especially when it comes to unit testing. I've alway felt that it was wrong to use the file system directly in source code, so when it came to the code generation task, I decided to tackle this head on.

There are 2 elements to the use of a virtual file system. The first is the use of an abstraction layer that forms a facade over the native file system which allows the underlying one to be exchanged for another without affecting the source code. The second is the use of an in-memory fs so that unit tests can be developed that are entirely isolated from each other without having to resort to the native fs.

The work on this is being seen as a prototype and a guide on how to do this in other projects. There is little need to use a 3rd party solution for a facade over the native file system. This is easily implemented by delegating file system calls to 'os.'. The internal ___storage___ package achieves this in ___native-fs.go___.

When it comes to the in-memory file system, this is a little more involved so I decided not to try a rebuild this, rather, I opted to use the one provided by [avfs](https://github.com/avfs/avfs). ___avfs___ actually does implement its own facade over the native file system, but I will not aim to use this, as there is not point; not much is saved and besides, implementing delegation to a custom virtual file system is a simple undertaking.

I guess in the long term, [extendio](https://github.com/snivilised/extendio) will provide a fully implemented virtual file system to be used across snivilised projects.
