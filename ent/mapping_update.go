// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/mapping"
	"github.com/in-toto/archivista/ent/omnitrail"
	"github.com/in-toto/archivista/ent/posix"
	"github.com/in-toto/archivista/ent/predicate"
)

// MappingUpdate is the builder for updating Mapping entities.
type MappingUpdate struct {
	config
	hooks    []Hook
	mutation *MappingMutation
}

// Where appends a list predicates to the MappingUpdate builder.
func (mu *MappingUpdate) Where(ps ...predicate.Mapping) *MappingUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetPath sets the "path" field.
func (mu *MappingUpdate) SetPath(s string) *MappingUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mu *MappingUpdate) SetNillablePath(s *string) *MappingUpdate {
	if s != nil {
		mu.SetPath(*s)
	}
	return mu
}

// SetType sets the "type" field.
func (mu *MappingUpdate) SetType(s string) *MappingUpdate {
	mu.mutation.SetType(s)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MappingUpdate) SetNillableType(s *string) *MappingUpdate {
	if s != nil {
		mu.SetType(*s)
	}
	return mu
}

// SetSha1 sets the "sha1" field.
func (mu *MappingUpdate) SetSha1(s string) *MappingUpdate {
	mu.mutation.SetSha1(s)
	return mu
}

// SetNillableSha1 sets the "sha1" field if the given value is not nil.
func (mu *MappingUpdate) SetNillableSha1(s *string) *MappingUpdate {
	if s != nil {
		mu.SetSha1(*s)
	}
	return mu
}

// SetSha256 sets the "sha256" field.
func (mu *MappingUpdate) SetSha256(s string) *MappingUpdate {
	mu.mutation.SetSha256(s)
	return mu
}

// SetNillableSha256 sets the "sha256" field if the given value is not nil.
func (mu *MappingUpdate) SetNillableSha256(s *string) *MappingUpdate {
	if s != nil {
		mu.SetSha256(*s)
	}
	return mu
}

// SetGitoidSha1 sets the "gitoidSha1" field.
func (mu *MappingUpdate) SetGitoidSha1(s string) *MappingUpdate {
	mu.mutation.SetGitoidSha1(s)
	return mu
}

// SetNillableGitoidSha1 sets the "gitoidSha1" field if the given value is not nil.
func (mu *MappingUpdate) SetNillableGitoidSha1(s *string) *MappingUpdate {
	if s != nil {
		mu.SetGitoidSha1(*s)
	}
	return mu
}

// SetGitoidSha256 sets the "gitoidSha256" field.
func (mu *MappingUpdate) SetGitoidSha256(s string) *MappingUpdate {
	mu.mutation.SetGitoidSha256(s)
	return mu
}

// SetNillableGitoidSha256 sets the "gitoidSha256" field if the given value is not nil.
func (mu *MappingUpdate) SetNillableGitoidSha256(s *string) *MappingUpdate {
	if s != nil {
		mu.SetGitoidSha256(*s)
	}
	return mu
}

// AddPosixIDs adds the "posix" edge to the Posix entity by IDs.
func (mu *MappingUpdate) AddPosixIDs(ids ...uuid.UUID) *MappingUpdate {
	mu.mutation.AddPosixIDs(ids...)
	return mu
}

// AddPosix adds the "posix" edges to the Posix entity.
func (mu *MappingUpdate) AddPosix(p ...*Posix) *MappingUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPosixIDs(ids...)
}

// SetOmnitrailID sets the "omnitrail" edge to the Omnitrail entity by ID.
func (mu *MappingUpdate) SetOmnitrailID(id uuid.UUID) *MappingUpdate {
	mu.mutation.SetOmnitrailID(id)
	return mu
}

// SetOmnitrail sets the "omnitrail" edge to the Omnitrail entity.
func (mu *MappingUpdate) SetOmnitrail(o *Omnitrail) *MappingUpdate {
	return mu.SetOmnitrailID(o.ID)
}

// Mutation returns the MappingMutation object of the builder.
func (mu *MappingUpdate) Mutation() *MappingMutation {
	return mu.mutation
}

// ClearPosix clears all "posix" edges to the Posix entity.
func (mu *MappingUpdate) ClearPosix() *MappingUpdate {
	mu.mutation.ClearPosix()
	return mu
}

// RemovePosixIDs removes the "posix" edge to Posix entities by IDs.
func (mu *MappingUpdate) RemovePosixIDs(ids ...uuid.UUID) *MappingUpdate {
	mu.mutation.RemovePosixIDs(ids...)
	return mu
}

// RemovePosix removes "posix" edges to Posix entities.
func (mu *MappingUpdate) RemovePosix(p ...*Posix) *MappingUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePosixIDs(ids...)
}

