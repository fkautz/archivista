// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/predicate"
	"github.com/testifysec/archivist/ent/statement"
	"github.com/testifysec/archivist/ent/subject"
)

// StatementUpdate is the builder for updating Statement entities.
type StatementUpdate struct {
	config
	hooks    []Hook
	mutation *StatementMutation
}

// Where appends a list predicates to the StatementUpdate builder.
func (su *StatementUpdate) Where(ps ...predicate.Statement) *StatementUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetStatement sets the "statement" field.
func (su *StatementUpdate) SetStatement(s string) *StatementUpdate {
	su.mutation.SetStatement(s)
	return su
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (su *StatementUpdate) AddSubjectIDs(ids ...int) *StatementUpdate {
	su.mutation.AddSubjectIDs(ids...)
	return su
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (su *StatementUpdate) AddSubjects(s ...*Subject) *StatementUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSubjectIDs(ids...)
}

// Mutation returns the StatementMutation object of the builder.
func (su *StatementUpdate) Mutation() *StatementMutation {
	return su.mutation
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (su *StatementUpdate) ClearSubjects() *StatementUpdate {
	su.mutation.ClearSubjects()
	return su
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (su *StatementUpdate) RemoveSubjectIDs(ids ...int) *StatementUpdate {
	su.mutation.RemoveSubjectIDs(ids...)
	return su
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (su *StatementUpdate) RemoveSubjects(s ...*Subject) *StatementUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSubjectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StatementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatementUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatementUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatementUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StatementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statement.Table,
			Columns: statement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statement.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Statement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statement.FieldStatement,
		})
	}
	if su.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !su.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StatementUpdateOne is the builder for updating a single Statement entity.
type StatementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatementMutation
}

// SetStatement sets the "statement" field.
func (suo *StatementUpdateOne) SetStatement(s string) *StatementUpdateOne {
	suo.mutation.SetStatement(s)
	return suo
}

// AddSubjectIDs adds the "subjects" edge to the Subject entity by IDs.
func (suo *StatementUpdateOne) AddSubjectIDs(ids ...int) *StatementUpdateOne {
	suo.mutation.AddSubjectIDs(ids...)
	return suo
}

// AddSubjects adds the "subjects" edges to the Subject entity.
func (suo *StatementUpdateOne) AddSubjects(s ...*Subject) *StatementUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSubjectIDs(ids...)
}

// Mutation returns the StatementMutation object of the builder.
func (suo *StatementUpdateOne) Mutation() *StatementMutation {
	return suo.mutation
}

// ClearSubjects clears all "subjects" edges to the Subject entity.
func (suo *StatementUpdateOne) ClearSubjects() *StatementUpdateOne {
	suo.mutation.ClearSubjects()
	return suo
}

// RemoveSubjectIDs removes the "subjects" edge to Subject entities by IDs.
func (suo *StatementUpdateOne) RemoveSubjectIDs(ids ...int) *StatementUpdateOne {
	suo.mutation.RemoveSubjectIDs(ids...)
	return suo
}

// RemoveSubjects removes "subjects" edges to Subject entities.
func (suo *StatementUpdateOne) RemoveSubjects(s ...*Subject) *StatementUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSubjectIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StatementUpdateOne) Select(field string, fields ...string) *StatementUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Statement entity.
func (suo *StatementUpdateOne) Save(ctx context.Context) (*Statement, error) {
	var (
		err  error
		node *Statement
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatementUpdateOne) SaveX(ctx context.Context) *Statement {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StatementUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatementUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StatementUpdateOne) sqlSave(ctx context.Context) (_node *Statement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statement.Table,
			Columns: statement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statement.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Statement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statement.FieldID)
		for _, f := range fields {
			if !statement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != statement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Statement(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statement.FieldStatement,
		})
	}
	if suo.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSubjectsIDs(); len(nodes) > 0 && !suo.mutation.SubjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statement.SubjectsTable,
			Columns: statement.SubjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Statement{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
