package gtv

import "github.com/peggypig/gtv/validators"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:34
**/
type IField interface {
	GetFieldName() string
	GetFieldValue() interface{}
	SetFieldValue(interface{})
	GetValidators() []IValidator
	GetFields() []IField
}

type ValueField struct {
	FieldName  string
	FieldValue interface{}
	Validators []IValidator
}

func (field *ValueField) GetFieldName() string {
	return field.FieldName
}

func (field *ValueField) GetFieldValue() interface{} {
	return field.FieldValue
}

func (field *ValueField) SetFieldValue(v interface{}) {
	field.FieldValue = v
}

func (field *ValueField) GetValidators() []IValidator {
	return field.Validators
}

func (field *ValueField) GetFields() []IField {
	return nil
}

type TableField struct {
	FieldName  string
	FieldValue map[string]interface{}
	Fields     []IField
}

func (field *TableField) GetFieldName() string {
	return field.FieldName
}

func (field *TableField) GetFieldValue() interface{} {
	return field.GetFieldValue()
}

func (field *TableField) SetFieldValue(v interface{}) {
	if field.FieldValue == nil {
		field.FieldValue = make(map[string]interface{})
	}
	if mv, ok := v.(map[string]interface{}); ok {
		for k, v := range mv {
			field.FieldValue[k] = v
		}
	}
}

func (field *TableField) GetValidators() []IValidator {
	return nil
}

func (field *TableField) GetFields() []IField {
	return field.Fields
}

type SliceField struct {
	FieldName  string
	FieldValue []interface{}
	Field      IField
	Validator  *validators.SliceValidator
}

func (field *SliceField) GetFieldName() string {
	return field.FieldName
}

func (field *SliceField) GetFieldValue() interface{} {
	return nil
}

func (field *SliceField) SetFieldValue(v interface{}) {
	if sv, ok := v.([]interface{}); ok {
		field.FieldValue = sv
	} else {
		field.FieldValue = append(field.FieldValue, v)
	}
}

func (field *SliceField) GetValidators() []IValidator {
	return nil
}

func (field *SliceField) GetFields() []IField {
	return []IField{field.Field}
}
