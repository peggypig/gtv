package gtv

import (
	"errors"
	"net/http"
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

func (table *Table) FillTable(request *http.Request) (err error,
	requestValues map[string]string) {
	if request != nil {
		err = request.ParseForm()
	} else {
		err = errors.New("request is nil")
	}
	if err == nil {
		requestValues = make(map[string]string)
		for _, field := range table.Fields {
			fillTable(field, request.Form.Get(field.GetFieldName()), requestValues)
		}
	}
	return
}

func fillTable(field IField, value string,
	requestValues map[string]string) (ok bool) {
	if field != nil {
		if field.IsTableField() {
			for _, v := range field.GetFields() {
				if fillTable(v, value, requestValues) {
					break
				}
			}
		} else {
			field.SetFieldValue(value)
			requestValues[field.GetFieldName()] = value
			ok = true
		}
	}
	return
}
