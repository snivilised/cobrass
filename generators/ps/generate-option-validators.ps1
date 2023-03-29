
$types = @{
  "Enum"     = [PSCustomObject]@{
    TypeName           = "Enum"
    GoType             = "string"
    DisplayType        = "enum"
    UnderlyingTypeName = "String"
    FlagName           = "Format"
    Short              = "f"
    Def                = "xml"
    Assign             = "outputFormatEnum := outputFormatEnumInfo.NewValue()"
    Setup              = "paramSet.Native.Format = XMLFormatEn"
    BindTo             = "&outputFormatEnum.Source"
    Assert             = "Expect(value).To(Equal(""xml""))"
    QuoteExpect        = $true
    Equate             = "Equal"
    Validatable        = $true
    ForeignValidatorFn = $true # validation does occur but using a foreign type, ie string
    GenerateSlice      = $false # EnumSlice is implemented as StringSlice
    SliceFlagName      = "Formats"
    SliceShort         = "F"
    DefSliceVal        = "[]string{}"
    ExpectSlice        = "[]string{""xml"", ""json"", ""text""}"
    SliceValue         = "xml,json,text"
    OptionValue        = "json"
    #
    PsTcEntry          = [PSCustomObject]@{
      AssertFn = "func() { Expect(outputFormatEnum.Source).To(Equal(""json"")) }"
    }
    #
    BindDoc            = @"

// Note that normally the client would bind to a member of the native parameter
// set. However, since there is a discrepancy between the type of the native int
// based pseudo enum member and the equivalent acceptable string value typed by
// the user on the command line (idiomatically stored on the enum info), the
// client needs to extract the enum value from the enum info, something like this:
//
// paramSet.Native.Format = OutputFormatEnumInfo.Value()
//
// The best place to put this would be inside the PreRun/PreRunE function, assuming the
// param set and the enum info are both in scope. Actually, every int based enum
// flag, would need to have this assignment performed.
"@

    BindValidatedDoc   = @"

// Custom enum types created via the generic 'EnumInfo'/'EnumValue' come with a 'IsValid' method.
// The client can utilise this method inside a custom function passed into 'BindValidatedEnum'.
// The implementation would simply call this method, either on the EnumInfo or the EnumValue.
// Please see the readme for more details.
"@

    #
    # Currently 'Comparable' for enum disabled because enum comparison would be
    # performed in the string domain but it might make more sense to the use if
    # it was in the int domain. We don't want to commit to publish this particular
    # api, if it's not clear how this would be implemented, so that it makes sense.

    Containable        = $true
    #
    BhTests            = @{
      "Contains" = @{
        First  = "[]string{""json"", ""text"", ""xml""}"
        Second = """null"""
        Assign = "outputFormatEnum.Source = value"
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          DoesContain    = @("""xml""", "true")
          DoesNotContain = @("""scr""", "false")
        }
      }
    }
  }

  "String"   = [PSCustomObject]@{
    TypeName      = "String"
    GoType        = "string"
    FlagName      = "Pattern"
    Short         = "p"
    Def           = "default-pattern"
    Setup         = "paramSet.Native.Pattern = ""{{OPTION-VALUE}}"""
    Assert        = "Expect(value).To(Equal(""{{OPTION-VALUE}}""))"
    QuoteExpect   = $true
    Equate        = "Equal"
    Validatable   = $true
    GenerateSlice = $true
    SliceFlagName = "Directories"
    SliceShort    = "C"
    DefSliceVal   = "[]string{}"
    ExpectSlice   = "[]string{""alpha"", ""beta"", ""delta""}"
    SliceValue    = "alpha,beta,delta"
    OptionValue   = "*music.infex*"
    #
    PsTcEntry     = [PSCustomObject]@{
      # AssertFn function is optional, but is the first item checked for
      # Next we assume Expect(value).To(something),
      # where something can be
      # - Equal[type](value) => type is optional template variable
      # - BeTrue()
      # or any other matcher
    }
    #
    Comparable    = $true
    Containable   = $true
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
    Setup         = "paramSet.Native.Offset = {{OPTION-VALUE}}"
    Assert        = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate        = "Equal"
    Validatable   = $true
    GenerateSlice = $true
    SliceFlagName = "Offsets"
    SliceShort    = "D"
    DefSliceVal   = "[]int{}"
    ExpectSlice   = "[]int{2, 4, 6, 8}"
    SliceValue    = "2,4,6,8"
    OptionValue   = "-9"
    #
    PsTcEntry     = [PSCustomObject]@{}
    #
    Comparable    = $true
    Containable   = $true
    #
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
    Setup          = "paramSet.Native.Offset8 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "-99"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
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
    Setup          = "paramSet.Native.Offset16 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "-999"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
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
    Setup          = "paramSet.Native.Offset32 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "-9999"
    #
    PsTcEntry      = [PSCustomObject]@{}
    GenerateSlice  = $true

    SliceFlagName  = "Offsets32"
    SliceShort     = "O"
    DefSliceVal    = "[]int32{}"
    ExpectSlice    = "[]int32{2, 4, 6, 8}"
    SliceValue     = "2,4,6,8"
    #
    Comparable     = $true
    Containable    = $true
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
    Setup          = "paramSet.Native.Offset64 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "-99999"
    #
    PsTcEntry      = [PSCustomObject]@{}
    GenerateSlice  = $true

    SliceFlagName  = "Offsets64"
    SliceShort     = "O"
    DefSliceVal    = "[]int64{}"
    ExpectSlice    = "[]int64{2, 4, 6, 8}"
    SliceValue     = "2,4,6,8"
    #
    Comparable     = $true
    Containable    = $true
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
    Setup          = "paramSet.Native.Count = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "99999"
    #
    PsTcEntry      = [PSCustomObject]@{}
    GenerateSlice  = $true
    SliceFlagName  = "Counts"
    SliceShort     = "P"
    DefSliceVal    = "[]uint{}"
    ExpectSlice    = "[]uint{2, 4, 6, 8}"
    SliceValue     = "2,4,6,8"
    #
    Comparable     = $true
    Containable    = $true
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
    Def            = "uint8(0)"
    Setup          = "paramSet.Native.Count8 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "33"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
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
    Def            = "uint16(0)"
    Setup          = "paramSet.Native.Count16 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "333"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
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
    Def            = "uint32(0)"
    Setup          = "paramSet.Native.Count32 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "3333"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint32"
    #
    BhTests        = $null
  }

  "Uint64"   = [PSCustomObject]@{
    TypeName       = "Uint64"
    GoType         = "uint64"
    FlagName       = "Count64"
    Short          = "c"
    Def            = "uint64(0)"
    Setup          = "paramSet.Native.Count64 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "33333"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    Comparable     = $true
    Containable    = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "uint64"
    #
    BhTests        = $null
  }

  "Float32"  = [PSCustomObject]@{
    TypeName       = "Float32"
    GoType         = "float32"
    FlagName       = "Gradientf32"
    Short          = "t"
    Def            = "float32(0)"
    Setup          = "paramSet.Native.Gradientf32 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "32.0"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    GenerateSlice  = $true
    SliceFlagName  = "Gradientsf32"
    SliceShort     = "G"
    DefSliceVal    = "[]float32{}"
    ExpectSlice    = "[]float32{3.0, 5.0, 7.0, 9.0}"
    SliceValue     = "3.0,5.0,7.0,9.0"
    #
    Comparable     = $true
    Containable    = $true
    #
    BhParent       = "Int"
    CastLiteralsAs = "float32"
    #
    BhTests        = $null
  }

  "Float64"  = [PSCustomObject]@{
    TypeName       = "Float64"
    GoType         = "float64"
    FlagName       = "Gradientf64"
    Short          = "t"
    Def            = "float64(0)"
    Setup          = "paramSet.Native.Gradientf64 = {{OPTION-VALUE}}"
    Assert         = "Expect(value).To(Equal({{OPTION-VALUE}}))"
    Equate         = "Equal"
    Validatable    = $true
    OptionValue    = "64.1234"
    #
    PsTcEntry      = [PSCustomObject]@{}
    #
    GenerateSlice  = $true
    SliceFlagName  = "Gradientsf64"
    SliceShort     = "G"
    DefSliceVal    = "[]float64{}"
    ExpectSlice    = "[]float64{4.0, 6.0, 8.0, 10.0}"
    SliceValue     = "4.0,6.0,8.0,10.0"
    #
    Comparable     = $true
    Containable    = $true
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
    Setup         = "paramSet.Native.Concise = {{OPTION-VALUE}}"
    Assert        = "Expect(value).To(BeTrue())"
    Equate        = "Equal"
    # bool is not Validatable, because there's not much to validate,
    # can only be true or false
    Validatable   = $false
    GenerateSlice = $true
    SliceFlagName = "Switches"
    SliceShort    = "S"
    DefSliceVal   = "[]bool{}"
    ExpectSlice   = "[]bool{true, false, true, false}"
    SliceValue    = "true,false,true,false"
    IsOptionLess  = $true
    OptionValue   = "true"
    #
    PsTcEntry     = [PSCustomObject]@{}
  }

  "Duration" = [PSCustomObject]@{
    TypeName         = "Duration"
    GoType           = "time.Duration"
    FlagName         = "Latency"
    Short            = "l"
    Def              = "duration(""0ms"")"
    Setup            = "paramSet.Native.Latency = {{OPTION-VALUE}}"
    Assert           = @"
    expect := {{OPTION-VALUE}}
    Expect(value).To(BeEquivalentTo(expect))
"@
    Equate           = "BeEquivalentTo"
    Validatable      = $true
    OptionValue      = "300ms"
    QuoteOptionValue = $true
    #
    PsTcEntry        = [PSCustomObject]@{}
    #
    GenerateSlice    = $true
    SliceFlagName    = "Latencies"
    SliceShort       = "L"
    DefSliceVal      = "[]time.Duration{}"
    ExpectSlice      = "[]time.Duration{duration(""1s""), duration(""2s""), duration(""3s"")}"
    SliceValue       = "1s,2s,3s"
    Comparable       = $true
    #
    # 'duration' is a function defined in the test suite, that is syntactically the
    # same as a type cast.
    # 
    CastLiteralsAs   = "duration"
    BhTests          = @{
      "Within"      = @{
        First  = """3s"""
        Second = """5s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below   = @("""2s""", "false")
          EqualLo = @("""3s""", "true")
          Inside  = @("""4s""", "true")
          EqualHi = @("""5s""", "true")
          Above   = @("""6s""", "false")
        }
      }

      "Contains"    = @{
        # Any test data that contains a type spec, needs to be defined by a template
        #
        First  = "[]time.Duration{duration(""1s""), duration(""2s""), duration(""3s"")}"
        Second = """0s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          DoesContain    = @("""1s""", "true")
          DoesNotContain = @("""99s""", "false")
        }
      }

      "GreaterThan" = @{
        First  = """3s"""
        Second = """0s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""2s""", "false")
          Equal = @("""3s""", "false")
          Above = @("""4s""", "true")
        }
      }

      "AtLeast"     = @{
        First  = """3s"""
        Second = """0s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""2s""", "false")
          Equal = @("""3s""", "true")
          Above = @("""4s""", "true")
        }
      }

      "LessThan"    = @{
        First  = """3s"""
        Second = """0s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""2s""", "true")
          Equal = @("""3s""", "false")
          Above = @("""4s""", "false")
        }
      }

      "AtMost"      = @{
        First  = """3s"""
        Second = """0s"""
        Entry  = [PSCustomObject]@{
          # array: @(Value, ExpectNil)
          Below = @("""2s""", "true")
          Equal = @("""3s""", "true")
          Above = @("""4s""", "false")
        }
      }
    }
  }

  "IPNet"    = [PSCustomObject]@{
    TypeName         = "IPNet"
    GoType           = "net.IPNet"
    FlagName         = "IPAddress"
    Short            = "i"
    Def              = "ipnet(""default"")"
    Setup            = "paramSet.Native.IPAddress = {{OPTION-VALUE}}"
    Assert           = "Expect(value).To(BeEquivalentTo({{OPTION-VALUE}}))"
    Equate           = "BeEquivalentTo"
    Validatable      = $true
    OptionValue      = "ipnet(""orion.net"")"
    CommandLineValue = "172.16.0.0"
    #
    PsTcEntry        = [PSCustomObject]@{
      AssertFn = "func() { Expect(paramSet.Native.IPAddress).ToNot(BeNil()) }"
    }
  }

  "IPMask"   = [PSCustomObject]@{
    TypeName         = "IPMask"
    GoType           = "net.IPMask"
    FlagName         = "IPMask"
    Short            = "m"
    Def              = "ipmask(""default"")"
    Setup            = "paramSet.Native.IPMask = {{OPTION-VALUE}}"
    Assert           = "Expect(value).To(BeEquivalentTo({{OPTION-VALUE}}))"
    Equate           = "BeEquivalentTo"
    Validatable      = $true
    OptionValue      = "ipmask(""orion.net"")"
    CommandLineValue = "255.255.255.0"
    #
    PsTcEntry        = [PSCustomObject]@{}
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
    ErrorTempl         = "New{{Not}}WithinOptValidationError"
    Comment            = "option value must be within the range"
    #
    Negate             = $true
    ExcludeTypes       = @("Bool", "Enum", "IPMask", "IPNet")
    NegateErrorMessage = "is within range"
    NegateComment      = "option value must not be within the range"
  }

  , [PSCustomObject]@{
    Name                 = "Contains"
    Documentation        = "fails validation if the option value is not a member of the 'collection' slice"
    Container            = $true
    MethodTemplate       = "{{OpName}}{{TypeName}}"
    Args                 = "collection"
    Condition            = "lo.IndexOf(collection, value) >= 0"
    ErrorMessage         = "not a member of"
    ArgsPlaceholder      = "[%v]"
    ErrorArgs            = "collection"
    ErrorTempl           = "New{{Not}}ContainsOptValidationError"
    Comment              = "option value must be a member of collection"
    #
    Negate               = $true
    ExcludeTypes         = @("Bool", "IPMask", "IPNet")
    NegateMethodTemplate = "Not{{OpName}}{{TypeName}}"
    NegateErrorMessage   = "is a member of"
    NegateComment        = "option value must not be a member of collection"
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
    ErrorTempl           = "New{{Not}}MatchOptValidationError"
    Comment              = "option value must match regex pattern"
    #
    Negate               = $true
    ExcludeTypes         = @()
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
    ErrorTempl      = "NewGreaterThanOptValidationError"
    Comment         = "option value must be greater than threshold"
    #
    # Relatable is the op equivalent of spec.Comparable. Operations that are relatable
    # are compatible with spec types that are Comparable.
    #
    Relatable       = $true
    ExcludeTypes    = @("Bool", "IPNet", "IPMask", "Enum")
  }

  , [PSCustomObject]@{
    Name            = "AtLeast"
    Documentation   = "fails validation if the option value is not comparably greater than or equal to 'threshold'"
    Args            = "threshold"
    Condition       = "value >= threshold"
    ErrorMessage    = "at least"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    ErrorTempl      = "NewAtLeastOptValidationError"
    Comment         = "option value must be greater than or equal to threshold"
    #
    Relatable       = $true
    ExcludeTypes    = @("Bool", "IPNet", "IPMask", "Enum")
  }

  , [PSCustomObject]@{
    Name            = "LessThan"
    Documentation   = "fails validation if the option value is not comparably less than 'threshold'"
    Args            = "threshold"
    Condition       = "value < threshold"
    ErrorMessage    = "not less than"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    ErrorTempl      = "NewLessThanOptValidationError"
    Comment         = "option value must be less than threshold"
    #
    Relatable       = $true
    ExcludeTypes    = @("Bool", "IPNet", "IPMask", "Enum")
  }

  , [PSCustomObject]@{
    Name            = "AtMost"
    Documentation   = "fails validation if the option value is not comparably less than or equal to 'threshold'"
    Args            = "threshold"
    Condition       = "value <= threshold"
    ErrorMessage    = "not at most"
    ArgsPlaceholder = "[%v]"
    ErrorArgs       = "threshold"
    ErrorTempl      = "NewAtMostOptValidationError"
    Comment         = "option value must be less than or equal to threshold"
    #
    Relatable       = $true
    ExcludeTypes    = @("Bool", "IPNet", "IPMask", "Enum")
  }
)

