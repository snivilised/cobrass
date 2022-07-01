
$types = @{
  "Enum"     = [PSCustomObject]@{
    TypeName           = "Enum"
    GoType             = "string"
    UnderlyingTypeName = "String"
    FlagName           = "Format"
    Short              = "f"
    Def                = "xml"
    Assign             = "outputFormatEnum := outputFormatEnumInfo.NewValue()"
    Setup              = "paramSet.Native.Format = XmlFormatEn"
    BindTo             = "&outputFormatEnum.Source"
    Assert             = "Expect(value).To(Equal(""xml""))"
    QuoteExpect        = $true
    Equate             = "Equal"
    GenerateSlice      = $false
    SliceFlagName      = "Formats"
    SliceShort         = "F"
    DefSliceVal        = "[]string{}"
    ExpectSlice        = "[]string{""xml"", ""json"", ""text""}"
    #
    # Currently 'Comparable' for enum disabled because enum comparison would be
    # performed in the string domain but it might make more sense to the use if
    # it was in the int domain. We don't want to commit to publish this particluar
    # api, if it's not clear how this would be implemented, so that it makes sense.
    #
    # Comparable         = $true
  }

  "String"   = [PSCustomObject]@{
    TypeName      = "String"
    GoType        = "string"
    FlagName      = "Pattern"
    Short         = "p"
    Def           = "default-pattern"
    Setup         = "paramSet.Native.Pattern = ""*music.infex*"""
    Assert        = "Expect(value).To(Equal(""*music.infex*""))"
    QuoteExpect   = $true
    Equate        = "Equal"
    GenerateSlice = $true
    SliceFlagName = "Categories"
    SliceShort    = "C"
    DefSliceVal   = "[]string{}"
    ExpectSlice   = "[]string{""alpha"", ""beta"", ""delta""}"
    #
    Comparable    = $true
    #
    BhTests       = @{
      "Within"      = @{
        First  = """c"""
        Second = """e"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below   = @("""b""", "false")
          EqualLo = @("""c""", "true")
          Inside  = @("""d""", "true")
          EqualHi = @("""e""", "true")
          Above   = @("""f""", "false")
        }
      }

      "Contains"    = @{
        First  = "[]string{""a"", ""b"", ""c""}"
        Second = """null"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          DoesContain    = @("""a""", "true")
          DoesNotContain = @("""x""", "false")
        }
      }

      # for string only!
      #
      "IsMatch"     = @{
        First  = """\\d{2}-\\d{2}-\\d{4}"""
        Second = """null"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          DoesMatch    = @("""18-10-1997""", "true")
          DoesNotMatch = @("""foo-bar""", "false")
        }
      }

      "GreaterThan" = @{
        First  = """c"""
        Second = """"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""b""", "false")
          Equal = @("""c""", "false")
          Above = @("""d""", "true")
        }
      }

      "AtLeast"     = @{
        First  = """c"""
        Second = """"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""b""", "false")
          Equal = @("""c""", "true")
          Above = @("""d""", "true")
        }
      }

      "LessThan"    = @{
        First  = """c"""
        Second = """"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""b""", "true")
          Equal = @("""c""", "false")
          Above = @("""d""", "false")
        }
      }

      "AtMost"      = @{
        First  = """c"""
        Second = """"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""b""", "true")
          Equal = @("""c""", "true")
          Above = @("""d""", "false")
        }
      }
    }
  }

  "Int"      = [PSCustomObject]@{
    TypeName      = "Int"
    GoType        = "int"
    FlagName      = "Offset"
    Short         = "o"
    Def           = -1
    Setup         = "paramSet.Native.Offset = -9"
    Assert        = "Expect(value).To(Equal(-9))"
    Equate        = "Equal"
    GenerateSlice = $true
    SliceFlagName = "Dimensions"
    SliceShort    = "D"
    DefSliceVal   = "[]int{}"
    ExpectSlice   = "[]int{2, 4, 6, 8}"
    #
    Comparable    = $true
    Threshold     = 10
    Value         = 9
    GtExpectNil   = "false"
    #
    BhTests       = @{
      "Within"      = @{
        First  = "3"
        Second = "5"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below   = @("2", "false")
          EqualLo = @("3", "true")
          Inside  = @("4", "true")
          EqualHi = @("5", "true")
          Above   = @("6", "false")
        }
      }

      "Contains"    = @{
        # Any test data that contains a type spec, needs to be defined by a template
        #
        First  = "[]{{SLICE-TYPE}}{1, 2, 3}"
        Second = "0"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          DoesContain    = @("1", "true")
          DoesNotContain = @("99", "false")
        }
      }

      "GreaterThan" = @{
        First  = "3"
        Second = "0"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("2", "false")
          Equal = @("3", "false")
          Above = @("4", "true")
        }
      }

      "AtLeast"     = @{
        First  = "3"
        Second = "0"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("2", "false")
          Equal = @("3", "true")
          Above = @("4", "true")
        }
      }

      "LessThan"    = @{
        First  = "3"
        Second = "0"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("2", "true")
          Equal = @("3", "false")
          Above = @("4", "false")
        }
      }

      "AtMost"      = @{
        First  = "3"
        Second = "0"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("2", "true")
          Equal = @("3", "true")
          Above = @("4", "false")
        }
      }
    }
  }

  "Int8"     = [PSCustomObject]@{
    TypeName       = "Int8"
    GoType         = "int8"
    FlagName       = "Offset8"
    Short          = "o"
    Def            = "int8(-1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Offset8 = int8(-99)"
    Assert         = "Expect(value).To(Equal(int8(-99)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "int8"
    #
    BhTests        = $null
  }

  "Int16"    = [PSCustomObject]@{
    TypeName       = "Int16"
    GoType         = "int16"
    FlagName       = "Offset16"
    Short          = "o"
    Def            = "int16(-1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Offset16 = int16(-999)"
    Assert         = "Expect(value).To(Equal(int16(-999)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "int16"
    #
    BhTests        = $null
  }

  "Int32"    = [PSCustomObject]@{
    TypeName       = "Int32"
    GoType         = "int32"
    FlagName       = "Offset32"
    Short          = "o"
    Def            = "int32(-1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Offset32 = int32(-9999)"
    Assert         = "Expect(value).To(Equal(int32(-9999)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "int32"
    #
    BhTests        = $null
  }

  "Int64"    = [PSCustomObject]@{
    TypeName       = "Int64"
    GoType         = "int64"
    FlagName       = "Offset64"
    Short          = "o"
    Def            = "int64(-1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Offset64 = int64(-99999)"
    Assert         = "Expect(value).To(Equal(int64(-99999)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "int64"
    #
    BhTests        = $null
  }

  "Unit"     = [PSCustomObject]@{
    TypeName       = "Uint"
    GoType         = "uint"
    FlagName       = "Count"
    Short          = "c"
    Def            = "uint(0)"
    CastDef        = $true
    Setup          = "paramSet.Native.Count = uint(99999)"
    Assert         = "Expect(value).To(Equal(uint(99999)))"
    Equate         = "Equal"
    GenerateSlice  = $true
    SliceFlagName  = "Points"
    SliceShort     = "P"
    DefSliceVal    = "[]uint{}"
    ExpectSlice    = "[]uint{2, 4, 6, 8}"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint"
    #
    BhTests        = $null
  }

  "Uint8"    = [PSCustomObject]@{
    TypeName       = "Uint8"
    GoType         = "uint8"
    FlagName       = "Count8"
    Short          = "c"
    Def            = "uint8(1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Count8 = uint8(33)"
    Assert         = "Expect(value).To(Equal(uint8(33)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint8"
    #
    BhTests        = $null
  }

  "Uint16"   = [PSCustomObject]@{
    TypeName       = "Uint16"
    GoType         = "uint16"
    FlagName       = "Count16"
    Short          = "c"
    Def            = "uint16(1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Count16 = uint16(333)"
    Assert         = "Expect(value).To(Equal(uint16(333)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint16"
    #
    BhTests        = $null
  }

  "Uint32"   = [PSCustomObject]@{
    TypeName       = "Uint32"
    GoType         = "uint32"
    FlagName       = "Count32"
    Short          = "c"
    Def            = "uint32(1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Count32 = uint32(3333)"
    Assert         = "Expect(value).To(Equal(uint32(3333)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint32"
    #
    BhTests        = $null
  }

  "Uit64"    = [PSCustomObject]@{
    TypeName       = "Uint64"
    GoType         = "uint64"
    FlagName       = "Count64"
    Short          = "c"
    Def            = "uint64(1)"
    CastDef        = $true
    Setup          = "paramSet.Native.Count64 = uint64(33333)"
    Assert         = "Expect(value).To(Equal(uint64(33333)))"
    Equate         = "Equal"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint64"
    #
    BhTests        = $null
  }

  "Float32"  = [PSCustomObject]@{
    TypeName       = "Float32"
    GoType         = "float32"
    FlagName       = "Gradient"
    Short          = "t"
    Def            = "float32(999.123)"
    CastDef        = $true
    Setup          = "paramSet.Native.Gradient = float32(32.1234)"
    Assert         = "Expect(value).To(Equal(float32(32.1234)))"
    Equate         = "Equal"
    GenerateSlice  = $true
    SliceFlagName  = "Temperatures"
    SliceShort     = "T"
    DefSliceVal    = "[]float32{}"
    ExpectSlice    = "[]float32{2.99, 4.99, 6.99, 8.99}"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "float32"
    #
    BhTests        = $null
  }

  "Float64"  = [PSCustomObject]@{
    TypeName       = "Float64"
    GoType         = "float64"
    FlagName       = "Threshold"
    Short          = "t"
    Def            = "float64(999.123)"
    Setup          = "paramSet.Native.Threshold = float64(64.1234)"
    Assert         = "Expect(value).To(Equal(float64(64.1234)))"
    Equate         = "Equal"
    GenerateSlice  = $true
    SliceFlagName  = "Scales"
    SliceShort     = "S"
    DefSliceVal    = "[]float64{}"
    ExpectSlice    = "[]float64{3.99, 5.99, 7.99, 9.99}"
    #
    Comparable     = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "float64"
    #
    BhTests        = $null
  }

  # Bool requires the manual definition of the not set test case
  # There must be a way to implement this automatically as an exception.
  #
  "Bool"     = [PSCustomObject]@{
    TypeName      = "Bool"
    GoType        = "bool"
    FlagName      = "Concise"
    Short         = "c"
    Def           = "false"
    Setup         = "paramSet.Native.Concise = true"
    Assert        = "Expect(value).To(BeTrue())"
    Equate        = "Equal"
    GenerateSlice = $true
    SliceFlagName = "Switches"
    SliceShort    = "S"
    DefSliceVal   = "[]bool{}"
    ExpectSlice   = "[]bool{true, false, true, false}"
  }

  "Duration" = [PSCustomObject]@{
    TypeName   = "Duration"
    GoType     = "time.Duration"
    FlagName   = "Latency"
    Short      = "l"
    Def        = "temp"
    Assign     = "temp, _ := time.ParseDuration(""0ms"")"
    Setup      = "paramSet.Native.Latency, _ = time.ParseDuration(""300ms"")"
    Assert     = @"
    expect, _ := time.ParseDuration("300ms")
    Expect(value).To(BeEquivalentTo(expect))
"@
    Equate     = "BeEquivalentTo"
    # SliceShort    = "?"
    # DefSliceVal   = "[]time.Duration{}"
    # ExpectSlice   = "[]bool{true, false, true, false}"
    # missing Durations slice on WidgetParameterSet
    #
    Comparable = $true
  }

  "IPNet"    = [PSCustomObject]@{
    TypeName = "IPNet"
    GoType   = "net.IPNet"
    FlagName = "IpAddress"
    Short    = "i"
    Def      = "net.IPNet{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPMask([]byte{0, 0, 0, 0}) }"
    Setup    = "paramSet.Native.IpAddress = net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPMask([]byte{255, 255, 255, 0}) }"
    Assert   = "Expect(value).To(BeEquivalentTo(net.IPNet{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPMask([]byte{255, 255, 255, 0}) }))"
    Equate   = "BeEquivalentTo"
  }

  "IPMask"   = [PSCustomObject]@{
    TypeName = "IPMask"
    GoType   = "net.IPMask"
    FlagName = "IpMask"
    Short    = "m"
    Def      = "net.IPMask([]byte{0, 0, 0, 0})"
    Setup    = "paramSet.Native.IpMask = net.IPMask([]byte{255, 255, 255, 0})"
    Assert   = "Expect(value).To(BeEquivalentTo(net.IPMask([]byte{255, 255, 255, 0})))"
    Equate   = "BeEquivalentTo"
  }
}

