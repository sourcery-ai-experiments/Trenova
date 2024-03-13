// Code generated by ent, DO NOT EDIT.

package businessunit

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the businessunit type in the database.
	Label = "business_unit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEntityKey holds the string denoting the entity_key field in the database.
	FieldEntityKey = "entity_key"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldTaxID holds the string denoting the tax_id field in the database.
	FieldTaxID = "tax_id"
	// FieldSubscriptionPlan holds the string denoting the subscription_plan field in the database.
	FieldSubscriptionPlan = "subscription_plan"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLegalName holds the string denoting the legal_name field in the database.
	FieldLegalName = "legal_name"
	// FieldContactName holds the string denoting the contact_name field in the database.
	FieldContactName = "contact_name"
	// FieldContactEmail holds the string denoting the contact_email field in the database.
	FieldContactEmail = "contact_email"
	// FieldPaidUntil holds the string denoting the paid_until field in the database.
	FieldPaidUntil = "paid_until"
	// FieldSettings holds the string denoting the settings field in the database.
	FieldSettings = "settings"
	// FieldFreeTrial holds the string denoting the free_trial field in the database.
	FieldFreeTrial = "free_trial"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the businessunit in the database.
	Table = "business_units"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "business_units"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_id"
)

// Columns holds all SQL columns for businessunit fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldEntityKey,
	FieldPhoneNumber,
	FieldAddress,
	FieldCity,
	FieldState,
	FieldCountry,
	FieldPostalCode,
	FieldTaxID,
	FieldSubscriptionPlan,
	FieldDescription,
	FieldLegalName,
	FieldContactName,
	FieldContactEmail,
	FieldPaidUntil,
	FieldSettings,
	FieldFreeTrial,
	FieldParentID,
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
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EntityKeyValidator is a validator for the "entity_key" field. It is called by the builders before save.
	EntityKeyValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
	// TaxIDValidator is a validator for the "tax_id" field. It is called by the builders before save.
	TaxIDValidator func(string) error
	// SubscriptionPlanValidator is a validator for the "subscription_plan" field. It is called by the builders before save.
	SubscriptionPlanValidator func(string) error
	// LegalNameValidator is a validator for the "legal_name" field. It is called by the builders before save.
	LegalNameValidator func(string) error
	// DefaultFreeTrial holds the default value on creation for the "free_trial" field.
	DefaultFreeTrial bool
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
		return fmt.Errorf("businessunit: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the BusinessUnit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEntityKey orders the results by the entity_key field.
func ByEntityKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityKey, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByTaxID orders the results by the tax_id field.
func ByTaxID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaxID, opts...).ToFunc()
}

// BySubscriptionPlan orders the results by the subscription_plan field.
func BySubscriptionPlan(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriptionPlan, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByLegalName orders the results by the legal_name field.
func ByLegalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLegalName, opts...).ToFunc()
}

// ByContactName orders the results by the contact_name field.
func ByContactName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactName, opts...).ToFunc()
}

// ByContactEmail orders the results by the contact_email field.
func ByContactEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactEmail, opts...).ToFunc()
}

// ByPaidUntil orders the results by the paid_until field.
func ByPaidUntil(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaidUntil, opts...).ToFunc()
}

// ByFreeTrial orders the results by the free_trial field.
func ByFreeTrial(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFreeTrial, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ParentTable, ParentColumn),
	)
}
