package gtv

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:31
**/

func Validator(table Table) (err error) {
	for _, field := range table.Fields {
		err = validator(field)
		if err != nil {
			break
		}
	}
	return
}


func validator(field IField) (err error) {
	if field != nil {
		if field.IsTableField() {
			for _, v := range field.GetFields() {
				err = validator(v)
				if err != nil {
					break
				}
			}
		} else {
			validators := field.GetValidators()
			for _, v := range validators {
				err = v.Validator(field.GetFieldName(), field.GetFieldValue())
				if err != nil {
					break
				}
			}
		}
	}
	return
}
