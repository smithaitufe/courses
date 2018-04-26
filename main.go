package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
	ccontext "github.com/smithaitufe/courses/context"
	h "github.com/smithaitufe/courses/handlers"
	ck "github.com/smithaitufe/courses/keys"
	"github.com/smithaitufe/courses/loaders"
	"github.com/smithaitufe/courses/resolvers"
	"github.com/smithaitufe/courses/schema"
	"github.com/smithaitufe/courses/services"
)

func main() {

	config := ccontext.LoadConfig()
	db := ccontext.OpenDB(config)

	courseService := services.NewCourseService(db)
	roleService := services.NewRoleService(db)
	companyService := services.NewCompanyService(db, courseService)
	categoryService := services.NewCategoryService(db, courseService)
	enrollmentService := services.NewEnrollmentService(db, courseService)
	userService := services.NewUserService(db, roleService, enrollmentService)

	ctx := context.Background()
	log := services.NewLogger(config)

	ctx = context.WithValue(ctx, ck.RoleServiceKey, roleService)
	ctx = context.WithValue(ctx, ck.CourseServiceKey, courseService)
	ctx = context.WithValue(ctx, ck.CategoryServiceKey, categoryService)
	ctx = context.WithValue(ctx, ck.CompanyServiceKey, companyService)
	ctx = context.WithValue(ctx, ck.EnrollmentServiceKey, enrollmentService)
	ctx = context.WithValue(ctx, ck.UserServiceKey, userService)

	schemaFile := schema.GetSchema()
	parsedSchema, err := graphql.ParseSchema(schemaFile, &resolvers.SchemaResolver{})
	if err != nil {
		log.Fatal(err)
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.43.45:8081"},
	})
	loggerHandler := &h.LoggerHandler{DebugMode: config.Logger.DebugMode}

	http.Handle("/", c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, err := ioutil.ReadFile("graphiql.html")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(page)
	})))

	handlerWithContext := h.AppendContext(ctx, &h.Handler{Schema: parsedSchema, Loaders: loaders.NewLoaderCollection()})
	http.Handle("/graphql", c.Handler(loggerHandler.Logging(handlerWithContext)))
	http.Handle("/graphql/", c.Handler(loggerHandler.Logging(handlerWithContext)))
	fmt.Printf("Server is listening at %s:%s", config.App.Host, config.App.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), nil))
}
