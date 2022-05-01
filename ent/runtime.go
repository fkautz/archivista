// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/testifysec/archivist/ent/digest"
	"github.com/testifysec/archivist/ent/schema"
	"github.com/testifysec/archivist/ent/subject"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	digestFields := schema.Digest{}.Fields()
	_ = digestFields
	// digestDescAlgorithm is the schema descriptor for algorithm field.
	digestDescAlgorithm := digestFields[0].Descriptor()
	// digest.AlgorithmValidator is a validator for the "algorithm" field. It is called by the builders before save.
	digest.AlgorithmValidator = digestDescAlgorithm.Validators[0].(func(string) error)
	// digestDescValue is the schema descriptor for value field.
	digestDescValue := digestFields[1].Descriptor()
	// digest.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	digest.ValueValidator = digestDescValue.Validators[0].(func(string) error)
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescName is the schema descriptor for name field.
	subjectDescName := subjectFields[0].Descriptor()
	// subject.NameValidator is a validator for the "name" field. It is called by the builders before save.
	subject.NameValidator = subjectDescName.Validators[0].(func(string) error)
}
