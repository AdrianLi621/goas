/**
 * @Author adrianli
 * @Description //TODO $
 * @Date 2021/12/6 17:54
 **/
package unit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"runtime"
)

var (
	db *sql.DB
)
type User struct {
	Id int
	Name string
}
/**
 *  InitDb
 *  @Description:
 *  @param cfg
 **/
func InitDb(cfg *viper.Viper) {
	user := cfg.Get("database.username")
	password := cfg.Get("database.password")
	host := cfg.Get("database.host")
	port := cfg.Get("database.port")
	dbname := cfg.Get("database.dbname")
	dsn := fmt.Sprintf("%s:%d@tcp(%s:%d)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func InitModel(conf *viper.Viper) {
	InitDb(conf)
	user:=new(User)
	row:=db.QueryRow("select id,name from `user` limit 1")
	fmt.Println(row.Err())
	runtime.Goexit()
	if err :=row.Scan(&user.Id,&user.Name); err != nil{
		fmt.Printf("scan failed, err:%v",err)
		return
	}
	fmt.Println(user.Id,user.Name)
}
