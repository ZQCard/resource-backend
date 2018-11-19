package models

type AuthRoleRoutes struct {
	ID          uint8 `gorm:"primary_key"`
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
