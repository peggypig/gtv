package gtv

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
	required         bool
	minLen           int
	boolMinLen       bool
	maxLen           int
	boolMaxLen       bool
	minELen          int
	boolMinELen      bool
	maxELen          int
	boolMaxELen      bool
	simpleValidators []IValidator
}

func (sv *SliceValidator) Validator(fieldName string, value interface{}) error {
	var err error
	if value == nil {
		err = errors.New(fieldName + "'s value is nil")
	}
	var sliceFlag = false
	if intSliceValue, ok := value.([]int); err == nil && ok {
		sliceFlag = true
		if err == nil && sv.required && len(intSliceValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		}
		if err == nil && sv.boolMinLen && len(intSliceValue) <= sv.minLen {
			err = errors.New(fieldName + "'s value'len should > " + strconv.Itoa(sv.minLen))
		}
		if err == nil && sv.boolMinELen && len(intSliceValue) < sv.minELen {
			err = errors.New(fieldName + "'s value'len should >= " + strconv.Itoa(sv.minELen))
		}
		if err == nil && sv.boolMaxLen && len(intSliceValue) >= sv.maxLen {
			err = errors.New(fieldName + "'s value'len should < " + strconv.Itoa(sv.maxLen))
		}
		if err == nil && sv.boolMaxELen && len(intSliceValue) > sv.maxELen {
			err = errors.New(fieldName + "'s value'len should <= " + strconv.Itoa(sv.maxELen))
		}
		if err == nil {
			for _, val := range intSliceValue {
				for _, vv := range sv.simpleValidators {
					err = vv.Validator(fieldName, val)
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
			}
		}
	}
	if strSliceValue, ok := value.([]string); err == nil && ok {
		sliceFlag = true
		if err == nil && sv.required && len(strSliceValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		}
		if err == nil && sv.boolMinLen && len(strSliceValue) <= sv.minLen {
			err = errors.New(fieldName + "'s value'len should > " + strconv.Itoa(sv.minLen))
		}
		if err == nil && sv.boolMinELen && len(strSliceValue) < sv.minELen {
			err = errors.New(fieldName + "'s value'len should >= " + strconv.Itoa(sv.minELen))
		}
		if err == nil && sv.boolMaxLen && len(strSliceValue) >= sv.maxLen {
			err = errors.New(fieldName + "'s value'len should < " + strconv.Itoa(sv.maxLen))
		}
		if err == nil && sv.boolMaxELen && len(strSliceValue) > sv.maxELen {
			err = errors.New(fieldName + "'s value'len should <= " + strconv.Itoa(sv.maxELen))
		}
		if err == nil {
			for _, val := range strSliceValue {
				for _, vv := range sv.simpleValidators {
					err = vv.Validator(fieldName, val)
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
			}
		}
	}
	if float64SliceValue, ok := value.([]float64); err == nil && ok {
		sliceFlag = true
		if err == nil && sv.required && len(float64SliceValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		}
		if err == nil && sv.boolMinLen && len(float64SliceValue) <= sv.minLen {
			err = errors.New(fieldName + "'s value'len should > " + strconv.Itoa(sv.minLen))
		}
		if err == nil && sv.boolMinELen && len(float64SliceValue) < sv.minELen {
			err = errors.New(fieldName + "'s value'len should >= " + strconv.Itoa(sv.minELen))
		}
		if err == nil && sv.boolMaxLen && len(float64SliceValue) >= sv.maxLen {
			err = errors.New(fieldName + "'s value'len should < " + strconv.Itoa(sv.maxLen))
		}
		if err == nil && sv.boolMaxELen && len(float64SliceValue) > sv.maxELen {
			err = errors.New(fieldName + "'s value'len should <= " + strconv.Itoa(sv.maxELen))
		}
		if err == nil {
			for _, val := range float64SliceValue {
				for _, vv := range sv.simpleValidators {
					err = vv.Validator(fieldName, val)
					if err != nil {
						break
					}
				}
				if err != nil {
					break
				}
			}
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

func (sv *SliceValidator) Min(min int) *SliceValidator {
	sv.minLen = min
	sv.boolMinLen = true
	return sv
}

func (sv *SliceValidator) MinE(minE int) *SliceValidator {
	sv.minELen = minE
	sv.boolMinELen = true
	return sv
}

func (sv *SliceValidator) Max(max int) *SliceValidator {
	sv.maxLen = max
	sv.boolMaxLen = true
	return sv
}

func (sv *SliceValidator) MaxE(maxE int) *SliceValidator {
	sv.maxELen = maxE
	sv.boolMaxELen = true
	return sv
}

func (sv *SliceValidator) SetSimpleValidators(validators []IValidator) {
	sv.simpleValidators = validators
}
