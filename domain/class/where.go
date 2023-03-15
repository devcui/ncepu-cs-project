// Code generated by ent, DO NOT EDIT.

package class

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/devcui/ncepu-cs-project/domain/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDescription, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldType, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldType, v))
}

// HasMajor applies the HasEdge predicate on the "major" edge.
func HasMajor() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MajorTable, MajorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMajorWith applies the HasEdge predicate on the "major" edge with a given conditions (other predicates).
func HasMajorWith(preds ...predicate.Major) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MajorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MajorTable, MajorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCampus applies the HasEdge predicate on the "campus" edge.
func HasCampus() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CampusTable, CampusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCampusWith applies the HasEdge predicate on the "campus" edge with a given conditions (other predicates).
func HasCampusWith(preds ...predicate.Campus) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CampusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, CampusTable, CampusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassLeader applies the HasEdge predicate on the "class_leader" edge.
func HasClassLeader() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ClassLeaderTable, ClassLeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassLeaderWith applies the HasEdge predicate on the "class_leader" edge with a given conditions (other predicates).
func HasClassLeaderWith(preds ...predicate.ClassLeader) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassLeaderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ClassLeaderTable, ClassLeaderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTutor applies the HasEdge predicate on the "tutor" edge.
func HasTutor() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TutorTable, TutorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTutorWith applies the HasEdge predicate on the "tutor" edge with a given conditions (other predicates).
func HasTutorWith(preds ...predicate.Tutor) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TutorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TutorTable, TutorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMajorDirection applies the HasEdge predicate on the "major_direction" edge.
func HasMajorDirection() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MajorDirectionTable, MajorDirectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMajorDirectionWith applies the HasEdge predicate on the "major_direction" edge with a given conditions (other predicates).
func HasMajorDirectionWith(preds ...predicate.MajorDirection) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MajorDirectionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MajorDirectionTable, MajorDirectionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
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
func Not(p predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		p(s.Not())
	})
}
