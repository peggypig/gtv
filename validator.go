package gtv

import "github.com/peggypig/gtv/validators"

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

func NewFloat64Validator() (validator *validators.Float64Validator) {
	return &validators.Float64Validator{}
}

func NewSliceValidator() (validator *validators.SliceValidator) {
	return &validators.SliceValidator{}
}