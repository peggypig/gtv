package validators

import (
	"github.com/peggypig/gtv/gerror"
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

func (sv *SliceValidator) Validator(fieldName string, value interface{})  *gerror.GError {
	var err = &gerror.GError{
		Key:   fieldName,
		Value: value,
	}
	if value == nil {
		err.Msg = "value is nil"
	}
	var sliceFlag = false
	if err.IsNull() {
		if sliceValue, ok := value.([]interface{}); ok {
			sliceFlag = true
			if err.IsNull() && sv.required && len(sliceValue) <= 0 {
				err.Msg = "value is required"
			}
			if err.IsNull() && sv.boolMinLen && len(sliceValue) <= sv.minLen {
				err.Msg = "value's len should > " + strconv.Itoa(sv.minLen)
			}
			if err.IsNull() && sv.boolMinELen && len(sliceValue) < sv.minELen {
				err.Msg = "value's len should >= " + strconv.Itoa(sv.minELen)
			}
			if err.IsNull() && sv.boolMaxLen && len(sliceValue) >= sv.maxLen {
				err.Msg = "value's len should < " + strconv.Itoa(sv.maxLen)
			}
			if err.IsNull() && sv.boolMaxELen && len(sliceValue) > sv.maxELen {
				err.Msg = "value's len should <= " + strconv.Itoa(sv.maxELen)
			}
		}
	}
	if !sliceFlag {
		err.Msg = "data type is not slice"
	}
	if err.IsNull() {
		err = nil
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