Write-Host "🤖 Build-Validators(gen-ov) ✨ => option-validator-auto.go"
Write-Host "🤖 Build-ParamSet(gen-ps) ✨ => param-set-auto.go"
Write-Host "🤖 Build-PsTestEntry(gen-ps-t) ✨ => param-set-auto_test.go"
Write-Host "🤖 Build-TestEntry(gen-ov-t) 🧪 => option-validator-auto_test.go"
Write-Host "🤖 Build-BinderHelpers(gen-help) 🎁 => param-set-binder-helpers-auto.go"
Write-Host "🤖 Build-BinderHelperTests(gen-help-t) 🧪 => param-set-binder-helpers-auto_test.go"

function Build-Validators {
  # (option-validator-auto.go)
  #
  [Alias("gen-ov")]
  param(
    [Parameter()]
    [switch]$NoClip
  )

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types only
      #
      $spec = $types[$_]

      $validatorType = $spec.TypeName
      $validatorStruct = "$($validatorType)OptionValidator"
      $validatorFn = $("$($spec.TypeName)ValidatorFn")
      $displayType = [string]::IsNullOrEmpty($spec.DisplayType) ? $spec.GoType : $spec.DisplayType

      # generate (XXXX = Type)
      # - type XXXXValidatorFn
      # - type XXXXOptionValidator
      # - func (validator XXXXOptionValidator) Validate()
      #
      if ($spec.Validatable) {
        @"
// $($validatorFn) defines the validator function for $($displayType) type.
type $($validatorFn) func($($spec.GoType), *pflag.Flag) error

// $($validatorStruct) defines the struct that wraps the client defined validator function
// $($validatorFn) for $($displayType) type. This is the instance that is returned by
// validated binder function BindValidated$($spec.TypeName).
type $($validatorStruct) GenericOptionValidatorWrapper[$($spec.GoType)]
        
"@
        if (-not($spec.ForeignValidatorFn)) {
          @"
// Validate invokes the client defined validator function for $($displayType) type.
func (validator $($validatorStruct)) Validate() error {
  return validator.Fn(*validator.Value, validator.Flag)
}

"@
        }

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
type $($sliceValidatorFn) func($($sliceType), *pflag.Flag) error

// $($sliceValidatorStruct) wraps the client defined validator function for type $($sliceType).
type $($sliceValidatorStruct) GenericOptionValidatorWrapper[$($sliceType)]

// Validate invokes the client defined validator function for $($sliceType) type.
func (validator $($sliceValidatorStruct)) Validate() error {
return validator.Fn(*validator.Value, validator.Flag)
}

"@
        }
      }
    })
    
  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'option-validator-auto.go'"
    $content | Set-Clipboard  
  }
}

