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
* @create : 2019-02-25 18:59
**/
type IntValidator struct {
	min      int
	boolMin  bool
	minE     int
	boolMinE bool
	max      int
	boolMax  bool
	maxE     int
	boolMaxE bool
}

func (iv *IntValidator) Validator(fieldName string, value interface{}) error {
	var err error
	if intValue, ok := value.(int); ok {
		if iv.boolMin && intValue <= iv.min {
			err = errors.New(fieldName + "'s value should > " + strconv.Itoa(iv.min))
		}
		if err == nil && iv.boolMinE && intValue < iv.minE {
			err = errors.New(fieldName + "'s value should >= " + strconv.Itoa(iv.min))
		}
		if err == nil && iv.boolMax && intValue >= iv.max {
			err = errors.New(fieldName + "'s value should < " + strconv.Itoa(iv.max))
		}
		if err == nil && iv.boolMaxE && intValue > iv.maxE {
			err = errors.New(fieldName + "'s value should <= " + strconv.Itoa(iv.maxE))
		}
	} else {
		err = errors.New("data type is not int:" + fieldName)
	}
	return err
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
