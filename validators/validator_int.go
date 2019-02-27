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
* @create : 2019-02-25 18:59
**/
type IntValidator struct {
	required bool
	min      int
	boolMin  bool
	minE     int
	boolMinE bool
	max      int
	boolMax  bool
	maxE     int
	boolMaxE bool
}

func (iv *IntValidator) Validator(fieldName string, value interface{})  *gerror.GError {
	var err  = &gerror.GError{
		Key:   fieldName,
		Value: value,
	}
	if value == nil {
		err.Msg = "value is nil"
	}
	dataTypeFlag := false
	if err.IsNull() {
		if intSliceValue, ok := value.([]int); ok {
			dataTypeFlag = true
			value = intSliceValue[0]
		}
	}
	if err.IsNull() {
		if strSliceValue, ok := value.([]string); ok {
			value = strSliceValue[0]
		}
	}

	if err.IsNull() {
		if strValue, ok := value.(string); ok {
			if iv.required && len(strValue) <= 0 {
				err.Msg = "value is required"
			} else {
				valueTemp, errTemp := strconv.Atoi(strValue)
				if errTemp != nil {
					err.Msg = "data type is not int"
				} else {
					dataTypeFlag = true
					value = valueTemp
				}
			}
		}
	}

	// json中只有float64
	if err.IsNull() {
		if float64Value, ok := value.(float64); ok {
			dataTypeFlag = true
			value = int(float64Value)
		}
	}

	if err.IsNull() {
		if intValue, ok := value.(int); ok {
			dataTypeFlag = true
			if iv.boolMin && intValue <= iv.min {
				err.Msg = "value should > " + strconv.Itoa(iv.min)
			}
			if err.IsNull() && iv.boolMinE && intValue < iv.minE {
				err.Msg = "value should >= " + strconv.Itoa(iv.minE)
			}
			if err.IsNull() && iv.boolMax && intValue >= iv.max {
				err.Msg = "value should < " + strconv.Itoa(iv.max)
			}
			if err.IsNull() && iv.boolMaxE && intValue > iv.maxE {
				err.Msg = "value should <= " + strconv.Itoa(iv.maxE)
			}
		}
	}
	if !dataTypeFlag && err.IsNull() {
		err.Msg = "data type is not int"
	}
	if err.IsNull() {
		err = nil
	}
	return err
}

func (iv *IntValidator) Required() *IntValidator {
	iv.required = true
	return iv
}

func (iv *IntValidator) Min(min int) *IntValidator {
	iv.min = min
	iv.boolMin = true
	return iv
}

func (iv *IntValidator) MinE(minE int) *IntValidator {
	iv.minE = minE
	iv.boolMinE = true
	return iv
}

func (iv *IntValidator) Max(max int) *IntValidator {
	iv.max = max
	iv.boolMax = true
	return iv
}

func (iv *IntValidator) MaxE(maxE int) *IntValidator {
	iv.maxE = maxE
	iv.boolMaxE = true
	return iv
}
