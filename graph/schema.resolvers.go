package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/joseporres/api_uni_graphql/graph/models"
	"github.com/joseporres/api_uni_graphql/graph/generated"
	"github.com/joseporres/api_uni_graphql/graph/model"
)

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	db := models.FetchConnection()
    defer db.Close()

	student := model.Student{Nombre: nombre, Dni: dni, Direccion: direccion, FechaNacimiento: fechaNacimiento}

	db.Create(&student)

	return &student,nil;

}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {
	db := models.FetchConnection()
    defer db.Close()

	course := model.Course{Nombre: nombre, Descripcion: descripcion, Temas: temas}

	db.Create(&course)
	
	return &course,nil;
}

// CreateRecord is the resolver for the createRecord field.
func (r *mutationResolver) CreateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, nombre *string, dni string, direccion *string, fechaNacimiento *string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, nombre string, descripcion *string, temas *string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateRecord is the resolver for the updateRecord field.
func (r *mutationResolver) UpdateRecord(ctx context.Context, student string, course string, startdate *string, finishdate *string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, dni string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, nombre string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteRecord is the resolver for the deleteRecord field.
func (r *mutationResolver) DeleteRecord(ctx context.Context, student string, course string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetStudents is the resolver for the getStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetRecords is the resolver for the getRecords field.
func (r *queryResolver) GetRecords(ctx context.Context) ([]*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetStudent is the resolver for the getStudent field.
func (r *queryResolver) GetStudent(ctx context.Context, dni string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, nombre string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetRecord is the resolver for the getRecord field.
func (r *queryResolver) GetRecord(ctx context.Context, student string, course string) (*model.Record, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver}
type queryResolver struct{ *Resolver }