[array]$operators = @(
  [PSCustomObject]@{
    Name               = "Within"
    Documentation      = "fails validation if the option value does not lie within 'lo' and 'hi' (inclusive)"
    Dual               = $true
    Args               = "lo, hi"
    Condition          = "value >= lo && value <= hi"
    ErrorMessage       = "out of range"
    ArgsPlaceholder    = "[%v]..[%v]"
    ErrorArgs          = "lo, hi"
    Comment            = "option value must be within the range"
    #
    Negate             = $true
    NegateErrorMessage = "is within range"
    NegateComment      = "option value must not be within the range"
  }

  , [PSCustomObject]@{
    Name               = "Contains"
    Documentation      = "fails validation if the option value is not a member of the 'collection' slice"
    Contains           = $true
    MethodTemplate     = "{{OpName}}{{TypeName}}"
    Args               = "collection"
    Condition          = "lo.IndexOf(collection, value) >= 0"
    ErrorMessage       = "not a member of"
    ArgsPlaceholder    = "[%v]"
    ErrorArgs          = "collection"
    Comment            = "option value must be a member of collection"
    #
    Negate             = $true
    NegateErrorMessage = "is a member of"
    NegateComment      = "option value must not be a member of collection"
  }

  , [PSCustomObject]@{
    Name                 = "IsMatch"
    Documentation        = "fails validation if the option value does not match the regular expression denoted by 'pattern'"
    AppliesOnlyTo        = "String"
    Args                 = "pattern"
    Condition            = "regexp.MustCompile(pattern).Match([]byte(value))"
    ErrorMessage         = "does not match"
    ArgsPlaceholder      = "[%v]"
    ErrorArgs            = "pattern"
    Comment              = "option value must match regex pattern"
    #
    Negate               = $true
    NegateMethodTemplate = "{{TypeName}}IsNotMatch"
    NegateErrorMessage   = "matches"
    NegateComment        = "option value must not match regex pattern"
  }

  , [PSCustomObject]@{
    Name            = "GreaterThan"
    Documentation   = "fails validation if the option value is not comparably greater than 'threshold'"
    Args            = "threshold"
    Condition       = "value > threshold"
    ErrorMessage    = "not greater than"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    Comment         = "option value must be greater than threshold"
  }

  , [PSCustomObject]@{
    Name            = "AtLeast"
    Documentation   = "fails validation if the option value is not comparably greater than or equal to 'threshold'"
    Args            = "threshold"
    Condition       = "value >= threshold"
    ErrorMessage    = "not at least"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    Comment         = "option value must be greater than or equal to threshold"
  }

  , [PSCustomObject]@{
    Name            = "LessThan"
    Documentation   = "fails validation if the option value is not comparably less than 'threshold'"
    Args            = "threshold"
    Condition       = "value < threshold"
    ErrorMessage    = "not less than"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    Comment         = "option value must be less than threshold"
  }

  , [PSCustomObject]@{
    Name            = "AtMost"
    Documentation   = "fails validation if the option value is not comparably less than or equal to 'threshold'"
    Args            = "threshold"
    Condition       = "value <= threshold"
    ErrorMessage    = "not at most"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    Comment         = "option value must be less than or equal to threshold"
  }
)

