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
	"github.com/devcui/ncepu-cs-project/domain/campus"
	"github.com/devcui/ncepu-cs-project/domain/class"
	"github.com/devcui/ncepu-cs-project/domain/classleader"
	"github.com/devcui/ncepu-cs-project/domain/department"
	"github.com/devcui/ncepu-cs-project/domain/major"
	"github.com/devcui/ncepu-cs-project/domain/majordirection"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
	"github.com/devcui/ncepu-cs-project/domain/student"
	"github.com/devcui/ncepu-cs-project/domain/tutor"
)

// ClassQuery is the builder for querying Class entities.
type ClassQuery struct {
	config
	ctx                *QueryContext
	order              []OrderFunc
	inters             []Interceptor
	predicates         []predicate.Class
	withMajor          *MajorQuery
	withDepartment     *DepartmentQuery
	withCampus         *CampusQuery
	withStudent        *StudentQuery
	withClassLeader    *ClassLeaderQuery
	withTutor          *TutorQuery
	withMajorDirection *MajorDirectionQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ClassQuery builder.
func (cq *ClassQuery) Where(ps ...predicate.Class) *ClassQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ClassQuery) Limit(limit int) *ClassQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ClassQuery) Offset(offset int) *ClassQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ClassQuery) Unique(unique bool) *ClassQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ClassQuery) Order(o ...OrderFunc) *ClassQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryMajor chains the current query on the "major" edge.
func (cq *ClassQuery) QueryMajor() *MajorQuery {
	query := (&MajorClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(major.Table, major.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, class.MajorTable, class.MajorColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartment chains the current query on the "department" edge.
func (cq *ClassQuery) QueryDepartment() *DepartmentQuery {
	query := (&DepartmentClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, class.DepartmentTable, class.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCampus chains the current query on the "campus" edge.
func (cq *ClassQuery) QueryCampus() *CampusQuery {
	query := (&CampusClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(campus.Table, campus.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, class.CampusTable, class.CampusColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudent chains the current query on the "student" edge.
func (cq *ClassQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, class.StudentTable, class.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassLeader chains the current query on the "class_leader" edge.
func (cq *ClassQuery) QueryClassLeader() *ClassLeaderQuery {
	query := (&ClassLeaderClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(classleader.Table, classleader.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, class.ClassLeaderTable, class.ClassLeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTutor chains the current query on the "tutor" edge.
func (cq *ClassQuery) QueryTutor() *TutorQuery {
	query := (&TutorClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(tutor.Table, tutor.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, class.TutorTable, class.TutorColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMajorDirection chains the current query on the "major_direction" edge.
func (cq *ClassQuery) QueryMajorDirection() *MajorDirectionQuery {
	query := (&MajorDirectionClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(class.Table, class.FieldID, selector),
			sqlgraph.To(majordirection.Table, majordirection.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, class.MajorDirectionTable, class.MajorDirectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Class entity from the query.
// Returns a *NotFoundError when no Class was found.
func (cq *ClassQuery) First(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{class.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ClassQuery) FirstX(ctx context.Context) *Class {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Class ID from the query.
// Returns a *NotFoundError when no Class ID was found.
func (cq *ClassQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{class.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ClassQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Class entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Class entity is found.
// Returns a *NotFoundError when no Class entities are found.
func (cq *ClassQuery) Only(ctx context.Context) (*Class, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{class.Label}
	default:
		return nil, &NotSingularError{class.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ClassQuery) OnlyX(ctx context.Context) *Class {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Class ID in the query.
// Returns a *NotSingularError when more than one Class ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ClassQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{class.Label}
	default:
		err = &NotSingularError{class.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ClassQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Classes.
func (cq *ClassQuery) All(ctx context.Context) ([]*Class, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Class, *ClassQuery]()
	return withInterceptors[[]*Class](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ClassQuery) AllX(ctx context.Context) []*Class {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Class IDs.
func (cq *ClassQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(class.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ClassQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ClassQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ClassQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ClassQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ClassQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("domain: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ClassQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ClassQuery) Clone() *ClassQuery {
	if cq == nil {
		return nil
	}
	return &ClassQuery{
		config:             cq.config,
		ctx:                cq.ctx.Clone(),
		order:              append([]OrderFunc{}, cq.order...),
		inters:             append([]Interceptor{}, cq.inters...),
		predicates:         append([]predicate.Class{}, cq.predicates...),
		withMajor:          cq.withMajor.Clone(),
		withDepartment:     cq.withDepartment.Clone(),
		withCampus:         cq.withCampus.Clone(),
		withStudent:        cq.withStudent.Clone(),
		withClassLeader:    cq.withClassLeader.Clone(),
		withTutor:          cq.withTutor.Clone(),
		withMajorDirection: cq.withMajorDirection.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithMajor tells the query-builder to eager-load the nodes that are connected to
// the "major" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithMajor(opts ...func(*MajorQuery)) *ClassQuery {
	query := (&MajorClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withMajor = query
	return cq
}

// WithDepartment tells the query-builder to eager-load the nodes that are connected to
// the "department" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithDepartment(opts ...func(*DepartmentQuery)) *ClassQuery {
	query := (&DepartmentClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDepartment = query
	return cq
}

// WithCampus tells the query-builder to eager-load the nodes that are connected to
// the "campus" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithCampus(opts ...func(*CampusQuery)) *ClassQuery {
	query := (&CampusClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCampus = query
	return cq
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithStudent(opts ...func(*StudentQuery)) *ClassQuery {
	query := (&StudentClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withStudent = query
	return cq
}

// WithClassLeader tells the query-builder to eager-load the nodes that are connected to
// the "class_leader" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithClassLeader(opts ...func(*ClassLeaderQuery)) *ClassQuery {
	query := (&ClassLeaderClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withClassLeader = query
	return cq
}

// WithTutor tells the query-builder to eager-load the nodes that are connected to
// the "tutor" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithTutor(opts ...func(*TutorQuery)) *ClassQuery {
	query := (&TutorClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTutor = query
	return cq
}

// WithMajorDirection tells the query-builder to eager-load the nodes that are connected to
// the "major_direction" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ClassQuery) WithMajorDirection(opts ...func(*MajorDirectionQuery)) *ClassQuery {
	query := (&MajorDirectionClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withMajorDirection = query
	return cq
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
//	client.Class.Query().
//		GroupBy(class.FieldName).
//		Aggregate(domain.Count()).
//		Scan(ctx, &v)
func (cq *ClassQuery) GroupBy(field string, fields ...string) *ClassGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ClassGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = class.Label
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
//	client.Class.Query().
//		Select(class.FieldName).
//		Scan(ctx, &v)
func (cq *ClassQuery) Select(fields ...string) *ClassSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ClassSelect{ClassQuery: cq}
	sbuild.label = class.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ClassSelect configured with the given aggregations.
func (cq *ClassQuery) Aggregate(fns ...AggregateFunc) *ClassSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ClassQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("domain: uninitialized interceptor (forgotten import domain/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !class.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("domain: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Class, error) {
	var (
		nodes       = []*Class{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [7]bool{
			cq.withMajor != nil,
			cq.withDepartment != nil,
			cq.withCampus != nil,
			cq.withStudent != nil,
			cq.withClassLeader != nil,
			cq.withTutor != nil,
			cq.withMajorDirection != nil,
		}
	)
	if cq.withMajor != nil || cq.withDepartment != nil || cq.withCampus != nil || cq.withClassLeader != nil || cq.withTutor != nil || cq.withMajorDirection != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, class.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Class).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Class{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withMajor; query != nil {
		if err := cq.loadMajor(ctx, query, nodes, nil,
			func(n *Class, e *Major) { n.Edges.Major = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withDepartment; query != nil {
		if err := cq.loadDepartment(ctx, query, nodes, nil,
			func(n *Class, e *Department) { n.Edges.Department = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCampus; query != nil {
		if err := cq.loadCampus(ctx, query, nodes, nil,
			func(n *Class, e *Campus) { n.Edges.Campus = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withStudent; query != nil {
		if err := cq.loadStudent(ctx, query, nodes,
			func(n *Class) { n.Edges.Student = []*Student{} },
			func(n *Class, e *Student) { n.Edges.Student = append(n.Edges.Student, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withClassLeader; query != nil {
		if err := cq.loadClassLeader(ctx, query, nodes, nil,
			func(n *Class, e *ClassLeader) { n.Edges.ClassLeader = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withTutor; query != nil {
		if err := cq.loadTutor(ctx, query, nodes, nil,
			func(n *Class, e *Tutor) { n.Edges.Tutor = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withMajorDirection; query != nil {
		if err := cq.loadMajorDirection(ctx, query, nodes, nil,
			func(n *Class, e *MajorDirection) { n.Edges.MajorDirection = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ClassQuery) loadMajor(ctx context.Context, query *MajorQuery, nodes []*Class, init func(*Class), assign func(*Class, *Major)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].class_major == nil {
			continue
		}
		fk := *nodes[i].class_major
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(major.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_major" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadDepartment(ctx context.Context, query *DepartmentQuery, nodes []*Class, init func(*Class), assign func(*Class, *Department)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].class_department == nil {
			continue
		}
		fk := *nodes[i].class_department
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
			return fmt.Errorf(`unexpected foreign-key "class_department" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadCampus(ctx context.Context, query *CampusQuery, nodes []*Class, init func(*Class), assign func(*Class, *Campus)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].campus_class == nil {
			continue
		}
		fk := *nodes[i].campus_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(campus.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "campus_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*Class, init func(*Class), assign func(*Class, *Student)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Class)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Student(func(s *sql.Selector) {
		s.Where(sql.InValues(class.StudentColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.student_class
		if fk == nil {
			return fmt.Errorf(`foreign-key "student_class" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "student_class" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ClassQuery) loadClassLeader(ctx context.Context, query *ClassLeaderQuery, nodes []*Class, init func(*Class), assign func(*Class, *ClassLeader)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].class_leader_class == nil {
			continue
		}
		fk := *nodes[i].class_leader_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(classleader.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "class_leader_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadTutor(ctx context.Context, query *TutorQuery, nodes []*Class, init func(*Class), assign func(*Class, *Tutor)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].tutor_class == nil {
			continue
		}
		fk := *nodes[i].tutor_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tutor.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tutor_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ClassQuery) loadMajorDirection(ctx context.Context, query *MajorDirectionQuery, nodes []*Class, init func(*Class), assign func(*Class, *MajorDirection)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Class)
	for i := range nodes {
		if nodes[i].major_direction_class == nil {
			continue
		}
		fk := *nodes[i].major_direction_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(majordirection.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "major_direction_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *ClassQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for i := range fields {
			if fields[i] != class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(class.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = class.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ClassGroupBy is the group-by builder for Class entities.
type ClassGroupBy struct {
	selector
	build *ClassQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ClassGroupBy) Aggregate(fns ...AggregateFunc) *ClassGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ClassGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ClassGroupBy) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ClassSelect is the builder for selecting fields of Class entities.
type ClassSelect struct {
	*ClassQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ClassSelect) Aggregate(fns ...AggregateFunc) *ClassSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ClassSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ClassQuery, *ClassSelect](ctx, cs.ClassQuery, cs, cs.inters, v)
}

func (cs *ClassSelect) sqlScan(ctx context.Context, root *ClassQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
