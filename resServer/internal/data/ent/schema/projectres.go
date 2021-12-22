package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ProjectRes holds the schema definition for the ProjectRes entity.
type ProjectRes struct {
	ent.Schema
}

// Fields of the ProjectRes.
func (ProjectRes) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("project_id"),
		field.String("name"),
		field.String("code"),

		field.Int("cpu_in_use"),
		field.Int("gpu_in_use"),
		field.Int("storage_in_use"),

		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the ProjectRes.
func (ProjectRes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("alloc", Alloc.Type),
		edge.To("alloc_records", AllocRecord.Type),
	}
}
