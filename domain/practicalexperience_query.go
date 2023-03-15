// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/practicalexperience"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
	"github.com/devcui/ncepu-cs-project/domain/student"
)

// PracticalExperienceQuery is the builder for querying PracticalExperience entities.
type PracticalExperienceQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.PracticalExperience
	withStudent *StudentQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PracticalExperienceQuery builder.
func (peq *PracticalExperienceQuery) Where(ps ...predicate.PracticalExperience) *PracticalExperienceQuery {
	peq.predicates = append(peq.predicates, ps...)
	return peq
}

// Limit the number of records to be returned by this query.
func (peq *PracticalExperienceQuery) Limit(limit int) *PracticalExperienceQuery {
	peq.ctx.Limit = &limit
	return peq
}

// Offset to start from.
func (peq *PracticalExperienceQuery) Offset(offset int) *PracticalExperienceQuery {
	peq.ctx.Offset = &offset
	return peq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (peq *PracticalExperienceQuery) Unique(unique bool) *PracticalExperienceQuery {
	peq.ctx.Unique = &unique
	return peq
}

// Order specifies how the records should be ordered.
func (peq *PracticalExperienceQuery) Order(o ...OrderFunc) *PracticalExperienceQuery {
	peq.order = append(peq.order, o...)
	return peq
}

// QueryStudent chains the current query on the "student" edge.
func (peq *PracticalExperienceQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: peq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := peq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := peq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(practicalexperience.Table, practicalexperience.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, practicalexperience.StudentTable, practicalexperience.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(peq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PracticalExperience entity from the query.
// Returns a *NotFoundError when no PracticalExperience was found.
func (peq *PracticalExperienceQuery) First(ctx context.Context) (*PracticalExperience, error) {
	nodes, err := peq.Limit(1).All(setContextOp(ctx, peq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{practicalexperience.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (peq *PracticalExperienceQuery) FirstX(ctx context.Context) *PracticalExperience {
	node, err := peq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PracticalExperience ID from the query.
// Returns a *NotFoundError when no PracticalExperience ID was found.
func (peq *PracticalExperienceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = peq.Limit(1).IDs(setContextOp(ctx, peq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{practicalexperience.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (peq *PracticalExperienceQuery) FirstIDX(ctx context.Context) int {
	id, err := peq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PracticalExperience entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PracticalExperience entity is found.
// Returns a *NotFoundError when no PracticalExperience entities are found.
func (peq *PracticalExperienceQuery) Only(ctx context.Context) (*PracticalExperience, error) {
	nodes, err := peq.Limit(2).All(setContextOp(ctx, peq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{practicalexperience.Label}
	default:
		return nil, &NotSingularError{practicalexperience.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (peq *PracticalExperienceQuery) OnlyX(ctx context.Context) *PracticalExperience {
	node, err := peq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PracticalExperience ID in the query.
// Returns a *NotSingularError when more than one PracticalExperience ID is found.
// Returns a *NotFoundError when no entities are found.
func (peq *PracticalExperienceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = peq.Limit(2).IDs(setContextOp(ctx, peq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{practicalexperience.Label}
	default:
		err = &NotSingularError{practicalexperience.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (peq *PracticalExperienceQuery) OnlyIDX(ctx context.Context) int {
	id, err := peq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PracticalExperiences.
func (peq *PracticalExperienceQuery) All(ctx context.Context) ([]*PracticalExperience, error) {
	ctx = setContextOp(ctx, peq.ctx, "All")
	if err := peq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PracticalExperience, *PracticalExperienceQuery]()
	return withInterceptors[[]*PracticalExperience](ctx, peq, qr, peq.inters)
}

// AllX is like All, but panics if an error occurs.
func (peq *PracticalExperienceQuery) AllX(ctx context.Context) []*PracticalExperience {
	nodes, err := peq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PracticalExperience IDs.
func (peq *PracticalExperienceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if peq.ctx.Unique == nil && peq.path != nil {
		peq.Unique(true)
	}
	ctx = setContextOp(ctx, peq.ctx, "IDs")
	if err = peq.Select(practicalexperience.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (peq *PracticalExperienceQuery) IDsX(ctx context.Context) []int {
	ids, err := peq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (peq *PracticalExperienceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, peq.ctx, "Count")
	if err := peq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, peq, querierCount[*PracticalExperienceQuery](), peq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (peq *PracticalExperienceQuery) CountX(ctx context.Context) int {
	count, err := peq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (peq *PracticalExperienceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, peq.ctx, "Exist")
	switch _, err := peq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("domain: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (peq *PracticalExperienceQuery) ExistX(ctx context.Context) bool {
	exist, err := peq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PracticalExperienceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (peq *PracticalExperienceQuery) Clone() *PracticalExperienceQuery {
	if peq == nil {
		return nil
	}
	return &PracticalExperienceQuery{
		config:      peq.config,
		ctx:         peq.ctx.Clone(),
		order:       append([]OrderFunc{}, peq.order...),
		inters:      append([]Interceptor{}, peq.inters...),
		predicates:  append([]predicate.PracticalExperience{}, peq.predicates...),
		withStudent: peq.withStudent.Clone(),
		// clone intermediate query.
		sql:  peq.sql.Clone(),
		path: peq.path,
	}
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (peq *PracticalExperienceQuery) WithStudent(opts ...func(*StudentQuery)) *PracticalExperienceQuery {
	query := (&StudentClient{config: peq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	peq.withStudent = query
	return peq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PracticalExperience.Query().
//		GroupBy(practicalexperience.FieldName).
//		Aggregate(domain.Count()).
//		Scan(ctx, &v)
func (peq *PracticalExperienceQuery) GroupBy(field string, fields ...string) *PracticalExperienceGroupBy {
	peq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PracticalExperienceGroupBy{build: peq}
	grbuild.flds = &peq.ctx.Fields
	grbuild.label = practicalexperience.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.PracticalExperience.Query().
//		Select(practicalexperience.FieldName).
//		Scan(ctx, &v)
func (peq *PracticalExperienceQuery) Select(fields ...string) *PracticalExperienceSelect {
	peq.ctx.Fields = append(peq.ctx.Fields, fields...)
	sbuild := &PracticalExperienceSelect{PracticalExperienceQuery: peq}
	sbuild.label = practicalexperience.Label
	sbuild.flds, sbuild.scan = &peq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PracticalExperienceSelect configured with the given aggregations.
func (peq *PracticalExperienceQuery) Aggregate(fns ...AggregateFunc) *PracticalExperienceSelect {
	return peq.Select().Aggregate(fns...)
}

func (peq *PracticalExperienceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range peq.inters {
		if inter == nil {
			return fmt.Errorf("domain: uninitialized interceptor (forgotten import domain/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, peq); err != nil {
				return err
			}
		}
	}
	for _, f := range peq.ctx.Fields {
		if !practicalexperience.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("domain: invalid field %q for query", f)}
		}
	}
	if peq.path != nil {
		prev, err := peq.path(ctx)
		if err != nil {
			return err
		}
		peq.sql = prev
	}
	return nil
}

func (peq *PracticalExperienceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PracticalExperience, error) {
	var (
		nodes       = []*PracticalExperience{}
		withFKs     = peq.withFKs
		_spec       = peq.querySpec()
		loadedTypes = [1]bool{
			peq.withStudent != nil,
		}
	)
	if peq.withStudent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, practicalexperience.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PracticalExperience).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PracticalExperience{config: peq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, peq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := peq.withStudent; query != nil {
		if err := peq.loadStudent(ctx, query, nodes, nil,
			func(n *PracticalExperience, e *Student) { n.Edges.Student = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (peq *PracticalExperienceQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*PracticalExperience, init func(*PracticalExperience), assign func(*PracticalExperience, *Student)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PracticalExperience)
	for i := range nodes {
		if nodes[i].practical_experience_student == nil {
			continue
		}
		fk := *nodes[i].practical_experience_student
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(student.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "practical_experience_student" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (peq *PracticalExperienceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := peq.querySpec()
	_spec.Node.Columns = peq.ctx.Fields
	if len(peq.ctx.Fields) > 0 {
		_spec.Unique = peq.ctx.Unique != nil && *peq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, peq.driver, _spec)
}

func (peq *PracticalExperienceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(practicalexperience.Table, practicalexperience.Columns, sqlgraph.NewFieldSpec(practicalexperience.FieldID, field.TypeInt))
	_spec.From = peq.sql
	if unique := peq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if peq.path != nil {
		_spec.Unique = true
	}
	if fields := peq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, practicalexperience.FieldID)
		for i := range fields {
			if fields[i] != practicalexperience.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := peq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := peq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := peq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := peq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (peq *PracticalExperienceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(peq.driver.Dialect())
	t1 := builder.Table(practicalexperience.Table)
	columns := peq.ctx.Fields
	if len(columns) == 0 {
		columns = practicalexperience.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if peq.sql != nil {
		selector = peq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if peq.ctx.Unique != nil && *peq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range peq.predicates {
		p(selector)
	}
	for _, p := range peq.order {
		p(selector)
	}
	if offset := peq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := peq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PracticalExperienceGroupBy is the group-by builder for PracticalExperience entities.
type PracticalExperienceGroupBy struct {
	selector
	build *PracticalExperienceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pegb *PracticalExperienceGroupBy) Aggregate(fns ...AggregateFunc) *PracticalExperienceGroupBy {
	pegb.fns = append(pegb.fns, fns...)
	return pegb
}

// Scan applies the selector query and scans the result into the given value.
func (pegb *PracticalExperienceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pegb.build.ctx, "GroupBy")
	if err := pegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PracticalExperienceQuery, *PracticalExperienceGroupBy](ctx, pegb.build, pegb, pegb.build.inters, v)
}

func (pegb *PracticalExperienceGroupBy) sqlScan(ctx context.Context, root *PracticalExperienceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pegb.fns))
	for _, fn := range pegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pegb.flds)+len(pegb.fns))
		for _, f := range *pegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PracticalExperienceSelect is the builder for selecting fields of PracticalExperience entities.
type PracticalExperienceSelect struct {
	*PracticalExperienceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pes *PracticalExperienceSelect) Aggregate(fns ...AggregateFunc) *PracticalExperienceSelect {
	pes.fns = append(pes.fns, fns...)
	return pes
}

// Scan applies the selector query and scans the result into the given value.
func (pes *PracticalExperienceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pes.ctx, "Select")
	if err := pes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PracticalExperienceQuery, *PracticalExperienceSelect](ctx, pes.PracticalExperienceQuery, pes, pes.inters, v)
}

func (pes *PracticalExperienceSelect) sqlScan(ctx context.Context, root *PracticalExperienceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pes.fns))
	for _, fn := range pes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
