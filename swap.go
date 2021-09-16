package common

import "encoding/json"

//通过json tag 进行结构体赋值
func SwapTo(req,target interface{})(err error){
	dataByte,err := json.Marshal(req)
	if err!=nil{
		return
	}
	err = json.Unmarshal(dataByte,target)
	return
}