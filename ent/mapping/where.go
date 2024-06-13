// Code generated by ent, DO NOT EDIT.

package mapping

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldID, id))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldPath, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldType, v))
}

// Sha1 applies equality check predicate on the "sha1" field. It's identical to Sha1EQ.
func Sha1(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldSha1, v))
}

// Sha256 applies equality check predicate on the "sha256" field. It's identical to Sha256EQ.
func Sha256(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldSha256, v))
}

// GitoidSha1 applies equality check predicate on the "gitoidSha1" field. It's identical to GitoidSha1EQ.
func GitoidSha1(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldGitoidSha1, v))
}

// GitoidSha256 applies equality check predicate on the "gitoidSha256" field. It's identical to GitoidSha256EQ.
func GitoidSha256(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldGitoidSha256, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldPath, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldType, v))
}

// Sha1EQ applies the EQ predicate on the "sha1" field.
func Sha1EQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldSha1, v))
}

// Sha1NEQ applies the NEQ predicate on the "sha1" field.
func Sha1NEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldSha1, v))
}

// Sha1In applies the In predicate on the "sha1" field.
func Sha1In(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldSha1, vs...))
}

// Sha1NotIn applies the NotIn predicate on the "sha1" field.
func Sha1NotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldSha1, vs...))
}

// Sha1GT applies the GT predicate on the "sha1" field.
func Sha1GT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldSha1, v))
}

// Sha1GTE applies the GTE predicate on the "sha1" field.
func Sha1GTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldSha1, v))
}

// Sha1LT applies the LT predicate on the "sha1" field.
func Sha1LT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldSha1, v))
}

// Sha1LTE applies the LTE predicate on the "sha1" field.
func Sha1LTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldSha1, v))
}

// Sha1Contains applies the Contains predicate on the "sha1" field.
func Sha1Contains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldSha1, v))
}

// Sha1HasPrefix applies the HasPrefix predicate on the "sha1" field.
func Sha1HasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldSha1, v))
}

// Sha1HasSuffix applies the HasSuffix predicate on the "sha1" field.
func Sha1HasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldSha1, v))
}

// Sha1EqualFold applies the EqualFold predicate on the "sha1" field.
func Sha1EqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldSha1, v))
}

// Sha1ContainsFold applies the ContainsFold predicate on the "sha1" field.
func Sha1ContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldSha1, v))
}

// Sha256EQ applies the EQ predicate on the "sha256" field.
func Sha256EQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldSha256, v))
}

// Sha256NEQ applies the NEQ predicate on the "sha256" field.
func Sha256NEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldSha256, v))
}

// Sha256In applies the In predicate on the "sha256" field.
func Sha256In(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldSha256, vs...))
}

// Sha256NotIn applies the NotIn predicate on the "sha256" field.
func Sha256NotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldSha256, vs...))
}

// Sha256GT applies the GT predicate on the "sha256" field.
func Sha256GT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldSha256, v))
}

// Sha256GTE applies the GTE predicate on the "sha256" field.
func Sha256GTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldSha256, v))
}

// Sha256LT applies the LT predicate on the "sha256" field.
func Sha256LT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldSha256, v))
}

// Sha256LTE applies the LTE predicate on the "sha256" field.
func Sha256LTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldSha256, v))
}

// Sha256Contains applies the Contains predicate on the "sha256" field.
func Sha256Contains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldSha256, v))
}

// Sha256HasPrefix applies the HasPrefix predicate on the "sha256" field.
func Sha256HasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldSha256, v))
}

// Sha256HasSuffix applies the HasSuffix predicate on the "sha256" field.
func Sha256HasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldSha256, v))
}

// Sha256EqualFold applies the EqualFold predicate on the "sha256" field.
func Sha256EqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldSha256, v))
}

// Sha256ContainsFold applies the ContainsFold predicate on the "sha256" field.
func Sha256ContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldSha256, v))
}

// GitoidSha1EQ applies the EQ predicate on the "gitoidSha1" field.
func GitoidSha1EQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldGitoidSha1, v))
}

// GitoidSha1NEQ applies the NEQ predicate on the "gitoidSha1" field.
func GitoidSha1NEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldGitoidSha1, v))
}

// GitoidSha1In applies the In predicate on the "gitoidSha1" field.
func GitoidSha1In(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldGitoidSha1, vs...))
}

// GitoidSha1NotIn applies the NotIn predicate on the "gitoidSha1" field.
func GitoidSha1NotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldGitoidSha1, vs...))
}

// GitoidSha1GT applies the GT predicate on the "gitoidSha1" field.
func GitoidSha1GT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldGitoidSha1, v))
}

