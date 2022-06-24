// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/payloaddigest"
	"github.com/testifysec/archivist/ent/predicate"
)

// PayloadDigestQuery is the builder for querying PayloadDigest entities.
type PayloadDigestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PayloadDigest
	// eager-loading edges.
	withDsse *DsseQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PayloadDigestQuery builder.
func (pdq *PayloadDigestQuery) Where(ps ...predicate.PayloadDigest) *PayloadDigestQuery {
	pdq.predicates = append(pdq.predicates, ps...)
	return pdq
}

// Limit adds a limit step to the query.
func (pdq *PayloadDigestQuery) Limit(limit int) *PayloadDigestQuery {
	pdq.limit = &limit
	return pdq
}

// Offset adds an offset step to the query.
func (pdq *PayloadDigestQuery) Offset(offset int) *PayloadDigestQuery {
	pdq.offset = &offset
	return pdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pdq *PayloadDigestQuery) Unique(unique bool) *PayloadDigestQuery {
	pdq.unique = &unique
	return pdq
}

// Order adds an order step to the query.
func (pdq *PayloadDigestQuery) Order(o ...OrderFunc) *PayloadDigestQuery {
	pdq.order = append(pdq.order, o...)
	return pdq
}

// QueryDsse chains the current query on the "dsse" edge.
func (pdq *PayloadDigestQuery) QueryDsse() *DsseQuery {
	query := &DsseQuery{config: pdq.config}
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
	nodes, err := pdq.Limit(1).All(ctx)
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
	if ids, err = pdq.Limit(1).IDs(ctx); err != nil {
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
	nodes, err := pdq.Limit(2).All(ctx)
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
	if ids, err = pdq.Limit(2).IDs(ctx); err != nil {
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
	if err := pdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pdq.sqlAll(ctx)
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
func (pdq *PayloadDigestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pdq.Select(payloaddigest.FieldID).Scan(ctx, &ids); err != nil {
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
	if err := pdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pdq.sqlCount(ctx)
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
	if err := pdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pdq.sqlExist(ctx)
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
		limit:      pdq.limit,
		offset:     pdq.offset,
		order:      append([]OrderFunc{}, pdq.order...),
		predicates: append([]predicate.PayloadDigest{}, pdq.predicates...),
		withDsse:   pdq.withDsse.Clone(),
		// clone intermediate query.
		sql:    pdq.sql.Clone(),
		path:   pdq.path,
		unique: pdq.unique,
	}
}

// WithDsse tells the query-builder to eager-load the nodes that are connected to
// the "dsse" edge. The optional arguments are used to configure the query builder of the edge.
func (pdq *PayloadDigestQuery) WithDsse(opts ...func(*DsseQuery)) *PayloadDigestQuery {
	query := &DsseQuery{config: pdq.config}
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
//
func (pdq *PayloadDigestQuery) GroupBy(field string, fields ...string) *PayloadDigestGroupBy {
	group := &PayloadDigestGroupBy{config: pdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pdq.sqlQuery(ctx), nil
	}
	return group
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
//
func (pdq *PayloadDigestQuery) Select(fields ...string) *PayloadDigestSelect {
	pdq.fields = append(pdq.fields, fields...)
	return &PayloadDigestSelect{PayloadDigestQuery: pdq}
}

func (pdq *PayloadDigestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pdq.fields {
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

func (pdq *PayloadDigestQuery) sqlAll(ctx context.Context) ([]*PayloadDigest, error) {
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
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PayloadDigest{config: pdq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pdq.withDsse; query != nil {
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
		query.Where(dsse.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "dsse_payload_digests" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Dsse = n
			}
		}
	}

	return nodes, nil
}

func (pdq *PayloadDigestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pdq.querySpec()
	_spec.Node.Columns = pdq.fields
	if len(pdq.fields) > 0 {
		_spec.Unique = pdq.unique != nil && *pdq.unique
	}
	return sqlgraph.CountNodes(ctx, pdq.driver, _spec)
}

func (pdq *PayloadDigestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pdq *PayloadDigestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payloaddigest.Table,
			Columns: payloaddigest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payloaddigest.FieldID,
			},
		},
		From:   pdq.sql,
		Unique: true,
	}
	if unique := pdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pdq.fields; len(fields) > 0 {
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
	if limit := pdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pdq.offset; offset != nil {
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
	columns := pdq.fields
	if len(columns) == 0 {
		columns = payloaddigest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pdq.sql != nil {
		selector = pdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pdq.unique != nil && *pdq.unique {
		selector.Distinct()
	}
	for _, p := range pdq.predicates {
		p(selector)
	}
	for _, p := range pdq.order {
		p(selector)
	}
	if offset := pdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PayloadDigestGroupBy is the group-by builder for PayloadDigest entities.
type PayloadDigestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pdgb *PayloadDigestGroupBy) Aggregate(fns ...AggregateFunc) *PayloadDigestGroupBy {
	pdgb.fns = append(pdgb.fns, fns...)
	return pdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pdgb *PayloadDigestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pdgb.path(ctx)
	if err != nil {
		return err
	}
	pdgb.sql = query
	return pdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pdgb.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) StringsX(ctx context.Context) []string {
	v, err := pdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) StringX(ctx context.Context) string {
	v, err := pdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pdgb.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) IntsX(ctx context.Context) []int {
	v, err := pdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) IntX(ctx context.Context) int {
	v, err := pdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pdgb.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pdgb.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pdgb *PayloadDigestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pdgb *PayloadDigestGroupBy) BoolX(ctx context.Context) bool {
	v, err := pdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pdgb *PayloadDigestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pdgb.fields {
		if !payloaddigest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pdgb *PayloadDigestGroupBy) sqlQuery() *sql.Selector {
	selector := pdgb.sql.Select()
	aggregation := make([]string, 0, len(pdgb.fns))
	for _, fn := range pdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pdgb.fields)+len(pdgb.fns))
		for _, f := range pdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pdgb.fields...)...)
}

// PayloadDigestSelect is the builder for selecting fields of PayloadDigest entities.
type PayloadDigestSelect struct {
	*PayloadDigestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pds *PayloadDigestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pds.prepareQuery(ctx); err != nil {
		return err
	}
	pds.sql = pds.PayloadDigestQuery.sqlQuery(ctx)
	return pds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pds *PayloadDigestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pds.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pds *PayloadDigestSelect) StringsX(ctx context.Context) []string {
	v, err := pds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pds *PayloadDigestSelect) StringX(ctx context.Context) string {
	v, err := pds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pds.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pds *PayloadDigestSelect) IntsX(ctx context.Context) []int {
	v, err := pds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pds *PayloadDigestSelect) IntX(ctx context.Context) int {
	v, err := pds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pds.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pds *PayloadDigestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pds *PayloadDigestSelect) Float64X(ctx context.Context) float64 {
	v, err := pds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pds.fields) > 1 {
		return nil, errors.New("ent: PayloadDigestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pds *PayloadDigestSelect) BoolsX(ctx context.Context) []bool {
	v, err := pds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pds *PayloadDigestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payloaddigest.Label}
	default:
		err = fmt.Errorf("ent: PayloadDigestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pds *PayloadDigestSelect) BoolX(ctx context.Context) bool {
	v, err := pds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pds *PayloadDigestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pds.sql.Query()
	if err := pds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
