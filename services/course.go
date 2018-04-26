package services

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type CourseService struct {
	db *sqlx.DB
}

func NewCourseService(db *sqlx.DB) *CourseService {
	return &CourseService{db: db}
}
func (c *CourseService) GetCourse(id *string) (*models.Course, error) {
	course := &models.Course{}
	query := "SELECT courses.* FROM courses WHERE id = $1"
	row := c.db.QueryRowx(query, id)
	err := row.StructScan(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}
func (c *CourseService) GetCourses(title *string) ([]*models.Course, error) {
	courses := make([]*models.Course, 0)
	query := "SELECT courses.* FROM courses"
	var err error
	if title != nil {
		query = query + " WHERE LOWER(title) LIKE CONCAT('%', LOWER($1), '%')"
		err = c.db.Select(&courses, query, &title)
	} else {
		err = c.db.Select(&courses, query)
	}
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (c *CourseService) GetCoursesByCompany(companyId *string) ([]*models.Course, error) {
	courses := make([]*models.Course, 0)
	query := "SELECT courses.* FROM courses WHERE company_id = $1"
	err := c.db.Select(&courses, query, &companyId)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (c *CourseService) GetCoursesByCategory(categoryId *string) ([]*models.Course, error) {
	courses := make([]*models.Course, 0)
	query := "SELECT courses.* FROM courses WHERE category_id = $1"
	err := c.db.Select(&courses, query, &categoryId)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
func (c *CourseService) CreateCourse(course *models.Course) (*models.Course, error) {
	id := xid.New()
	course.ID = id.String()
	query := "INSERT INTO courses(id, company_id, category_id, code, title, overview, description, amount, hours, created_at, updated_at) VALUES(:id,:company_id, :category_id, :code, :title, :overview, :description, :amount, :hours, :created_at, :updated_at)"
	_, err := c.db.NamedExec(query, course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *CourseService) UpdateCourse(course *models.Course, id string) (*models.Course, error) {
	query := `UPDATE courses
	SET company_id = :company_id, category_id=:category_id, code := :code, title = :title, hours = :hours, amount = :amountupdated_at = :updated_at
	WHERE id = :id RETURNING courses.*`
	params := make(map[string]interface{}, 0)

	params["company_id"] = course.CompanyID
	params["category_id"] = course.CategoryID
	params["code"] = course.Code
	params["title"] = course.Title
	params["overview"] = course.Overview
	params["description"] = course.Description
	params["hours"] = course.Hours
	params["amount"] = course.Amount
	params["updated_at"] = time.Now()
	params["id"] = id

	result, err := c.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	result.StructScan(&course)
	return course, nil
}

func (c *CourseService) RemoveCourse(id string) {
	query := "DELETE courses WHERE id = $1"
	_, err := c.db.Exec(query, &id)
	if err != nil {
		panic(err)
	}
}
