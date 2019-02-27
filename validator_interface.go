package gtv

import "github.com/peggypig/gtv/gerror"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:33
**/

type IValidator interface {
	Validator(fieldName string, value interface{})  *gerror.GError
}
