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
