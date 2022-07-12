# 🐲 ___Powershell Code Generation___

Powershell is being used to generate go code and this article documents how this works.

There are 6 main functions that perform code generation, 3 that generate source code and the remaining generating ginkgo/gomega based test suites. The reason why code generation was needed mainly stems from the `Cobra` api. Since, they are type based functions, in order to fully integrate with it, `Cobrass` must also provide a type based api. During the initial design of cobrass, functionality was built manually, with a lot of copy and pasting. When it was discovered how laborious is was just to build the initial binder functions, it was decided that another more automated approach would be desirable. This way, any future (non-breaking!) changes can be implemented relatively swiftly, without one having to tear ones hair out due to the monotony of code modifications and it's inherent vulnerability to cut and paste errors.

The 6 generator functions (aliases in brackets) are:

- ___Build-Validators(gen-ov)___: `option-validator-auto.go` generates the core validator api types
- ___Build-ParamSet(gen-ps)___: `param-set-auto.go` generates binder functions
- ___Build-PsTestEntry(gen-ps-t)___: `param-set-auto_test.go` generates unit tests for binder functions
- ___Build-TestEntry(gen-ov-t)___: `option-validator-auto_test.go` generates unit tests for validated binder functions
- ___Build-BinderHelpers(gen-help)___: `param-set-binder-helpers-auto.go` generates validation helper functions
- ___Build-BinderHelperTests(gen-help-t)___: `param-set-binder-helpers-auto_test.go` generates unit tests for validation helper functions

The order of the above list is significant, as it reflects the order in which commands should be run (there are exceptions to this depending on the nature of the change being made).

There are an additional 2 functions that have been built to protect against accidental api breaking changes, afterall _with great power comes great responsibility_, that is to say, with code generation in play, it could be easy to release a breaking change and the possibility of this has to be stamped out as much as possible.

These 2 commands are:

+ ___Checkpoint-ParamSetSignatures(check-sig)___: invokes all 3 source code generators and creates a hash from the generated api surface
+ ___Show-ParamSetSignatures___(show-sig): displays the result of the calculated hash from check-sig. Also compares with the previous hash to provide a quick and easy mechanism to determine if any api changes have occured

## 🤖 Using the Code Generators

### 👉 Invoking

To help insertion of generated code, the scripts make use of the system clipboard. Once a command is run, the contents are copied to the clipboard. The user then has to insert that content into the source file indicated. There are markers inside the code files which indicates where the new code should be inserted.

When running for the first time, it is recommended to set the current hash inside the powershell session. This is done via an environment variable (___$env:COBRASS_API_HASH___) and this should be set inside the `$profile` script (see [$profile](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.2) for more details). The initial hash can be retrieved from file [`signature.HASH.txt`](generators/signature.HASH.txt). With this hash in place, `source` the code generation script:

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

Everytime a change occurs, both files [`signature.OUTPUT.txt`](generators/signature.OUTPUT.txt) and [`signature.HASH.txt`](generators/signature.HASH.txt) should be kept up to date and the `$env:COBRASS_API_HASH` should also be updated in the powershell profile.

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

At the highest level of abstraction, 2 collections of entites have been defined. 1 represents _Types_ and the other represents _Operations_. Conceptually, these are combined to form a matrix. However, due to some exceptions, not every element of this matrix is a valid code generation point, giving rise to exceptions where custom functionality is employed. An example of this is the _Bool_ type, which does not need a validator, because clearly there is little or practically nothing that can be validated for a true or false value. Another example is the pseudo _Enum_ type which performs its validation in the string domain therefore validating in a type different to its own.

(sides concept)

### ⚜️ The Types

Inserted into a map object keyed by a logical type name. This logical type name is not the underlying Go type, rather it is a a name similar to the go type but reflects the name used in the cobra api. Eg on the ___FlagSet___, there is a type based api ___BindInt64___, so there is a corresponding entry under ___Int64___.

#### ☑️ Fields of note

- ___TypeName___: same as the type key just described above
- ___GoType___: the Go type, usually in lowercase. For those types that are explicitly imported, includes the package name (eg ___time.Duration___)
- ___UnderlyingTypeName___: only applicable to `Enum` and defines the foreign type in which validation actually occurs
- ___FlagName___: defines which member of the test ___WidgetParameterSet___ the option value is bound to
- ___Short___: single letter short code of the flag
- ___Def___: the default value passed to the binder function for that flag
- ___Validatable___: defines whether a validator (in `option-validator-auto.go`) is to be defined for that type. The ___Bool___ type does not have a validator, so its ___Validatable___ is $false
- ___ForeignValidatorFn___: ___Enum___ is the only tyoe with this set to true and it indicates that validation occurs in a different type (`string`) than its own
- ___SliceFlagName___: defines which slice member of the test ___WidgetParameterSet___ the option value is bound to
- ___SliceShort___: single letter short code of the slice flag
- ___SliceValue___: the value used on the command line to represent slice option valies, in tests 
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
- ___Args___: arguement(s) required by the operator
- ___Condition___: most of the operators behave in a similar manner. The part that distinuishes them is the condition, which is the core of the operation
- ___Negate___: defines whether the operation is `notable`. The relational operators are not `notable`, because the opposite operation can be defined using opposing functions
- ___ExcludeTypes___: some types are not compatible with some operations. This value defines which types are not compatible with this operation


## ✨ The Generators

### 💎 Generic Concepts

Describes cross cutting concepts spanning different generators.

#### 🪁 Casting

`CastLiteralsAs` is defined on `Types` and controls how values are cast in Go. It is usually the same as the Go type, but this is not always the case, hence it is a different setting. An example of where this difference arises is with `time.Duration`. For duration literals, a conversion has to be performed to translate from its literal form to an actual duration instance. This can only be done using a function call ___time.ParseDuration___, but because it also returns an error, this has been wrapped into a helper function (duration(), see next section).

#### 💪 Helper functions

Some types are less straight forward to work with, eg `IPNet` and `IPMask`. To simply working with these types, helper functions have been defined (in `param-set-data_test.go`). This means the gory details are hidden away and don't have to be present in the generation process:

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
  Validator: func() adapters.OptionValidator {
    $($spec.Assign)
    return paramSet.BindValidated$($spec.TypeName)(
      adapters.NewFlagInfo("$($lowerFlagName)", "$($spec.Short)", $default),
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
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
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

`BhTests` (binder helper tests) is property that provides the test data for tests. It turns out that this test data can be shared among different types as long as values are chosen that work with all types. So there is only a single instance of `BhTests` that just so happens to have been defined in `Int`, but all other compatible types reference this via their ___BhParent___ property.

`BhEntry` is a dictionary object mapping the `Operator` name into a test data object (let's call this the test data source). All operations take either 1 or 2 validation arguments and these are represented by ___First___ and ___Second___. For those operations that require just a single argument, the second is considered a dummy and is not referenced, but provided to make code generation logic simpler.

Test cases are generated into a `Gomega` test table structure. The ___Entry___ member of the test data source is another object that defines which defines which entries are inserted into the test table structure. The values inside these entries are the values to be validated according the the validation arguments as they appear in ___First___ and ___Second___. The other value in each entry describes whether nil is expected as a result of the validation operation.
