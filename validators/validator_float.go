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
* @create : 2019-02-26 10:20
**/

type Float64Validator struct {
	required bool
	min      float64
	boolMin  bool
	minE     float64
	boolMinE bool
	max      float64
	boolMax  bool
	maxE     float64
	boolMaxE bool
}

func (fv *Float64Validator) Validator(fieldName string, value interface{}) error {
	var err error
	if strValue, ok := value.(string); ok {
		if fv.required && len(strValue) <= 0 {
			err = errors.New(fieldName + "'s value is required")
		} else {
			value, err = strconv.ParseFloat(strValue, 64)
		}
	}
	if float64Value, ok := value.(float64); err == nil && ok {
		if fv.boolMin && float64Value <= fv.min {
			err = errors.New(fieldName + "'s value should > " + strconv.FormatFloat(fv.min, 'f', -1, 64))
		}
		if err == nil && fv.boolMinE && float64Value < fv.minE {
			err = errors.New(fieldName + "'s value should >= " + strconv.FormatFloat(fv.minE, 'f', -1, 64))
		}
		if err == nil && fv.boolMax && float64Value >= fv.max {
			err = errors.New(fieldName + "'s value should < " + strconv.FormatFloat(fv.max, 'f', -1, 64))
		}
		if err == nil && fv.boolMaxE && float64Value > fv.maxE {
			err = errors.New(fieldName + "'s value should <= " + strconv.FormatFloat(fv.maxE, 'f', -1, 64))
		}
	} else {
		err = errors.New("data type is not float64:" + fieldName)
	}
	return err
}

func (fv *Float64Validator) Required() *Float64Validator {
	fv.required = true
	return fv
}

func (fv *Float64Validator) Min(min float64) *Float64Validator {
	fv.min = min
	fv.boolMin = true
	return fv
}

func (fv *Float64Validator) MinE(minE float64) *Float64Validator {
	fv.minE = minE
	fv.boolMinE = true
	return fv
}

func (fv *Float64Validator) Max(max float64) *Float64Validator {
	fv.max = max
	fv.boolMax = true
	return fv
}

func (fv *Float64Validator) MaxE(maxE float64) *Float64Validator {
	fv.maxE = maxE
	fv.boolMaxE = true
	return fv
}
