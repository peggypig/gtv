package gerror

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-02-27 11:19
**/

type GError struct {
	Key   string
	Value interface{}
	Msg   string
}

func (err *GError) Error() string {
	return err.Msg
}

func (err *GError) IsNull() (null bool){
	null = true
	if len(err.Msg) > 0 {
		null = false
	}
	return
}
