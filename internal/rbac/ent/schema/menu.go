package schema

import "entgo.io/ent"

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return nil
}

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return nil
}