function Build-ParamSet {
  # (param-set-auto.go)
  #
  [Alias("gen-ps")]
  param(
    [Parameter()]
    [switch]$NoClip
  )

  # each operation defined independently
  #
  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types only
      #
      $spec = $types[$_]

      $validatorFn = $("$($spec.TypeName)ValidatorFn")
      $actualTypeName = [string]::IsNullOrEmpty($spec.UnderlyingTypeName) ? $spec.TypeName : $spec.UnderlyingTypeName
      $displayType = [string]::IsNullOrEmpty($spec.DisplayType) ? $spec.GoType : $spec.DisplayType

      # generate BindXXXX
      #
      @"
// Bind$($spec.TypeName) binds $($displayType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name.
//$($spec.BindDoc)
func (params *ParamSet[N]) Bind$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType)) *ParamSet[N] {
  flagSet := params.ResolveFlagSet(info)
  if info.Short == "" {
    flagSet.$($actualTypeName)Var(to, info.FlagName(), info.Default.($($spec.GoType)), info.Usage)
  } else {
    flagSet.$($actualTypeName)VarP(to, info.FlagName(), info.Short, info.Default.($($spec.GoType)), info.Usage)
  }

  return params
}

"@

      if ($spec.Validatable) {
        # generate BindValidatedXXXX
        #
        @"
// BindValidated$($spec.TypeName) binds $($displayType) slice flag with a shorthand if
// 'info.Short' has been set otherwise binds without a short name. Client can provide a
// function to validate option values of $($displayType) type.
//$($spec.BindValidatedDoc)
func (params *ParamSet[N]) BindValidated$($spec.TypeName)(info *FlagInfo, to *$($spec.GoType), validator $($validatorFn)) OptionValidator {
  params.Bind$($spec.TypeName)(info, to)

  wrapper := $($actualTypeName)OptionValidator{
    Fn:    validator,
    Value: to,
    Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
  }
  params.validators.Add(info.FlagName(), wrapper)

  return wrapper
}

"@
      }

      # generate BindXXXXSlice
      #
      if ($spec.GenerateSlice) {
        $sliceTypeName = "$($spec.TypeName)Slice"
        $sliceType = "[]$($spec.GoType)"
        $defaultSlice = $("[]$($spec.GoType)")

        @"
// Bind$($sliceTypeName) binds $($sliceType) slice flag with a shorthand if 'info.Short' has been set
// otherwise binds without a short name.
func (params *ParamSet[N]) Bind$($sliceTypeName)(info *FlagInfo, to *$($sliceType)) *ParamSet[N] {
  flagSet := params.ResolveFlagSet(info)
  if info.Short == "" {
    flagSet.$($sliceTypeName)Var(to, info.FlagName(), info.Default.($($sliceType)), info.Usage)
  } else {
    flagSet.$($sliceTypeName)VarP(to, info.FlagName(), info.Short, info.Default.($($defaultSlice)), info.Usage)
  }

  return params
}

"@

        if ($spec.Validatable) {
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
func (params *ParamSet[N]) BindValidated$($sliceTypeName)(info *FlagInfo, to *$($sliceType), validator $($sliceValidatorFn)) OptionValidator {
  params.Bind$($sliceTypeName)(info, to)

  wrapper := $($sliceTypeName)OptionValidator{
    Fn:    validator,
    Value: to,
    Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
  }
  params.validators.Add(info.FlagName(), wrapper)

  return wrapper
}

"@
        }
      }
    })

  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'param-set-auto.go'"
    $content | Set-Clipboard  
  }
}

