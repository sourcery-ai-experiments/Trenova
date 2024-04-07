// Code generated by ent, DO NOT EDIT.

package organizationfeatureflag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldOrganizationID, v))
}

// FeatureFlagID applies equality check predicate on the "feature_flag_id" field. It's identical to FeatureFlagIDEQ.
func FeatureFlagID(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldFeatureFlagID, v))
}

// IsEnabled applies equality check predicate on the "is_enabled" field. It's identical to IsEnabledEQ.
func IsEnabled(v bool) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldIsEnabled, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// FeatureFlagIDEQ applies the EQ predicate on the "feature_flag_id" field.
func FeatureFlagIDEQ(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldFeatureFlagID, v))
}

// FeatureFlagIDNEQ applies the NEQ predicate on the "feature_flag_id" field.
func FeatureFlagIDNEQ(v uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldFeatureFlagID, v))
}

// FeatureFlagIDIn applies the In predicate on the "feature_flag_id" field.
func FeatureFlagIDIn(vs ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldIn(FieldFeatureFlagID, vs...))
}

// FeatureFlagIDNotIn applies the NotIn predicate on the "feature_flag_id" field.
func FeatureFlagIDNotIn(vs ...uuid.UUID) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNotIn(FieldFeatureFlagID, vs...))
}

// IsEnabledEQ applies the EQ predicate on the "is_enabled" field.
func IsEnabledEQ(v bool) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldEQ(FieldIsEnabled, v))
}

// IsEnabledNEQ applies the NEQ predicate on the "is_enabled" field.
func IsEnabledNEQ(v bool) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.FieldNEQ(FieldIsEnabled, v))
}

// HasFeatureFlag applies the HasEdge predicate on the "feature_flag" edge.
func HasFeatureFlag() predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FeatureFlagTable, FeatureFlagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeatureFlagWith applies the HasEdge predicate on the "feature_flag" edge with a given conditions (other predicates).
func HasFeatureFlagWith(preds ...predicate.FeatureFlag) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(func(s *sql.Selector) {
		step := newFeatureFlagStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrganizationFeatureFlag) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrganizationFeatureFlag) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrganizationFeatureFlag) predicate.OrganizationFeatureFlag {
	return predicate.OrganizationFeatureFlag(sql.NotPredicates(p))
}
