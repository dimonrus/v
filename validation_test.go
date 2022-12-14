package v

import (
	"reflect"
	"testing"
	"time"
)

type ComplexStruct struct {
	Cool bool
}

type AliasOnTypeString string

type Nested struct {
	Foo int32 `json:"foo"`
	Bar *bool `json:"bar" valid:"required"`
}

type TestValidationStruct struct {
	Name      string            `json:"name" valid:"required;exp~[a-z]+"`
	Number    int               `json:"number" valid:"exp~[0-5]+;range~1:50;enum~[5,10,15,20,25]"`
	IsTrue    *bool             `json:"isTrue"`
	Complex   *ComplexStruct    `json:"complex" valid:"required"`
	Sl        []int64           `json:"sl"`
	SuperName AliasOnTypeString `json:"superName" valid:"required"`
	Nested    Nested            `json:"nested"`
}

func TestValidateStructPointer(t *testing.T) {
	v := TestValidationStruct{Complex: &ComplexStruct{}}
	e := ValidateStruct(&v)
	if e != nil {
		for _, iError := range e.GetDetails() {
			t.Log(iError.Error())
		}
		t.Log(e.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	n := Nested{Bar: new(bool)}
	e := ValidateStruct(n)
	if e == nil {
		t.Fatal("must have en error")
	}
	t.Log(e.Error())
}

//switch val.Kind() {
//	case reflect.Ptr:
//		if val.IsNil() || val.IsZero() {
//			return false
//		}
//		return IsRequiredValid(val.Elem())
//	case reflect.String:
//		return val.String() != ""
//	case reflect.Float64:
//		fallthrough
//	case reflect.Float32:
//		return val.Float() != 0
//	case reflect.Int:
//		fallthrough
//	case reflect.Int8:
//		fallthrough
//	case reflect.Int16:
//		fallthrough
//	case reflect.Int32:
//		fallthrough
//	case reflect.Int64:
//		return val.Int() != 0
//	case reflect.Uint:
//		fallthrough
//	case reflect.Uint8:
//		fallthrough
//	case reflect.Uint16:
//		fallthrough
//	case reflect.Uint32:
//		fallthrough
//	case reflect.Uint64:
//		return val.Uint() != 0
//	case reflect.Slice:
//		fallthrough
//	case reflect.Array:
//		return val.Len() > 0
//	case reflect.Chan:
//		return val.Len() > 0
//	case reflect.Struct:
//		return !val.IsZero()
//	}
//	return true

func TestValidateStruct(t *testing.T) {
	v := TestValidationStruct{Name: "foo", Complex: &ComplexStruct{}}
	e := ValidateStruct(v)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.Origin()
			t.Log(er.Name, er.Message)
		}
		t.Log(e.Error())
	}
}

func TestValidationRequired(t *testing.T) {
	var val int32 = 12
	v := reflect.ValueOf(val)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	val = 0
	v = reflect.ValueOf(val)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pval = new(int32)
	*pval = 12
	v = reflect.ValueOf(pval)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pval = new(int32)
	v = reflect.ValueOf(pval)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var val32 uint32 = 12
	v = reflect.ValueOf(val32)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	val32 = 0
	v = reflect.ValueOf(val32)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pval32 = new(uint32)
	*pval32 = 12
	v = reflect.ValueOf(pval32)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pval32 = new(uint32)
	v = reflect.ValueOf(pval32)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var str string
	v = reflect.ValueOf(str)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pstr *string
	v = reflect.ValueOf(pstr)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	str = "okey"
	v = reflect.ValueOf(str)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pstr = new(string)
	*pstr = "okey"
	v = reflect.ValueOf(pstr)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var boolean bool
	v = reflect.ValueOf(boolean)
	if IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	boolean = true
	v = reflect.ValueOf(boolean)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var pboolean *bool
	v = reflect.ValueOf(pboolean)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var fl float32
	v = reflect.ValueOf(fl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	fl = 0.33
	v = reflect.ValueOf(fl)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var pfl *float32
	v = reflect.ValueOf(pfl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	pfl = new(float32)
	v = reflect.ValueOf(pfl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	*pfl = 12.1231
	v = reflect.ValueOf(pfl)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var array []int
	v = reflect.ValueOf(array)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	marray := [3]int{1, 2, 3}
	v = reflect.ValueOf(marray)
	if !IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	mnarray := [0]int{}
	v = reflect.ValueOf(mnarray)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch chan int
	v = reflect.ValueOf(ch)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch1 = make(chan int, 0)
	v = reflect.ValueOf(ch1)
	if !IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch2 = make(chan int, 1)
	ch2 <- 2
	v = reflect.ValueOf(ch2)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}
	close(ch2)

	var st = struct{ a int }{}
	v = reflect.ValueOf(st)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	st = struct{ a int }{a: 1}
	v = reflect.ValueOf(st)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}
}

type TestEnumStruct struct {
	Foo     string  `json:"foo" valid:"enum~empty,base,value"`
	Number  float32 `json:"number" valid:"enum~0.1,0.5,0.9"`
	Bar     int64   `json:"bar" valid:"enum~200,500,9000,100"`
	PNumber *int64  `json:"pNumber" valid:"enum~100,50,20,10"`
}

func TestIsEnumValid(t *testing.T) {
	s := TestEnumStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 1001
	e := ValidateStruct(s)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.Origin()
			t.Log(er.Name, er.Message)
		}
	} else {
		t.Fatal("must be an error")
	}
}

