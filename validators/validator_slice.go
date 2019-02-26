package validators

import (
	"errors"
	"strconv"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-26 16:32
**/

type SliceValidator struct {
	required    bool
	minLen      int
	boolMinLen  bool
	maxLen      int
	boolMaxLen  bool
	minELen     int
	boolMinELen bool
	maxELen     int
	boolMaxELen bool
}

func (sv *SliceValidator) Validator(fieldName string, value interface{}) error {
	var err error
	if value == nil {
		err = errors.New(fieldName + "'s value is nil")
	}
	var sliceFlag = false
	if sliceValue, ok := value.([]interface{}); err == nil && ok {
		sliceFlag = true
		if err == nil && sv.required && len(sliceValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		}
		if err == nil && sv.boolMinLen && len(sliceValue) <= sv.minLen {
			err = errors.New(fieldName + "'s value'len should > " + strconv.Itoa(sv.minLen))
		}
		if err == nil && sv.boolMinELen && len(sliceValue) < sv.minELen {
			err = errors.New(fieldName + "'s value'len should >= " + strconv.Itoa(sv.minELen))
		}
		if err == nil && sv.boolMaxLen && len(sliceValue) >= sv.maxLen {
			err = errors.New(fieldName + "'s value'len should < " + strconv.Itoa(sv.maxLen))
		}
		if err == nil && sv.boolMaxELen && len(sliceValue) > sv.maxELen {
			err = errors.New(fieldName + "'s value'len should <= " + strconv.Itoa(sv.maxELen))
		}
	}
	if !sliceFlag {
		err = errors.New("data type is not slice:" + fieldName)
	}
	return err
}

func (sv *SliceValidator) Required() *SliceValidator {
	sv.required = true
	return sv
}

func (sv *SliceValidator) MinLen(min int) *SliceValidator {
	sv.minLen = min
	sv.boolMinLen = true
	return sv
}

func (sv *SliceValidator) MinELen(minE int) *SliceValidator {
	sv.minELen = minE
	sv.boolMinELen = true
	return sv
}

func (sv *SliceValidator) MaxLen(max int) *SliceValidator {
	sv.maxLen = max
	sv.boolMaxLen = true
	return sv
}

func (sv *SliceValidator) MaxELen(maxE int) *SliceValidator {
	sv.maxELen = maxE
	sv.boolMaxELen = true
	return sv
}
