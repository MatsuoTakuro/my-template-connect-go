package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("store_cd"),
		field.Int("company_cd"),
		field.String("store_name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(50)",
			}),
		field.String("address").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(255)",
			}),
		field.Float("latitude").
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(7,4)",
			}),
		field.Float("longitude").
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(7,4)",
			}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return nil
}

func (Store) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("store_cd", "company_cd").Unique(),
	}
}
