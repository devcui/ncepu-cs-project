// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/student"
	"github.com/devcui/ncepu-cs-project/domain/tutor"
)

// TutorCreate is the builder for creating a Tutor entity.
type TutorCreate struct {
	config
	mutation *TutorMutation
	hooks    []Hook
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (tc *TutorCreate) SetClassID(id int) *TutorCreate {
	tc.mutation.SetClassID(id)
	return tc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (tc *TutorCreate) SetNillableClassID(id *int) *TutorCreate {
	if id != nil {
		tc = tc.SetClassID(*id)
	}
	return tc
}

// SetClass sets the "class" edge to the Class entity.
func (tc *TutorCreate) SetClass(c *Class) *TutorCreate {
	return tc.SetClassID(c.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (tc *TutorCreate) SetStudentID(id int) *TutorCreate {
	tc.mutation.SetStudentID(id)
	return tc
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (tc *TutorCreate) SetNillableStudentID(id *int) *TutorCreate {
	if id != nil {
		tc = tc.SetStudentID(*id)
	}
	return tc
}

// SetStudent sets the "student" edge to the Student entity.
func (tc *TutorCreate) SetStudent(s *Student) *TutorCreate {
	return tc.SetStudentID(s.ID)
}

// Mutation returns the TutorMutation object of the builder.
func (tc *TutorCreate) Mutation() *TutorMutation {
	return tc.mutation
}

// Save creates the Tutor in the database.
func (tc *TutorCreate) Save(ctx context.Context) (*Tutor, error) {
	return withHooks[*Tutor, TutorMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TutorCreate) SaveX(ctx context.Context) *Tutor {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TutorCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TutorCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TutorCreate) check() error {
	return nil
}

func (tc *TutorCreate) sqlSave(ctx context.Context) (*Tutor, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TutorCreate) createSpec() (*Tutor, *sqlgraph.CreateSpec) {
	var (
		_node = &Tutor{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tutor.Table, sqlgraph.NewFieldSpec(tutor.FieldID, field.TypeInt))
	)
	if nodes := tc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tutor.ClassTable,
			Columns: []string{tutor.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tutor.StudentTable,
			Columns: []string{tutor.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TutorCreateBulk is the builder for creating many Tutor entities in bulk.
type TutorCreateBulk struct {
	config
	builders []*TutorCreate
}

// Save creates the Tutor entities in the database.
func (tcb *TutorCreateBulk) Save(ctx context.Context) ([]*Tutor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tutor, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TutorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TutorCreateBulk) SaveX(ctx context.Context) []*Tutor {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TutorCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TutorCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}