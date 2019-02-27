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

func (fv *Float64Validator) Validator(fieldName string, value interface{})  *gerror.GError {
	var err = &gerror.GError{
		Value: value,
		Key:   fieldName,
	}
	if value == nil {
		err.Msg = "value is nil"
	}
	dataTypeFlag := false
	if err.IsNull() {
		if float64SliceValue, ok := value.([]float64); ok {
			dataTypeFlag = true
			value = float64SliceValue[0]
		}
	}
	if err.IsNull() {
		if strSliceValue, ok := value.([]string); ok {
			value = strSliceValue[0]
		}
	}
	if err.IsNull() {
		if strValue, ok := value.(string); ok {
			if fv.required && len(strValue) <= 0 {
				err.Msg = "value is required"
			} else {
				valueTemp, errTemp := strconv.ParseFloat(strValue, 64)
				if errTemp != nil {
					err.Msg = "data type is not float64"
				} else {
					dataTypeFlag = true
					value = valueTemp
				}
			}
		}
	}

	if err.IsNull() {
		if float64Value, ok := value.(float64); ok {
			dataTypeFlag = true
			if fv.boolMin && float64Value <= fv.min {
				err.Msg = "value should > " + strconv.FormatFloat(fv.min, 'f', -1, 64)
			}
			if err.IsNull() && fv.boolMinE && float64Value < fv.minE {
				err.Msg = "value should >= " + strconv.FormatFloat(fv.minE, 'f', -1, 64)
			}
			if err.IsNull() && fv.boolMax && float64Value >= fv.max {
				err.Msg = "value should < " + strconv.FormatFloat(fv.max, 'f', -1, 64)
			}
			if err.IsNull() && fv.boolMaxE && float64Value > fv.maxE {
				err.Msg = "value should <= " + strconv.FormatFloat(fv.maxE, 'f', -1, 64)
			}
		}
	}

	if !dataTypeFlag && err.IsNull() {
		err.Msg = "data type is not float64"
	}
	if err.IsNull() {
		err = nil
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