Write-Host "🤖 Build-Validators(gen-ov) ✨ => option-validator-auto.go"
Write-Host "🤖 Build-ParamSet(gen-ps) ✨ => param-set-auto.go"
Write-Host "🤖 Build-TestEntry(gen-ov-t) 🧪 => option-validator-auto_test.go"
Write-Host "🤖 Build-Predefined(gen-help) 🎁 => paramset-binder-helpers-auto.go"
Write-Host "🤖 Build-BinderHelperTests(gen-help-t) 🧪 => paramset-binder-helpers-auto_test.go"

function Build-Validators {
  # (option-validator-auto.go)
  #
  [Alias("gen-ov")]
  param()

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      $spec = $types[$_]

      $validatorType = $spec.TypeName
      $validatorStruct = "$($validatorType)OptionValidator"
      $validatorFn = $("$($spec.TypeName)ValidatorFn")

      # generate
      # - type XXXXValidatorFn
      # - type XXXXOptionValidator
      # - func (validator XXXXOptionValidator) Validate()
      #
      @"
// $($validatorFn) defines the validator function for $($spec.GoType) type.
//
type $($validatorFn) func(value $($spec.GoType)) error

// $($validatorStruct) defines the struct that wraps the client defined validator function
// $($validatorFn) for $($spec.GoType) type. This is the instance that is returned by
// validated binder function BindValidated$($spec.TypeName). If not using the ParamSet
// (which is recommended), the client should add this instance to a self managed
// ValidatorContainer.
//
type $($validatorStruct) GenericOptionValidatorWrapper[$($spec.GoType)]

// Validate invokes the client defined validator function for $($spec.GoType) type.
//
func (validator $($validatorStruct)) Validate() error {
	return validator.Fn(*validator.Value)
}

"@
      if ($spec.GenerateSlice) {
        # generate
        # - type XXXXSliceValidatorFn
        # - type XXXXSliceOptionValidator
        # - func (validator XXXXSliceOptionValidator) Validate()
        #
        $sliceTypeName = "$($spec.TypeName)Slice"
        $typeName = "$($sliceTypeName)OptionValidator"
        $sliceType = "[]$($spec.GoType)"
        $sliceValidatorStruct = "$($sliceTypeName)OptionValidator"
        $sliceValidatorFn = $("$($spec.TypeName)SliceValidatorFn")
        @"
// $($typeName) defines the validator function for $($sliceTypeName) type.
//
type $($sliceValidatorFn) func(value $($sliceType)) error

// $($sliceValidatorStruct) wraps the client defined validator function for type $($sliceType).
//
type $($sliceValidatorStruct) GenericOptionValidatorWrapper[$($sliceType)]

// Validate invokes the client defined validator function for $($sliceType) type.
//
func (validator $($sliceValidatorStruct)) Validate() error {
	return validator.Fn(*validator.Value)
}

"@
      }
    })
  $content | Set-Clipboard

  Write-Host "🎯 Paste into ---> 'option-validator-auto.go'"
}

