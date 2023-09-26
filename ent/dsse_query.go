// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivista/ent/dsse"
	"github.com/testifysec/archivista/ent/payloaddigest"
	"github.com/testifysec/archivista/ent/predicate"
	"github.com/testifysec/archivista/ent/signature"
	"github.com/testifysec/archivista/ent/statement"
)

// DsseQuery is the builder for querying Dsse entities.
type DsseQuery struct {
	config
	ctx                     *QueryContext
	order                   []dsse.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Dsse
	withStatement           *StatementQuery
	withSignatures          *SignatureQuery
	withPayloadDigests      *PayloadDigestQuery
	withFKs                 bool
	modifiers               []func(*sql.Selector)
	loadTotal               []func(context.Context, []*Dsse) error
	withNamedSignatures     map[string]*SignatureQuery
	withNamedPayloadDigests map[string]*PayloadDigestQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DsseQuery builder.
func (dq *DsseQuery) Where(ps ...predicate.Dsse) *DsseQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DsseQuery) Limit(limit int) *DsseQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DsseQuery) Offset(offset int) *DsseQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DsseQuery) Unique(unique bool) *DsseQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DsseQuery) Order(o ...dsse.OrderOption) *DsseQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryStatement chains the current query on the "statement" edge.
func (dq *DsseQuery) QueryStatement() *StatementQuery {
	query := (&StatementClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dsse.Table, dsse.FieldID, selector),
			sqlgraph.To(statement.Table, statement.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, dsse.StatementTable, dsse.StatementColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySignatures chains the current query on the "signatures" edge.
func (dq *DsseQuery) QuerySignatures() *SignatureQuery {
	query := (&SignatureClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dsse.Table, dsse.FieldID, selector),
			sqlgraph.To(signature.Table, signature.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dsse.SignaturesTable, dsse.SignaturesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayloadDigests chains the current query on the "payload_digests" edge.
func (dq *DsseQuery) QueryPayloadDigests() *PayloadDigestQuery {
	query := (&PayloadDigestClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dsse.Table, dsse.FieldID, selector),
			sqlgraph.To(payloaddigest.Table, payloaddigest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dsse.PayloadDigestsTable, dsse.PayloadDigestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Dsse entity from the query.
// Returns a *NotFoundError when no Dsse was found.
func (dq *DsseQuery) First(ctx context.Context) (*Dsse, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dsse.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DsseQuery) FirstX(ctx context.Context) *Dsse {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Dsse ID from the query.
// Returns a *NotFoundError when no Dsse ID was found.
func (dq *DsseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dsse.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DsseQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Dsse entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Dsse entity is found.
// Returns a *NotFoundError when no Dsse entities are found.
func (dq *DsseQuery) Only(ctx context.Context) (*Dsse, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dsse.Label}
	default:
		return nil, &NotSingularError{dsse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DsseQuery) OnlyX(ctx context.Context) *Dsse {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Dsse ID in the query.
// Returns a *NotSingularError when more than one Dsse ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DsseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dsse.Label}
	default:
		err = &NotSingularError{dsse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DsseQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Dsses.
func (dq *DsseQuery) All(ctx context.Context) ([]*Dsse, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Dsse, *DsseQuery]()
	return withInterceptors[[]*Dsse](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DsseQuery) AllX(ctx context.Context) []*Dsse {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Dsse IDs.
func (dq *DsseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(dsse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DsseQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DsseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DsseQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DsseQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DsseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DsseQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DsseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DsseQuery) Clone() *DsseQuery {
	if dq == nil {
		return nil
	}
	return &DsseQuery{
		config:             dq.config,
		ctx:                dq.ctx.Clone(),
		order:              append([]dsse.OrderOption{}, dq.order...),
		inters:             append([]Interceptor{}, dq.inters...),
		predicates:         append([]predicate.Dsse{}, dq.predicates...),
		withStatement:      dq.withStatement.Clone(),
		withSignatures:     dq.withSignatures.Clone(),
		withPayloadDigests: dq.withPayloadDigests.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithStatement tells the query-builder to eager-load the nodes that are connected to
// the "statement" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DsseQuery) WithStatement(opts ...func(*StatementQuery)) *DsseQuery {
	query := (&StatementClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withStatement = query
	return dq
}

// WithSignatures tells the query-builder to eager-load the nodes that are connected to
// the "signatures" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DsseQuery) WithSignatures(opts ...func(*SignatureQuery)) *DsseQuery {
	query := (&SignatureClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withSignatures = query
	return dq
}

// WithPayloadDigests tells the query-builder to eager-load the nodes that are connected to
// the "payload_digests" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DsseQuery) WithPayloadDigests(opts ...func(*PayloadDigestQuery)) *DsseQuery {
	query := (&PayloadDigestClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withPayloadDigests = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GitoidSha256 string `json:"gitoid_sha256,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Dsse.Query().
//		GroupBy(dsse.FieldGitoidSha256).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DsseQuery) GroupBy(field string, fields ...string) *DsseGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DsseGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = dsse.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GitoidSha256 string `json:"gitoid_sha256,omitempty"`
//	}
//
//	client.Dsse.Query().
//		Select(dsse.FieldGitoidSha256).
//		Scan(ctx, &v)
func (dq *DsseQuery) Select(fields ...string) *DsseSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DsseSelect{DsseQuery: dq}
	sbuild.label = dsse.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DsseSelect configured with the given aggregations.
func (dq *DsseQuery) Aggregate(fns ...AggregateFunc) *DsseSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DsseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !dsse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DsseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Dsse, error) {
	var (
		nodes       = []*Dsse{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withStatement != nil,
			dq.withSignatures != nil,
			dq.withPayloadDigests != nil,
		}
	)
	if dq.withStatement != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, dsse.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Dsse).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Dsse{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withStatement; query != nil {
		if err := dq.loadStatement(ctx, query, nodes, nil,
			func(n *Dsse, e *Statement) { n.Edges.Statement = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withSignatures; query != nil {
		if err := dq.loadSignatures(ctx, query, nodes,
			func(n *Dsse) { n.Edges.Signatures = []*Signature{} },
			func(n *Dsse, e *Signature) { n.Edges.Signatures = append(n.Edges.Signatures, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withPayloadDigests; query != nil {
		if err := dq.loadPayloadDigests(ctx, query, nodes,
			func(n *Dsse) { n.Edges.PayloadDigests = []*PayloadDigest{} },
			func(n *Dsse, e *PayloadDigest) { n.Edges.PayloadDigests = append(n.Edges.PayloadDigests, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedSignatures {
		if err := dq.loadSignatures(ctx, query, nodes,
			func(n *Dsse) { n.appendNamedSignatures(name) },
			func(n *Dsse, e *Signature) { n.appendNamedSignatures(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedPayloadDigests {
		if err := dq.loadPayloadDigests(ctx, query, nodes,
			func(n *Dsse) { n.appendNamedPayloadDigests(name) },
			func(n *Dsse, e *PayloadDigest) { n.appendNamedPayloadDigests(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DsseQuery) loadStatement(ctx context.Context, query *StatementQuery, nodes []*Dsse, init func(*Dsse), assign func(*Dsse, *Statement)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Dsse)
	for i := range nodes {
		if nodes[i].dsse_statement == nil {
			continue
		}
		fk := *nodes[i].dsse_statement
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(statement.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "dsse_statement" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DsseQuery) loadSignatures(ctx context.Context, query *SignatureQuery, nodes []*Dsse, init func(*Dsse), assign func(*Dsse, *Signature)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Dsse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Signature(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dsse.SignaturesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dsse_signatures
		if fk == nil {
			return fmt.Errorf(`foreign-key "dsse_signatures" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dsse_signatures" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DsseQuery) loadPayloadDigests(ctx context.Context, query *PayloadDigestQuery, nodes []*Dsse, init func(*Dsse), assign func(*Dsse, *PayloadDigest)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Dsse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PayloadDigest(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dsse.PayloadDigestsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dsse_payload_digests
		if fk == nil {
			return fmt.Errorf(`foreign-key "dsse_payload_digests" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dsse_payload_digests" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DsseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DsseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dsse.Table, dsse.Columns, sqlgraph.NewFieldSpec(dsse.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dsse.FieldID)
		for i := range fields {
			if fields[i] != dsse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DsseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(dsse.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = dsse.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSignatures tells the query-builder to eager-load the nodes that are connected to the "signatures"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DsseQuery) WithNamedSignatures(name string, opts ...func(*SignatureQuery)) *DsseQuery {
	query := (&SignatureClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedSignatures == nil {
		dq.withNamedSignatures = make(map[string]*SignatureQuery)
	}
	dq.withNamedSignatures[name] = query
	return dq
}

// WithNamedPayloadDigests tells the query-builder to eager-load the nodes that are connected to the "payload_digests"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DsseQuery) WithNamedPayloadDigests(name string, opts ...func(*PayloadDigestQuery)) *DsseQuery {
	query := (&PayloadDigestClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedPayloadDigests == nil {
		dq.withNamedPayloadDigests = make(map[string]*PayloadDigestQuery)
	}
	dq.withNamedPayloadDigests[name] = query
	return dq
}

// DsseGroupBy is the group-by builder for Dsse entities.
type DsseGroupBy struct {
	selector
	build *DsseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DsseGroupBy) Aggregate(fns ...AggregateFunc) *DsseGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DsseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DsseQuery, *DsseGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DsseGroupBy) sqlScan(ctx context.Context, root *DsseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DsseSelect is the builder for selecting fields of Dsse entities.
type DsseSelect struct {
	*DsseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DsseSelect) Aggregate(fns ...AggregateFunc) *DsseSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DsseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DsseQuery, *DsseSelect](ctx, ds.DsseQuery, ds, ds.inters, v)
}

func (ds *DsseSelect) sqlScan(ctx context.Context, root *DsseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
