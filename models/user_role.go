package models

type UserRole struct {
	UserID string `db:"user_id",json:"user_id"`
	User   *User
	RoleID string `db:"role_id",json:"role_id"`
	Role   *Role
}
