package database

import (
	"fmt"
	"time"

	gsb "github.com/huandu/go-sqlbuilder"
)

func SqlInsert() {
	insertBuilder := gsb.NewInsertBuilder()
	insertBuilder = insertBuilder.InsertInto("student").Cols("name", "province", "city", "entrollment").Values("小明", "河南", "郑州", "2015-01-01") //除InsertInto外也支持ReplaceInto
	for i := range 3 {
		insertBuilder = insertBuilder.Values(RandStringRunes(4), "浙江", "杭州", time.Now().Add(time.Hour*24*time.Duration(i)).Format("2006-01-02"))
	}
	sql, args := insertBuilder.Build()
	fmt.Println(sql)     //INSERT INTO student (name, province, city, enrollment) VALUES (?, ?, ?, ?), (?, ?, ?, ?), (?, ?, ?, ?), (?, ?, ?, ?)
	fmt.Println(args...) //[小明 河南 郑州 2015-01-01 bZ72 河南 郑州 2025-01-16 CeKQ 河南 郑州 2025-01-17 Ip1p 河南 郑州 2025-01-18]
}

func SqlDelete() {
	deleteBuilder := gsb.NewDeleteBuilder()
	deleteBuilder = deleteBuilder.DeleteFrom("student").Where(
		deleteBuilder.Equal("city", "杭州"),
	)
	sql, args := deleteBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}

func SqlRead() {
	selectBuilder := gsb.NewSelectBuilder()
	selectBuilder.SetFlavor(gsb.MySQL) // 不同的数据库sql语法可能有略微的差异，通过Flavor指定使用哪种数据库的语法
	selectBuilder = selectBuilder.Select("name", "province").From("student")
	selectBuilder = selectBuilder.Where("score<80")
	selectBuilder = selectBuilder.Where( //多个where默认是and关系
		selectBuilder.Or(
			selectBuilder.Equal("province", "杭州"),
			selectBuilder.GE("enrollment", "2015-01-01"),
		),
		selectBuilder.In("杭州", "上海", "北京", "宁波"),
	)
	selectBuilder = selectBuilder.OrderByDesc("name") //Asc是升序，Desc是降序
	selectBuilder = selectBuilder.Limit(3).Offset(10)
	sql, args := selectBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}

func SqlUpdate() {
	updateBuilder := gsb.NewUpdateBuilder()
	updateBuilder = updateBuilder.Update("student").Set("name=zcy", "score=50").Where("score=0")
	sql, args := updateBuilder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}