# Build-PsTestEntry(gen-ps-t)
function Build-PsTestEntry {
  # (param-set-auto_test.go)
  #
  [Alias("gen-ps-t")]
  param(
    [Parameter()]
    [switch]$NoClip
  )

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types only
      #
      $spec = $types[$_]

      if ($null -ne $spec.PsTcEntry) {
        $lowerFlagName = $spec.FlagName.ToLower()
        $lowerSliceFlagName = $spec.GenerateSlice ? $spec.SliceFlagName.ToLower() : [string]::IsNullOrEmpty
        $default = $spec.QuoteExpect ? $('"' + $spec.Def + '"') : $spec.Def
        $defaultSlice = $spec.GenerateSlice ? $spec.DefSliceVal : [string]::IsNullOrEmpty
        $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo
  
        $sides = @(
          [PSCustomObject]@{
            Short             = $true
            FlagInfoArgs      = """$lowerFlagName"", ""$($spec.Short)"", $default"
            SliceFlagInfoArgs = """$lowerSliceFlagName"", ""$($spec.SliceShort)"", $defaultSlice"
            MessageAdornments = ", (with-short)"
          },
          [PSCustomObject]@{
            Short             = $false
            FlagInfoArgs      = """$lowerFlagName"", """", $default"
            SliceFlagInfoArgs = """$lowerSliceFlagName"", """", $defaultSlice"
            MessageAdornments = ", (without-short)"
          }
        )
  
        $optionValue = $spec.QuoteExpect ? $("""$($spec.OptionValue)""") : $spec.OptionValue
        if ($spec.QuoteOptionValue) {
          $optionValue = """$($optionValue)"""
        }

        if (-not([string]::IsNullOrEmpty($spec.CastLiteralsAs))) {
          $optionValue = $("$($spec.CastLiteralsAs)($($optionValue))")
        }

        $commandLineValue = [string]::IsNullOrEmpty($spec.CommandLineValue) ? $spec.OptionValue : $spec.CommandLineValue
        $commandLineValue = $commandLineValue.Replace("{{OPTION-VALUE}}", $spec.OptionValue)

        $assertion = if (-not([string]::IsNullOrEmpty($spec.PsTcEntry.AssertFn))) {
          $spec.PsTcEntry.AssertFn
        }
        elseif (-not([string]::IsNullOrEmpty($spec.Equate))) {
          "func() { Expect(paramSet.Native.$($spec.FlagName)).To($($spec.Equate)($optionValue)) }"
        }
        else {
          "func() { Fail(""❌ MISSING ASSERTION FUNCTION OR EQUATE"") }"
        }

        $commandLine = $spec.IsOptionLess ? "--$($lowerFlagName)" : "--$($lowerFlagName)=$($commandLineValue)"
        $displayType = [string]::IsNullOrEmpty($spec.DisplayType) ? $spec.GoType : $spec.DisplayType

        $sliceTypeName = "$($spec.TypeName)Slice"
        $sliceType = "[]$($displayType)"
        $defaultSlice = $("[]$($spec.GoType)")

        $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo

        foreach ($side in $sides) {
          $constructFlagInfo = $("assistant.NewFlagInfo($($side.FlagInfoArgs))")

          @"
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

"@
          if ($spec.GenerateSlice) {
            $sliceCommandLine = "--$($lowerSliceFlagName)=$($spec.SliceValue)"
            $constructSliceFlagInfo = $("assistant.NewFlagInfo($($side.SliceFlagInfoArgs))")
            $sliceAssertion = if (-not([string]::IsNullOrEmpty($spec.PsTcEntry.AssertSliceFn))) {
              $spec.PsTcEntry.AssertFn
            }
            else {
              "func() { Expect(paramSet.Native.$($spec.SliceFlagName)).To(BeEquivalentTo($($spec.ExpectSlice))) }"
            }

            $sliceBindTo = $("&paramSet.Native.$($spec.SliceFlagName)")
    
            @"
Entry(nil, TcEntry{
  Message: "$($sliceType) slice type$($side.MessageAdornments) (auto)",
  Binder: func() {
    paramSet.Bind$($sliceTypeName)(
      $($constructSliceFlagInfo),
      $($sliceBindTo),
    )
  },
  CommandLine: "$sliceCommandLine",
  Assert: $($sliceAssertion),
}),
            
"@            
          }
        }
      }
    })

  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'param-set-auto_test.go'"
    $content | Set-Clipboard  
  }
}

