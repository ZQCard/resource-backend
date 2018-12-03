package models

import (
	"strings"
)

type AuthRoleAssignment struct {
	ID     uint `gorm:"primary_key"`
	UserId uint
	Role   string
}

func (AuthRoleAssignment) TableName() string {
	return "auth_role_assignment"
}

func FindRoleByUserId(id uint) (roles []string) {
	var assignments []AuthRoleAssignment
	db.Model(AuthRoleAssignment{}).Where("user_id = ?", id).Select("role").Find(&assignments)
	for _, v := range assignments {
		roles = append(roles, v.Role)
	}
	return
}

func AssignRemoveRoles(userId uint, assign []string, remove []string) error {
	var err error
	tx := db.Begin()

	// 删除要免去的角色
	err = tx.Exec("DELETE FROM auth_role_assignment WHERE user_id = ? AND role in (?)", userId, remove).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if assign == nil{
		tx.Commit()
		return nil
	}
	// 插入新授予的角色
	insertSql := "INSERT INTO auth_role_assignment(user_id,role) VALUES"
	vals := []interface{}{}
	const rowSQL = "(?,?)"
	var inserts []string
	for _,role := range assign {
		inserts = append(inserts, rowSQL)
		vals = append(vals,userId,role)
	}
	insertSql = insertSql + strings.Join(inserts, ",")
	err = tx.Exec(insertSql, vals...).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func FindUserByRole(role string) (users []AuthRoleAssignment) {
	db.Model(AuthRoleAssignment{}).Where("role = ?", role).Select("user_id").Find(&users)
	return
}