// ClearOmnitrail clears the "omnitrail" edge to the Omnitrail entity.
func (mu *MappingUpdate) ClearOmnitrail() *MappingUpdate {
	mu.mutation.ClearOmnitrail()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MappingUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MappingUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MappingUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MappingUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MappingUpdate) check() error {
	if v, ok := mu.mutation.Path(); ok {
		if err := mapping.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Mapping.path": %w`, err)}
		}
	}
	if v, ok := mu.mutation.GetType(); ok {
		if err := mapping.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Mapping.type": %w`, err)}
		}
	}
	if v, ok := mu.mutation.GitoidSha1(); ok {
		if err := mapping.GitoidSha1Validator(v); err != nil {
			return &ValidationError{Name: "gitoidSha1", err: fmt.Errorf(`ent: validator failed for field "Mapping.gitoidSha1": %w`, err)}
		}
	}
	if v, ok := mu.mutation.GitoidSha256(); ok {
		if err := mapping.GitoidSha256Validator(v); err != nil {
			return &ValidationError{Name: "gitoidSha256", err: fmt.Errorf(`ent: validator failed for field "Mapping.gitoidSha256": %w`, err)}
		}
	}
	if _, ok := mu.mutation.OmnitrailID(); mu.mutation.OmnitrailCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mapping.omnitrail"`)
	}
	return nil
}

func (mu *MappingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(mapping.Table, mapping.Columns, sqlgraph.NewFieldSpec(mapping.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.SetField(mapping.FieldPath, field.TypeString, value)
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(mapping.FieldType, field.TypeString, value)
	}
	if value, ok := mu.mutation.Sha1(); ok {
		_spec.SetField(mapping.FieldSha1, field.TypeString, value)
	}
	if value, ok := mu.mutation.Sha256(); ok {
		_spec.SetField(mapping.FieldSha256, field.TypeString, value)
	}
	if value, ok := mu.mutation.GitoidSha1(); ok {
		_spec.SetField(mapping.FieldGitoidSha1, field.TypeString, value)
	}
	if value, ok := mu.mutation.GitoidSha256(); ok {
		_spec.SetField(mapping.FieldGitoidSha256, field.TypeString, value)
	}
	if mu.mutation.PosixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPosixIDs(); len(nodes) > 0 && !mu.mutation.PosixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PosixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.OmnitrailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mapping.OmnitrailTable,
			Columns: []string{mapping.OmnitrailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(omnitrail.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OmnitrailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mapping.OmnitrailTable,
			Columns: []string{mapping.OmnitrailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(omnitrail.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MappingUpdateOne is the builder for updating a single Mapping entity.
type MappingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MappingMutation
}

// SetPath sets the "path" field.
func (muo *MappingUpdateOne) SetPath(s string) *MappingUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillablePath(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetPath(*s)
	}
	return muo
}

// SetType sets the "type" field.
func (muo *MappingUpdateOne) SetType(s string) *MappingUpdateOne {
	muo.mutation.SetType(s)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillableType(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetType(*s)
	}
	return muo
}

// SetSha1 sets the "sha1" field.
func (muo *MappingUpdateOne) SetSha1(s string) *MappingUpdateOne {
	muo.mutation.SetSha1(s)
	return muo
}

// SetNillableSha1 sets the "sha1" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillableSha1(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetSha1(*s)
	}
	return muo
}

// SetSha256 sets the "sha256" field.
func (muo *MappingUpdateOne) SetSha256(s string) *MappingUpdateOne {
	muo.mutation.SetSha256(s)
	return muo
}

// SetNillableSha256 sets the "sha256" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillableSha256(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetSha256(*s)
	}
	return muo
}

// SetGitoidSha1 sets the "gitoidSha1" field.
func (muo *MappingUpdateOne) SetGitoidSha1(s string) *MappingUpdateOne {
	muo.mutation.SetGitoidSha1(s)
	return muo
}

// SetNillableGitoidSha1 sets the "gitoidSha1" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillableGitoidSha1(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetGitoidSha1(*s)
	}
	return muo
}

// SetGitoidSha256 sets the "gitoidSha256" field.
func (muo *MappingUpdateOne) SetGitoidSha256(s string) *MappingUpdateOne {
	muo.mutation.SetGitoidSha256(s)
	return muo
}

// SetNillableGitoidSha256 sets the "gitoidSha256" field if the given value is not nil.
func (muo *MappingUpdateOne) SetNillableGitoidSha256(s *string) *MappingUpdateOne {
	if s != nil {
		muo.SetGitoidSha256(*s)
	}
	return muo
}

// AddPosixIDs adds the "posix" edge to the Posix entity by IDs.
func (muo *MappingUpdateOne) AddPosixIDs(ids ...uuid.UUID) *MappingUpdateOne {
	muo.mutation.AddPosixIDs(ids...)
	return muo
}

