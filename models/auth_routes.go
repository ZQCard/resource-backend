package models

import (
	"strings"
)

type Routes struct {
	Path   string
	Method string
}

func (Routes)TableName() string {
	return "auth_routes"
}

func Refresh(routes []Routes) bool {
	tx := db.Begin()
	db.Exec("DELETE FROM auth_routes")
	sql := "INSERT INTO auth_routes(`route`, `method`) VALUES "
	for _, v := range routes{
		sql += "('"+v.Path+"','"+v.Method+"'),"
	}
	sql = strings.TrimSuffix(sql, ",")
	db.Exec(sql)
	tx.Commit()
	return true
}

func RoutesList() (routes []Routes) {
	db.Model(Routes{}).Find(&routes)
	return
}
