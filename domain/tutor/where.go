// Code generated by ent, DO NOT EDIT.

package tutor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tutor {
	return predicate.Tutor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tutor {
	return predicate.Tutor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tutor {
	return predicate.Tutor(sql.FieldLTE(FieldID, id))
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tutor) predicate.Tutor {
	return predicate.Tutor(func(s *sql.Selector) {
		p(s.Not())
	})
}