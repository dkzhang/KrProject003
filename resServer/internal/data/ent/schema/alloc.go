package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Alloc holds the schema definition for the Alloc entity.
type Alloc struct {
	ent.Schema
}

// Fields of the Alloc.
func (Alloc) Fields() []ent.Field {
	return []ent.Field{
		field.Time("alloc_at"),
		field.Int("node_count"),
		field.Int64("ctrl_id"),
		field.String("ctrl_name"),
		field.String("remarks"),
	}
}

// Edges of the Alloc.
func (Alloc) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", ProjectRes.Type).
			Ref("alloc"),
		edge.To("nodes", Node.Type),
		edge.To("amends", AllocRecordAmend.Type),
	}
}
