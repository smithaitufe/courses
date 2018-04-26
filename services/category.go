package services

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type CategoryService struct {
	db            *sqlx.DB
	courseService *CourseService
}

func NewCategoryService(db *sqlx.DB, courseService *CourseService) *CategoryService {
	return &CategoryService{db: db, courseService: courseService}
}

func (c *CategoryService) GetCategories() ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	query := "SELECT categories.* FROM categories ORDER BY id ASC"
	err := c.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}
	// for _, category := categories {
	// 	cateory.Courses = c.courseService.GetCoursesByCategory(category.ID)
	// }
	return categories, nil
}

func (c *CategoryService) GetCategory(id *string) (*models.Category, error) {
	category := &models.Category{}
	query := "SELECT categories.* FROM categories WHERE id =$1"
	row := c.db.QueryRowx(query, id)
	err := row.StructScan(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	id := xid.New()
	category.ID = id.String()
	query := "INSERT INTO categories(id, name, created_at, updated_at) VALUES(:id, :name, :created_at, :updated_at)"
	result, err := c.db.NamedQuery(query, category)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	result.StructScan(&category)
	return category, nil
}

func (c *CategoryService) UpdateCategory(category *models.Category, id string) (*models.Category, error) {
	query := `UPDATE categories
	SET name = :name, updated_at = :updated_at
	WHERE id = :id RETURNING categories.*`
	params := make(map[string]interface{}, 0)

	params["name"] = category.Name
	params["updated_at"] = time.Now()
	params["id"] = id

	result, err := c.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	result.StructScan(&category)
	return category, nil

}

func (c *CategoryService) RemoveCategory(id string) {
	query := "DELETE categories WHERE id = $1"
	_, err := c.db.Exec(query, &id)
	if err != nil {
		panic(err)
	}
}