function Build-ParamSet {
  # (param-set-auto.go)
  #
  [Alias("gen-ps")]
  param()

  # each operation defined independently
  #
  $content = ($types.Keys | Sort-Object | ForEach-Object {
      $spec = $types[$_]

      $validatorFn = $("$($spec.TypeName)ValidatorFn")
      $actualTypeName = [string]::IsNullOrEmpty($spec.UnderlyingTypeName) ? $spec.TypeName : $spec.UnderlyingTypeName

      # generate BindXXXX
      #
      @"
// Bind$($spec.TypeName) binds $($spec.GoType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//
func (params *ParamSet[N]) Bind$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType)) *ParamSet[N] {
  if info.Short == "" {
    params.FlagSet.$($actualTypeName)Var(to, info.FlagName(), info.Default.($($spec.GoType)), info.Usage)
  } else {
    params.FlagSet.$($actualTypeName)VarP(to, info.FlagName(), info.Short, info.Default.($($spec.GoType)), info.Usage)
  }

  return params
}

"@

      # generate BindValidatedXXXX
      #
      @"
// BindValidated$($spec.TypeName) binds $($spec.GoType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of $($spec.GoType) type.
//
func (params *ParamSet[N]) BindValidated$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType), validator $($validatorFn)) OptionValidator {

  params.Bind$($spec.TypeName)(info, to)
  wrapper := $($actualTypeName)OptionValidator{
    Fn:    validator,
    Value: to,
  }
  params.validators.Add(info.FlagName(), wrapper)
  return wrapper
}