function Build-TestEntry {
  # (option-validator-auto_test.go)
  #
  [Alias("gen-ov-t")]
  param(
    [Parameter()]
    [switch]$NoClip
  )
  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types only
      #
      $spec = $types[$_]

      $lowerFlagName = $spec.FlagName.ToLower()
      $default = $spec.QuoteExpect ? $('"' + $spec.Def + '"') : $spec.Def
      $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo

      if ($spec.Validatable) {
        $optionValue = $spec.OptionValue
        if ($spec.QuoteOptionValue) {
          $optionValue = """$($optionValue)"""
        }

        if (-not([string]::IsNullOrEmpty($spec.CastLiteralsAs))) {
          $optionValue = $("$($spec.CastLiteralsAs)($($optionValue))")
        }

        $setup = $spec.Setup.Replace("{{OPTION-VALUE}}", $optionValue)
        $assert = $spec.Assert.Replace("{{OPTION-VALUE}}", $optionValue)

        # generate BindValidatedXXX OvEntry
        #
        @"
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
      func(value $($spec.GoType), flag *pflag.Flag) error {
        $($assert)
        return nil
      },
    )
  },
}),

"@
      }

      if ($spec.GenerateSlice) {
        if ($spec.Validatable) {
          $sliceTypeName = "$($spec.TypeName)Slice"
          $sliceType = "[]$($spec.GoType)"

          # generate BindValidatedXXXSlice OvEntry
          #
          @"
Entry(nil, OvEntry{
  Message: "$($sliceType) type (auto)",
  Setup: func() {
    paramSet.Native.$($spec.SliceFlagName) = $($spec.ExpectSlice)
  },
  Validator: func() assistant.OptionValidator {
    return paramSet.BindValidated$($sliceTypeName)(
      assistant.NewFlagInfo("$($spec.SliceFlagName)", "$($spec.SliceShort)", $($spec.DefSliceVal)),
      &paramSet.Native.$($spec.SliceFlagName),
      func(value $($sliceType), flag *pflag.Flag) error {
        Expect(value).To($($spec.Equate)($($spec.ExpectSlice)))
        return nil
      },
    )
  },
}),

"@
        }

      }
    })
  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'option-validator-auto_test.go'"
    $content | Set-Clipboard  
  }
}

