# Validation library
Allows define custom validation rule

1. Implement callback according to the __ValidationCallback__ signature 
    ```
    func(val reflect.Value, args ...string) bool
    ```
2. Define custom validation map
    ```
    var MyValidation = map[string]ValidationCallback{"custom": func(val reflect.Value, args ...string) bool{return true}}
    ```
3. Call __PrepareActualValidationRules__ at start of application 
    ```
    v.PrepareActualValidationRules(MyValidation)
    ```

## How to use with validation rules
Form validation use combination of validation rules.
Possible rules as parts of 'valid' tag:
- required. Filed is required
- rx. Regular expression
- range. Range if values
- enum. Predefined enum
- min. Minimum value or length
- max. Maximum value or length
- digit. Only digits in value. Can specify length
- notnull. Filed must be not null

Example: `valid:"required;rx~[0-5]+;range~1:50;enum~5,10,15,20,25;digit~4,10;min~3;max~10"`

You can ignore field for validation specify valid tag as "-"
Example: `valid:"-"`

```
type ComplexStruct struct {
	Cool bool
}
type AliasOnTypeString string
type Nested struct {
	Foo int32 `json:"foo"`
	Bar *bool `json:"bar" valid:"required"`
}
type TestValidationStruct struct {
	Name      string            `json:"name" valid:"required;rx~[a-z]+"`
	Number    int               `json:"number" valid:"notnull;rx~[0-5]+;range~1:50;enum~5,10,15,20,25"`
	IsTrue    *bool             `json:"isTrue"`
	Complex   *ComplexStruct    `json:"complex" valid:"required"`
	Sl        []int64           `json:"sl"`
	SuperName AliasOnTypeString `json:"superName" valid:"required"`
	Nested    Nested            `json:"nested"`
}
v := TestValidationStruct{Complex: &ComplexStruct{}}
e := ValidateStruct(&v)
if e != nil {
   panic(e)
}
```

#### If you find this project useful or want to support the author, you can send tokens to any of these wallets
- Bitcoin: bc1qgx5c3n7q26qv0tngculjz0g78u6mzavy2vg3tf
- Ethereum: 0x62812cb089E0df31347ca32A1610019537bbFe0D
- Dogecoin: DET7fbNzZftp4sGRrBehfVRoi97RiPKajV
