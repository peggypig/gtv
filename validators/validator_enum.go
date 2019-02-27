package validators

import (
	"github.com/peggypig/gtv/gerror"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-27 13:36
**/
type EnumValidator struct {
	required   bool
	caseCap    bool
	enumValues interface{}
}

func (ev *EnumValidator) Validator(fieldName string, value interface{}) *gerror.GError {
	var err = &gerror.GError{
		Key:   fieldName,
		Value: value,
	}
	if value == nil {
		err.Msg = "value is nil"
	}
	if ev.enumValues == nil {
		err.Msg = "enum values is nil"
	}
	if err.IsNull() {
		if strValue, ok := value.(string); ok {
			ev.validatorStrEnum(strValue, ev.enumValues, err)
		}
		if intValue, ok := value.(int); ok {
			ev.validatorIntEnum(intValue, ev.enumValues, err)
		}
	}
	if err.IsNull() {
		err = nil
	}
	return err
}

func (ev *EnumValidator) validatorIntEnum(i int, enumValues interface{}, err *gerror.GError) {
	if intEnum, ok := enumValues.([]int); ok {
		valid := false
		for _, ev := range intEnum {
			if ev == i {
				valid = true
				break
			}
		}
		if !valid {
			err.Msg = "value is not exist in enum value"
		}
	} else {
		err.Msg = "enum value is not []int"
	}
}

func (ev *EnumValidator) validatorStrEnum(s string, enumValues interface{}, err *gerror.GError) {
	if strEnum, ok := enumValues.([]string); ok {
		if !ev.caseCap {
			s = strings.ToLower(s)
		}
		valid := false
		for _, ev := range strEnum {
			if strings.ToLower(ev) == s {
				valid = true
				break
			}
		}
		if !valid {
			err.Msg = "value is not exist in enum value"
		}
	} else {
		err.Msg = "enum value is not []string"
	}
}

func (ev *EnumValidator) Enum(enumValue interface{}) *EnumValidator {
	ev.enumValues = enumValue
	return ev
}

func (ev *EnumValidator) CaseCap(caseCap bool) *EnumValidator {
	ev.caseCap = caseCap
	return ev
}
