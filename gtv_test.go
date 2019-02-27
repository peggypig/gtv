package gtv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 19:38
**/

func TestValidator(t *testing.T) {

	err := Validator(Table{
		[]IField{
			&ValueField{
				FieldName:  "Name",
				FieldValue: "zhangSan",
				Validators: []IValidator{
					NewStringValidator().Required(),
				},
			},
			&ValueField{
				FieldName:  "Alias",
				FieldValue: "zhangSan",
				Validators: nil,
			},
			&ValueField{
				FieldName:  "Age",
				FieldValue: 10,
				Validators: []IValidator{
					NewIntValidator().MaxE(10),
				},
			},
			&TableField{
				Fields: []IField{
					&ValueField{
						FieldName:  "Like",
						FieldValue: "ping pang",
						Validators: []IValidator{
							NewStringValidator().MaxLen(5),
						},
					},
					&TableField{
						Fields: []IField{
							&ValueField{
								FieldName:  "LikeType",
								FieldValue: "desktop",
								Validators: []IValidator{
									NewStringValidator().MaxLen(5),
								},
							},
						},
					},
				},
			},
		},
	})
	fmt.Println(err)
	assert.NotNil(t, err)
}

type fakeHttpBody struct {
	body io.ReadSeeker
}

func (body *fakeHttpBody) Read(p []byte) (n int, err error) {
	n, err = body.body.Read(p)
	if err == io.EOF {
		body.body.Seek(0, 0)
	}
	return n, err
}

func (body *fakeHttpBody) Close() error {
	return nil
}
func TestTable_FillTable_Body(t *testing.T) {
	request := &http.Request{
		Body:   &fakeHttpBody{bytes.NewReader([]byte(`Name=zhangSan&Age=12`))},
		Method: "POST",
		Header: map[string][]string{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		},
	}
	table := Table{
		Fields: []IField{
			&ValueField{
				FieldName: "Name",
				Validators: []IValidator{
					NewStringValidator().Required(),
				},
			},
			&ValueField{
				FieldName: "Age",
				Validators: []IValidator{
					NewIntValidator().Max(20),
				},
			},
		},
	}
	requestValues, err := table.FillTable(request)
	err = Validator(table)
	fmt.Println(err)
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"Name": []string{"zhangSan"},
		"Age":  []string{"12"},
	}, requestValues)
}

func TestTable_FillTable_JSON(t *testing.T) {
	request := &http.Request{
		Body: &fakeHttpBody{bytes.NewReader([]byte(`{
"Name":"zhangSan",
"Age":10,
"Class":{
	"ClassNo":12,
	"ClassName":"class"
},
"Likes":[
"Ping",
"Dance"
],
"Phones":[
	{
	"PhoneNum":"110",
	"PhoneText":"110"
	},
	{
	"PhoneNum":"119",
	"PhoneText":"119"
	}
]
}`))},
		Method: "POST",
		Header: map[string][]string{
			"Content-Type": []string{"application/json"},
		},
	}
	table := Table{
		Fields: []IField{
			&ValueField{
				FieldName: "Name",
				Validators: []IValidator{
					NewStringValidator().Required(),
				},
			},
			&ValueField{
				FieldName: "Age",
				Validators: []IValidator{
					NewIntValidator().Max(20),
				},
			},
			&SliceField{
				FieldName: "Likes",
				Field: &ValueField{
					FieldName: "Likes",
					Validators: []IValidator{
						NewStringValidator().MaxELen(5),
					},
				},
			},
			&SliceField{
				FieldName: "Phones",
				Field: &TableField{
					FieldName: "Phone",
					Fields: []IField{
						&ValueField{
							FieldName: "PhoneNum",
						},
						&ValueField{
							FieldName: "PhoneText",
						},
					},
				},
				Validator:NewSliceValidator().MaxLen(3),
			},
			&TableField{
				FieldName: "Class",
				Fields: []IField{
					&ValueField{
						FieldName: "ClassName",
					},
					&ValueField{
						FieldName: "ClassNo",
						Validators: []IValidator{
							NewIntValidator().Max(20).Min(12),
						},
					},
				},
			},
		},
	}
	requestValues, err := table.FillTable(request)
	err = Validator(table)
	assert.Nil(t, err)
	target := map[string]interface{}{
		"Name": "zhangSan",
		"Age":  10,
		"Class": map[string]interface{}{
			"ClassNo":   12,
			"ClassName": "class",
		},
		"Likes": []interface{}{
			"Ping",
			"Dance",
		},
		"Phones": []map[string]interface{}{
			map[string]interface{}{
				"PhoneNum":  "110",
				"PhoneText": "110",
			},
			map[string]interface{}{
				"PhoneNum":  "119",
				"PhoneText": "119",
			},
		},
	}
	requestValuesJson, _ := json.Marshal(requestValues)
	targetJson, _ := json.Marshal(target)
	assert.Equal(t, requestValuesJson, targetJson)
}
