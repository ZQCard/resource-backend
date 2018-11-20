package models

import (
	"resource-backend/pkg/logging"
)

type AuthRoleAssignment struct {
	ID     uint `gorm:"primary_key"`
	UserId uint
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

func FindOrCreateAssignment(userId uint, role string) bool {
	var assignment AuthRoleAssignment
	assignment.UserId = userId
	assignment.Role = role
	if err := db.Model(AuthRoleRoutes{}).Where(AuthRoleAssignment{UserId: userId, Role: role}).FirstOrCreate(&assignment).Error; err != nil {
		logging.Error(err.Error())
		return false
	}
	if assignment.ID > 0{
		return true
	}
	return true
}
