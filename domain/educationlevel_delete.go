// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/educationlevel"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
)

// EducationLevelDelete is the builder for deleting a EducationLevel entity.
type EducationLevelDelete struct {
	config
	hooks    []Hook
	mutation *EducationLevelMutation
}

// Where appends a list predicates to the EducationLevelDelete builder.
func (eld *EducationLevelDelete) Where(ps ...predicate.EducationLevel) *EducationLevelDelete {
	eld.mutation.Where(ps...)
	return eld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eld *EducationLevelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EducationLevelMutation](ctx, eld.sqlExec, eld.mutation, eld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eld *EducationLevelDelete) ExecX(ctx context.Context) int {
	n, err := eld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eld *EducationLevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(educationlevel.Table, sqlgraph.NewFieldSpec(educationlevel.FieldID, field.TypeInt))
	if ps := eld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eld.mutation.done = true
	return affected, err
}

// EducationLevelDeleteOne is the builder for deleting a single EducationLevel entity.
type EducationLevelDeleteOne struct {
	eld *EducationLevelDelete
}

// Where appends a list predicates to the EducationLevelDelete builder.
func (eldo *EducationLevelDeleteOne) Where(ps ...predicate.EducationLevel) *EducationLevelDeleteOne {
	eldo.eld.mutation.Where(ps...)
	return eldo
}

// Exec executes the deletion query.
func (eldo *EducationLevelDeleteOne) Exec(ctx context.Context) error {
	n, err := eldo.eld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{educationlevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eldo *EducationLevelDeleteOne) ExecX(ctx context.Context) {
	if err := eldo.Exec(ctx); err != nil {
		panic(err)
	}
}