function Build-BinderHelpers {
  # (param-set-binder-helpers-auto.go)
  #
  [Alias("gen-help")]
  [CmdletBinding()]
  param(
    [Parameter()]
    [Switch]$Indicate,

    [Parameter()]
    [switch]$NoClip
  )
  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types and Operations
      #
      $spec = $types[$_]

      foreach ($op in $operators) {
        if (-not(Test-IsCompatibleCombo -TypeSpec $spec -Operation $op -Indicate:$Indicate.IsPresent)) {
          continue
        }

        # assuming all args have the same type
        #
        $argumentsStmt = if ($op.Container) {
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

        $errorFn = $op.ErrorTempl.Replace("{{Not}}", [string]::Empty)
        $errorF = "i18n.$($errorFn)(info.FlagName(), value, $($op.ErrorArgs))"

        # generate BuildValidatedXXXXOp/BuildValidatedOpXXXX
        #
        @"
// BindValidated$($methodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method $($op.Documentation).
func (params *ParamSet[N]) BindValidated$($methodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {
  params.Bind$($spec.TypeName)(info, to)

  wrapper := GenericOptionValidatorWrapper[$($spec.GoType)]{
    Fn: func(value $($spec.GoType), flag *pflag.Flag) error {
      if $($op.Condition) {
        return nil
      }
      return $($errorF)
    },
    Value: to,
    Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
  }
  params.validators.Add(info.FlagName(), wrapper)

  return wrapper
}

"@

        if (-not($op.Relatable)) {

          $notMethodSubStmt = if (-not([string]::IsNullOrEmpty($op.NegateMethodTemplate))) {
            $op.NegateMethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
          }
          else {
            # Default is TypeNameNotOpName, eg: StringNotGreaterThan
            #
            $("$($spec.TypeName)Not$($op.Name)")
          }
          $negatedCondition = $("!($($op.Condition))")

          $errorFn = $op.ErrorTempl.Replace("{{Not}}", "Not")
          $errorF = "i18n.$($errorFn)(info.FlagName(), value, $($op.ErrorArgs))"
  
          # generate NOT method
          #
          @"
// BindValidated$($notMethodSubStmt) is an alternative to using BindValidated$($spec.TypeName). Instead of providing
// a function, the client passes in argument(s): '$($op.Args)' to utilise predefined functionality as a helper.
// This method performs the inverse of 'BindValidated$($methodSubStmt)'.
func (params *ParamSet[N]) BindValidated$($notMethodSubStmt)(info *FlagInfo, to *$($spec.GoType), $($argumentsStmt)) OptionValidator {
  params.Bind$($spec.TypeName)(info, to)

  wrapper := GenericOptionValidatorWrapper[$($spec.GoType)]{
    Fn: func(value $($spec.GoType), flag *pflag.Flag) error {
      if $($negatedCondition) {
        return nil
      }
      return $($errorF)
    },
    Value: to,
    Flag:  params.ResolveFlagSet(info).Lookup(info.Name),
  }
  params.validators.Add(info.FlagName(), wrapper)

  return wrapper
}

"@
        }
      }
    })

  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'param-set-binder-helpers-auto.go'"
    $content | Set-Clipboard  
  }
}

