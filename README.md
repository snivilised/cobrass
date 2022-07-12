# 🐲 ___Cobrass: assistant for cli applications using cobra___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/cobrass.svg)](https://pkg.go.dev/github.com/snivilised/cobrass)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/cobrass)](https://goreportcard.com/report/github.com/snivilised/cobrass)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/cobrass/badge.svg?branch=master)](https://coveralls.io/github/snivilised/cobrass?branch=master)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

_[Cobra](https://cobra.dev/) is an excellent framework for the development of command line applications, but there are few aspects that could do with being made easier to work with. This package aims to fullfil this purpose, especially in regards to creation of commands, encapulating commands into a container and providing an export mechanism to re-create cli data in a form that is free from cobra (and indeed cobrass) abstractions. The aim of this last aspect to to be able to inject data into the core of an application in a way that removes tight coupling to the cobra framework, which is achieved by representing data only in terms of client defined (native) abstractions. Currently, Cobra does not provide a mechanism for validating option values, this is also implemented by_ ___Cobrass___.

___Status___: 💤 not yet published

## 🔨 Usage

To install _cobrass_ into an application:

> go get github.com/snivilised/cobrass@latest

Most of the functionality is defined in the _adapters_ package so import as:

> import "github.com/snivilised/cobrass/src/adapters"

## 🎀 Features

- Cobra container; collection of cobra commands that can be independently referenced by name as opposed to via child/parent relationship. The container also takes care of adding commands to the root or any other as required.
- A `parameter set` groups together all the flag option values, so that they don't have to be handled separately. A single entity (the ___ParamSet___) can be created and passed into the core of the client application.
- Pseudo int based enum; provides a mapping between user specifed enum string values and the their internal int based representation.
- Option value validation; a user defined function can be provided for each option value to be validated
- Option validator helpers; as an alternative to providing a function to perform option validation, the client can invoke any of the predefined validator helpers for various types.


### 🎁 Cobra Container

The container serves as a repository for `Cobra` commands and `Cobrass` parameter sets. Commands in `Cobra` are related to each other via parent child relationships. The container, flattens this hierarchy so that a command can be queried for, simply by its name, as opposed to getting the commands by parent command, ie ___parentCommand.Commands()___.

The methods on the container, should not fail. Any failures that occur are due to programming errors. For this reason, when an error scenario occurs, a panic is raised.

Regsitering commands/parameter sets with the container, obviates the need to use specific `Cobra` api calls as they are handled on the clients behalf by the container. For parameter sets, the type specific methods on the various ___FlagSet___ definitions, such as ___Float32Var___, do not have to be called by the client. For commands, ___AddCommand___ does not have to be called explicitly either.

### 💎 Param Set

The rationale behind the concept of a parameter set came from initial discovery of how the `Cobra` api worked. Capturing user defined command line input requires binding option values into disparate variables. Having to manage independently defined variables usually at a package level could lead to a scattering of these variables on an adhoc basis. Having to then pass all these items independently into the core of a client application could easily become disorganised.

To manage this, the concept of a `parameter set` was introduced to bring about a consistency of design to the implemenation of multiple cli applications. The aim of this is to reduce the number package level global variables that have to be managed. Instead of handling multiple option variables independently, the client can group them together into a parameter set.

Each `Cobra` command can define multiple parameter sets which reflects the different ways that a particular command can be invoked by the user. However, to reduce complexity, it's probably best to stick with a single parameter set per command. Option values not defined by the user can already be defaulted by the `Cobra` api itself, but it may be, that distinguishing the way that a command is invoked (ie what combination of flags/options appear on the command line) may be significant to the application, in which case the client can define multiple parameter sets.

The ___ParamSet___ also handles flag definition on each command. The client defines the flag info and passes this into the appropriate `binder` method depending on the option value type. There are 3 forms of binder methods:

- 1️⃣ ___Bind\<Type>___ : where ___\<Type>___ represents the type, (eg ___BindString___), the client passes in '___info___', a ___FlagInfo___ object and '___to___' a pointer to a variable to which `Cobra` will bind the option value to.

- 2️⃣ ___BindValidated\<Type>___: (eg BindValidatedString) same as 1️⃣, except the client can also pass in a function whose signature reflects the type of the option value to be bound to (See [Option Validators](#option-validators)).

- 3️⃣ ___BindValidated\<Type>\<Op>___: (eg BindValidatedStringWithin) same as 2️⃣, except client passes in operation specific parameters (See [Validation Helpers](#validation-helpers)).

📌 The names of the ___BindValidated\<Type>\<Op>___ methods are not always strictly in this form as sometime it reads better with _Op_ and _Type_ are swapped around especially when one considers that there are _Not_ versions of some commands. The reader is invited to review the go package documentation to see the exact names.

### 💠 Pseudo Enum

Since Go does not have built in support for enums, this feature has to be faked by the use of custom definitions. Typically these would be via int based type definitions. However, when developing a cli, attention has to be paid into how the user specifies discreet values and how they are interpreted as options.

There is a disparity between what the user would want to specify and how these values are represented internally by the application. Typically in code, we'd want to represent these values with longer more expressive names, but this is not necessarily user friendly. For example given the following pseudo enum definition:

```go
type OutputFormatEnum int

const (
  _ OutputFormatEnum = iota
	XmlFormatEn
	JsonFormatEn
	TextFormatEn
	ScribbleFormatEn
)
```

... how would we allow the user represent these values as options on the command line? We could require that the user specify the names exactly as above, but those names are not user friendly. Rather, we would prefer something simple like 'xml' to represent ___XmlFormatEn___, but that would be unwise in code, because the name 'xml' is too generic and would more than likely clash with another identifier named xml in the package.

This is where the type ___EnumInfo___ comes into play. It allows us to provide a mapping between what the user would type in and how this value is represented internally.

####  🍑 Enum Info

An ___EnumInfo___ instance for our psuedo enum type ___OutputFormatEnum___ can be created with ___NewEnumInfo___ as follows:

```go
OutputFormatEnumInfo = adapters.NewEnumInfo(adapters.AcceptableEnumValues[OutputFormatEnum]{
  XmlFormatEn:      []string{"xml", "x"},
  JsonFormatEn:     []string{"json", "j"},
  TextFormatEn:     []string{"text", "tx"},
  ScribbleFormatEn: []string{"scribble", "scribbler", "scr"},
})
```

Points to note from the above:

- The argument passed into ___NewEnumInfo___ is a map of our enum value to a slice of 'acceptable' strings. We define a slice for each enum value so that we can define multiple ways of representing that value to aid usability. So for ___XmlFormatEn___, the user can type this either as 'xml' or even just 'x'.

- The return value of ___NewEnumInfo___ is an instance that represents the `meta` data for the pseudo enum type.

- Each application need only create a single instance of ___EnumInfo___ for each enum entity so logically this should be treated as a singleton, although it hasnt been enforced as a singleton in code.


#### 🍉 Enum Value<a name="enum-value"></a>

The client can create ___EnumValue___ variables from the ___EnumInfo___ as follows:

```go
outputFormatEnum := OutputFormatEnumInfo.NewValue()
```

Points to note from the above:

- As many enum values as needed in the client can be created

- A string value can be checked to determine if it is a valid value (as defined by the acceptable values passed into ___NewEnumInfo___), to the ___IsValid___ method on the ___EnumInfo___  or we can simply call the same method on ___EnumValue___ without passing in a string value; in this case, the check is performed on it's member variable 'Source' which can be assigned at any time.

- The ___EnumInfo___ struct contains a ___String___ method to support printing. It is provided because passing in the ___int___ form of the enum value to a printing function just results in the numeric value being displayed, which is not very useful. Instead, when there is a need to print an ___EnumValue___, it's custom ___String___ method should be invoked. Since that method retrieves the first acceptable value defined for the enum value, the user should specify a longer more expressive form as the first entry, followed by 1 or more shorter forms. Actually, to be clear, as long as the first item is expressive enough when displayed in isolation, it doesn't really matter if the first item is the longest or not. 

#### 🍈 Enum Slice

If an option value needs to be defined as a collection of enum values, then the client can make use of ___EnumSlice___.

📌 ___An enum slice is not the same as defining a slice of enums, eg []MyCustomEnum___, because doing so in that manner would incorrectly replicate the 'parent' EnumInfo reference. Using ___EnumSlice___, ensures that there is just a single EnumInfo reference for multiple enum values.

In the same way an ___EnumValue___ can be created off the ___EnumInfo___, an ___EnumSlice___ can be created by invoking the ___NewSlice___ method off ___EnumInfo___, eg:

```go
outputFormatSlice := OutputFormatEnumInfo.NewSlice()
```

___NewSlice___ contains various _collection_ methods equilavent to it's value based (___EnumValue___) counterpart.

## ☂️ Option Binding and Validation

The following sections describe the validation process, option validators and the helpers.

📌 ___When using the option validators, there is no need to use the `Cobra` flag set methods (eg cmd.Flags().StringVarP) directly to define the flags for the command. This is taken care of on the client's behalf___.

### ✅ Validation Sequencing<a name="validation-sequencing"></a>

The following is a checklist of items that need to be performed:

- 1️⃣ _create cobra container_: typically in the same place where the root command is defined. The root command should then be passed into ___Container___ constructor function ___NewCobraContainer___ eg:

```go
var Container = adapters.NewCobraContainer(&cobra.Command{
	Use:   "root",
	Short: "foo bar",
	Long: "This is the root command.",
})
var rootCommand = Container.Root()
```

- 2️⃣ _register sub commands_: for each sub command directly decended from the root, on the ___Container___ instance, invoke ___RegisterRootedCommand___ eg:

```go
  Container.RegisterRootedCommand(widgetCommand)
```

If a command is a descendent of a command other than the root, then this command should be registered using ___RegisterCommand___ instead. eg:

assuming a command with the name "foo", has already been registered

```go
  Container.RegisterCommand("foo", widgetCommand)
```

📌 ___Note, when using the Cobra Container to register commands, you do not need to use Cobra's AddCommand. The container takes care of this for you.___

- 3️⃣ _define native parameter set_: for each parameter set associated with each command eg:

```go
type WidgetParameterSet struct {
	Directory string
	Format    OutputFormatEnum
	Concise   bool
	Pattern   string
}
```

- 4️⃣ _create the ParamSet_: for each native parameter set using ___NewParamSet___ eg:

```go
  paramSet = adapters.NewParamSet[WidgetParameterSet](widgetCommand)
```

The result of ___NewParamSet___ is an object that contains a member ___Native___. This `native` member is the type of the paramter set that was defined, in this case `WidgetParameterSet`.

- 5️⃣ _define the flags_: use the `binder` methods on the _ParamSet_ to declare the commands flags.

The members of an instance of this `native` param set will be used to `bind` to when binding values, eg:

```go
  paramSet.BindValidatedString(
    adapters.NewFlagInfo("directory", "d", "/foo-bar"),
    &paramSet.Native.Directory,
    func(value string) error {
      if _, err := os.Stat(value); err != nil {
        if os.IsNotExist(err) {
            return err
        }
      }
      return nil
    },
  )
```

... and a specialisation for enum members:

```go
  outputFormatEnum := outputFormatEnumInfo.NewValue()
  paramSet.BindValidatedEnum(
    adapters.NewFlagInfo("format", "f", "xml"),
    &outputFormatEnum.Source,
    func(value string) error {
      Expect(value).To(Equal("xml"))
      return nil
    },
  )
```

Note, because we can't bind directly to the `native` member of WidgetParameterSet, (that being ___Format___ in this case), since the user will be typing in a string value that is internally represented as an int based `enum`, we have to bind to ___Source___, a string member of an ___EnumValue___, ie ___&outputFormatEnum.Source___ in the above code snippet. Later on we'll simply copy the value over from ___outputFormatEnum.Source___ to where its supposed to be, ___paramSet.Native.Format___.

- 6️⃣ _register param set_: this is optional, but doing do means that the param set can easily be retrieved at a later point. The param set is registered (typically after all the flags have been bound in) as follows:

```go
  Container.RegisterParamSet("widget-ps", paramSet)
```

- 7️⃣ _rebind enum values_: in the function defined as the Run/RunE member of the command, the entry point of application execution, we now need to 'rebind' the enum members. In the previous code snippet, we can see that a new ___EnumValue___ was created from the ___EnumInfo___, ie ___outputFormatEnum___. We can set the value of the enum to the appropriate `native` member, so in this case it would be:

```go
  paramSet.Native.Format = outputFormatEnum.Value()
```

- 8️⃣ _invoke option validation_: also inside the command's ___Run/RunE___ run function, before entering into the core of the application, we need to invoke option validation:

```go
  RunE: func(command *cobra.Command, args []string) error {

    ps := container.ParamSet("widget-ps").(*adapters.ParamSet[WidgetParameterSet])

    if err := ps.Validate(); err == nil {
      native = ps.Native

      // rebind enum into native member
      //
      native.Format = outputFormatEnum.Value()

      // optionally invoke cross field validation
      //
      xv := ps.CrossValidate(func(ps *WidgetParameterSet) error {
        condition := (ps.Format == XmlFormatEn)
        if condition {
          return nil
        }
        return fmt.Errorf("format: '%v' is invalid", ps.Format)
      })

      if (xv == nil) {
        // ---> execute application core with the parameter set (native)
        //
        // runApplication(native)
        //
      } else {
        return xv
      }

    } else {
      return err
    }
  },
```

The validation may occur in 2 stages depending on whether cross field valiation is required. To proceed, we need to obtain both the wrapper parameter set (ie ___container.ParamSet___ in this example) and the native parameter set ___native = ps.Native___).

Also note how we retrieve the parameter set previously registered from the cobra container using the ___Native___ method. Since ___Native___ returns ___any___, a type assertion has to be performed to get back the `native` type. If the param set you created using ___NewParamSet___ is in scope, then there is no need to query the container for it by name. It is just shown here this way, to illustrate how to proceed if parameter set was created in a local function/method and is therefore no longer in scope.

Option validation occurs first (___ps.Validate()___), then rebinding of enum members, if any (___native.Format = outputFormatEnum.Value()___), then cross field validation (___xv := ps.CrossValidate___).

If we have no errors at this point, we can enter the application, passing in the native parameters set.

The validation process will fail on the first error encountered and return that error. It is not mandatory to register the parameter set this way, it is there to help minimise the number of package global variables.

- 9️⃣ _invoke cross field validation_ (optional): see [Cross Field Validation](#cross-field-validation)

### 🎭 Alternative Flag Set

By default, binding a flag is performed on the default flag set. This flag set is the one you get from calling ___command.Flags()___ (this is performed automatically by ___NewFlagInfo___). However, there are a few more options for defining flags in `Cobra`. There are multiple flag set methods on the `Cobra` command, eg ___command.PersistentFlags()___. To utilise an alternative flag set, the client should use ___NewFlagInfoOnFlagSet___ instead of ___NewFlagInfo___. ___NewFlagInfoOnFlagSet___ requires that an extra parameter be provided and that is the alternative flag set, which can be derived from calling the appropriate method on the `command`, eg:

```go
  paramSet.BindString(
    adapters.NewFlagInfoOnFlagSet("pattern", "p", "default-pattern",
      widgetCommand.PersistentFlags()), &paramSet.Native.Pattern,
  )
```

The flag set defined for the flag (in the above case 'pattern'), will always override the default one defined on the parameter set.

### ⛔ Option Validators<a name="option-validators"></a>

As previously described, the validator is a client defined type specific function that takes a single argument representing the option value to be validated. The function should return nil if valid, or an error describing the reason for validation failure.

There are multiple ___BindValidated___ methods on the ___ParamSet___, all which relate to the different types supported. The `binder` method simply adds a wrapper around the function to be invoked later and adds that to an internal collection. The wrapper object is returned, but need not be consumed.

For `enum` validation, ___ParamSet___ contains a validator ___BindValidatedEnum___. It is important to be aware that the validation occurs in the `string` domain not in the `int` domain as the reader might expect. So when a `enum` validator is defined, the function has to take a string parameter, not the native `enum` type.

The following is an example of how to define an `enum` validator:

```go
  outputFormatEnum := outputFormatEnumInfo.NewValue()

  wrapper := paramSet.BindValidatedEnum(
    adapters.NewFlagInfo("format", "f", "xml"),
    &outputFormatEnum.Source,
    func(value string) error {
      return lo.Ternary(outputFormatEnumInfo.IsValid(value), nil,
        fmt.Errorf("Enum value: '%v' is not valid", value))
    },
  )
  outputFormatEnum.Source = "xml"
```

The following points should be noted:

- validation is implemented using the ___EnumInfo___ instance. This could easily have been implemented using an ___EnumValue___ instance instead.
- the manual assignment of __'outputFormatEnum.Source'___ is a synthetic operation just done for the purposes of illustration. When used within the context of a cobra cli, it's cobra that would perform this assignment as it parses the command line, assuming the corresponding flag has been bound in as is peformed here using ___BindValidatedEnum___.
- the client would convert this string to the enum type and set on the appropriate native member (ie ___paramSet.Native.Format = outputFormatEnum.Value()___)

To bind a flag without a short name, the client can either:

- pass in an empty string for the ___Short___ parameter of ___NewFlagInfo___ eg:

```go
  adapters.NewFlagInfo("format", "", "xml"),
```

or

- not use the ___NewFlagInfo___ constructor function at all and pass in a literal struct without setting the ___Short___ member. Note in this case, make sure that the ___Name___ property is set properly, ie it should be the first word of ___Usage___ eg:

```go
  paramSet.BindValidatedEnum(
    &adapters.FlagInfo{
      Name: "format",
      Usage: "format usage",
      Default: "xml",
		},
    &outputFormatEnum.Source,
    func(value string) error {
      return lo.Ternary(outputFormatEnumInfo.IsValid(value), nil,
        fmt.Errorf("Enum value: '%v' is not valid", value))
    },
  )
```

### 🛡️ Validator Helpers<a name="validation-helpers"></a>

An alternative way of implementing option validation, the client can use the validation helpers defined based on type.

The following are the categories of helpers that have been provided:

- _comparison(threshold)_: ___GreaterThan(\> threshold)___, ___AtLeast(\>= threshold)___, ___LessThan(\< threshold)___, ___AtMost(\<= threshold)___
- _range(lo, hi)_: ___Within(\>= lo and \<= hi)___
- _collection(collection)_: ___Contains(is member of collection)___

Specialised for type:
- _string_: ___'BindValidatedStringIsMatch'___

`Not` versions of most methods have also been provided, so for example to get string not match, use ___'BindValidatedStringIsNotMatch'___. The `Not` functions that have been omitted are the ones which can easily be implemented by using the opposite operator. There are no `Not` versions of the _comparison_ helpers, eg there is no ___'BindValidatedIntNotGreaterThan'___ because that can be easily acheieved using ___'BindValidatedIntAtMost'___.

There are also `slice` versions of some of the validators, to allow an option value to be defined as a collection of values. An example of a `slice` version is ___'BindValidatedStringSlice'___.

Our pseudo enums are a special case, because it is not possible to define generic versions of the binder methods where a generic parameter would be the client defined int based enum, there are no option validator helpers for `enum` types.

### ⚔️ Cross Field Validation<a name="cross-field-validation"></a>

When the client needs to perform cross field validation, then ___ParamSet.CrossValidate___ should be invoked. Cross field validation is meant for checking option values of different flags, so that cross field constraints can be imposed. Contrary to `option validators` and `validator helpers` which are based upon checking values compare favourably against static boundaries, `cross field validation` is concerned with checking the dynamic value of options of different flags. The reader should be aware this is not about enforcing that all flags in a group are present or not. Those kinds of checks are already enforceable via `Cobra's` group checks. It may be that 1 option value must constrain the range of another option value. This is where cross field validation can be utilised.

The client should pass in a validator function, whose signature contains a pointer to the native parameterset, eg:

```go
  result := paramSet.CrossValidate(func(ps *WidgetParameterSet) error {
    condition := (ps.Strike >= ps.Lower) && (ps.Strike <= ps.Higher)

    if condition {
      return nil
    }
    return fmt.Errorf("strike: '%v' is out of range", ps.Strike)
  })
```

The native parameter set, should be in its 'finalised' state. This means that all parameters should be bound in. So in the case of pseudo enum types, they should have been populated from temporary placeholder enum values. Recall from step 7️⃣ _rebind enum values_ of [Validation Sequencing](#validation-sequencing), that enum members have to be rebound. Well this is what is meant by finalisation. Before cross field validation is invoked, make sure that the enum members are correctly set. This way, you can be sure that the cross field validator is working with the correct state of the native parameter set. The validator can work in the 'enum domain' as opposed to checking raw string values, eg:

```go
  result := paramSet.CrossValidate(func(ps *WidgetParameterSet) error {
    condition := (ps.Format == XmlFormatEn)
    if condition {
      return nil
    }
    return fmt.Errorf("format: '%v' is invalid", ps.Format)
  })
```

This is a rather contrived example, but the important part of it is the use of the enum field ___ps.Format___.

## 🧰 Developer Info

### 🥇 Task Runner

<p align="left">
  <a href="https://taskfile.dev/"><img src="https://taskfile.dev/img/logo.svg" width="50" /></a>
</p>

Uses [Taskfile](https://taskfile.dev/). A simple `Taskfile.yml` has been defined in the root dir of the repo and defines tasks that make building and running [Ginkgo](https://onsi.github.io/ginkgo/) commands easier to perform.

### ✨ Code Generation

Please see [Powershell Code Generation](CODE-GEN.md)

### 🧪 Unit Testing

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" /></a>
</p>

[Ginkgo](https://onsi.github.io/ginkgo/) is the bbd testing style of choice used in `Cobrass`. I have found it to be a total revelation to work work with, in all aspects except 1, which was discovered well after I had gone all in on `Ginkgo`. I am using the Ginkgo test explorer in `vscode` and while it is good at exploring tests, running them and even generating coverage with little fuss, the single fly in the ointment is that debuging test cases is currently difficult to achieve:

```
Starting: /home/plastikfan/go/bin/dlv dap --check-go-version=false --listen=127.0.0.1:40849 --log-dest=3 from /home/plastikfan/dev/github/go/snivilised/cobrass/src/adapters
DAP server listening at: 127.0.0.1:40849
Type 'dlv help' for list of commands.
Running Suite: Adapters Suite - /home/plastikfan/dev/github/go/snivilised/cobrass/src/adapters
==============================================================================================
Random Seed: 1657619476

Will run 0 of 504 specs
SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS

Ran 0 of 504 Specs in 0.016 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 504 Skipped
You're using deprecated Ginkgo functionality:
=============================================
  --ginkgo.debug is deprecated
  Learn more at: https://onsi.github.io/ginkgo/MIGRATING_TO_V2#removed--debug
  --ginkgo.reportFile is deprecated, use --ginkgo.junit-report instead
  Learn more at: https://onsi.github.io/ginkgo/MIGRATING_TO_V2#improved-reporting-infrastructure

To silence deprecations that can be silenced set the following environment variable:
  ACK_GINKGO_DEPRECATIONS=2.1.4
```

vscode, debugging is an issue
