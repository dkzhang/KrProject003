package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AllocRecord holds the schema definition for the AllocRecord entity.
type AllocRecord struct {
	ent.Schema
}

// Fields of the AllocRecord.
func (AllocRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Time("alloc_at"),
		field.Time("end_at"),
		field.String("node_str"),
		field.Int("node_count"),
		field.Int64("ctrl_id"),
		field.String("ctrl_name"),
		field.String("remarks"),
	}
}

// Edges of the AllocRecord.
func (AllocRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amends", AllocRecordAmend.Type),
	}
}