func TestIsEmptyEnumValid(t *testing.T) {
	s := TestEnumStruct{}
	e := ValidateStruct(s)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.Origin()
			t.Fatal(er.Name, er.Message)
		}
	}
}

func BenchmarkEnumStruct(b *testing.B) {
	s := TestEnumStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 1001
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

type TestRequiredStruct struct {
	Foo     string  `json:"foo" valid:"required"`
	Number  float32 `json:"number" valid:"required"`
	Bar     int64   `json:"bar" valid:"required"`
	PNumber *int64  `json:"pNumber" valid:"required"`
}

func BenchmarkRequired(b *testing.B) {
	s := TestRequiredStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 100
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

func BenchmarkCheckNative(b *testing.B) {
	s := TestRequiredStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

func TestParseValidTag(t *testing.T) {
	rules := ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25")
	t.Log(rules)
}

func BenchmarkParseValidTag(b *testing.B) {
	rules := ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25;other~cool;cool~231231")
	b.Log(rules)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25")
	}
	b.ReportAllocs()
}

type TestRangeValidation struct {
	RangeInt     int     `json:"rangeInt" valid:"range~-1:50"`
	RangeFloat32 float32 `json:"rangeFloat32" valid:"range~9.5:10.5"`
	RangeUint    uint    `json:"rangeUint" valid:"range~0:10"`
}

func TestIsRangeValid(t *testing.T) {
	s := TestRangeValidation{
		RangeInt:     -1,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	e := ValidateStruct(s)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.Origin()
			t.Log(er.Name, er.Message)
		}
		t.Fatal(e)
	}
}

func TestIsRangeInValid(t *testing.T) {
	s := TestRangeValidation{
		RangeInt:     70,
		RangeFloat32: 9,
		RangeUint:    12,
	}
	e := ValidateStruct(s)
	if e == nil {
		t.Fatal("Must be an error")
	} else {
		for _, iError := range e.GetDetails() {
			er := iError.Origin()
			t.Log(er.Name, er.Message)
		}
	}
}

func BenchmarkIsRangeValid(b *testing.B) {
	s := TestRangeValidation{
		RangeInt:     22,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateStruct(s)
	}
	b.ReportAllocs()
}

func BenchmarkSeveralTag(b *testing.B) {
	s := TestRangeValidation{
		RangeInt:     22,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	t := reflect.TypeOf(s)
	var tag reflect.StructTag
	for i := 0; i < b.N; i++ {
		for j := 0; j < t.NumField(); j++ {
			tag = t.Field(j).Tag
			_ = tag.Get("valid")
			_ = tag.Get("json")
		}
	}
	b.ReportAllocs()
}

type TestReg struct {
	Name string `json:"name" valid:"rx~^\\d+$;rx~[0-8]+"`
	//Name string `json:"name"`
}

func TestIsRegularValid(t *testing.T) {
	rs := TestReg{
		Name: "12291",
	}
	e := ValidateStruct(rs)
	if e != nil {
		t.Fatal(e.GetDetails())
	}

	v := "1"
	e = ValidateStruct(v)
	if e == nil {
		t.Fatal("must be error& wrong type")
	}
}

func BenchmarkIsRegularValid(b *testing.B) {
	rs := TestReg{
		Name: "1221",
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateStruct(rs)
	}
	b.ReportAllocs()
}

type TestValidStruct struct {
	RangeInt *int `json:"rangeInt" valid:"required"`
}

type WrapStruct struct {
	*TestValidStruct
	When *time.Time `json:"when"`
}

func TestWraped(t *testing.T) {
	now := time.Now()
	rng := 50
	s := &WrapStruct{
		TestValidStruct: &TestValidStruct{RangeInt: &rng},
		When:            &now,
	}
	e := ValidateStruct(s)
	if e != nil {
		t.Fatal(e)
	}
}

func TestVTestValidStruct(t *testing.T) {
	tag := "required"
	rules := ParseValidTag(tag)
	for _, rule := range rules {
		t.Log("name:", rule.Name)
		t.Log("arg:", rule.Args)
		t.Log("####")
	}
}

func BenchmarkRequiredTag(b *testing.B) {
	tag := "required;"
	for i := 0; i < b.N; i++ {
		_ = ParseValidTag(tag)
	}
	b.ReportAllocs()
}

func BenchmarkWrapStruct(b *testing.B) {
	now := time.Now()
	rng := 50
	s := &WrapStruct{
		TestValidStruct: &TestValidStruct{RangeInt: &rng},
		When:            &now,
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateStruct(s)
	}
	b.ReportAllocs()
}

func TestIsMinValid(t *testing.T) {
	t.Run("min_error_expected", func(t *testing.T) {
		minErr := struct {
			Min string `valid:"min~5"`
		}{Min: "some"}
		e := ValidateStruct(&minErr)
		if e == nil {
			t.Fatal("min error expected")
		} else {
			t.Log(e.GetDetails())
		}
	})
	t.Run("not_min_error", func(t *testing.T) {
		minErr := struct {
			Min string `valid:"min~5"`
		}{Min: "something"}
		e := ValidateStruct(&minErr)
		if e != nil {
			t.Fatal("min error unexpected")
		}
	})
	t.Run("not_min_error_float", func(t *testing.T) {
		minErr := struct {
			Min float64 `valid:"min~5"`
		}{Min: 5.01}
		e := ValidateStruct(&minErr)
		if e != nil {
			t.Fatal("min error unexpected float")
		}
	})
	t.Run("min_error_float", func(t *testing.T) {
		minErr := struct {
			Min float64 `valid:"min~5"`
		}{Min: 4.01}
		e := ValidateStruct(&minErr)
		if e == nil {
			t.Fatal("min error expected float")
		}
	})
	t.Run("not_min_error_int", func(t *testing.T) {
		minErr := struct {
			Min int32 `valid:"min~5"`
		}{Min: 5}
		e := ValidateStruct(&minErr)
		if e != nil {
			t.Fatal("min error unexpected int")
		}
	})
	t.Run("min_error_int", func(t *testing.T) {
		minErr := struct {
			Min float64 `valid:"min~5"`
		}{Min: 4}
		e := ValidateStruct(&minErr)
		if e == nil {
			t.Fatal("min error expected int")
		}
	})
	t.Run("not_min_error_uint", func(t *testing.T) {
		minErr := struct {
			Min uint `valid:"min~5"`
		}{Min: 5}
		e := ValidateStruct(&minErr)
		if e != nil {
			t.Fatal("min error unexpected uint")
		}
	})
	t.Run("min_error_uint", func(t *testing.T) {
		minErr := struct {
			Min uint `valid:"min~5"`
		}{Min: 4}
		e := ValidateStruct(&minErr)
		if e == nil {
			t.Fatal("min error expected uint")
		}
	})
}

func TestIsMaxValid(t *testing.T) {
	t.Run("max_error_expected", func(t *testing.T) {
		maxErr := struct {
			Max string `valid:"max~4"`
		}{Max: "something"}
		e := ValidateStruct(&maxErr)
		if e == nil {
			t.Fatal("max error expected")
		} else {
			t.Log(e.GetDetails())
		}
	})
	t.Run("not_max_error", func(t *testing.T) {
		maxErr := struct {
			Min string `valid:"max~5"`
		}{Min: "thing"}
		e := ValidateStruct(&maxErr)
		if e != nil {
			t.Fatal("max error unexpected")
		}
	})
	t.Run("not_max_error_float", func(t *testing.T) {
		maxErr := struct {
			Max float64 `valid:"max~6"`
		}{Max: 5.01}
		e := ValidateStruct(&maxErr)
		if e != nil {
			t.Fatal("max error unexpected float")
		}
	})
	t.Run("max_error_float", func(t *testing.T) {
		maxErr := struct {
			Max float64 `valid:"max~5"`
		}{Max: 5.01}
		e := ValidateStruct(&maxErr)
		if e == nil {
			t.Fatal("max error expected float")
		}
	})
	t.Run("not_max_error_int", func(t *testing.T) {
		maxErr := struct {
			Max int32 `valid:"max~5"`
		}{Max: 5}
		e := ValidateStruct(&maxErr)
		if e != nil {
			t.Fatal("max error unexpected int")
		}
	})
	t.Run("max_error_int", func(t *testing.T) {
		maxErr := struct {
			Max float64 `valid:"max~5"`
		}{Max: 6}
		e := ValidateStruct(&maxErr)
		if e == nil {
			t.Fatal("max error expected int")
		}
	})
	t.Run("not_max_error_uint", func(t *testing.T) {
		maxErr := struct {
			Max uint `valid:"max~5"`
		}{Max: 5}
		e := ValidateStruct(&maxErr)
		if e != nil {
			t.Fatal("max error unexpected uint")
		}
	})
	t.Run("max_error_uint", func(t *testing.T) {
		maxErr := struct {
			Max uint `valid:"max~5"`
		}{Max: 6}
		e := ValidateStruct(&maxErr)
		if e == nil {
			t.Fatal("max error expected uint")
		}
	})
}

func TestIsDigits(t *testing.T) {
	t.Run("digits_string_ok", func(t *testing.T) {
		st := struct {
			Value string `valid:"digit"`
		}{Value: "1234567890"}
		e := ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_string_ok")
		}
	})
	t.Run("digits_string_nok", func(t *testing.T) {
		st := struct {
			Value string `valid:"digit"`
		}{Value: "1234567890a"}
		e := ValidateStruct(&st)
		if e == nil {
			t.Fatal("digits_string_nok")
		}
	})
	t.Run("digits_string_ok_len", func(t *testing.T) {
		st := struct {
			Value string `valid:"digit~5,10"`
		}{Value: "1234567890"}
		e := ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_string_ok")
		}
		st = struct {
			Value string `valid:"digit~5,10"`
		}{Value: "12345"}
		e = ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_string_ok")
		}
	})
	t.Run("digits_string_nok_len", func(t *testing.T) {
		st := struct {
			Value string `valid:"digit~5,10"`
		}{Value: "123456789"}
		e := ValidateStruct(&st)
		if e == nil {
			t.Fatal("digits_string_nok_len")
		}
		st = struct {
			Value string `valid:"digit~5,10"`
		}{Value: "1234v"}
		e = ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_string_nok_len")
		}
	})
	t.Run("digits_int_ok", func(t *testing.T) {
		st := struct {
			Value int `valid:"digit"`
		}{Value: 9100001111}
		e := ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_int_ok")
		}
	})
	t.Run("digits_int_nok", func(t *testing.T) {
		st := struct {
			Value int `valid:"digit~9"`
		}{Value: 9100001111}
		e := ValidateStruct(&st)
		if e == nil {
			t.Fatal("digits_int_nok")
		}
	})
	t.Run("digits_uint_ok", func(t *testing.T) {
		st := struct {
			Value uint `valid:"digit~4"`
		}{Value: 1223}
		e := ValidateStruct(&st)
		if e != nil {
			t.Fatal("digits_uint_ok")
		}
	})
	t.Run("digits_uint_nok", func(t *testing.T) {
		st := struct {
			Value uint `valid:"digit~3"`
		}{Value: 1123}
		e := ValidateStruct(&st)
		if e == nil {
			t.Fatal("digits_uint_nok")
		}
	})
}

func TestIsNotNullValid(t *testing.T) {
	t.Run("is_not_null_ok", func(t *testing.T) {
		v := new(uint)
		*v = 11
		st := struct {
			Value *uint `valid:"notnull"`
		}{Value: v}
		e := ValidateStruct(&st)
		if e != nil {
			t.Fatal("is_not_null_ok", e.Error())
		}
	})
	t.Run("is_not_null_nok", func(t *testing.T) {
		st := struct {
			Value *uint `valid:"notnull"`
		}{Value: nil}
		e := ValidateStruct(&st)
		if e == nil {
			t.Fatal("is_not_null_nok")
		}
	})
}
