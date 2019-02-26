package gtv

import (
	"reflect"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:31
**/

func Validator(table Table) (err error) {
	for _, field := range table.Fields {
		err = validator(field)
		if err != nil {
			break
		}
	}
	return
}

func validator(field IField) (err error) {
	if field != nil {
		switch reflect.TypeOf(field).String() {
		case "*gtv.ValueField":
			err = validatorValueField(field)
		case "*gtv.SliceField":
			err = validatorSliceField(field)
		case "*gtv.TableField":
			err = validatorTableField(field)
		}
	}
	return
}

func validatorSliceField(field IField) (err error) {
	sliceField := field.(*SliceField)
	sliceFieldValue := sliceField.FieldValue
	if sliceField.Validator != nil {
		err = sliceField.Validator.Validator(sliceField.GetFieldName(), sliceFieldValue)
	}
	if err == nil {
		for _, sonOfSliceField := range sliceField.GetFields() {
			switch reflect.TypeOf(sonOfSliceField).String() {
			case "*gtv.ValueField":
				err = validatorValueFieldOfSliceField(sonOfSliceField, sliceFieldValue)
			case "*gtv.TableField":
				err = validatorTableFieldOfSliceField(sonOfSliceField)
			}
		}
	}
	return
}

func validatorTableFieldOfSliceField(sonOfSliceField IField) (err error) {
	err = validatorTableField(sonOfSliceField)
	return
}

func validatorValueFieldOfSliceField(sonOfSliceField IField, sliceFieldValue []interface{}) (err error) {
	validators := sonOfSliceField.GetValidators()
	for _, value := range sliceFieldValue {
		for _, vd := range validators {
			err = vd.Validator(sonOfSliceField.GetFieldName(), value)
		}
	}
	return
}

func validatorValueField(field IField) (err error) {
	validators := field.GetValidators()
	for _, v := range validators {
		err = v.Validator(field.GetFieldName(), field.GetFieldValue())
		if err != nil {
			break
		}
	}
	return
}

func validatorTableField(field IField) (err error) {
	tableField := field.(*TableField)
	tableFieldValue := tableField.FieldValue
	for _, v := range tableField.GetFields() {
		v.SetFieldValue(tableFieldValue[v.GetFieldName()])
		err = validator(v)
		if err != nil {
			break
		}
		v.SetFieldValue(nil)
	}
	return
}
