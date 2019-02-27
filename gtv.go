package gtv

import (
	"github.com/peggypig/gtv/gerror"
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

func Validator(table Table) (err *gerror.GError) {
	for _, field := range table.Fields {
		err = validator(field)
		if !reflect.ValueOf(err).IsNil() {
			break
		}
	}
	return
}

func validator(field IField) (err  *gerror.GError) {
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

func validatorSliceField(field IField) (err  *gerror.GError) {
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

func validatorTableFieldOfSliceField(sonOfSliceField IField) (err  *gerror.GError) {
	err = validatorTableField(sonOfSliceField)
	return
}

func validatorValueFieldOfSliceField(sonOfSliceField IField, sliceFieldValue []interface{}) (err  *gerror.GError) {
	validators := sonOfSliceField.GetValidators()
	for _, value := range sliceFieldValue {
		for _, vd := range validators {
			err = vd.Validator(sonOfSliceField.GetFieldName(), value)
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}
	}
	return
}

func validatorValueField(field IField) (err  *gerror.GError) {
	validators := field.GetValidators()
	for _, v := range validators {
		err = v.Validator(field.GetFieldName(), field.GetFieldValue())
		if err != nil {
			break
		}
	}
	return
}

func validatorTableField(field IField) (err  *gerror.GError) {
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
