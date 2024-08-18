package main

import (
	"fmt"
	"log"
	"time"

	"go_test/utils"
)

/*
最后，返回到命令行
执行命令: go run main.go
如果你创建了其他文件并想同时执行它们，可以使用这个命令:
go run . */

type User struct {
	UserId     string
	username   string
	password   string
	CreateDate *time.Time
	ExpireDate *time.Time
}

func main() {

	// user := User{
	// 	UserId:     uuid.NewString(),
	// 	username:   "张三",
	// 	password:   fmt.Sprint(rand.Int() * 1000),
	// 	CreateDate: time.Now(),
	// }

	// INSERTD, err := AddUser(&user)
	// if err != nil {
	// 	log.Fatalf("failed to get rows affected: %v", err)
	// }
	// fmt.Printf("insertd row is ,%d", INSERTD)

	// userId := "b9feaf27-10c3-4110-bd4c-4d75088af11d"
	// DELETED, err := DelUser(userId)

	// if err != nil {
	// 	log.Fatalf("failed to get rows affected: %v", err)
	// }
	// fmt.Printf("insertd row is ,%d", DELETED)

	users := QueryAllUser()

	fmt.Printf("users is %+v", users)

}

// DelUser is delete user by user_id
// params id
func DelUser(id string) (int64, error) {
	db, _ := utils.ConnectMysql()

	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM T_SYS_USER WHERE USER_ID = ?")

	if err != nil {
		return 0, fmt.Errorf("failed to prepare SQL statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	// Attempt to get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Handle the case where RowsAffected() is not supported
		// You might want to log the error or return a specific error type
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return rowsAffected, nil

}

func AddUser(u *User) (int64, error) {
	db, _ := utils.ConnectMysql()

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO t_sys_user(user_id,username,password,create_date) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.UserId, u.username, u.password, u.CreateDate)
	if err != nil {
		return 0, err
	}

	// Attempt to get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Handle the case where RowsAffected() is not supported
		// You might want to log the error or return a specific error type
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	return rowsAffected, nil
}

func QueryAllUser() []User {
	// 调用db包ConnMySQL()
	db, nil := utils.ConnectMysql()
	// 预编译查询sql创建 statement
	stmt, err := db.Prepare("SELECT * from t_sys_user")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	// 执行查询sql，返回查询结果rows
	rows, err := stmt.Query()
	if err != nil {
		// 打印错误信息
		log.Fatal(err)
		// 抛出错误信息，阻止程序继续运行
		panic(err)
	}
	// 定义User切片
	s := make([]User, 0)
	// 遍历rows
	for rows.Next() {
		u := User{}
		// 扫描rows的每一列并保存数据到User对应字段
		err := rows.Scan(&u.UserId, &u.username, &u.password, &u.CreateDate, &u.ExpireDate)
		if err != nil {
			// 打印错误信息
			log.Fatal(err)
			// 抛出错误信息，阻止程序继续运行
			panic(err)
		}
		// 扫描后的user加入到切片
		s = append(s, u)
	}
	return s
}
