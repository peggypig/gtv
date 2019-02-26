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
	//IsTableField() bool
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

//func (field *ValueField) IsTableField() bool {
//	return false
//}

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
	FieldValue interface{}
	Fields     []IField
}

//func (field *TableField) IsTableField() bool {
//	return true
//}

func (field *TableField) GetFieldName() string {
	return field.FieldName
}

func (field *TableField) GetFieldValue() interface{} {
	return field.GetFieldValue()
}

func (field *TableField) SetFieldValue(v interface{}) {
	field.FieldValue = v
}

func (field *TableField) GetValidators() []IValidator {
	return nil
}

func (field *TableField) GetFields() []IField {
	return field.Fields
}

type SliceField struct {
	FieldName string
	Field     IField
}

func (field *SliceField) GetFieldName() string {
	return field.FieldName
}

func (field *SliceField) GetFieldValue() interface{} {
	return nil
}

func (field *SliceField) SetFieldValue(v interface{}) {
}

func (field *SliceField) GetValidators() []IValidator {
	return nil
}

func (field *SliceField) GetFields() []IField {
	return []IField{field.Field}
}