function Build-BinderHelperTests {
  # (param-set-binder-helpers-auto_test.go)
  #
  [Alias("gen-help-t")]
  [CmdletBinding()]
  param(
    [Parameter()]
    [Switch]$Indicate,

    [Parameter()]
    [switch]$NoClip
  )

  $content = ($types.Keys | Sort-Object | ForEach-Object {
      # iterate over spec Types and Operations
      #
      $spec = $types[$_]

      $bindTo = [string]::IsNullOrEmpty($spec.BindTo) ? $("&paramSet.Native.$($spec.FlagName)") : $spec.BindTo
      $default = $spec.QuoteExpect ? $('"' + $spec.Def + '"') : $spec.Def
      [int]$ValueIndex = 0
      [int]$ExpectNilIndex = 1

      foreach ($op in $operators) {
        if (-not(Test-IsCompatibleCombo -TypeSpec $spec -Operation $op -Indicate:$Indicate.IsPresent)) {
          continue
        }

        if (($null -eq $spec.BhTests)) {
          if (-not([string]::IsNullOrEmpty($spec.BhParent)) -and ($types.ContainsKey($spec.BhParent))) {
            $spec.BhTests = $types[$spec.BhParent].BhTests
  
            if ($null -eq $spec.BhTests) {
              Write-Host "===> 🔥🔥🔥 BhParent '$($spec.BhParent)' of '$($spec.TypeName)' does not have a valid BhTests, skipping ..."
              continue
            }
          }
          else {
            Write-Host "===> 🔥🔥🔥 '$($spec.TypeName)' does not have a valid BhTests, skipping ..."
            continue
          }
        }

        $cast = $spec.CastLiteralsAs

        if (-not($spec.BhTests.ContainsKey($op.Name))) {
          continue
        }

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

        $notMethodSubStmt = if (-not([string]::IsNullOrEmpty($op.NegateMethodTemplate))) {
          $op.NegateMethodTemplate.Replace("{{OpName}}", $op.Name).Replace("{{TypeName}}", $spec.TypeName)
        }
        else {
          # Default is TypeNameNotOpName, eg: StringNotGreaterThan
          #
          $("$($spec.TypeName)Not$($op.Name)")
        }

        $sides = @(
          [PSCustomObject]@{
            Expectation = "expectNil"
            Method      = $methodSubStmt
          },
          [PSCustomObject]@{
            Expectation = "!expectNil"
            Method      = $notMethodSubStmt
          }
        )

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

          # generate Within/NotWithin testcases
          #
          # For the NOT scenario test cases, we should still be able to use the exact same test data,
          # but just tweak the test code to reverse the logic. Eg, so we simply negate the expectNil
          # and everything else stays the same.
          #
          foreach ($side in $sides) {
            
            $testTable = @"
DescribeTable("BindValidated$($side.Method)",
  func(given, should string, value $($spec.GoType), expectNil bool, low, high $($spec.GoType)) {
    validator := paramSet.BindValidated$($side.Method)(
      assistant.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), low, high,
    )
    paramSet.Native.$($spec.FlagName) = value

    if $($side.Expectation) {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }
  },
  func(given, should string, value $($spec.GoType), expectNil bool, low, high $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍋 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value is below range", "return error", $($belowEntryArgs)),
  Entry(nil, "value is equal to low end of range", "return error", $($equalLoEntryArgs)),
  Entry(nil, "value is inside range", "return error", $($insideEntryArgs)),
  Entry(nil, "value is equal to high end of range", "return error", $($equalHiEntryArgs)),
  Entry(nil, "value is above range", "NOT return error", $($aboveEntryArgs)),
)

"@
            $testTable
          }
        }
        elseif ($op.Container) {
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

          $assign = ([string]::IsNullOrEmpty($testOp.Assign)) ? "paramSet.Native.$($spec.FlagName) = value": $testOp.Assign

          # generate Contains/NotContains testcases
          #
          # For the NOT scenario test cases, we should still be able to use the exact same test data,
          # but just tweak the test code to reverse the logic. Eg, so we simply negate the expectNil
          # and everything else stays the same.
          #
          foreach ($side in $sides) {
            $testTable = @"
DescribeTable("BindValidated$($side.Method)",
  func(given, should string, value $($spec.GoType), expectNil bool, collection []$($spec.GoType), dummy $($spec.GoType)) {
    validator := paramSet.BindValidated$($side.Method)(
      assistant.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), collection,
    )
    $($assign)

    if $($side.Expectation) {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }
  },
  func(given, should string, value $($spec.GoType), expectNil bool, collection []$($spec.GoType), dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍎 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "collection contains member", "return error", $($doesContainArgs)),
  Entry(nil, "collection does not contain member", "return error", $($doesNotContainArgs)),
)

"@
            $testTable
          }
        }
        elseif ($op.Name -eq "IsMatch") {
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

          # generate IsMatch/IsNotMatch testcases
          #
          # For the NOT scenario test cases, we should still be able to use the exact same test data,
          # but just tweak the test code to reverse the logic. Eg, so we simply negate the expectNil
          # and everything else stays the same.
          #
          foreach ($side in $sides) {

            $testTable = @"
DescribeTable("BindValidated$($side.Method)",
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy $($spec.GoType)) {
    validator := paramSet.BindValidated$($side.Method)(
      assistant.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), pattern,
    )
    paramSet.Native.$($spec.FlagName) = value

    if $($side.Expectation) {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }
  },
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍇 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value matches pattern", "return error", $($doesMatchArgs)),
  Entry(nil, "value does not match pattern", "return error", $($doesNotMatchArgs)),
)

"@
            $testTable
          }
        }
        elseif ($spec.Comparable) {
          if ([string]::IsNullOrEmpty($cast)) {
            $belowThresholdArgs = $("$($testOp.Entry.Below[$ValueIndex]), $($testOp.Entry.Below[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
            $equalThresholdArgs = $("$($testOp.Entry.Equal[$ValueIndex]), $($testOp.Entry.Equal[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")
            $aboveThresholdArgs = $("$($testOp.Entry.Above[$ValueIndex]), $($testOp.Entry.Above[$ExpectNilIndex]), $($testOp.First), $($testOp.Second)")  
          }
          else {
            $belowThresholdArgs = $("$($cast)($($testOp.Entry.Below[$ValueIndex])), $($testOp.Entry.Below[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
            $equalThresholdArgs = $("$($cast)($($testOp.Entry.Equal[$ValueIndex])), $($testOp.Entry.Equal[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
            $aboveThresholdArgs = $("$($cast)($($testOp.Entry.Above[$ValueIndex])), $($testOp.Entry.Above[$ExpectNilIndex]), $($cast)($($testOp.First)), $($cast)($($testOp.Second))")
          }

          @"
DescribeTable("BindValidated$($methodSubStmt)",
  func(given, should string, value $($spec.GoType), expectNil bool, threshold, dummy $($spec.GoType)) {
    validator := paramSet.BindValidated$($methodSubStmt)(
      assistant.NewFlagInfo("$($spec.FlagName.ToLower())", "$($spec.Short)", $($default)),
      $($bindTo), threshold,
    )
    paramSet.Native.$($spec.FlagName) = value

    if expectNil {
      Expect(validator.Validate()).Error().To(BeNil())
    } else {
      Expect(validator.Validate()).Error().ToNot(BeNil())
    }
  },
  func(given, should string, value $($spec.GoType), expectNil bool, pattern, dummy $($spec.GoType)) string {
    return fmt.Sprintf("🧪 --> 🍌 given: '%v', should: '%v'",
      given, should)
  },
  Entry(nil, "value is below threshold", "return error", $($belowThresholdArgs)),
  Entry(nil, "value is equal threshold", "return error", $($equalThresholdArgs)),
  Entry(nil, "value is above threshold", "NOT return error", $($aboveThresholdArgs)),
)

"@
        }
        else {
          Write-Host "!!!! 👽 SKIPPING OP: '$($op.Name)' for Type: '$($spec.TypeName)' (coverage gap likely)"
        }
      }
    })

  if ($NoClip.IsPresent) {
    return $content
  }
  else {
    Write-Host "🎯 Paste into ---> 'param-set-binder-helpers-auto_test.go'"
    $content | Set-Clipboard  
  }
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
    [PSCustomObject]$Operator,

    [Parameter()]
    [switch]$NoClip

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


