package gtv

import (
	"bytes"
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
				[]IField{
					&ValueField{
						FieldName:  "Like",
						FieldValue: "ping pang",
						Validators: []IValidator{
							NewStringValidator().MaxLen(5),
						},
					},
					&TableField{
						[]IField{
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

type fakeHttpResponseBody struct {
	body io.ReadSeeker
}

func (body *fakeHttpResponseBody) Read(p []byte) (n int, err error) {
	n, err = body.body.Read(p)
	if err == io.EOF {
		body.body.Seek(0, 0)
	}
	return n, err
}

func (body *fakeHttpResponseBody) Close() error {
	return nil
}
func TestTable_FillTable(t *testing.T) {
	request := &http.Request{
		Body:   &fakeHttpResponseBody{bytes.NewReader([]byte(`Name=zhangSan&Age=12`))},
		Method: "POST",
		Header: map[string][]string{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		},
	}
	table := Table{
		Fields: []IField{
			&ValueField{
				FieldName: "Name",
			},
			&ValueField{
				FieldName: "Age",
			},
		},
	}
	err, requestValues := table.FillTable(request)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{
		"Name": "zhangSan",
		"Age":  "12",
	}, requestValues)
}
