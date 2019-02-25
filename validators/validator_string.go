package validators

import (
	"errors"
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

func (sv *StringValidator) Validator(fieldName string, value interface{}) error {
	var err error
	if strValue, ok := value.(string); ok {
		if sv.required && len(strValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		}
		if err == nil && sv.boolMinELen && len(strValue) < sv.minELen {
			err = errors.New(fieldName + "'s value's len should >= " + strconv.Itoa(sv.minELen))
		}
		if err == nil && sv.boolMinLen && len(strValue) <= sv.minELen {
			err = errors.New(fieldName + "'s value's len should > " + strconv.Itoa(sv.minLen))
		}
		if err == nil && sv.boolMaxELen && len(strValue) > sv.maxELen {
			err = errors.New(fieldName + "'s value's len should <= " + strconv.Itoa(sv.maxELen))
		}
		if err == nil && sv.boolMaxLen && len(strValue) >= sv.maxLen {
			err = errors.New(fieldName + "'s value's len should < " + strconv.Itoa(sv.maxLen))
		}
		if err == nil && sv.boolRegexp {
			compile, errCompile := regexp.Compile(sv.regexp)
			if errCompile != nil {
				err = errCompile
			}
			if err == nil && !compile.MatchString(strValue) {
				err = errors.New(fieldName + "'s value is not match rule")
			}
		}
	} else {
		err = errors.New("data type is not string:" + fieldName)
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
