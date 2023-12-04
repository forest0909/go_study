package main

import (
	"encoding/json"
	"fmt"
)

//func plus(n1, n2 int) int {
//	return n1 + n2
//}
//func plus1(n1, n2 uint) uint {
//	return n1 + n2
//}

//func plusT[T int | uint](n1, n2 T) T {
//	return n1 + n2
//}

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func plusT[T Number](n1, n2 T) T {
	return n1 + n2
}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {

	//plus(1, 20)
	//var u1, u2 = uint(2), uint(3)
	////plus1(u1, u2)
	//plusT(u1, u2)

	type User struct {
		Name string `json:"name"`
	}
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//user := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: User{
	//		Name: "李四",
	//	},
	//}
	//byteData, _ := json.Marshal(user)
	//println(string(byteData))
	//
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: UserInfo{
	//		Name: "张三",
	//		Age:  18,
	//	},
	//}
	//byteData, _ = json.Marshal(userInfo)
	//println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"李四"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)

	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"张三","age":18}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}
