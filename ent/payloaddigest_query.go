// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivista/ent/dsse"
	"github.com/testifysec/archivista/ent/payloaddigest"
	"github.com/testifysec/archivista/ent/predicate"
)

// PayloadDigestQuery is the builder for querying PayloadDigest entities.
type PayloadDigestQuery struct {
	config
	ctx        *QueryContext
	order      []payloaddigest.OrderOption
	inters     []Interceptor
	predicates []predicate.PayloadDigest
	withDsse   *DsseQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*PayloadDigest) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PayloadDigestQuery builder.
func (pdq *PayloadDigestQuery) Where(ps ...predicate.PayloadDigest) *PayloadDigestQuery {
	pdq.predicates = append(pdq.predicates, ps...)
	return pdq
}

// Limit the number of records to be returned by this query.
func (pdq *PayloadDigestQuery) Limit(limit int) *PayloadDigestQuery {
	pdq.ctx.Limit = &limit
	return pdq
}

// Offset to start from.
func (pdq *PayloadDigestQuery) Offset(offset int) *PayloadDigestQuery {
	pdq.ctx.Offset = &offset
	return pdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pdq *PayloadDigestQuery) Unique(unique bool) *PayloadDigestQuery {
	pdq.ctx.Unique = &unique
	return pdq
}

// Order specifies how the records should be ordered.
func (pdq *PayloadDigestQuery) Order(o ...payloaddigest.OrderOption) *PayloadDigestQuery {
	pdq.order = append(pdq.order, o...)
	return pdq
}

