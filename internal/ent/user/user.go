// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBusinessUnitID holds the string denoting the business_unit_id field in the database.
	FieldBusinessUnitID = "business_unit_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// FieldProfilePicURL holds the string denoting the profile_pic_url field in the database.
	FieldProfilePicURL = "profile_pic_url"
	// FieldThumbnailURL holds the string denoting the thumbnail_url field in the database.
	FieldThumbnailURL = "thumbnail_url"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldIsAdmin holds the string denoting the is_admin field in the database.
	FieldIsAdmin = "is_admin"
	// FieldIsSuperAdmin holds the string denoting the is_super_admin field in the database.
	FieldIsSuperAdmin = "is_super_admin"
	// FieldLastLogin holds the string denoting the last_login field in the database.
	FieldLastLogin = "last_login"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeUserFavorites holds the string denoting the user_favorites edge name in mutations.
	EdgeUserFavorites = "user_favorites"
	// EdgeShipments holds the string denoting the shipments edge name in mutations.
	EdgeShipments = "shipments"
	// EdgeShipmentComments holds the string denoting the shipment_comments edge name in mutations.
	EdgeShipmentComments = "shipment_comments"
	// EdgeShipmentCharges holds the string denoting the shipment_charges edge name in mutations.
	EdgeShipmentCharges = "shipment_charges"
	// EdgeReports holds the string denoting the reports edge name in mutations.
	EdgeReports = "reports"
	// Table holds the table name of the user in the database.
	Table = "users"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "users"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "users"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// UserFavoritesTable is the table that holds the user_favorites relation/edge.
	UserFavoritesTable = "user_favorites"
	// UserFavoritesInverseTable is the table name for the UserFavorite entity.
	// It exists in this package in order to avoid circular dependency with the "userfavorite" package.
	UserFavoritesInverseTable = "user_favorites"
	// UserFavoritesColumn is the table column denoting the user_favorites relation/edge.
	UserFavoritesColumn = "user_id"
	// ShipmentsTable is the table that holds the shipments relation/edge.
	ShipmentsTable = "shipments"
	// ShipmentsInverseTable is the table name for the Shipment entity.
	// It exists in this package in order to avoid circular dependency with the "shipment" package.
	ShipmentsInverseTable = "shipments"
	// ShipmentsColumn is the table column denoting the shipments relation/edge.
	ShipmentsColumn = "created_by"
	// ShipmentCommentsTable is the table that holds the shipment_comments relation/edge.
	ShipmentCommentsTable = "shipment_comments"
	// ShipmentCommentsInverseTable is the table name for the ShipmentComment entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentcomment" package.
	ShipmentCommentsInverseTable = "shipment_comments"
	// ShipmentCommentsColumn is the table column denoting the shipment_comments relation/edge.
	ShipmentCommentsColumn = "created_by"
	// ShipmentChargesTable is the table that holds the shipment_charges relation/edge.
	ShipmentChargesTable = "shipment_charges"
	// ShipmentChargesInverseTable is the table name for the ShipmentCharges entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentcharges" package.
	ShipmentChargesInverseTable = "shipment_charges"
	// ShipmentChargesColumn is the table column denoting the shipment_charges relation/edge.
	ShipmentChargesColumn = "created_by"
	// ReportsTable is the table that holds the reports relation/edge.
	ReportsTable = "user_reports"
	// ReportsInverseTable is the table name for the UserReport entity.
	// It exists in this package in order to avoid circular dependency with the "userreport" package.
	ReportsInverseTable = "user_reports"
	// ReportsColumn is the table column denoting the reports relation/edge.
	ReportsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldName,
	FieldUsername,
	FieldPassword,
	FieldEmail,
	FieldTimezone,
	FieldProfilePicURL,
	FieldThumbnailURL,
	FieldPhoneNumber,
	FieldIsAdmin,
	FieldIsSuperAdmin,
	FieldLastLogin,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultIsAdmin holds the default value on creation for the "is_admin" field.
	DefaultIsAdmin bool
	// DefaultIsSuperAdmin holds the default value on creation for the "is_super_admin" field.
	DefaultIsSuperAdmin bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusA is the default value of the Status enum.
const DefaultStatus = StatusA

// Status values.
const (
	StatusA Status = "A"
	StatusI Status = "I"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusA, StatusI:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for status field: %q", s)
	}
}

// Timezone defines the type for the "timezone" enum field.
type Timezone string

// TimezoneAmericaLosAngeles is the default value of the Timezone enum.
const DefaultTimezone = TimezoneAmericaLosAngeles

// Timezone values.
const (
	TimezoneAmericaLosAngeles Timezone = "AmericaLosAngeles"
	TimezoneAmericaDenver     Timezone = "AmericaDenver"
	TimezoneAmericaChicago    Timezone = "AmericaChicago"
	TimezoneAmericaNewYork    Timezone = "AmericaNewYork"
)

func (t Timezone) String() string {
	return string(t)
}

// TimezoneValidator is a validator for the "timezone" field enum values. It is called by the builders before save.
func TimezoneValidator(t Timezone) error {
	switch t {
	case TimezoneAmericaLosAngeles, TimezoneAmericaDenver, TimezoneAmericaChicago, TimezoneAmericaNewYork:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for timezone field: %q", t)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBusinessUnitID orders the results by the business_unit_id field.
func ByBusinessUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessUnitID, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByTimezone orders the results by the timezone field.
func ByTimezone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimezone, opts...).ToFunc()
}

// ByProfilePicURL orders the results by the profile_pic_url field.
func ByProfilePicURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfilePicURL, opts...).ToFunc()
}

// ByThumbnailURL orders the results by the thumbnail_url field.
func ByThumbnailURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailURL, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByIsAdmin orders the results by the is_admin field.
func ByIsAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAdmin, opts...).ToFunc()
}

// ByIsSuperAdmin orders the results by the is_super_admin field.
func ByIsSuperAdmin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSuperAdmin, opts...).ToFunc()
}

// ByLastLogin orders the results by the last_login field.
func ByLastLogin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLogin, opts...).ToFunc()
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserFavoritesCount orders the results by user_favorites count.
func ByUserFavoritesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserFavoritesStep(), opts...)
	}
}

// ByUserFavorites orders the results by user_favorites terms.
func ByUserFavorites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserFavoritesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByShipmentsCount orders the results by shipments count.
func ByShipmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShipmentsStep(), opts...)
	}
}

// ByShipments orders the results by shipments terms.
func ByShipments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByShipmentCommentsCount orders the results by shipment_comments count.
func ByShipmentCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShipmentCommentsStep(), opts...)
	}
}

// ByShipmentComments orders the results by shipment_comments terms.
func ByShipmentComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByShipmentChargesCount orders the results by shipment_charges count.
func ByShipmentChargesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShipmentChargesStep(), opts...)
	}
}

// ByShipmentCharges orders the results by shipment_charges terms.
func ByShipmentCharges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentChargesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReportsCount orders the results by reports count.
func ByReportsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReportsStep(), opts...)
	}
}

// ByReports orders the results by reports terms.
func ByReports(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReportsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
	)
}
func newUserFavoritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserFavoritesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserFavoritesTable, UserFavoritesColumn),
	)
}
func newShipmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ShipmentsTable, ShipmentsColumn),
	)
}
func newShipmentCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ShipmentCommentsTable, ShipmentCommentsColumn),
	)
}
func newShipmentChargesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentChargesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ShipmentChargesTable, ShipmentChargesColumn),
	)
}
func newReportsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReportsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReportsTable, ReportsColumn),
	)
}
