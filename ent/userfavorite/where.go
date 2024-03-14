// Code generated by ent, DO NOT EDIT.

package userfavorite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUpdatedAt, v))
}

// PageLink applies equality check predicate on the "page_link" field. It's identical to PageLinkEQ.
func PageLink(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldPageLink, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUserID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldUpdatedAt, v))
}

// PageLinkEQ applies the EQ predicate on the "page_link" field.
func PageLinkEQ(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldPageLink, v))
}

// PageLinkNEQ applies the NEQ predicate on the "page_link" field.
func PageLinkNEQ(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldPageLink, v))
}

// PageLinkIn applies the In predicate on the "page_link" field.
func PageLinkIn(vs ...string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldPageLink, vs...))
}

// PageLinkNotIn applies the NotIn predicate on the "page_link" field.
func PageLinkNotIn(vs ...string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldPageLink, vs...))
}

// PageLinkGT applies the GT predicate on the "page_link" field.
func PageLinkGT(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGT(FieldPageLink, v))
}

// PageLinkGTE applies the GTE predicate on the "page_link" field.
func PageLinkGTE(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldGTE(FieldPageLink, v))
}

// PageLinkLT applies the LT predicate on the "page_link" field.
func PageLinkLT(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLT(FieldPageLink, v))
}

// PageLinkLTE applies the LTE predicate on the "page_link" field.
func PageLinkLTE(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldLTE(FieldPageLink, v))
}

// PageLinkContains applies the Contains predicate on the "page_link" field.
func PageLinkContains(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldContains(FieldPageLink, v))
}

// PageLinkHasPrefix applies the HasPrefix predicate on the "page_link" field.
func PageLinkHasPrefix(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldHasPrefix(FieldPageLink, v))
}

// PageLinkHasSuffix applies the HasSuffix predicate on the "page_link" field.
func PageLinkHasSuffix(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldHasSuffix(FieldPageLink, v))
}

// PageLinkEqualFold applies the EqualFold predicate on the "page_link" field.
func PageLinkEqualFold(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEqualFold(FieldPageLink, v))
}

// PageLinkContainsFold applies the ContainsFold predicate on the "page_link" field.
func PageLinkContainsFold(v string) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldContainsFold(FieldPageLink, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserFavorite {
	return predicate.UserFavorite(sql.FieldNotIn(FieldUserID, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserFavorite {
	return predicate.UserFavorite(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserFavorite) predicate.UserFavorite {
	return predicate.UserFavorite(sql.NotPredicates(p))
}