// QueryDsse chains the current query on the "dsse" edge.
func (pdq *PayloadDigestQuery) QueryDsse() *DsseQuery {
	query := (&DsseClient{config: pdq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payloaddigest.Table, payloaddigest.FieldID, selector),
			sqlgraph.To(dsse.Table, dsse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payloaddigest.DsseTable, payloaddigest.DsseColumn),
		)
		fromU = sqlgraph.SetNeighbors(pdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PayloadDigest entity from the query.
// Returns a *NotFoundError when no PayloadDigest was found.
func (pdq *PayloadDigestQuery) First(ctx context.Context) (*PayloadDigest, error) {
	nodes, err := pdq.Limit(1).All(setContextOp(ctx, pdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{payloaddigest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pdq *PayloadDigestQuery) FirstX(ctx context.Context) *PayloadDigest {
	node, err := pdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PayloadDigest ID from the query.
// Returns a *NotFoundError when no PayloadDigest ID was found.
func (pdq *PayloadDigestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdq.Limit(1).IDs(setContextOp(ctx, pdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{payloaddigest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pdq *PayloadDigestQuery) FirstIDX(ctx context.Context) int {
	id, err := pdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PayloadDigest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PayloadDigest entity is found.
// Returns a *NotFoundError when no PayloadDigest entities are found.
func (pdq *PayloadDigestQuery) Only(ctx context.Context) (*PayloadDigest, error) {
	nodes, err := pdq.Limit(2).All(setContextOp(ctx, pdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{payloaddigest.Label}
	default:
		return nil, &NotSingularError{payloaddigest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pdq *PayloadDigestQuery) OnlyX(ctx context.Context) *PayloadDigest {
	node, err := pdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PayloadDigest ID in the query.
// Returns a *NotSingularError when more than one PayloadDigest ID is found.
// Returns a *NotFoundError when no entities are found.
func (pdq *PayloadDigestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pdq.Limit(2).IDs(setContextOp(ctx, pdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = &NotSingularError{payloaddigest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pdq *PayloadDigestQuery) OnlyIDX(ctx context.Context) int {
	id, err := pdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PayloadDigests.
func (pdq *PayloadDigestQuery) All(ctx context.Context) ([]*PayloadDigest, error) {
	ctx = setContextOp(ctx, pdq.ctx, "All")
	if err := pdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PayloadDigest, *PayloadDigestQuery]()
	return withInterceptors[[]*PayloadDigest](ctx, pdq, qr, pdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pdq *PayloadDigestQuery) AllX(ctx context.Context) []*PayloadDigest {
	nodes, err := pdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PayloadDigest IDs.
func (pdq *PayloadDigestQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pdq.ctx.Unique == nil && pdq.path != nil {
		pdq.Unique(true)
	}
	ctx = setContextOp(ctx, pdq.ctx, "IDs")
	if err = pdq.Select(payloaddigest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pdq *PayloadDigestQuery) IDsX(ctx context.Context) []int {
	ids, err := pdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pdq *PayloadDigestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pdq.ctx, "Count")
	if err := pdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pdq, querierCount[*PayloadDigestQuery](), pdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pdq *PayloadDigestQuery) CountX(ctx context.Context) int {
	count, err := pdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pdq *PayloadDigestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pdq.ctx, "Exist")
	switch _, err := pdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pdq *PayloadDigestQuery) ExistX(ctx context.Context) bool {
	exist, err := pdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PayloadDigestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pdq *PayloadDigestQuery) Clone() *PayloadDigestQuery {
	if pdq == nil {
		return nil
	}
	return &PayloadDigestQuery{
		config:     pdq.config,
		ctx:        pdq.ctx.Clone(),
		order:      append([]payloaddigest.OrderOption{}, pdq.order...),
		inters:     append([]Interceptor{}, pdq.inters...),
		predicates: append([]predicate.PayloadDigest{}, pdq.predicates...),
		withDsse:   pdq.withDsse.Clone(),
		// clone intermediate query.
		sql:  pdq.sql.Clone(),
		path: pdq.path,
	}
}

// WithDsse tells the query-builder to eager-load the nodes that are connected to
// the "dsse" edge. The optional arguments are used to configure the query builder of the edge.
func (pdq *PayloadDigestQuery) WithDsse(opts ...func(*DsseQuery)) *PayloadDigestQuery {
	query := (&DsseClient{config: pdq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pdq.withDsse = query
	return pdq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Algorithm string `json:"algorithm,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PayloadDigest.Query().
//		GroupBy(payloaddigest.FieldAlgorithm).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pdq *PayloadDigestQuery) GroupBy(field string, fields ...string) *PayloadDigestGroupBy {
	pdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PayloadDigestGroupBy{build: pdq}
	grbuild.flds = &pdq.ctx.Fields
	grbuild.label = payloaddigest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Algorithm string `json:"algorithm,omitempty"`
//	}
//
//	client.PayloadDigest.Query().
//		Select(payloaddigest.FieldAlgorithm).
//		Scan(ctx, &v)
func (pdq *PayloadDigestQuery) Select(fields ...string) *PayloadDigestSelect {
	pdq.ctx.Fields = append(pdq.ctx.Fields, fields...)
	sbuild := &PayloadDigestSelect{PayloadDigestQuery: pdq}
	sbuild.label = payloaddigest.Label
	sbuild.flds, sbuild.scan = &pdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PayloadDigestSelect configured with the given aggregations.
func (pdq *PayloadDigestQuery) Aggregate(fns ...AggregateFunc) *PayloadDigestSelect {
	return pdq.Select().Aggregate(fns...)
}

func (pdq *PayloadDigestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pdq); err != nil {
				return err
			}
		}
	}
	for _, f := range pdq.ctx.Fields {
		if !payloaddigest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pdq.path != nil {
		prev, err := pdq.path(ctx)
		if err != nil {
			return err
		}
		pdq.sql = prev
	}
	return nil
}

func (pdq *PayloadDigestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PayloadDigest, error) {
	var (
		nodes       = []*PayloadDigest{}
		withFKs     = pdq.withFKs
		_spec       = pdq.querySpec()
		loadedTypes = [1]bool{
			pdq.withDsse != nil,
		}
	)
	if pdq.withDsse != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, payloaddigest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PayloadDigest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PayloadDigest{config: pdq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pdq.modifiers) > 0 {
		_spec.Modifiers = pdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pdq.withDsse; query != nil {
		if err := pdq.loadDsse(ctx, query, nodes, nil,
			func(n *PayloadDigest, e *Dsse) { n.Edges.Dsse = e }); err != nil {
			return nil, err
		}
	}
	for i := range pdq.loadTotal {
		if err := pdq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pdq *PayloadDigestQuery) loadDsse(ctx context.Context, query *DsseQuery, nodes []*PayloadDigest, init func(*PayloadDigest), assign func(*PayloadDigest, *Dsse)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*PayloadDigest)
	for i := range nodes {
		if nodes[i].dsse_payload_digests == nil {
			continue
		}
		fk := *nodes[i].dsse_payload_digests
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(dsse.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dsse_payload_digests" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pdq *PayloadDigestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pdq.querySpec()
	if len(pdq.modifiers) > 0 {
		_spec.Modifiers = pdq.modifiers
	}
	_spec.Node.Columns = pdq.ctx.Fields
	if len(pdq.ctx.Fields) > 0 {
		_spec.Unique = pdq.ctx.Unique != nil && *pdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pdq.driver, _spec)
}

func (pdq *PayloadDigestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(payloaddigest.Table, payloaddigest.Columns, sqlgraph.NewFieldSpec(payloaddigest.FieldID, field.TypeInt))
	_spec.From = pdq.sql
	if unique := pdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pdq.path != nil {
		_spec.Unique = true
	}
	if fields := pdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payloaddigest.FieldID)
		for i := range fields {
			if fields[i] != payloaddigest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pdq *PayloadDigestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pdq.driver.Dialect())
	t1 := builder.Table(payloaddigest.Table)
	columns := pdq.ctx.Fields
	if len(columns) == 0 {
		columns = payloaddigest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pdq.sql != nil {
		selector = pdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pdq.ctx.Unique != nil && *pdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pdq.predicates {
		p(selector)
	}
	for _, p := range pdq.order {
		p(selector)
	}
	if offset := pdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PayloadDigestGroupBy is the group-by builder for PayloadDigest entities.
type PayloadDigestGroupBy struct {
	selector
	build *PayloadDigestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pdgb *PayloadDigestGroupBy) Aggregate(fns ...AggregateFunc) *PayloadDigestGroupBy {
	pdgb.fns = append(pdgb.fns, fns...)
	return pdgb
}

// Scan applies the selector query and scans the result into the given value.
func (pdgb *PayloadDigestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pdgb.build.ctx, "GroupBy")
	if err := pdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PayloadDigestQuery, *PayloadDigestGroupBy](ctx, pdgb.build, pdgb, pdgb.build.inters, v)
}

func (pdgb *PayloadDigestGroupBy) sqlScan(ctx context.Context, root *PayloadDigestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pdgb.fns))
	for _, fn := range pdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pdgb.flds)+len(pdgb.fns))
		for _, f := range *pdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PayloadDigestSelect is the builder for selecting fields of PayloadDigest entities.
type PayloadDigestSelect struct {
	*PayloadDigestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pds *PayloadDigestSelect) Aggregate(fns ...AggregateFunc) *PayloadDigestSelect {
	pds.fns = append(pds.fns, fns...)
	return pds
}

// Scan applies the selector query and scans the result into the given value.
func (pds *PayloadDigestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pds.ctx, "Select")
	if err := pds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PayloadDigestQuery, *PayloadDigestSelect](ctx, pds.PayloadDigestQuery, pds, pds.inters, v)
}

func (pds *PayloadDigestSelect) sqlScan(ctx context.Context, root *PayloadDigestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pds.fns))
	for _, fn := range pds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
