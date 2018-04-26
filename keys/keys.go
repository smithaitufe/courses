package keys

type Key string

const (
	CategoryLoaderKey   Key = "category"
	CompanyLoaderKey    Key = "company"
	CourseLoaderKey     Key = "course"
	UserLoaderKey       Key = "user"
	EnrollmentLoaderKey Key = "enrollment"
	RoleLoaderKey       Key = "role"

	RoleServiceKey       Key = "roleService"
	CourseServiceKey     Key = "courseService"
	CategoryServiceKey   Key = "categoryService"
	CompanyServiceKey    Key = "companyService"
	EnrollmentServiceKey Key = "enrollmentService"
	UserServiceKey       Key = "userService"
)
