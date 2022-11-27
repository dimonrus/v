package v

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// IsRequiredValid Required validation rule
func IsRequiredValid(val reflect.Value, args ...string) bool {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		return !val.Elem().IsZero()
	}
	return !val.IsZero()
}

// IsRegularValid check regular expression
func IsRegularValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.String {
		return false
	}
	for _, arg := range args {
		matched, err := regexp.MatchString(arg, val.String())
		if err != nil || !matched {
			return false
		}
	}
	return true
}

// IsEnumValid In list validation rule
func IsEnumValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 || val.IsZero() {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		val = val.Elem()
	}
	values := strings.Split(args[0], ",")
	if len(values) == 0 {
		return true
	}
	switch val.Kind() {
	case reflect.String:
		v := val.String()
		for _, value := range values {
			if v == value {
				return true
			}
		}
		return false
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		v := val.Float()
		for _, value := range values {
			comp, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return false
			}
			if math.Abs(v-comp) < 1e-7 {
				return true
			}
		}
		return false
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		v := val.Int()
		for _, value := range values {
			comp, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return false
			}
			if v == comp {
				return true
			}
		}
		return false
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		v := val.Uint()
		for _, value := range values {
			comp, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return false
			}
			if v == comp {
				return true
			}
		}
		return false
	}
	return true
}

// IsRangeValid Range list validation rule
func IsRangeValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 || val.IsZero() {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		val = val.Elem()
	}
	delim := strings.Index(args[0], ":")
	if delim < 0 {
		return false
	}
	left := args[0][:delim]
	right := args[0][delim+1:]
	switch val.Kind() {
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		v := val.Float()
		min, err := strconv.ParseFloat(left, 64)
		if err != nil {
			return false
		}
		max, err := strconv.ParseFloat(right, 64)
		if err != nil {
			return false
		}
		if v >= min && max >= v {
			return true
		}
		return false
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		v := val.Int()
		min, err := strconv.ParseInt(left, 10, 64)
		if err != nil {
			return false
		}
		max, err := strconv.ParseInt(right, 10, 64)
		if err != nil {
			return false
		}
		if v >= min && v <= max {
			return true
		}
		return false
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		v := val.Uint()
		min, err := strconv.ParseUint(left, 10, 64)
		if err != nil {
			return false
		}
		max, err := strconv.ParseUint(right, 10, 64)
		if err != nil {
			return false
		}
		if v >= min && v <= max {
			return true
		}
		return false
	}
	return true
}

// IsMinValid check min
func IsMinValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		val = val.Elem()
	}
	if len(args) == 0 {
		return true
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		return false
	}
	switch val.Kind() {
	case reflect.String:
		return len([]rune(val.String())) >= min
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		return val.Float() >= float64(min)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return val.Int() >= int64(min)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return val.Uint() >= uint64(min)
	}
	return true
}

// IsMaxValid check max
func IsMaxValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		val = val.Elem()
	}
	if len(args) == 0 {
		return true
	}
	max, err := strconv.Atoi(args[0])
	if err != nil {
		return false
	}
	switch val.Kind() {
	case reflect.String:
		return len([]rune(val.String())) <= max
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		return val.Float() <= float64(max)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return val.Int() <= int64(max)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return val.Uint() <= uint64(max)
	}
	return true
}

// IsDigits check for digits
func IsDigits(val reflect.Value, args ...string) bool {
	if val.IsZero() {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return true
		}
		val = val.Elem()
	}
	var value string
	switch val.Kind() {
	case reflect.String:
		value = val.String()
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value = strconv.FormatInt(val.Int(), 10)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value = strconv.FormatUint(val.Uint(), 10)
	}

	runes := []rune(value)
	lr := len(runes)
	var l int
	for _, s := range runes {
		if s >= '0' && s <= '9' {
			l++
		}
	}
	if len(args) > 0 {
		lengths := strings.Split(args[0], ",")
		if len(lengths) == 0 && lr == l {
			return true
		}
		for _, length := range lengths {
			ll, _ := strconv.Atoi(length)
			if ll == lr {
				return true
			}
		}
		return false
	}
	return lr == l
}
