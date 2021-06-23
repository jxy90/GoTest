package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/rand"
	"testing"
)

type Articles struct {
	Id      int64  `xorm:"pk autoincr"`
	Title   string `xorm:"varchar(50) notnull index(title) 'title'""`
	Content int    `xorm:"int notnull index(content)  'content'"`
}

func Test_Partition(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=UTC&multiStatements=true")
	if err != nil {
		println(err)
		return
	}
	//err = engine.Sync(Articles{})
	//if err != nil {
	//	println(err)
	//	return
	//}
	abcs := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := range abcs {
		for j := range abcs {
			users := make([]*Articles, 0)
			for k := 0; k < 1000; k++ {
				title := fmt.Sprintf("%v%v%v", string(abcs[i]), string(abcs[j]), rand.Intn(50))
				user := &Articles{
					Title:   title,
					Content: rand.Intn(50),
				}
				users = append(users, user)
			}
			for p := 0; p < 10; p++ {
				count, err := engine.Insert(users)
				if err != nil {
					println(fmt.Sprintf("插入失败%v", err))
					return
				}
				println(fmt.Sprintf("插入 : %v", count))
			}
		}
	}
}
