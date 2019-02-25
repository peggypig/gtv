package gtv

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
				FieldValue: "zhangsan",
				Validators: []IValidator{
					NewStringValidator().Required(),
				},
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
