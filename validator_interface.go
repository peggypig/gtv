package gtv

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:33
**/

type IValidator interface {
	Validator(fieldName string, value interface{}) error
}
