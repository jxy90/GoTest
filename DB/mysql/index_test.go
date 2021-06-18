package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/rand"
	"testing"
)

type User struct {
	Id   int64  `xorm:"pk autoincr"`
	Name string `xorm:"varchar(25) notnull index(name_age) 'name'""`
	Age  int    `xorm:"int notnull index(name_age)  'age'"`
}

func Test_Index(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3308)/test?charset=utf8&parseTime=True&loc=UTC&multiStatements=true")
	if err != nil {
		println(err)
		return
	}
	err = engine.Sync(User{})
	if err != nil {
		println(err)
		return
	}
	abcs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range abcs {
		for j := range abcs {
			users := make([]*User, 0)
			for k := 0; k < 100; k++ {
				user := &User{
					Name: fmt.Sprintf("%v%v", string(abcs[i]), string(abcs[j])),
					Age:  rand.Intn(50),
				}
				users = append(users, user)
			}
			count, err := engine.Insert(users)
			if err != nil {
				println(fmt.Sprintf("插入失败%v", err))
				return
			}
			println(fmt.Sprintf("插入 : %v", count))
		}
	}
}
