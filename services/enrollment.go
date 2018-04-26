package services

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/smithaitufe/courses/models"
)

type EnrollmentService struct {
	db            *sqlx.DB
	courseService *CourseService
}

func NewEnrollmentService(db *sqlx.DB, courseService *CourseService) *EnrollmentService {
	return &EnrollmentService{db: db, courseService: courseService}
}

func (e *EnrollmentService) GetEnrollments() ([]*models.Enrollment, error) {
	enrollments := make([]*models.Enrollment, 0)
	query := "SELECT enrollments.* FROM enrollments ORDER BY id ASC"
	err := e.db.Select(&enrollments, query)
	if err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (e *EnrollmentService) GetEnrollment(id string) (*models.Enrollment, error) {
	enrollment := &models.Enrollment{}
	query := "SELECT enrollments.* FROM enrollments WHERE id = $1"
	row := e.db.QueryRowx(query, id)
	err := row.StructScan(&enrollment)
	if err != nil {
		return nil, err
	}
	return enrollment, nil
}

func (e *EnrollmentService) CreateEnrollment(enrollment *models.Enrollment) (*models.Enrollment, error) {
	enrollment.ID = xid.New().String()
	query := "INSERT INTO enrollments(id, user_id, role_id, created_at, updated_at) VALUES(:id, :user_id, :role_id, :created_at, :updated_at)"
	result, err := e.db.NamedQuery(query, &enrollment)
	if err != nil {
		return nil, err
	}
	result.StructScan(&enrollment)
	return enrollment, nil
}

func (e *EnrollmentService) UpdateEnrollment(enrollment *models.Enrollment, id string) (*models.Enrollment, error) {
	query := `UPDATE enrollments
	SET
  user_id = :user_id,
  course_id = :course_id,
  updated_at = :updated_at
	WHERE id = :id RETURNING enrollments.*`
	params := make(map[string]interface{}, 0)

	params["user_id"] = enrollment.UserID
	params["course_id"] = enrollment.CourseID
	params["updated_at"] = time.Now()
	params["id"] = id

	result, err := e.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	result.StructScan(&enrollment)
	return enrollment, nil

}

func (e *EnrollmentService) FindEnrollmentsByUserId(userId string) ([]*models.Enrollment, error) {
	enrollments := make([]*models.Enrollment, 0)
	query := "SELECT enrollments.* FROM enrollments WHERE enrollments.user_id =$1"
	row := e.db.QueryRowx(query, userId)
	err := row.StructScan(enrollments)
	if err != nil {
		return nil, err
	}
	return enrollments, nil
}

func (e *EnrollmentService) FindEnrollmentsByCourseId(courseId string) ([]*models.Enrollment, error) {
	enrollments := make([]*models.Enrollment, 0)
	query := "SELECT enrollments.* FROM enrollments WHERE enrollments.course_id =$1"
	row := e.db.QueryRowx(query, courseId)
	err := row.StructScan(enrollments)
	if err != nil {
		return nil, err
	}
	return enrollments, nil
}
