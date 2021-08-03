package schema

import "entgo.io/ent"

// Route holds the schema definition for the Route entity.
type Route struct {
	ent.Schema
}

// Fields of the Route.
func (Route) Fields() []ent.Field {
	return nil
}

// Edges of the Route.
func (Route) Edges() []ent.Edge {
	return nil
}
