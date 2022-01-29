package helper

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/webpkg/mysql"
)

// create by td


// common ---

func AssertStringIsNotBlank(str string, errorMessage string) error {

	str = strings.Trim(str, " ")

	if str == "" {
		return errors.New(errorMessage)
	}

	return nil
}

// support type: int、int*、uint*、float*、string
func AssertNumbersMoreThenZero(number interface{}, errorMessage string) error {

	switch number.(type) {
	case int8:
		n := number.(int8)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case int16:
		n := number.(int16)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case int:
		n := number.(int)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case int32:
		n := number.(int32)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case int64:
		n := number.(int64)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case uint8:
		n := number.(uint8)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case uint16:
		n := number.(uint16)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case uint32:
		n := number.(uint32)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case uint64:
		n := number.(uint64)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case float32:
		n := number.(float32)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case float64:
		n := number.(float64)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	case string:
		s := number.(string)
		s = strings.Trim(s, " ")
		n, _ := strconv.ParseInt(s, 10, 64)
		if n <= 0 {
			return errors.New(errorMessage)
		}
	default:
		return errors.New("not support assert type")
	}

	return nil
}

// common struct example: AssertIsTrue(t != (Test{}), "message")
func AssertIsTrue(condition bool, errorMessage string) error {

	if !condition {
		return errors.New(errorMessage)
	}

	return nil
}

func AssertStructPointerIsNotNil(object interface{}, errorMessage string) error {

	rv := reflect.ValueOf(object)

	if rv.Kind() != reflect.Ptr {
		return errors.New("AssertStructPointerIsNotNil的断言参数类型必须为指针类型")
	}

	if rv.IsNil() {
		return errors.New(errorMessage)
	}

	return nil
}

func AssertSliceIsNotNil(slice interface{}, errorMessage string) error {

	rt := reflect.TypeOf(slice)

	if rt.Kind() != reflect.Slice && rt.Kind() != reflect.Array {
		return errors.New("AssertSliceIsNotNil的断言参数类型必须为Array或Slice")
	}

	rv := reflect.ValueOf(slice)

	if rv.Len() <= 0 {
		return errors.New(errorMessage)
	}

	return nil
}



// mysql error ---

type ErrorMessage struct {
	keyName, keyValue string
}

// AssertIsMySQLError
// successful: string
// failure: ""
func AssertIsMySQLError(err interface{}) string {
	fmt.Print(err)

	e, ok := err.(*mysql.MySQLError)
	if !ok {
		return ""
	}

	if e.Number == 1062 {
		em, ok := findErrorMessageByRegexp(e.Message)
		if ok {
			warpErrorMessage(e, em, "值不可重复")
		}
	}
	// order errors
	// TODO

	return e.Message
}

func warpErrorMessage(e *mysql.MySQLError, em *ErrorMessage, message string) {
	e.Message = message + "。字段名：" + em.keyName + "，字段值：" + em.keyValue
}

func findErrorMessageByRegexp(str string) (*ErrorMessage, bool) {

	compile, _ := regexp.Compile("'(\\d+)'.+'(\\w+)'")
	matchs := compile.FindStringSubmatch(str)

	if len(matchs) >= 3 {
		e := new(ErrorMessage)
		e.keyValue = matchs[1]
		e.keyName = matchs[2]

		return e, true
	}

	return nil, false
}








