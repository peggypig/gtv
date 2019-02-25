package gtv

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:34
**/
type IField interface {
	IsTableField() bool
	GetFieldName() string
	GetFieldValue() interface{}
	GetValidators() []IValidator
	GetFields() []IField
}

type ValueField struct {
	FieldName  string
	FieldValue interface{}
	Validators []IValidator
}

func (field *ValueField) IsTableField() bool {
	return false
}

func (field *ValueField) GetFieldName() string {
	return field.FieldName
}

func (field *ValueField) GetFieldValue() interface{} {
	return field.FieldValue
}

func (field *ValueField) GetValidators() []IValidator {
	return field.Validators
}

func (field *ValueField) GetFields() []IField {
	return nil
}

type TableField struct {
	Fields []IField
}

func (field *TableField) IsTableField() bool {
	return true
}

func (field *TableField) GetFieldName() string {
	return ""
}

func (field *TableField) GetFieldValue() interface{} {
	return nil
}

func (field *TableField) GetValidators() []IValidator {
	return nil
}

func (field *TableField) GetFields() []IField {
	return field.Fields
}
