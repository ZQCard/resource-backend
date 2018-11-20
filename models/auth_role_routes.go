package models

import "errors"

type AuthRoleRoutes struct {
	ID          uint8 `gorm:"primary_key"`
	Role        string
	RoutePath   string
	RouteMethod string
}

func (AuthRoleRoutes) TableName() string {
	return "auth_role_routes"
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

func AddRoleRoute(role, route, method string) error {
	var roleRoute AuthRoleRoutes

	if err := db.Model(AuthRoleRoutes{}).Where(AuthRoleRoutes{Role:role, RoutePath:route, RouteMethod:method}).FirstOrCreate(&roleRoute).Error; err != nil {
		return err
	}
	if roleRoute.ID > 0{
		return nil
	}
	return errors.New("创建数据失败")
}

func GetRoutesByUserId(userId uint) []string {
	// 这里查两次表,然后在查拥有路由
	var authRoleAssignment []AuthRoleAssignment
	db.Model(AuthRoleAssignment{}).Where("user_id = ?", userId).Select("role").Find(&authRoleAssignment)

	var roles []string
	for _, v := range authRoleAssignment{
		roles = append(roles, v.Role)
	}

	var routes []string
	var authRoleRoutes []AuthRoleRoutes
	db.Model(AuthRoleRoutes{}).Where("role in (?)", roles).Select("route_path, route_method").Find(&authRoleRoutes)
	for _, v := range authRoleRoutes{
		routes = append(routes, v.RouteMethod + ":" + v.RoutePath)
	}
	return routes
}
