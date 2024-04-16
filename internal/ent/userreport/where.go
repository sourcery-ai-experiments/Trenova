// Code generated by entc, DO NOT EDIT.

package userreport

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldVersion, v))
}

// ReportURL applies equality check predicate on the "report_url" field. It's identical to ReportURLEQ.
func ReportURL(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldReportURL, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldUserID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserReport {
	return predicate.UserReport(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.UserReport {
	return predicate.UserReport(sql.FieldLTE(FieldVersion, v))
}

// ReportURLEQ applies the EQ predicate on the "report_url" field.
func ReportURLEQ(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldReportURL, v))
}

// ReportURLNEQ applies the NEQ predicate on the "report_url" field.
func ReportURLNEQ(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldReportURL, v))
}

// ReportURLIn applies the In predicate on the "report_url" field.
func ReportURLIn(vs ...string) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldReportURL, vs...))
}

// ReportURLNotIn applies the NotIn predicate on the "report_url" field.
func ReportURLNotIn(vs ...string) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldReportURL, vs...))
}

// ReportURLGT applies the GT predicate on the "report_url" field.
func ReportURLGT(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldGT(FieldReportURL, v))
}

// ReportURLGTE applies the GTE predicate on the "report_url" field.
func ReportURLGTE(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldGTE(FieldReportURL, v))
}

// ReportURLLT applies the LT predicate on the "report_url" field.
func ReportURLLT(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldLT(FieldReportURL, v))
}

// ReportURLLTE applies the LTE predicate on the "report_url" field.
func ReportURLLTE(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldLTE(FieldReportURL, v))
}

// ReportURLContains applies the Contains predicate on the "report_url" field.
func ReportURLContains(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldContains(FieldReportURL, v))
}

// ReportURLHasPrefix applies the HasPrefix predicate on the "report_url" field.
func ReportURLHasPrefix(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldHasPrefix(FieldReportURL, v))
}

// ReportURLHasSuffix applies the HasSuffix predicate on the "report_url" field.
func ReportURLHasSuffix(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldHasSuffix(FieldReportURL, v))
}

// ReportURLEqualFold applies the EqualFold predicate on the "report_url" field.
func ReportURLEqualFold(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldEqualFold(FieldReportURL, v))
}

// ReportURLContainsFold applies the ContainsFold predicate on the "report_url" field.
func ReportURLContainsFold(v string) predicate.UserReport {
	return predicate.UserReport(sql.FieldContainsFold(FieldReportURL, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserReport {
	return predicate.UserReport(sql.FieldNotIn(FieldUserID, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserReport {
	return predicate.UserReport(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserReport) predicate.UserReport {
	return predicate.UserReport(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserReport) predicate.UserReport {
	return predicate.UserReport(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserReport) predicate.UserReport {
	return predicate.UserReport(sql.NotPredicates(p))
}
