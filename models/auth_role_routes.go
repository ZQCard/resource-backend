package models

import (
	"strings"
)

type AuthRoleRoutes struct {
	ID          uint8 `gorm:"primary_key"`
	Role        string
	RoutePath   string
	RouteMethod string
}

func (AuthRoleRoutes) TableName() string {
	return "auth_role_routes"
}

func RoleList() (roles []string) {
	var allocate []AuthRoleRoutes
	db.Raw("SELECT DISTINCT role FROM auth_role_routes").Scan(&allocate)
	for _, v := range allocate {
		roles = append(roles, v.Role)
	}
	return
}

func FindRoutesByRole(role string) (routes []string) {
	var roleRoutes []AuthRoleRoutes
	db.Model(AuthRoleRoutes{}).Select("route_path, route_method").Where("role = ?", role).Find(&roleRoutes)
	for _, v := range roleRoutes {
		routes = append(routes, v.RouteMethod+":"+v.RoutePath)
	}
	return
}

func CheckRoleExist(role string) bool {
	var authRole AuthRoleRoutes
	if err := db.Model(AuthRoleRoutes{}).Where("role = ?", role).First(&authRole).Error; err != nil {
		return false
	}
	return true
}

func AssignRemoveRoutes(role string, assign []string, remove []string) error {
	var err error
	tx := db.Begin()

	// 删除要免去的路由
	for _, item := range remove {
		temp := strings.Split(item, ":")
		err = db.Where("role = ? and route_method = ? and route_path = ?", role, temp[0], temp[1]).Delete(&AuthRoleRoutes{}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if assign == nil {
		tx.Commit()
		return nil
	}
	// 插入新分配的路由
	insertSql := "INSERT INTO auth_role_routes(role,route_method,route_path) VALUES"
	vals := []interface{}{}
	const rowSQL = "(?,?,?)"
	var inserts []string
	for _, item := range assign {
		temp := strings.Split(item, ":")
		inserts = append(inserts, rowSQL)
		vals = append(vals, role, temp[0], temp[1])
	}
	insertSql = insertSql + strings.Join(inserts, ",")
	err = tx.Debug().Exec(insertSql, vals...).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetRoutesByUserId(userId uint) []string {
	// 这里查两次表,然后在查拥有路由
	var authRoleAssignment []AuthRoleAssignment
	db.Model(AuthRoleAssignment{}).Where("user_id = ?", userId).Select("role").Find(&authRoleAssignment)

	var roles []string
	for _, v := range authRoleAssignment {
		roles = append(roles, v.Role)
	}

	var routes []string
	var authRoleRoutes []AuthRoleRoutes
	db.Model(AuthRoleRoutes{}).Where("role in (?)", roles).Select("route_path, route_method").Find(&authRoleRoutes)
	for _, v := range authRoleRoutes {
		routes = append(routes, v.RouteMethod+":"+v.RoutePath)
	}
	return routes
}

func DeleteRoutesByRole(name string) error {
	if err := db.Where("role = ?", name).Delete(&AuthRoleRoutes{}).Error; err != nil {
		return err
	}
	return nil
}