// GitoidSha1GTE applies the GTE predicate on the "gitoidSha1" field.
func GitoidSha1GTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldGitoidSha1, v))
}

// GitoidSha1LT applies the LT predicate on the "gitoidSha1" field.
func GitoidSha1LT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldGitoidSha1, v))
}

// GitoidSha1LTE applies the LTE predicate on the "gitoidSha1" field.
func GitoidSha1LTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldGitoidSha1, v))
}

// GitoidSha1Contains applies the Contains predicate on the "gitoidSha1" field.
func GitoidSha1Contains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldGitoidSha1, v))
}

// GitoidSha1HasPrefix applies the HasPrefix predicate on the "gitoidSha1" field.
func GitoidSha1HasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldGitoidSha1, v))
}

// GitoidSha1HasSuffix applies the HasSuffix predicate on the "gitoidSha1" field.
func GitoidSha1HasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldGitoidSha1, v))
}

// GitoidSha1EqualFold applies the EqualFold predicate on the "gitoidSha1" field.
func GitoidSha1EqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldGitoidSha1, v))
}

// GitoidSha1ContainsFold applies the ContainsFold predicate on the "gitoidSha1" field.
func GitoidSha1ContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldGitoidSha1, v))
}

// GitoidSha256EQ applies the EQ predicate on the "gitoidSha256" field.
func GitoidSha256EQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEQ(FieldGitoidSha256, v))
}

// GitoidSha256NEQ applies the NEQ predicate on the "gitoidSha256" field.
func GitoidSha256NEQ(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNEQ(FieldGitoidSha256, v))
}

// GitoidSha256In applies the In predicate on the "gitoidSha256" field.
func GitoidSha256In(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldIn(FieldGitoidSha256, vs...))
}

// GitoidSha256NotIn applies the NotIn predicate on the "gitoidSha256" field.
func GitoidSha256NotIn(vs ...string) predicate.Mapping {
	return predicate.Mapping(sql.FieldNotIn(FieldGitoidSha256, vs...))
}

// GitoidSha256GT applies the GT predicate on the "gitoidSha256" field.
func GitoidSha256GT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGT(FieldGitoidSha256, v))
}

// GitoidSha256GTE applies the GTE predicate on the "gitoidSha256" field.
func GitoidSha256GTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldGTE(FieldGitoidSha256, v))
}

// GitoidSha256LT applies the LT predicate on the "gitoidSha256" field.
func GitoidSha256LT(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLT(FieldGitoidSha256, v))
}

// GitoidSha256LTE applies the LTE predicate on the "gitoidSha256" field.
func GitoidSha256LTE(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldLTE(FieldGitoidSha256, v))
}

// GitoidSha256Contains applies the Contains predicate on the "gitoidSha256" field.
func GitoidSha256Contains(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContains(FieldGitoidSha256, v))
}

// GitoidSha256HasPrefix applies the HasPrefix predicate on the "gitoidSha256" field.
func GitoidSha256HasPrefix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasPrefix(FieldGitoidSha256, v))
}

// GitoidSha256HasSuffix applies the HasSuffix predicate on the "gitoidSha256" field.
func GitoidSha256HasSuffix(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldHasSuffix(FieldGitoidSha256, v))
}

// GitoidSha256EqualFold applies the EqualFold predicate on the "gitoidSha256" field.
func GitoidSha256EqualFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldEqualFold(FieldGitoidSha256, v))
}

// GitoidSha256ContainsFold applies the ContainsFold predicate on the "gitoidSha256" field.
func GitoidSha256ContainsFold(v string) predicate.Mapping {
	return predicate.Mapping(sql.FieldContainsFold(FieldGitoidSha256, v))
}

// HasPosix applies the HasEdge predicate on the "posix" edge.
func HasPosix() predicate.Mapping {
	return predicate.Mapping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PosixTable, PosixColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPosixWith applies the HasEdge predicate on the "posix" edge with a given conditions (other predicates).
func HasPosixWith(preds ...predicate.Posix) predicate.Mapping {
	return predicate.Mapping(func(s *sql.Selector) {
		step := newPosixStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOmnitrail applies the HasEdge predicate on the "omnitrail" edge.
func HasOmnitrail() predicate.Mapping {
	return predicate.Mapping(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OmnitrailTable, OmnitrailColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOmnitrailWith applies the HasEdge predicate on the "omnitrail" edge with a given conditions (other predicates).
func HasOmnitrailWith(preds ...predicate.Omnitrail) predicate.Mapping {
	return predicate.Mapping(func(s *sql.Selector) {
		step := newOmnitrailStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Mapping) predicate.Mapping {
	return predicate.Mapping(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Mapping) predicate.Mapping {
	return predicate.Mapping(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Mapping) predicate.Mapping {
	return predicate.Mapping(sql.NotPredicates(p))
}
