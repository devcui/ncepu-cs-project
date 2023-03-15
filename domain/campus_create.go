// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/campus"
	"github.com/devcui/ncepu-cs-project/domain/class"
)

// CampusCreate is the builder for creating a Campus entity.
type CampusCreate struct {
	config
	mutation *CampusMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CampusCreate) SetName(s string) *CampusCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *CampusCreate) SetAddress(s string) *CampusCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cc *CampusCreate) SetClassID(id int) *CampusCreate {
	cc.mutation.SetClassID(id)
	return cc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (cc *CampusCreate) SetNillableClassID(id *int) *CampusCreate {
	if id != nil {
		cc = cc.SetClassID(*id)
	}
	return cc
}

// SetClass sets the "class" edge to the Class entity.
func (cc *CampusCreate) SetClass(c *Class) *CampusCreate {
	return cc.SetClassID(c.ID)
}

// Mutation returns the CampusMutation object of the builder.
func (cc *CampusCreate) Mutation() *CampusMutation {
	return cc.mutation
}

// Save creates the Campus in the database.
func (cc *CampusCreate) Save(ctx context.Context) (*Campus, error) {
	return withHooks[*Campus, CampusMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CampusCreate) SaveX(ctx context.Context) *Campus {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CampusCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CampusCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CampusCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`domain: missing required field "Campus.name"`)}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`domain: missing required field "Campus.address"`)}
	}
	return nil
}

func (cc *CampusCreate) sqlSave(ctx context.Context) (*Campus, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CampusCreate) createSpec() (*Campus, *sqlgraph.CreateSpec) {
	var (
		_node = &Campus{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(campus.Table, sqlgraph.NewFieldSpec(campus.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(campus.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.SetField(campus.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if nodes := cc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   campus.ClassTable,
			Columns: []string{campus.ClassColumn},
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
	return _node, _spec
}

// CampusCreateBulk is the builder for creating many Campus entities in bulk.
type CampusCreateBulk struct {
	config
	builders []*CampusCreate
}

// Save creates the Campus entities in the database.
func (ccb *CampusCreateBulk) Save(ctx context.Context) ([]*Campus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Campus, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CampusMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CampusCreateBulk) SaveX(ctx context.Context) []*Campus {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CampusCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CampusCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
