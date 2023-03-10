// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/educationlevel"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
	"github.com/devcui/ncepu-cs-project/domain/student"
)

// EducationLevelUpdate is the builder for updating EducationLevel entities.
type EducationLevelUpdate struct {
	config
	hooks    []Hook
	mutation *EducationLevelMutation
}

// Where appends a list predicates to the EducationLevelUpdate builder.
func (elu *EducationLevelUpdate) Where(ps ...predicate.EducationLevel) *EducationLevelUpdate {
	elu.mutation.Where(ps...)
	return elu
}

// SetName sets the "name" field.
func (elu *EducationLevelUpdate) SetName(s string) *EducationLevelUpdate {
	elu.mutation.SetName(s)
	return elu
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (elu *EducationLevelUpdate) SetStudentID(id int) *EducationLevelUpdate {
	elu.mutation.SetStudentID(id)
	return elu
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (elu *EducationLevelUpdate) SetNillableStudentID(id *int) *EducationLevelUpdate {
	if id != nil {
		elu = elu.SetStudentID(*id)
	}
	return elu
}

// SetStudent sets the "student" edge to the Student entity.
func (elu *EducationLevelUpdate) SetStudent(s *Student) *EducationLevelUpdate {
	return elu.SetStudentID(s.ID)
}

// Mutation returns the EducationLevelMutation object of the builder.
func (elu *EducationLevelUpdate) Mutation() *EducationLevelMutation {
	return elu.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (elu *EducationLevelUpdate) ClearStudent() *EducationLevelUpdate {
	elu.mutation.ClearStudent()
	return elu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (elu *EducationLevelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EducationLevelMutation](ctx, elu.sqlSave, elu.mutation, elu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (elu *EducationLevelUpdate) SaveX(ctx context.Context) int {
	affected, err := elu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (elu *EducationLevelUpdate) Exec(ctx context.Context) error {
	_, err := elu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elu *EducationLevelUpdate) ExecX(ctx context.Context) {
	if err := elu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (elu *EducationLevelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(educationlevel.Table, educationlevel.Columns, sqlgraph.NewFieldSpec(educationlevel.FieldID, field.TypeInt))
	if ps := elu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := elu.mutation.Name(); ok {
		_spec.SetField(educationlevel.FieldName, field.TypeString, value)
	}
	if elu.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   educationlevel.StudentTable,
			Columns: []string{educationlevel.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := elu.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   educationlevel.StudentTable,
			Columns: []string{educationlevel.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, elu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{educationlevel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	elu.mutation.done = true
	return n, nil
}

// EducationLevelUpdateOne is the builder for updating a single EducationLevel entity.
type EducationLevelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EducationLevelMutation
}

// SetName sets the "name" field.
func (eluo *EducationLevelUpdateOne) SetName(s string) *EducationLevelUpdateOne {
	eluo.mutation.SetName(s)
	return eluo
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (eluo *EducationLevelUpdateOne) SetStudentID(id int) *EducationLevelUpdateOne {
	eluo.mutation.SetStudentID(id)
	return eluo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (eluo *EducationLevelUpdateOne) SetNillableStudentID(id *int) *EducationLevelUpdateOne {
	if id != nil {
		eluo = eluo.SetStudentID(*id)
	}
	return eluo
}

// SetStudent sets the "student" edge to the Student entity.
func (eluo *EducationLevelUpdateOne) SetStudent(s *Student) *EducationLevelUpdateOne {
	return eluo.SetStudentID(s.ID)
}

// Mutation returns the EducationLevelMutation object of the builder.
func (eluo *EducationLevelUpdateOne) Mutation() *EducationLevelMutation {
	return eluo.mutation
}

// ClearStudent clears the "student" edge to the Student entity.
func (eluo *EducationLevelUpdateOne) ClearStudent() *EducationLevelUpdateOne {
	eluo.mutation.ClearStudent()
	return eluo
}

// Where appends a list predicates to the EducationLevelUpdate builder.
func (eluo *EducationLevelUpdateOne) Where(ps ...predicate.EducationLevel) *EducationLevelUpdateOne {
	eluo.mutation.Where(ps...)
	return eluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eluo *EducationLevelUpdateOne) Select(field string, fields ...string) *EducationLevelUpdateOne {
	eluo.fields = append([]string{field}, fields...)
	return eluo
}

// Save executes the query and returns the updated EducationLevel entity.
func (eluo *EducationLevelUpdateOne) Save(ctx context.Context) (*EducationLevel, error) {
	return withHooks[*EducationLevel, EducationLevelMutation](ctx, eluo.sqlSave, eluo.mutation, eluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eluo *EducationLevelUpdateOne) SaveX(ctx context.Context) *EducationLevel {
	node, err := eluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eluo *EducationLevelUpdateOne) Exec(ctx context.Context) error {
	_, err := eluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eluo *EducationLevelUpdateOne) ExecX(ctx context.Context) {
	if err := eluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eluo *EducationLevelUpdateOne) sqlSave(ctx context.Context) (_node *EducationLevel, err error) {
	_spec := sqlgraph.NewUpdateSpec(educationlevel.Table, educationlevel.Columns, sqlgraph.NewFieldSpec(educationlevel.FieldID, field.TypeInt))
	id, ok := eluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`domain: missing "EducationLevel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, educationlevel.FieldID)
		for _, f := range fields {
			if !educationlevel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("domain: invalid field %q for query", f)}
			}
			if f != educationlevel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eluo.mutation.Name(); ok {
		_spec.SetField(educationlevel.FieldName, field.TypeString, value)
	}
	if eluo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   educationlevel.StudentTable,
			Columns: []string{educationlevel.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eluo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   educationlevel.StudentTable,
			Columns: []string{educationlevel.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EducationLevel{config: eluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{educationlevel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eluo.mutation.done = true
	return _node, nil
}
