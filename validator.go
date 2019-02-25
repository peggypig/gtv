package gtv

import "gtv/validators"

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-25 18:36
**/

func NewIntValidator() (validator *validators.IntValidator) {
	return &validators.IntValidator{}
}

func NewStringValidator() (validator *validators.StringValidator) {
	return &validators.StringValidator{}
}