// AddPosix adds the "posix" edges to the Posix entity.
func (muo *MappingUpdateOne) AddPosix(p ...*Posix) *MappingUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPosixIDs(ids...)
}

// SetOmnitrailID sets the "omnitrail" edge to the Omnitrail entity by ID.
func (muo *MappingUpdateOne) SetOmnitrailID(id uuid.UUID) *MappingUpdateOne {
	muo.mutation.SetOmnitrailID(id)
	return muo
}

// SetOmnitrail sets the "omnitrail" edge to the Omnitrail entity.
func (muo *MappingUpdateOne) SetOmnitrail(o *Omnitrail) *MappingUpdateOne {
	return muo.SetOmnitrailID(o.ID)
}

// Mutation returns the MappingMutation object of the builder.
func (muo *MappingUpdateOne) Mutation() *MappingMutation {
	return muo.mutation
}

// ClearPosix clears all "posix" edges to the Posix entity.
func (muo *MappingUpdateOne) ClearPosix() *MappingUpdateOne {
	muo.mutation.ClearPosix()
	return muo
}

// RemovePosixIDs removes the "posix" edge to Posix entities by IDs.
func (muo *MappingUpdateOne) RemovePosixIDs(ids ...uuid.UUID) *MappingUpdateOne {
	muo.mutation.RemovePosixIDs(ids...)
	return muo
}

// RemovePosix removes "posix" edges to Posix entities.
func (muo *MappingUpdateOne) RemovePosix(p ...*Posix) *MappingUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePosixIDs(ids...)
}

// ClearOmnitrail clears the "omnitrail" edge to the Omnitrail entity.
func (muo *MappingUpdateOne) ClearOmnitrail() *MappingUpdateOne {
	muo.mutation.ClearOmnitrail()
	return muo
}

// Where appends a list predicates to the MappingUpdate builder.
func (muo *MappingUpdateOne) Where(ps ...predicate.Mapping) *MappingUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MappingUpdateOne) Select(field string, fields ...string) *MappingUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Mapping entity.
func (muo *MappingUpdateOne) Save(ctx context.Context) (*Mapping, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MappingUpdateOne) SaveX(ctx context.Context) *Mapping {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MappingUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MappingUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MappingUpdateOne) check() error {
	if v, ok := muo.mutation.Path(); ok {
		if err := mapping.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Mapping.path": %w`, err)}
		}
	}
	if v, ok := muo.mutation.GetType(); ok {
		if err := mapping.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Mapping.type": %w`, err)}
		}
	}
	if v, ok := muo.mutation.GitoidSha1(); ok {
		if err := mapping.GitoidSha1Validator(v); err != nil {
			return &ValidationError{Name: "gitoidSha1", err: fmt.Errorf(`ent: validator failed for field "Mapping.gitoidSha1": %w`, err)}
		}
	}
	if v, ok := muo.mutation.GitoidSha256(); ok {
		if err := mapping.GitoidSha256Validator(v); err != nil {
			return &ValidationError{Name: "gitoidSha256", err: fmt.Errorf(`ent: validator failed for field "Mapping.gitoidSha256": %w`, err)}
		}
	}
	if _, ok := muo.mutation.OmnitrailID(); muo.mutation.OmnitrailCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Mapping.omnitrail"`)
	}
	return nil
}

func (muo *MappingUpdateOne) sqlSave(ctx context.Context) (_node *Mapping, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(mapping.Table, mapping.Columns, sqlgraph.NewFieldSpec(mapping.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Mapping.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mapping.FieldID)
		for _, f := range fields {
			if !mapping.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mapping.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Path(); ok {
		_spec.SetField(mapping.FieldPath, field.TypeString, value)
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(mapping.FieldType, field.TypeString, value)
	}
	if value, ok := muo.mutation.Sha1(); ok {
		_spec.SetField(mapping.FieldSha1, field.TypeString, value)
	}
	if value, ok := muo.mutation.Sha256(); ok {
		_spec.SetField(mapping.FieldSha256, field.TypeString, value)
	}
	if value, ok := muo.mutation.GitoidSha1(); ok {
		_spec.SetField(mapping.FieldGitoidSha1, field.TypeString, value)
	}
	if value, ok := muo.mutation.GitoidSha256(); ok {
		_spec.SetField(mapping.FieldGitoidSha256, field.TypeString, value)
	}
	if muo.mutation.PosixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPosixIDs(); len(nodes) > 0 && !muo.mutation.PosixCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PosixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   mapping.PosixTable,
			Columns: []string{mapping.PosixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(posix.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.OmnitrailCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mapping.OmnitrailTable,
			Columns: []string{mapping.OmnitrailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(omnitrail.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OmnitrailIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mapping.OmnitrailTable,
			Columns: []string{mapping.OmnitrailColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(omnitrail.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Mapping{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mapping.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