"@

      # generate BindXXXXSlice
      #
      if ($spec.GenerateSlice) {
        $sliceTypeName = "$($spec.TypeName)Slice"
        $sliceType = "[]$($spec.GoType)"
        $defaultSlice = $("[]$($spec.GoType)")

        @"
// Bind$($sliceTypeName) binds $($sliceType) slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
//
func (params *ParamSet[N]) Bind$($sliceTypeName)(info *FlagInfo, to *$($sliceType)) *ParamSet[N] {
  if info.Short == "" {
    params.FlagSet.$($sliceTypeName)Var(to, info.FlagName(), info.Default.($($sliceType)), info.Usage)
  } else {
    params.FlagSet.$($sliceTypeName)VarP(to, info.FlagName(), info.Short, info.Default.($($defaultSlice)), info.Usage)
  }

  return params
}

"@
        # generate BindValidatedXXXXSlice
        #
        $sliceTypeName = "$($spec.TypeName)Slice"
        $sliceType = "[]$($spec.GoType)"
        $defaultSlice = $("[]$($spec.GoType)")
        $sliceValidatorFn = $("$($sliceTypeName)ValidatorFn")

        @"
// BindValidated$($sliceTypeName) binds $($sliceType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.  Client can provide a
// function to validate option values of $sliceType type.
//
func (params *ParamSet[N]) BindValidated$($sliceTypeName)(info *FlagInfo, to *$($sliceType), validator $($sliceValidatorFn)) OptionValidator {

  params.Bind$($sliceTypeName)(info, to)
  wrapper := $($sliceTypeName)OptionValidator{
    Fn:    validator,
    Value: to,
  }
  params.validators.Add(info.FlagName(), wrapper)
  return wrapper
}

"@
      }
    })
  $content | Set-Clipboard

  Write-Host "🎯 Paste into ---> 'param-set-auto.go'"
}



function Build-TestEntry {
  # (option-validator-auto_test.go)
  #
  [Alias("gen-ov-t")]
  param()
  $content = ($types.Keys | Sort-Object | ForEach-Object {
      $spec = $types[$_]

      $lowerFlagName = $spec.FlagName.ToLower()
      $default = $spec.QuoteExpect ? $('"' + $spec.Def + '"') : $spec.Def
      $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo

      @"
Entry(nil, OvEntry{
  Message: "$($spec.GoType) type (auto)",
  Setup: func() {
    $($spec.Setup)
  },
  Validator: func() adapters.OptionValidator {
    $($spec.Assign)
    return paramSet.BindValidated$($spec.TypeName)(
      adapters.NewFlagInfo("$($lowerFlagName)", "$($spec.Short)", $default),
      $bindTo,
      func(value $($spec.GoType)) error {
        $($spec.Assert)
        return nil
      },
    )
  },
}),

"@
      # generate XXXXSlice
      #
      if ($spec.GenerateSlice) {
        $sliceTypeName = "$($spec.TypeName)Slice"
        $sliceType = "[]$($spec.GoType)"
        @"
Entry(nil, OvEntry{
  Message: "$($sliceType) type (auto)",
  Setup: func() {
    paramSet.Native.$($spec.SliceFlagName) = $($spec.ExpectSlice)
  },
  Validator: func() adapters.OptionValidator {
    return paramSet.BindValidated$($sliceTypeName)(
      adapters.NewFlagInfo("$($spec.SliceFlagName)", "$($spec.SliceShort)", $($spec.DefSliceVal)),
      &paramSet.Native.$($spec.SliceFlagName),
      func(value $($sliceType)) error {
        Expect(value).To($($spec.Equate)($($spec.ExpectSlice)))
        return nil
      },
    )
  },
}),

"@
      }
    })
  $content | Set-Clipboard

  Write-Host "🎯 Paste into ---> 'option-validator-auto_test.go'"
}

