package op

import (
	"database/sql"
	"errors"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	UserId   int32
	UserName string
	UserPwd  string
}

var db *sql.DB

func InitMySql() error {
	//连接MySQL数据库
	//第一个参数是数据库引擎
	//第二个参数是数据库DNS配置 go中没有统一DNS 不同数据库DNS配置可能不同 具体参考文档
	_db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	if err != nil {
		err := errors.New("Connect To MySQL Failed, Please Check Your DNS Config...")
		return err
	}

	db = _db

	return nil
}

//增加用户
func AddUser(user UserInfo) (int64, error) {
	//检查数据库是否已经打开
	if db == nil {
		err := errors.New("Please Call InitMySql First...")
		return -1, err
	}

	//准备插入数据
	stmt, err := db.Prepare("insert into UserInfo values(?, ?, ?);")
	if err != nil {
		return -1, err
	}

	//插入数据
	result, err := stmt.Exec(user.UserId, user.UserName, user.UserPwd)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, nil
	}

	return id, nil
}

//根据用户ID删除用户
func DelUser(userId int32) (int64, error) {
	if db == nil {
		err := errors.New("Please Call InitMySql First...")
		return -1, err
	}

	//准备删除用户
	stmt, err := db.Prepare("delete from UserInfo where UserId=?;")
	if err != nil {
		return -1, err
	}

	//删除用户
	result, err := stmt.Exec(userId)
	if err != nil {
		return -1, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return row, nil
}

//根据用户ID修改用户密码
func UpdateUser(userId int32, userPwd string) (int64, error) {
	if db == nil {
		err := errors.New("Please Call InitMySql First...")
		return -1, err
	}

	//准备更新用户
	stmt, err := db.Prepare("update UserInfo set UserPwd=? where UserId=?;")
	if err != nil {
		return -1, err
	}

	//更新用户
	result, err := stmt.Exec(userPwd, userId)
	if err != nil {
		return -1, err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return row, nil
}

//查找所有用户
func FindAllUser() ([]*UserInfo, error) {
	if db == nil {
		err := errors.New("Please Call InitMySql First...")
		return nil, err
	}

	//定义返回用户信息的切片 切片初始时为空
	users := make([]*UserInfo, 0)

	//查找用户
	rows, err := db.Query("select * from UserInfo")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userId int32
		var userName string
		var userPwd string

		err := rows.Scan(&userId, &userName, &userPwd)
		if err != nil {
			return nil, err
		}

		user := &UserInfo{UserId: userId, UserName: userName, UserPwd: userPwd}
		//向切片中增加值
		users = append(users, user)
	}

	return users, nil
}
