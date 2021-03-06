## Go Table Validator

    表验证器，用于参数合法性校验。  
    基本思想是把model转化成表（Table，包含多个Field，每个Field又包含FieldName,  
    FieldValue,Validators三个属性），最终对参数表进行校验。
    没有使用反射！没有使用反射！没有使用反射！
        
    
#### quick start

方式一：  
```go
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
```

方式二：
```go
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
    							NewIntValidator().Max(20).Min(1),
    						},
    					},
    				},
    			},
    		},
    	}
    	requestValues, err := table.FillTable(request)
    	err = Validator(table)
```
这个调用方式主要是简化了用户构造Table的操作，可以在构造出基本的表结构后，通过   
FillTable将http.request传入，表填充器将根据表的FieldName和请求参数的Key进行比对，  
将Value填充到表中。

#### Extra
1. 目前自带支持的校验器有：
    - StringValidator
    - IntValidator
    - Float64Validator
    - SliceValidator
    - EnumValidator
2. 在自带校验器不满足需求时，可以自定义校验器，实现IValidator接口即可。
3. 表结构支持以下：
    1. 表中嵌套表
    2. 表中嵌套切片
    3. 表中嵌套字段
    4. 切片中嵌套表
    5. 切片中嵌套字段
    