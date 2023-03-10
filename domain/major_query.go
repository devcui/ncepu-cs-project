// Code generated by ent, DO NOT EDIT.

package domain

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/department"
	"github.com/devcui/ncepu-cs-project/domain/major"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
	"github.com/devcui/ncepu-cs-project/domain/student"
)

// MajorQuery is the builder for querying Major entities.
type MajorQuery struct {
	config
	ctx            *QueryContext
	order          []OrderFunc
	inters         []Interceptor
	predicates     []predicate.Major
	withDepartment *DepartmentQuery
	withStudent    *StudentQuery
	withClass      *ClassQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MajorQuery builder.
func (mq *MajorQuery) Where(ps ...predicate.Major) *MajorQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MajorQuery) Limit(limit int) *MajorQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MajorQuery) Offset(offset int) *MajorQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MajorQuery) Unique(unique bool) *MajorQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MajorQuery) Order(o ...OrderFunc) *MajorQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryDepartment chains the current query on the "department" edge.
func (mq *MajorQuery) QueryDepartment() *DepartmentQuery {
	query := (&DepartmentClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, major.DepartmentTable, major.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudent chains the current query on the "student" edge.
func (mq *MajorQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, major.StudentTable, major.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClass chains the current query on the "class" edge.
func (mq *MajorQuery) QueryClass() *ClassQuery {
	query := (&ClassClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, major.ClassTable, major.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Major entity from the query.
// Returns a *NotFoundError when no Major was found.
func (mq *MajorQuery) First(ctx context.Context) (*Major, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{major.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MajorQuery) FirstX(ctx context.Context) *Major {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Major ID from the query.
// Returns a *NotFoundError when no Major ID was found.
func (mq *MajorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{major.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MajorQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Major entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Major entity is found.
// Returns a *NotFoundError when no Major entities are found.
func (mq *MajorQuery) Only(ctx context.Context) (*Major, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{major.Label}
	default:
		return nil, &NotSingularError{major.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MajorQuery) OnlyX(ctx context.Context) *Major {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Major ID in the query.
// Returns a *NotSingularError when more than one Major ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MajorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{major.Label}
	default:
		err = &NotSingularError{major.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MajorQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Majors.
func (mq *MajorQuery) All(ctx context.Context) ([]*Major, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Major, *MajorQuery]()
	return withInterceptors[[]*Major](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MajorQuery) AllX(ctx context.Context) []*Major {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Major IDs.
func (mq *MajorQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(major.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MajorQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MajorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MajorQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MajorQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MajorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("domain: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MajorQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MajorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MajorQuery) Clone() *MajorQuery {
	if mq == nil {
		return nil
	}
	return &MajorQuery{
		config:         mq.config,
		ctx:            mq.ctx.Clone(),
		order:          append([]OrderFunc{}, mq.order...),
		inters:         append([]Interceptor{}, mq.inters...),
		predicates:     append([]predicate.Major{}, mq.predicates...),
		withDepartment: mq.withDepartment.Clone(),
		withStudent:    mq.withStudent.Clone(),
		withClass:      mq.withClass.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithDepartment tells the query-builder to eager-load the nodes that are connected to
// the "department" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithDepartment(opts ...func(*DepartmentQuery)) *MajorQuery {
	query := (&DepartmentClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withDepartment = query
	return mq
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithStudent(opts ...func(*StudentQuery)) *MajorQuery {
	query := (&StudentClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withStudent = query
	return mq
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MajorQuery) WithClass(opts ...func(*ClassQuery)) *MajorQuery {
	query := (&ClassClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withClass = query
	return mq
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
//	client.Major.Query().
//		GroupBy(major.FieldName).
//		Aggregate(domain.Count()).
//		Scan(ctx, &v)
func (mq *MajorQuery) GroupBy(field string, fields ...string) *MajorGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MajorGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = major.Label
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
//	client.Major.Query().
//		Select(major.FieldName).
//		Scan(ctx, &v)
func (mq *MajorQuery) Select(fields ...string) *MajorSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MajorSelect{MajorQuery: mq}
	sbuild.label = major.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MajorSelect configured with the given aggregations.
func (mq *MajorQuery) Aggregate(fns ...AggregateFunc) *MajorSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MajorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("domain: uninitialized interceptor (forgotten import domain/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !major.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("domain: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MajorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Major, error) {
	var (
		nodes       = []*Major{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [3]bool{
			mq.withDepartment != nil,
			mq.withStudent != nil,
			mq.withClass != nil,
		}
	)
	if mq.withDepartment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, major.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Major).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Major{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withDepartment; query != nil {
		if err := mq.loadDepartment(ctx, query, nodes, nil,
			func(n *Major, e *Department) { n.Edges.Department = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withStudent; query != nil {
		if err := mq.loadStudent(ctx, query, nodes,
			func(n *Major) { n.Edges.Student = []*Student{} },
			func(n *Major, e *Student) { n.Edges.Student = append(n.Edges.Student, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withClass; query != nil {
		if err := mq.loadClass(ctx, query, nodes,
			func(n *Major) { n.Edges.Class = []*Class{} },
			func(n *Major, e *Class) { n.Edges.Class = append(n.Edges.Class, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MajorQuery) loadDepartment(ctx context.Context, query *DepartmentQuery, nodes []*Major, init func(*Major), assign func(*Major, *Department)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Major)
	for i := range nodes {
		if nodes[i].major_department == nil {
			continue
		}
		fk := *nodes[i].major_department
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(department.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "major_department" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MajorQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*Major, init func(*Major), assign func(*Major, *Student)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Major)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Student(func(s *sql.Selector) {
		s.Where(sql.InValues(major.StudentColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.student_major
		if fk == nil {
			return fmt.Errorf(`foreign-key "student_major" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "student_major" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MajorQuery) loadClass(ctx context.Context, query *ClassQuery, nodes []*Major, init func(*Major), assign func(*Major, *Class)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Major)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Class(func(s *sql.Selector) {
		s.Where(sql.InValues(major.ClassColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.class_major
		if fk == nil {
			return fmt.Errorf(`foreign-key "class_major" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_major" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MajorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MajorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(major.Table, major.Columns, sqlgraph.NewFieldSpec(major.FieldID, field.TypeInt))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, major.FieldID)
		for i := range fields {
			if fields[i] != major.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MajorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(major.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = major.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MajorGroupBy is the group-by builder for Major entities.
type MajorGroupBy struct {
	selector
	build *MajorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MajorGroupBy) Aggregate(fns ...AggregateFunc) *MajorGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MajorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MajorQuery, *MajorGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MajorGroupBy) sqlScan(ctx context.Context, root *MajorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MajorSelect is the builder for selecting fields of Major entities.
type MajorSelect struct {
	*MajorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MajorSelect) Aggregate(fns ...AggregateFunc) *MajorSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MajorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MajorQuery, *MajorSelect](ctx, ms.MajorQuery, ms, ms.inters, v)
}

func (ms *MajorSelect) sqlScan(ctx context.Context, root *MajorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