function Build-Predefined {
  # (paramset-binder-helpers-auto.go)
  #
  [Alias("gen-help")]
  param()

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      $spec = $types[$_]

      if ($spec.Comparable) {
        foreach ($op in $operators) {
          if (-not([string]::IsNullOrEmpty($op.AppliesOnlyTo)) -and ($op.AppliesOnlyTo -ne $spec.TypeName)) {
            continue
          }
          # assuming all args have the same type
          #
          $argumentsStmt = if ($op.Contains) {
            "$("$($op.Args) []$($spec.GoType)")"
          }
          else {
            "$("$($op.Args) $($spec.GoType)")"
          }

          # $methodSubStmt = Get-MethodSubStmtFromOperator -Spec $spec -Operator $op
          $methodSubStmt = if (-not([string]::IsNullOrEmpty($op.MethodTemplate))) {
            $op.MethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
          }
          else {
            # Default is TypeNameOpName, eg: StringGreaterThan
            #
            $("$($spec.TypeName)$($op.Name)")
          }

          $errorMessage = "$($op.ErrorMessage): $($op.ArgsPlaceholder)"
          # generate BuildValidatedXXXXOp/BuildValidatedOpXXXX
          #
          @"
// BindValidated$($methodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method $($op.Documentation).
// 
func (params *ParamSet[N]) BindValidated$($methodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {

  params.Bind$($spec.TypeName)(info, to)
  wrapper := GenericOptionValidatorWrapper[$($spec.GoType)]{
    Fn: func(value $($spec.GoType)) error {
      if $($op.Condition) {
        return nil
      }
      return fmt.Errorf("(%v): option validation failed, '%v', $($errorMessage)",
        info.FlagName(), value, $($op.ErrorArgs),
      )
    },
    Value: to,
  }
  params.validators.Add(info.FlagName(), wrapper)
  return wrapper
}

"@

          if ($op.Negate) {
            $negateErrorMessage = "$($op.NegateErrorMessage): $($op.ArgsPlaceholder)"

            $methodSubStmt = if (-not([string]::IsNullOrEmpty($op.MethodTemplate))) {
              $op.MethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
            }
            else {
              # Default is TypeNameOpName, eg: StringGreaterThan
              #
              $("$($spec.TypeName)$($op.Name)")
            }

            $notMethodSubStmt = if (-not([string]::IsNullOrEmpty($op.NegateMethodTemplate))) {
              $op.NegateMethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
            }
            else {
              # Default is TypeNameNotOpName, eg: StringNotGreaterThan
              #
              $("$($spec.TypeName)Not$($op.Name)")
            }
            $negatedCondition = $("!($($op.Condition))")
  
            # generate NOT method
            #
            @"
// BindValidated$($notMethodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidated$($methodSubStmt)'.
//
func (params *ParamSet[N]) BindValidated$($notMethodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {

  params.Bind$($spec.TypeName)(info, to)
  wrapper := GenericOptionValidatorWrapper[$($spec.GoType)]{
    Fn: func(value $($spec.GoType)) error {
      if $($negatedCondition) {
        return nil
      }
      return fmt.Errorf("(%v): option validation failed, '%v', $($negateErrorMessage)",
        info.FlagName(), value, $($op.ErrorArgs),
      )
    },
    Value: to,
  }
  params.validators.Add(info.FlagName(), wrapper)
  return wrapper
}
            
"@
          }
        }
      }
    })
  $content | Set-Clipboard

  Write-Host "🎯 Paste into ---> 'paramset-binder-helpers-auto.go'"
}

