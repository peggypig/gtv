package validators

import (
	"github.com/peggypig/gtv/gerror"
	"regexp"
	"strconv"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 19:01
**/

type StringValidator struct {
	required    bool
	minLen      int
	boolMinLen  bool
	minELen     int
	boolMinELen bool
	maxLen      int
	boolMaxLen  bool
	maxELen     int
	boolMaxELen bool
	regexp      string
	boolRegexp  bool
}

func (sv *StringValidator) Validator(fieldName string, value interface{})  *gerror.GError {
	var err = &gerror.GError{
		Key:   fieldName,
		Value: value,
	}
	if value == nil {
		err.Msg = "value is nil"
	}
	dataTypeFlag := false
	if err.IsNull() {
		if strSliceValue, ok := value.([]string); ok {
			dataTypeFlag = true
			value = strSliceValue[0]
		}
	}
	if err.IsNull() {
		if intSliceValue, ok := value.([]int); ok {
			dataTypeFlag = true
			value = strconv.Itoa(intSliceValue[0])
		}
	}

	if err.IsNull() {
		if strValue, ok := value.(string); ok {
			dataTypeFlag = true
			if sv.required && len(strValue) <= 0 {
				err.Msg = "value is required"
			}
			if err.IsNull() && sv.boolMinELen && len(strValue) < sv.minELen {
				err.Msg = "value's len should >= " + strconv.Itoa(sv.minELen)
			}
			if err.IsNull() && sv.boolMinLen && len(strValue) <= sv.minELen {
				err.Msg = "value's len should > " + strconv.Itoa(sv.minLen)
			}
			if err.IsNull() && sv.boolMaxELen && len(strValue) > sv.maxELen {
				err.Msg = "value's len should <= " + strconv.Itoa(sv.maxELen)
			}
			if err.IsNull() && sv.boolMaxLen && len(strValue) >= sv.maxLen {
				err.Msg = "value's len should < " + strconv.Itoa(sv.maxLen)
			}
			if err.IsNull() && sv.boolRegexp {
				compile, errCompile := regexp.Compile(sv.regexp)
				if errCompile != nil {
					err.Msg = errCompile.Error()
				}
				if err.IsNull() && !compile.MatchString(strValue) {
					err.Msg = "value not match rule"
				}
			}
		}
	}

	if !dataTypeFlag && err.IsNull() {
		err.Msg = "data type is not string"
	}
	if err.IsNull() {
		err = nil
	}
	return err
}

func (sv *StringValidator) Required() *StringValidator {
	sv.required = true
	return sv
}

func (sv *StringValidator) MinLen(minLen int) *StringValidator {
	sv.minLen = minLen
	sv.boolMinLen = true
	return sv
}

func (sv *StringValidator) MinELen(minELen int) *StringValidator {
	sv.minELen = minELen
	sv.boolMinELen = true
	return sv
}

func (sv *StringValidator) MaxLen(maxLen int) *StringValidator {
	sv.maxLen = maxLen
	sv.boolMaxLen = true
	return sv
}

func (sv *StringValidator) MaxELen(maxELen int) *StringValidator {
	sv.maxELen = maxELen
	sv.boolMaxELen = true
	return sv
}

func (sv *StringValidator) Regexp(regexp string) *StringValidator {
	sv.regexp = regexp
	sv.boolRegexp = true
	return sv
}