function Test-IsCompatibleCombo {
  [CmdletBinding()]
  [OutputType([bool])]
  param(
    [Parameter(Mandatory)]
    [PSCustomObject]$TypeSpec,

    [Parameter(Mandatory)]
    [PSCustomObject]$Operation,

    [Parameter()]
    [Switch]$Indicate
  )
  [bool]$result = $true;

  if (-not([string]::IsNullOrEmpty($Operation.AppliesOnlyTo)) -and ($Operation.AppliesOnlyTo -ne $TypeSpec.TypeName)) {
    $result = $false;
  }
  elseif ($Operation.ExcludeTypes -contains $TypeSpec.TypeName) {
    $result = $false;
  }

  if ($Indicate.IsPresent) {
    $indicator = $result ? "✔️" : "❌"
    Write-Host "===> Test-IsCompatibleCombo  $($indicator) | type: '$($TypeSpec.TypeName)', op: '$($Operation.Name)'"  
  }

  return $result
}

function Checkpoint-ParamSetSignatures {
  [Alias("check-sig")]
  [CmdletBinding()]
  [OutputType([PSCustomObject])]
  param(
    [Parameter()]
    $Sources = @("gen-ov", "gen-ps", "gen-help")
  )
  [System.Text.StringBuilder]$hBuilder = [System.Text.StringBuilder]::new()

  [hashtable]$metrics = @{}
  $capture = ""

  $totals = [PSCustomObject]@{
    Counters = [PSCustomObject]@{
      Functions = 0
      Types     = 0
    }
  }

  foreach ($producer in $Sources) {
    $capture = $(Invoke-Expression -Command "$producer -NoClip")

    $metrics[$producer] = [PSCustomObject]@{
      Counters = [PSCustomObject]@{
        Functions = 0
        Types     = 0
      }
    }

    $endings = $IsWindows ? "`r`n" : "`n"
    $capture -split $endings | ForEach-Object {
      $line = $_
      if ($line.StartsWith("func")) {
        $index = $line.LastIndexOf(" {")
        if ($index -ge 0) {
          $signature = $line.Substring(0, $index + 1)
          $null = $hBuilder.AppendLine($signature.Trim())
          $metrics[$producer].Counters.Functions++
        }
      }
      elseif ($line.StartsWith("type")) {
        $null = $hBuilder.AppendLine($line.Trim())
        $metrics[$producer].Counters.Types++
      }
    }
    $totals.Counters.Functions += $metrics[$producer].Counters.Functions
    $totals.Counters.Types += $metrics[$producer].Counters.Types
  }

  [System.Text.StringBuilder]$oBuilder = [System.Text.StringBuilder]::new()
  $metrics.Keys | ForEach-Object {
    $metric = $metrics[$_]
    $null = $oBuilder.AppendLine("---> 🍄 [$_] Signature Counts - 🍅functions: '$($metric.Counters.Functions)', 🥦types: '$($metric.Counters.Types)'")
  }
  $null = $oBuilder.AppendLine("---> 🍄 Total Counts - 🍅functions: '$($totals.Counters.Functions)', 🥦types: '$($totals.Counters.Types)'")

  $stream = [IO.MemoryStream]::new([byte[]][char[]]$hBuilder.ToString())
  $hash = Get-FileHash -InputStream $stream -Algorithm SHA256

  return [PSCustomObject]@{
    PreviousHash = $($env:COBRASS_API_HASH)
    Hash         = $hash.Hash
    Metrics      = $metrics
    Output       = $oBuilder.ToString() 
  }
}

function Show-ParamSetSignatures {
  [Alias("show-sig")]
  [CmdletBinding()]
  [OutputType([string])]
  param(
    [Parameter()]
    $Sources = @("gen-ov", "gen-ps", "gen-help")
  )
  $paramSigs = Checkpoint-ParamSetSignatures -Sources $Sources 

  Write-Host $paramSigs.Output -ForegroundColor Cyan
  Write-Host "===> [🤖]  THIS-HASH: '$($paramSigs.Hash)'" -ForegroundColor Green

  $status = if ([string]::IsNullOrEmpty($paramSigs.PreviousHash)) {
    "💤 Previous api hash not found"
  }
  elseif ($paramSigs.PreviousHash -eq $paramSigs.Hash) {
    "✔️ Hashes are equal"
  }
  else {
    Write-Host "===> [👾]   OLD-HASH: '$($paramSigs.PreviousHash)'" -ForegroundColor Magenta
    "💥 Api changes detected"
  }
  Write-Host "===> [🛡️] STATUS: '$($status)' (COBRASS_API_HASH)" -ForegroundColor Blue
}
