package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// AllocRecordAmend holds the schema definition for the AllocRecordAmend entity.
type AllocRecordAmend struct {
	ent.Schema
}

// Fields of the AllocRecordAmend.
func (AllocRecordAmend) Fields() []ent.Field {
	return []ent.Field{
		field.Time("amend_at").Default(time.Now),
		field.String("description"),
		field.Int64("ctrl_id"),
		field.String("ctrl_name"),
		field.String("remarks"),
	}
}

// Edges of the AllocRecordAmend.
func (AllocRecordAmend) Edges() []ent.Edge {
	return nil
}
