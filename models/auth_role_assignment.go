package models

type AuthRoleAssignment struct {
	ID     uint `gorm:"primary_key"`
	UserId int
	Role   string
}

func (AuthRoleAssignment) TableName() string {
	return "auth_role_assignment"
}

func RoleList() (roles []string) {
	var assignments []AuthRoleAssignment
	db.Table("auth_role_assignment").Raw("SELECT DISTINCT role FROM auth_role_assignment").Scan(&assignments)
	for _, v := range assignments {
		roles = append(roles, v.Role)
	}
	return
}

func FindRoleByUserId(id uint) (roles []string) {
	var assignments []AuthRoleAssignment
	db.Model(AuthRoleAssignment{}).Where("id = ?", id).Select("role").Find(&assignments)
	for _, v := range assignments {
		roles = append(roles, v.Role)
	}
	return
}
