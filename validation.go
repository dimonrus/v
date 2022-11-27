package v

import (
	"fmt"
	"github.com/dimonrus/porterr"
	"net/http"
	"reflect"
)

// ValidationCallback function that performs validation rule
type ValidationCallback func(val reflect.Value, args ...string) bool

// ValidationRule validation params
type ValidationRule struct {
	// Validator name
	Name string
	// Validator argument
	Args []string
}

// Basic validation rules
// You can override by using var CustomValidationRules
var basicValidationRules = map[string]ValidationCallback{
	// Required validator
	"required": IsRequiredValid,
	// Enum validator
	"enum": IsEnumValid,
	// Range validation
	"range": IsRangeValid,
	// Regular expression validation
	"rx": IsRegularValid,
	// Check if value or length <= min
	"min": IsMinValid,
	// Check if value or length >= max
	"max": IsMaxValid,
	// check for digits. can specify len
	"digit": IsDigits,
}

// Will be used in validation method
var actualValidationRules map[string]ValidationCallback

// PrepareActualValidationRules func to append basicValidationRules or replace existing rules
// customValidationRules If you want to use your own validation rules
// add the rules in to customValidationRules var
func PrepareActualValidationRules(customValidationRules map[string]ValidationCallback) {
	actualValidationRules = make(map[string]ValidationCallback)
	for s, callback := range basicValidationRules {
		actualValidationRules[s] = callback
	}
	for s, callback := range customValidationRules {
		actualValidationRules[s] = callback
	}
}

// ValidateStruct struct fields validation
func ValidateStruct(v interface{}) porterr.IError {
	e := porterr.New(porterr.PortErrorValidation, "Validation error").HTTP(http.StatusBadRequest)
	ve := reflect.ValueOf(v)
	te := reflect.TypeOf(v)

	if ve.Kind() == reflect.Ptr {
		ve = ve.Elem()
		te = te.Elem()
	}

	if ve.Kind() != reflect.Struct {
		e = e.PushDetail(porterr.PortErrorParam, "type", "Type struct required. Type "+ve.Kind().String()+" received")
		return e
	}

	var fieldName string
	var rules []ValidationRule

	var f reflect.Value
	var t reflect.StructField

	for i := 0; i < ve.NumField(); i++ {
		f = ve.Field(i)
		t = te.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			e = e.MergeDetails(ValidateStruct(f.Interface()))
		case reflect.Ptr:
			if !f.IsNil() {
				if _, ok := f.Elem().Interface().(fmt.Stringer); f.Elem().Kind() == reflect.Struct && !ok {
					e = e.MergeDetails(ValidateStruct(f.Interface()))
				}
			}
		}
		fieldName = t.Tag.Get("json")
		if fieldName == "" {
			fieldName = t.Name
		}
		rules = ParseValidTag(t.Tag.Get("valid"))
		for _, rule := range rules {
			if vRule, ok := actualValidationRules[rule.Name]; ok {
				if !vRule(f, rule.Args...) {
					e = e.PushDetail(porterr.PortErrorParam, fieldName, "Invalid validation for "+rule.Name+" rule")
				}
			}
		}
	}
	return e.IfDetails()
}

// ParseValidTag parse validation tag for rule and arguments
// Example
// valid:"exp~[0-5]+;range~1-50;enum~5,10,15,20,25"`
func ParseValidTag(validTag string) []ValidationRule {
	if validTag == "" {
		return nil
	}
	var result = make([]ValidationRule, 4)
	var ruleCount int
	var indexStart, i int

	for {
		if validTag[i] == ';' && result[ruleCount].Name == "" {
			result[ruleCount].Name = validTag[indexStart:i]
			ruleCount++
			indexStart = i + 1
		}
		if validTag[i] == '~' {
			if ruleCount == len(result) {
				result = append(result, make([]ValidationRule, 4)...)
			}
			result[ruleCount].Name = validTag[indexStart:i]
			i++
			indexStart = i
			for {
				if validTag[i] == ';' {
					result[ruleCount].Args = []string{validTag[indexStart:i]}
					ruleCount++
					indexStart = i + 1
					break
				}
				i++
				if i >= len(validTag) {
					break
				}
			}
		}
		i++
		if i >= len(validTag) {
			break
		}
	}

	if validTag[len(validTag)-1] != ';' {
		if result[ruleCount].Name == "" {
			result[ruleCount].Name = validTag[indexStart:i]
		} else {
			result[ruleCount].Args = []string{validTag[indexStart:]}
		}
		ruleCount++
	}

	return result[:ruleCount]
}

// Init default validators
func init() {
	PrepareActualValidationRules(nil)
}
