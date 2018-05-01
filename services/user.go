package services

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type UserService struct {
	db                *sqlx.DB
	roleService       *RoleService
	enrollmentService *EnrollmentService
}

func NewUserService(db *sqlx.DB, roleService *RoleService, enrollmentService *EnrollmentService) *UserService {
	return &UserService{db: db, roleService: roleService, enrollmentService: enrollmentService}
}

func (u *UserService) GetUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)
	query := "SELECT users.* FROM users ORDER BY id ASC"
	err := u.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) GetUser(id *string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT users.* FROM users WHERE id =$1"
	row := u.db.QueryRowx(query, id)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	id := xid.New()
	user.ID = id.String()
	query := "INSERT INTO users(id, first_name, last_name, email, country, dialing_code, phone_number, password, created_at, updated_at) VALUES(:id, :first_name, :last_name, :email, :country, :dialing_code, :password, :created_at, :updated_at)"
	result, err := u.db.NamedQuery(query, user)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	result.StructScan(&user)
	return user, nil
}

func (u *UserService) UpdateUser(user *models.User, id string) (*models.User, error) {
	query := `UPDATE users 	SET
			last_name = :last_name,
			first_name = :first_name,
			password = :password,
			email = :email,
			country = :country,
			dialing_code = :dialing_code,
			phone_number = :phone_number,
			updated_at = :updated_at
			WHERE id = :id RETURNING users.*`
	params := make(map[string]interface{}, 0)

	params["last_name"] = user.LastName
	params["first_name"] = user.FirstName
	params["password"] = user.Password
	params["email"] = user.Email
	params["country"] = user.Country
	params["dialing_code"] = user.DialingCode
	params["phone_number"] = user.PhoneNumber
	params["updated_at"] = time.Now()
	params["id"] = id

	result, err := u.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	result.StructScan(&user)
	return user, nil

}

func (u *UserService) RemoveUser(id string) {
	query := "DELETE users WHERE id = $1"
	_, err := u.db.Exec(query, &id)
	if err != nil {
		panic(err)
	}
}

func (u *UserService) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT users.* FROM users WHERE email =$1"
	row := u.db.QueryRowx(query, email)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) FindUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT users.* FROM users WHERE phone_number =$1"
	row := u.db.QueryRowx(query, phoneNumber)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
