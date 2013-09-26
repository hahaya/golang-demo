package main

import (
	"MySQLOp/op"
	"fmt"
)

func main() {
	//初始化数据库
	op.InitMySql()

	//新建要插入数据库中的用户
	user := op.UserInfo{UserId: 1, UserName: "hahaya", UserPwd: "1111"}

	id, err := op.AddUser(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Last Insert User Id: %d\n", id)

	//新建要插入数据库中的用户
	user_ := op.UserInfo{UserId: 2, UserName: "golang", UserPwd: "22222"}

	id_, err := op.AddUser(user_)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Last Insert User Id: %d\n", id_)

	//删除用户
	row, err := op.DelUser(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Delete The %d Row From UserInfo Table\n", row)

	//更新用户
	row_, err := op.UpdateUser(2, "ChangePwd")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Update The %d Row From UserInfo Table\n", row_)

	//查询所有用户
	users, err := op.FindAllUser()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, user := range users {
		fmt.Printf("UserId:%d\tUserName:%s\tUserPwd:%s\n", user.UserId, user.UserName, user.UserPwd)
	}
}
