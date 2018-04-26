package services

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type RoleService struct {
	db *sqlx.DB
}

func NewRoleService(db *sqlx.DB) *RoleService {
	return &RoleService{db: db}
}

func (r *RoleService) GetRoles() ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	query := "SELECT roles.* FROM roles ORDER BY id ASC"
	err := r.db.Select(&roles, query)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleService) GetRole(id string) (*models.Role, error) {
	role := &models.Role{}
	query := "SELECT roles.* FROM roles WHERE id =$1"
	row := r.db.QueryRowx(query, id)
	err := row.StructScan(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleService) CreateRole(role *models.Role) (*models.Role, error) {
	id := xid.New()
	role.ID = id.String()
	query := "INSERT INTO roles(id, first_name, last_name, email, country, dialing_code, phone_number, password, created_at, updated_at) VALUES(:id, :first_name, :last_name, :email, :country, :dialing_code, :password, :created_at, :updated_at)"
	result, err := r.db.NamedQuery(query, role)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	result.StructScan(&role)
	return role, nil
}

func (r *RoleService) UpdateRole(role *models.Role, id string) (*models.Role, error) {
	query := `UPDATE roles
	SET
  name = :name,
  updated_at = :updated_at
	WHERE id = :id RETURNING roles.*`
	params := make(map[string]interface{}, 0)

	params["name"] = role.Name
	params["updated_at"] = time.Now()
	params["id"] = id

	result, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	result.StructScan(&role)
	return role, nil

}

func (r *RoleService) RemoveRole(id string) {
	query := "DELETE roles WHERE id = $1"
	_, err := r.db.Exec(query, &id)
	if err != nil {
		panic(err)
	}
}

func (r *RoleService) FindRolesByUserId(userId string) ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	query := "SELECT r.* FROM roles r INNER JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id =$1"
	row := r.db.QueryRowx(query, userId)
	err := row.StructScan(roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