function Build-BinderHelperTests {
  # (paramset-binder-helpers-auto_test.go)
  #
  [Alias("gen-help-t")]
  param()

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      $spec = $types[$_]

      if ($spec.Comparable) {
        $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo
        $default = $spec.QuoteExpect ? $('"' + $spec.Def + '"') : $spec.Def
        [int]$ValueIndex = 0
        [int]$ExpectNilIndex = 1

        foreach ($current in $operators) {
          if (($null -eq $spec.BhTests)) {
            if (-not([string]::IsNullOrEmpty($spec.BhParent)) -and ($types.ContainsKey($spec.BhParent))) {
              $spec.BhTests = $types[$spec.BhParent].BhTests

              if ($null -eq $spec.BhTests) {
                Write-Host "===> 🔥🔥🔥 BhParent '$($spec.BhParent)' of '$($spec.TypeName)' does not have a valid BhTests, skipping ..."
                continue
              }
            }
            else {
              continue
            }
          }
          $cast = $spec.CastLiteralsAs

          if (-not($spec.BhTests.ContainsKey($current.Name))) {
            continue
          }

          $op = $current
          $testOp = $($spec.BhTests[$op.Name])

          # $methodSubStmt = Get-MethodSubStmtFromOperator -Spec $spec -Operator $op
          $methodSubStmt = if (-not([string]::IsNullOrEmpty($op.MethodTemplate))) {
            $op.MethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
          }
          else {
            # Default is TypeNameOpName, eg: StringGreaterThan
            #
            $("$($spec.TypeName)$($op.Name)")
          }


          if ($op.Dual) {
            # dual means we need hi and lo
            #
            if ([string]::IsNullOrEmpty($cast)) {
              $belowEntryArgs = $("$($testOp.Entry.Below[$ValueIndex]), $($testOp.Entry.Below[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $equalLoEntryArgs = $("$($testOp.Entry.EqualLo[$ValueIndex]), $($testOp.Entry.EqualLo[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $insideEntryArgs = $("$($testOp.Entry.Inside[$ValueIndex]), $($testOp.Entry.Inside[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $equalHiEntryArgs = $("$($testOp.Entry.EqualHi[$ValueIndex]), $($testOp.Entry.EqualHi[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $aboveEntryArgs = $("$($testOp.Entry.Above[$ValueIndex]), $($testOp.Entry.Above[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")  
            }
            else {
              $belowEntryArgs = $("$($cast)($($testOp.Entry.Below[$ValueIndex])), $($testOp.Entry.Below[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $equalLoEntryArgs = $("$($cast)($($testOp.Entry.EqualLo[$ValueIndex])), $($testOp.Entry.EqualLo[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $insideEntryArgs = $("$($cast)($($testOp.Entry.Inside[$ValueIndex])), $($testOp.Entry.Inside[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $equalHiEntryArgs = $("$($cast)($($testOp.Entry.EqualHi[$ValueIndex])), $($testOp.Entry.EqualHi[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $aboveEntryArgs = $("$($cast)($($testOp.Entry.Above[$ValueIndex])), $($testOp.Entry.Above[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
            }

            @"
DescribeTable("BindValidated$($methodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, low, high $($spec.GoType)) {
    validator := paramSet.BindValidated$($methodSubStmt)(
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), low, high,
    )
    paramSet.Native.$($spec.FlagName) = value

    if expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }

  },
  func(given, should string, value $($spec.GoType), expectNil bool, low, high $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value is below range", "return error", $($belowEntryArgs)),
  Entry(nil, "value is equal to low end of range", "return error", $($equalLoEntryArgs)),
  Entry(nil, "value is inside range", "return error", $($insideEntryArgs)),
  Entry(nil, "value is equal to high end of range", "return error", $($equalHiEntryArgs)),
  Entry(nil, "value is above range", "NOT return error", $($aboveEntryArgs)),
)

"@
          }
          elseif ($op.Contains) {
            # contains requires a sequence
            if ([string]::IsNullOrEmpty($cast)) {
              $literalCollection = $($testOp.First).Replace("{{SLICE-TYPE}}", $($spec.GoType))
              $doesContainArgs = $("$($testOp.Entry.DoesContain[$ValueIndex]), $($testOp.Entry.DoesContain[$ExpectNilIndex]), $literalCollection, $($testOp.Second)")
              $doesNotContainArgs = $("$($testOp.Entry.DoesNotContain[$ValueIndex]), $($testOp.Entry.DoesNotContain[$ExpectNilIndex]), $literalCollection, $($testOp.Second)")
            }
            else {
              $literalCollection = $($testOp.First).Replace("{{SLICE-TYPE}}", $($cast))
              $doesContainArgs = $("$($cast)($($testOp.Entry.DoesContain[$ValueIndex])), $($testOp.Entry.DoesContain[$ExpectNilIndex]), $literalCollection, $($cast)($($testOp.Second))")
              $doesNotContainArgs = $("$($cast)($($testOp.Entry.DoesNotContain[$ValueIndex])), $($testOp.Entry.DoesNotContain[$ExpectNilIndex]), $literalCollection, $($cast)($($testOp.Second))")
            }

            @"
DescribeTable("BindValidated$($methodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, collection []$($spec.GoType), dummy $($spec.GoType)) {
    validator := paramSet.BindValidated$($methodSubStmt)(
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), collection,
    )
    paramSet.Native.$($spec.FlagName) = value

    if expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }

  },
  func(given, should string, value $($spec.GoType), expectNil bool, collection []$($spec.GoType), dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "collection contains member", "return error", $($doesContainArgs)),
  Entry(nil, "collection does not contain member", "return error", $($doesNotContainArgs)),
)

"@
          }
          elseif (($op.AppliesOnlyTo -eq "String") -and ($op.Name -eq "IsMatch")) {
            # generate DescribeTable tests for BindValidatedStringIsMatch
            #
            if ([string]::IsNullOrEmpty($cast)) {
              $doesMatchArgs = $("$($testOp.Entry.DoesMatch[$ValueIndex]), $($testOp.Entry.DoesMatch[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $doesNotMatchArgs = $("$($testOp.Entry.DoesNotMatch[$ValueIndex]), $($testOp.Entry.DoesNotMatch[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
            }
            else {
              $doesMatchArgs = $("$($cast)($($testOp.Entry.DoesMatch[$ValueIndex])), $($testOp.Entry.DoesMatch[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $doesNotMatchArgs = $("$($cast)($($testOp.Entry.DoesNotMatch[$ValueIndex])), $($testOp.Entry.DoesNotMatch[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
            }

            @"
DescribeTable("BindValidated$($methodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy string) {
    validator := paramSet.BindValidated$($methodSubStmt)(
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), pattern,
    )
    paramSet.Native.$($spec.FlagName) = value

    if expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }

  },
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value matches pattern", "return error", $($doesMatchArgs)),
  Entry(nil, "value does not match pattern", "return error", $($doesNotMatchArgs)),
)

"@
            if ($op.Negate) {
              # generate Not Test
              $notMethodSubStmt = if (-not([string]::IsNullOrEmpty($op.NegateMethodTemplate))) {
                $op.NegateMethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
              }
              else {
                # Default is TypeNameNotOpName, eg: StringNotGreaterThan
                #
                $("$($spec.TypeName)Not$($op.Name)")
              }
              # For the NOT scenario test cases, we should still be able to use the exact same test data,
              # but just tweak the test code to reverse the logic. Eg, so we simply negate the expectNil
              # and everything else stays the same.
              #
              @"
DescribeTable("BindValidated$($notMethodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy string) {
    validator := paramSet.BindValidated$($notMethodSubStmt)(
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), pattern,
    )
    paramSet.Native.$($spec.FlagName) = value

    if !expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }

  },
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value matches pattern", "return error", $($doesMatchArgs)),
  Entry(nil, "value does not match pattern", "return error", $($doesNotMatchArgs)),
)

"@  
            }
          }
          else {
            if ([string]::IsNullOrEmpty($cast)) {
              $belowEntryArgs = $("$($testOp.Entry.Below[$ValueIndex]), $($testOp.Entry.Below[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $equalEntryArgs = $("$($testOp.Entry.Equal[$ValueIndex]), $($testOp.Entry.Equal[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
              $aboveEntryArgs = $("$($testOp.Entry.Above[$ValueIndex]), $($testOp.Entry.Above[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
            }
            else {
              $belowEntryArgs = $("$($cast)($($testOp.Entry.Below[$ValueIndex])), $($testOp.Entry.Below[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
              $equalEntryArgs = $("$($cast)($($testOp.Entry.Equal[$ValueIndex])), $($testOp.Entry.Equal[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")                
              $aboveEntryArgs = $("$($cast)($($testOp.Entry.Above[$ValueIndex])), $($testOp.Entry.Above[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
            }

            @"
DescribeTable("BindValidated$($methodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, threshold, dummy $($spec.GoType)) {
    validator := paramSet.BindValidated$($methodSubStmt)(
      adapters.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), threshold,
    )
    paramSet.Native.$($spec.FlagName) = value

    if expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }

  },
  func(given, should string, value $($spec.GoType), expectNil bool, threshold, dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍒 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value is below threshold", "return error", $($belowEntryArgs)),
  Entry(nil, "value is equal to threshold", "return error", $($equalEntryArgs)),
  Entry(nil, "value is above threshold", "NOT return error", $($aboveEntryArgs)),
)

"@
          }
        }
      }
    })
  $content | Set-Clipboard

  Write-Host "🎯 Paste into ---> 'paramset-binder-helpers-auto_test.go'"
}

# === Utilities

# WTF: Cannot process argument transformation on parameter 'Spec'.
# $methodSubStmt = Get-MethodSubStmtFromOperator -Spec $spec -Operator $op
#
function Get-MethodSubStmtFromOperator {
  [CmdletBinding()]
  [OutputType([string])]
  param(
    [Parameter]
    [PSCustomObject]$Spec,

    [Parameter]
    [PSCustomObject]$Operator
  )

  $result = if (-not([string]::IsNullOrEmpty($Operator.MethodTemplate))) {
    $Operator.MethodTemplate.Replace("{{OpName}}", $Operator.Name).Replace("{{TypeName}}", $Spec.TypeName)
  }
  else {
    # Default is TypeNameOpName, eg: StringGreaterThan
    #
    $("$($Spec.TypeName)$($Operator.Name)")
  }
  return $result
}
