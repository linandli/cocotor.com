package model

import (
	"cocotor.com/utils"
	"database/sql"
	"fmt"
	"log"

    _ "github.com/lib/pq"
)

const (
    host     = "121.42.160.108"
    port     = 5432
    user     = "postgres"
    password = "tuji2013"
    dbname   = "torch_service"
)

// User 用户类
type User struct {
	Id         string `json:"userId"`
	Name       string `json:"userName"`
	Gender     string `json:"gender"`
	Phone      string `json:"userMobile"`
	Pwd        string `json:"pwd"`
	Permission string `json:"permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
}

type RegistInfo struct {
	Name  string `json:"name"`
	Gender     string `json:"gender"`
	Phone string `json:"mobile"`
	Pwd   string `json:"pwd"`
	Permission string `json:"permission"`
}

// 序列化
func dumpUser(user User) []byte {
	dumped, _ := user.MarshalJSON()
	return dumped
}

// 反序列化
func loadUser(jsonByte []byte) User {
	res := User{}
	res.UnmarshalJSON(jsonByte)
	return res
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
}

// LoginCheck 登录验证
func LoginCheck(loginReq LoginReq) (bool, User, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
       log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	var resultUser User
	
	resultBool := false
	sqlStatement := fmt.Sprintf(`SELECT "Id", "Name", "Gender", "Phone", 
		"Pwd", "Permission" FROM "core"."USER" WHERE "Phone"='%s' AND 
		"Pwd"='%s';`, loginReq.Phone, utils.GetMd5String(loginReq.Pwd))
	err = db.QueryRow(sqlStatement).Scan(&resultUser.Id, &resultUser.Name, 
		&resultUser.Gender, &resultUser.Phone, &resultUser.Pwd, 
		&resultUser.Permission)
	
	if err == sql.ErrNoRows  {
		return resultBool, resultUser, err
	} else {
		resultBool = true
		return resultBool, resultUser, err
	}
}

// Register 插入用户，先检查是否存在用户，如果没有则存入
func Register(name string, phone string, pwd string, gender string, permission string) error {
	if CheckUser(phone) {
		return fmt.Errorf("用户已存在！")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	
    if err != nil {
        fmt.Println(err)
	}
	
	defer db.Close()

	uid := utils.UniqueId()
	stmt,err := db.Prepare(`INSERT INTO "core"."USER"("Id", "Phone", "Name", "Gender", "Pwd", "Permission") VALUES ($1, $2, $3, $4, $5, $6);`)
	
	if err != nil {
		log.Fatal(err)
	}

	_,err = stmt.Exec(uid, phone, name, gender, utils.GetMd5String(pwd), permission)

	if err != nil {
		log.Fatal(err)
	}else {
		log.Println("插入数据库表成功！")
	}
	stmt.Close()
	return err
}


// CheckUser 检查用户是否存在
func CheckUser(phone string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	
    if err != nil {
        log.Fatal(err)
	}
	
	defer db.Close()

	res_phone := 0
	err = db.QueryRow(`SELECT 1 FROM "core"."USER" WHERE "Phone"=$1;`, phone).Scan(&res_phone)
	
	if err != nil  {
		return false
	} else {
		return true
	}
}