package services

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type CompanyService struct {
	db            *sqlx.DB
	courseService *CourseService
}

func NewCompanyService(db *sqlx.DB, courseService *CourseService) *CompanyService {
	return &CompanyService{db: db, courseService: courseService}
}

func (c *CompanyService) GetCompanies() ([]*models.Company, error) {
	companies := make([]*models.Company, 0)
	query := "SELECT companies.* FROM companies"
	err := c.db.Select(&companies, query)
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		company.Courses, _ = c.courseService.GetCoursesByCompany(&company.ID)
	}
	return companies, nil
}

func (c *CompanyService) GetCompany(id *string) (*models.Company, error) {
	company := &models.Company{}
	query := "SELECT companies.* FROM companies WHERE id = $1"
	row := c.db.QueryRowx(query, id)
	err := row.StructScan(company)
	if err != nil {
		return nil, err
	}
	courses, _ := c.courseService.GetCoursesByCompany(id)
	company.Courses = courses
	return company, nil
}

func (c *CompanyService) CreateCompany(company *models.Company) (*models.Company, error) {
	id := xid.New()
	company.ID = id.String()
	query := "INSERT INTO companies(id, name, created_at, updated_at) VALUES(:id, :name, :created_at, :updated_at)"
	_, err := c.db.NamedExec(query, company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *CompanyService) UpdateCompany(company *models.Company, id *string) (*models.Company, error) {
	query := `UPDATE companies
	SET name = :name, updated_at = :updated_at
	WHERE id = :id`
	params := make(map[string]interface{}, 0)

	params["name"] = company.Name
	params["updated_at"] = time.Now()
	params["id"] = id

	_, err := c.db.NamedExec(query, params)
	if err != nil {
		return nil, err
	}

	return c.GetCompany(id)
}

func (c *CompanyService) RemoveCompany(id uint) {
	query := "DELETE companies WHERE id = $1"
	_, err := c.db.Exec(query, &id)
	if err != nil {
		panic(err)
	}
}
