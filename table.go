package gtv

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:32
**/

type Table struct {
	Fields []IField
}

func (table *Table) FillTable(request *http.Request) (
	requestValues map[string]interface{}, err error) {
	params, errGet := getRequestParam(request)
	if errGet != nil {
		err = errGet
	}
	if err == nil {
		requestValues = make(map[string]interface{})
		for _, field := range table.Fields {
			fillTable(nil, field, params[field.GetFieldName()], requestValues)
		}
	} else {
		err = errGet
	}
	return
}

func fillTable(parentField IField, field IField, value interface{},
	requestValues map[string]interface{}) {
	if field != nil {
		switch reflect.TypeOf(field).String() {
		case "*gtv.ValueField":
			handlerValueField(field, value, parentField, requestValues)
		case "*gtv.TableField":
			handlerTableField(field, value, parentField, requestValues)
		case "*gtv.SliceField":
			handlerSliceField(field, value, requestValues)
		}
	}
	return
}

func handlerSliceField(field IField, value interface{}, requestValues map[string]interface{}) {
	field.SetFieldValue(value)
	sliceField := field.(*SliceField)
	for _, fieldValue := range sliceField.GetFields() {
		switch reflect.TypeOf(fieldValue).String() {
		case "*gtv.ValueField":
			handlerValueFieldOfSliceField(fieldValue, value, requestValues, sliceField)
		case "*gtv.TableField":
			handlerTableFieldOfSliceField(fieldValue, value, requestValues, sliceField)
		}
	}
}

func handlerValueField(field IField, value interface{}, parentField IField, requestValues map[string]interface{}) {
	if parentField == nil {
		field.SetFieldValue(value)
		requestValues[field.GetFieldName()] = value
	} else {
		if _, ok := requestValues[parentField.GetFieldName()].(map[string]interface{}); ok {
			requestValues[parentField.GetFieldName()].(map[string]interface{})[field.GetFieldName()] = value
		}
		if _, ok := requestValues[parentField.GetFieldName()].([]map[string]interface{}); ok {
			handlerSliceMapValueOfValueField(requestValues, parentField, field, value)
		}
		if _, ok := requestValues[parentField.GetFieldName()].([]interface{}); ok {
			requestValues[parentField.GetFieldName()] =
				append(requestValues[parentField.GetFieldName()].([]interface{}), value)
		}
	}
}

func handlerTableField(field IField, value interface{}, parentField IField, requestValues map[string]interface{}) {
	if parentField == nil {
		field.SetFieldValue(value)
	}
	tableField := field.(*TableField)
	for _, fieldValue := range tableField.GetFields() {
		if valueValue, ok := value.(map[string]interface{}); ok {
			handlerMapValueOfTableField(parentField, requestValues, tableField, fieldValue, valueValue)
		}
	}
}

func handlerSliceMapValueOfValueField(requestValues map[string]interface{}, parentField IField, field IField, value interface{}) {
	temp := requestValues[parentField.GetFieldName()].([]map[string]interface{})
	if len(temp) == 0 {
		temp = append(temp, map[string]interface{}{})
	}
	var last = 0
	if len(temp) > 0 {
		last = len(temp) - 1
	}
	keyExist := false
	for k, _ := range temp[last] {
		if k == field.GetFieldName() {
			keyExist = true
			break
		}
	}
	// 值存在 则新增一个map[string]interface{}{}
	if keyExist {
		temp = append(temp, map[string]interface{}{
			field.GetFieldName(): value,
		})
	} else { // 值不存在
		temp[last][field.GetFieldName()] = value
	}
	requestValues[parentField.GetFieldName()] = temp
}

func handlerMapValueOfTableField(parentField IField, requestValues map[string]interface{}, tableField *TableField, fieldValue IField, valueValue map[string]interface{}) {
	if parentField == nil {
		if requestValues[tableField.GetFieldName()] == nil {
			requestValues[tableField.GetFieldName()] = make(map[string]interface{})
		}
		fillTable(tableField, fieldValue, valueValue[fieldValue.GetFieldName()], requestValues)
	} else {
		if _, ok := parentField.(*SliceField); ok {
			fillTable(parentField, fieldValue, valueValue[fieldValue.GetFieldName()], requestValues)
		} else {
			fillTable(tableField, fieldValue, valueValue[fieldValue.GetFieldName()], requestValues)
		}
	}
}

func handlerTableFieldOfSliceField(fieldValue IField, value interface{}, requestValues map[string]interface{}, sliceField *SliceField) {
	tableField := fieldValue.(*TableField)
	if valueValue, ok := value.([]interface{}); ok {
		if requestValues[sliceField.GetFieldName()] == nil {
			requestValues[sliceField.GetFieldName()] = make([]map[string]interface{}, 0)
		}
		for _, vv := range valueValue {
			fillTable(sliceField, tableField, vv, requestValues)
		}
	}
}

func handlerValueFieldOfSliceField(fieldValue IField, value interface{}, requestValues map[string]interface{}, sliceField *SliceField) {
	valueField := fieldValue.(*ValueField)
	if valueValue, ok := value.([]interface{}); ok {
		if requestValues[sliceField.GetFieldName()] == nil {
			requestValues[sliceField.GetFieldName()] = make([]interface{}, 0)
		}
		for _, vv := range valueValue {
			fillTable(sliceField, valueField, vv, requestValues)
		}
	}
}

func getRequestParam(request *http.Request) (param map[string]interface{}, err error) {
	if request != nil {
		param = make(map[string]interface{})
		contentType := strings.ToLower(request.Header.Get("Content-Type"))
		switch {
		case strings.Contains(contentType, "application/json"):
			param, err = getJsonParam(request)
		case strings.Contains(contentType, "application/x-www-form-urlencoded"):
			param, err = getFormAndUrlParam(request)
		}

	} else {
		err = errors.New("request is nil")
	}
	return
}

func getFormAndUrlParam(request *http.Request) (param map[string]interface{}, err error) {
	err = request.ParseForm()
	param = make(map[string]interface{})
	for k, _ := range request.Form {
		param[k] = request.Form[k]
	}
	return
}

func getJsonParam(request *http.Request) (param map[string]interface{}, err error) {
	bytes, errRead := ioutil.ReadAll(request.Body)
	if errRead != nil {
		err = errRead
	} else {
		errJson := json.Unmarshal(bytes, &param)
		if errJson != nil {
			err = errJson
			param = nil
		}
	}
	return
}
