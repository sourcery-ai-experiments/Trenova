// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/invoicecontrol"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// InvoiceControlUpdate is the builder for updating InvoiceControl entities.
type InvoiceControlUpdate struct {
	config
	hooks     []Hook
	mutation  *InvoiceControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InvoiceControlUpdate builder.
func (icu *InvoiceControlUpdate) Where(ps ...predicate.InvoiceControl) *InvoiceControlUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetUpdatedAt sets the "updated_at" field.
func (icu *InvoiceControlUpdate) SetUpdatedAt(t time.Time) *InvoiceControlUpdate {
	icu.mutation.SetUpdatedAt(t)
	return icu
}

// SetInvoiceNumberPrefix sets the "invoice_number_prefix" field.
func (icu *InvoiceControlUpdate) SetInvoiceNumberPrefix(s string) *InvoiceControlUpdate {
	icu.mutation.SetInvoiceNumberPrefix(s)
	return icu
}

// SetNillableInvoiceNumberPrefix sets the "invoice_number_prefix" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceNumberPrefix(s *string) *InvoiceControlUpdate {
	if s != nil {
		icu.SetInvoiceNumberPrefix(*s)
	}
	return icu
}

// SetCreditMemoNumberPrefix sets the "credit_memo_number_prefix" field.
func (icu *InvoiceControlUpdate) SetCreditMemoNumberPrefix(s string) *InvoiceControlUpdate {
	icu.mutation.SetCreditMemoNumberPrefix(s)
	return icu
}

// SetNillableCreditMemoNumberPrefix sets the "credit_memo_number_prefix" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableCreditMemoNumberPrefix(s *string) *InvoiceControlUpdate {
	if s != nil {
		icu.SetCreditMemoNumberPrefix(*s)
	}
	return icu
}

// SetInvoiceTerms sets the "invoice_terms" field.
func (icu *InvoiceControlUpdate) SetInvoiceTerms(s string) *InvoiceControlUpdate {
	icu.mutation.SetInvoiceTerms(s)
	return icu
}

// SetNillableInvoiceTerms sets the "invoice_terms" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceTerms(s *string) *InvoiceControlUpdate {
	if s != nil {
		icu.SetInvoiceTerms(*s)
	}
	return icu
}

// ClearInvoiceTerms clears the value of the "invoice_terms" field.
func (icu *InvoiceControlUpdate) ClearInvoiceTerms() *InvoiceControlUpdate {
	icu.mutation.ClearInvoiceTerms()
	return icu
}

// SetInvoiceFooter sets the "invoice_footer" field.
func (icu *InvoiceControlUpdate) SetInvoiceFooter(s string) *InvoiceControlUpdate {
	icu.mutation.SetInvoiceFooter(s)
	return icu
}

// SetNillableInvoiceFooter sets the "invoice_footer" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceFooter(s *string) *InvoiceControlUpdate {
	if s != nil {
		icu.SetInvoiceFooter(*s)
	}
	return icu
}

// ClearInvoiceFooter clears the value of the "invoice_footer" field.
func (icu *InvoiceControlUpdate) ClearInvoiceFooter() *InvoiceControlUpdate {
	icu.mutation.ClearInvoiceFooter()
	return icu
}

// SetInvoiceLogoURL sets the "invoice_logo_url" field.
func (icu *InvoiceControlUpdate) SetInvoiceLogoURL(s string) *InvoiceControlUpdate {
	icu.mutation.SetInvoiceLogoURL(s)
	return icu
}

// SetNillableInvoiceLogoURL sets the "invoice_logo_url" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceLogoURL(s *string) *InvoiceControlUpdate {
	if s != nil {
		icu.SetInvoiceLogoURL(*s)
	}
	return icu
}

// ClearInvoiceLogoURL clears the value of the "invoice_logo_url" field.
func (icu *InvoiceControlUpdate) ClearInvoiceLogoURL() *InvoiceControlUpdate {
	icu.mutation.ClearInvoiceLogoURL()
	return icu
}

// SetInvoiceDateFormat sets the "invoice_date_format" field.
func (icu *InvoiceControlUpdate) SetInvoiceDateFormat(idf invoicecontrol.InvoiceDateFormat) *InvoiceControlUpdate {
	icu.mutation.SetInvoiceDateFormat(idf)
	return icu
}

// SetNillableInvoiceDateFormat sets the "invoice_date_format" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceDateFormat(idf *invoicecontrol.InvoiceDateFormat) *InvoiceControlUpdate {
	if idf != nil {
		icu.SetInvoiceDateFormat(*idf)
	}
	return icu
}

// SetInvoiceDueAfterDays sets the "invoice_due_after_days" field.
func (icu *InvoiceControlUpdate) SetInvoiceDueAfterDays(u uint8) *InvoiceControlUpdate {
	icu.mutation.ResetInvoiceDueAfterDays()
	icu.mutation.SetInvoiceDueAfterDays(u)
	return icu
}

// SetNillableInvoiceDueAfterDays sets the "invoice_due_after_days" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceDueAfterDays(u *uint8) *InvoiceControlUpdate {
	if u != nil {
		icu.SetInvoiceDueAfterDays(*u)
	}
	return icu
}

// AddInvoiceDueAfterDays adds u to the "invoice_due_after_days" field.
func (icu *InvoiceControlUpdate) AddInvoiceDueAfterDays(u int8) *InvoiceControlUpdate {
	icu.mutation.AddInvoiceDueAfterDays(u)
	return icu
}

// SetInvoiceLogoWidth sets the "invoice_logo_width" field.
func (icu *InvoiceControlUpdate) SetInvoiceLogoWidth(u uint16) *InvoiceControlUpdate {
	icu.mutation.ResetInvoiceLogoWidth()
	icu.mutation.SetInvoiceLogoWidth(u)
	return icu
}

// SetNillableInvoiceLogoWidth sets the "invoice_logo_width" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableInvoiceLogoWidth(u *uint16) *InvoiceControlUpdate {
	if u != nil {
		icu.SetInvoiceLogoWidth(*u)
	}
	return icu
}

// AddInvoiceLogoWidth adds u to the "invoice_logo_width" field.
func (icu *InvoiceControlUpdate) AddInvoiceLogoWidth(u int16) *InvoiceControlUpdate {
	icu.mutation.AddInvoiceLogoWidth(u)
	return icu
}

// SetShowAmountDue sets the "show_amount_due" field.
func (icu *InvoiceControlUpdate) SetShowAmountDue(b bool) *InvoiceControlUpdate {
	icu.mutation.SetShowAmountDue(b)
	return icu
}

// SetNillableShowAmountDue sets the "show_amount_due" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableShowAmountDue(b *bool) *InvoiceControlUpdate {
	if b != nil {
		icu.SetShowAmountDue(*b)
	}
	return icu
}

// SetAttachPdf sets the "attach_pdf" field.
func (icu *InvoiceControlUpdate) SetAttachPdf(b bool) *InvoiceControlUpdate {
	icu.mutation.SetAttachPdf(b)
	return icu
}

// SetNillableAttachPdf sets the "attach_pdf" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableAttachPdf(b *bool) *InvoiceControlUpdate {
	if b != nil {
		icu.SetAttachPdf(*b)
	}
	return icu
}

// SetShowInvoiceDueDate sets the "show_invoice_due_date" field.
func (icu *InvoiceControlUpdate) SetShowInvoiceDueDate(b bool) *InvoiceControlUpdate {
	icu.mutation.SetShowInvoiceDueDate(b)
	return icu
}

// SetNillableShowInvoiceDueDate sets the "show_invoice_due_date" field if the given value is not nil.
func (icu *InvoiceControlUpdate) SetNillableShowInvoiceDueDate(b *bool) *InvoiceControlUpdate {
	if b != nil {
		icu.SetShowInvoiceDueDate(*b)
	}
	return icu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (icu *InvoiceControlUpdate) SetOrganizationID(id uuid.UUID) *InvoiceControlUpdate {
	icu.mutation.SetOrganizationID(id)
	return icu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (icu *InvoiceControlUpdate) SetOrganization(o *Organization) *InvoiceControlUpdate {
	return icu.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (icu *InvoiceControlUpdate) SetBusinessUnitID(id uuid.UUID) *InvoiceControlUpdate {
	icu.mutation.SetBusinessUnitID(id)
	return icu
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (icu *InvoiceControlUpdate) SetBusinessUnit(b *BusinessUnit) *InvoiceControlUpdate {
	return icu.SetBusinessUnitID(b.ID)
}

// Mutation returns the InvoiceControlMutation object of the builder.
func (icu *InvoiceControlUpdate) Mutation() *InvoiceControlMutation {
	return icu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (icu *InvoiceControlUpdate) ClearOrganization() *InvoiceControlUpdate {
	icu.mutation.ClearOrganization()
	return icu
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (icu *InvoiceControlUpdate) ClearBusinessUnit() *InvoiceControlUpdate {
	icu.mutation.ClearBusinessUnit()
	return icu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *InvoiceControlUpdate) Save(ctx context.Context) (int, error) {
	icu.defaults()
	return withHooks(ctx, icu.sqlSave, icu.mutation, icu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icu *InvoiceControlUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *InvoiceControlUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *InvoiceControlUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icu *InvoiceControlUpdate) defaults() {
	if _, ok := icu.mutation.UpdatedAt(); !ok {
		v := invoicecontrol.UpdateDefaultUpdatedAt()
		icu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icu *InvoiceControlUpdate) check() error {
	if v, ok := icu.mutation.InvoiceNumberPrefix(); ok {
		if err := invoicecontrol.InvoiceNumberPrefixValidator(v); err != nil {
			return &ValidationError{Name: "invoice_number_prefix", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_number_prefix": %w`, err)}
		}
	}
	if v, ok := icu.mutation.CreditMemoNumberPrefix(); ok {
		if err := invoicecontrol.CreditMemoNumberPrefixValidator(v); err != nil {
			return &ValidationError{Name: "credit_memo_number_prefix", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.credit_memo_number_prefix": %w`, err)}
		}
	}
	if v, ok := icu.mutation.InvoiceLogoURL(); ok {
		if err := invoicecontrol.InvoiceLogoURLValidator(v); err != nil {
			return &ValidationError{Name: "invoice_logo_url", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_logo_url": %w`, err)}
		}
	}
	if v, ok := icu.mutation.InvoiceDateFormat(); ok {
		if err := invoicecontrol.InvoiceDateFormatValidator(v); err != nil {
			return &ValidationError{Name: "invoice_date_format", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_date_format": %w`, err)}
		}
	}
	if v, ok := icu.mutation.InvoiceDueAfterDays(); ok {
		if err := invoicecontrol.InvoiceDueAfterDaysValidator(v); err != nil {
			return &ValidationError{Name: "invoice_due_after_days", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_due_after_days": %w`, err)}
		}
	}
	if v, ok := icu.mutation.InvoiceLogoWidth(); ok {
		if err := invoicecontrol.InvoiceLogoWidthValidator(v); err != nil {
			return &ValidationError{Name: "invoice_logo_width", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_logo_width": %w`, err)}
		}
	}
	if _, ok := icu.mutation.OrganizationID(); icu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceControl.organization"`)
	}
	if _, ok := icu.mutation.BusinessUnitID(); icu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceControl.business_unit"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (icu *InvoiceControlUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InvoiceControlUpdate {
	icu.modifiers = append(icu.modifiers, modifiers...)
	return icu
}

func (icu *InvoiceControlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := icu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicecontrol.Table, invoicecontrol.Columns, sqlgraph.NewFieldSpec(invoicecontrol.FieldID, field.TypeUUID))
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicecontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := icu.mutation.InvoiceNumberPrefix(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceNumberPrefix, field.TypeString, value)
	}
	if value, ok := icu.mutation.CreditMemoNumberPrefix(); ok {
		_spec.SetField(invoicecontrol.FieldCreditMemoNumberPrefix, field.TypeString, value)
	}
	if value, ok := icu.mutation.InvoiceTerms(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceTerms, field.TypeString, value)
	}
	if icu.mutation.InvoiceTermsCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceTerms, field.TypeString)
	}
	if value, ok := icu.mutation.InvoiceFooter(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceFooter, field.TypeString, value)
	}
	if icu.mutation.InvoiceFooterCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceFooter, field.TypeString)
	}
	if value, ok := icu.mutation.InvoiceLogoURL(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceLogoURL, field.TypeString, value)
	}
	if icu.mutation.InvoiceLogoURLCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceLogoURL, field.TypeString)
	}
	if value, ok := icu.mutation.InvoiceDateFormat(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceDateFormat, field.TypeEnum, value)
	}
	if value, ok := icu.mutation.InvoiceDueAfterDays(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceDueAfterDays, field.TypeUint8, value)
	}
	if value, ok := icu.mutation.AddedInvoiceDueAfterDays(); ok {
		_spec.AddField(invoicecontrol.FieldInvoiceDueAfterDays, field.TypeUint8, value)
	}
	if value, ok := icu.mutation.InvoiceLogoWidth(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceLogoWidth, field.TypeUint16, value)
	}
	if value, ok := icu.mutation.AddedInvoiceLogoWidth(); ok {
		_spec.AddField(invoicecontrol.FieldInvoiceLogoWidth, field.TypeUint16, value)
	}
	if value, ok := icu.mutation.ShowAmountDue(); ok {
		_spec.SetField(invoicecontrol.FieldShowAmountDue, field.TypeBool, value)
	}
	if value, ok := icu.mutation.AttachPdf(); ok {
		_spec.SetField(invoicecontrol.FieldAttachPdf, field.TypeBool, value)
	}
	if value, ok := icu.mutation.ShowInvoiceDueDate(); ok {
		_spec.SetField(invoicecontrol.FieldShowInvoiceDueDate, field.TypeBool, value)
	}
	if icu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   invoicecontrol.OrganizationTable,
			Columns: []string{invoicecontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   invoicecontrol.OrganizationTable,
			Columns: []string{invoicecontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if icu.mutation.BusinessUnitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicecontrol.BusinessUnitTable,
			Columns: []string{invoicecontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicecontrol.BusinessUnitTable,
			Columns: []string{invoicecontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(icu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicecontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	icu.mutation.done = true
	return n, nil
}

// InvoiceControlUpdateOne is the builder for updating a single InvoiceControl entity.
type InvoiceControlUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InvoiceControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (icuo *InvoiceControlUpdateOne) SetUpdatedAt(t time.Time) *InvoiceControlUpdateOne {
	icuo.mutation.SetUpdatedAt(t)
	return icuo
}

// SetInvoiceNumberPrefix sets the "invoice_number_prefix" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceNumberPrefix(s string) *InvoiceControlUpdateOne {
	icuo.mutation.SetInvoiceNumberPrefix(s)
	return icuo
}

// SetNillableInvoiceNumberPrefix sets the "invoice_number_prefix" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceNumberPrefix(s *string) *InvoiceControlUpdateOne {
	if s != nil {
		icuo.SetInvoiceNumberPrefix(*s)
	}
	return icuo
}

// SetCreditMemoNumberPrefix sets the "credit_memo_number_prefix" field.
func (icuo *InvoiceControlUpdateOne) SetCreditMemoNumberPrefix(s string) *InvoiceControlUpdateOne {
	icuo.mutation.SetCreditMemoNumberPrefix(s)
	return icuo
}

// SetNillableCreditMemoNumberPrefix sets the "credit_memo_number_prefix" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableCreditMemoNumberPrefix(s *string) *InvoiceControlUpdateOne {
	if s != nil {
		icuo.SetCreditMemoNumberPrefix(*s)
	}
	return icuo
}

// SetInvoiceTerms sets the "invoice_terms" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceTerms(s string) *InvoiceControlUpdateOne {
	icuo.mutation.SetInvoiceTerms(s)
	return icuo
}

// SetNillableInvoiceTerms sets the "invoice_terms" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceTerms(s *string) *InvoiceControlUpdateOne {
	if s != nil {
		icuo.SetInvoiceTerms(*s)
	}
	return icuo
}

// ClearInvoiceTerms clears the value of the "invoice_terms" field.
func (icuo *InvoiceControlUpdateOne) ClearInvoiceTerms() *InvoiceControlUpdateOne {
	icuo.mutation.ClearInvoiceTerms()
	return icuo
}

// SetInvoiceFooter sets the "invoice_footer" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceFooter(s string) *InvoiceControlUpdateOne {
	icuo.mutation.SetInvoiceFooter(s)
	return icuo
}

// SetNillableInvoiceFooter sets the "invoice_footer" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceFooter(s *string) *InvoiceControlUpdateOne {
	if s != nil {
		icuo.SetInvoiceFooter(*s)
	}
	return icuo
}

// ClearInvoiceFooter clears the value of the "invoice_footer" field.
func (icuo *InvoiceControlUpdateOne) ClearInvoiceFooter() *InvoiceControlUpdateOne {
	icuo.mutation.ClearInvoiceFooter()
	return icuo
}

// SetInvoiceLogoURL sets the "invoice_logo_url" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceLogoURL(s string) *InvoiceControlUpdateOne {
	icuo.mutation.SetInvoiceLogoURL(s)
	return icuo
}

// SetNillableInvoiceLogoURL sets the "invoice_logo_url" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceLogoURL(s *string) *InvoiceControlUpdateOne {
	if s != nil {
		icuo.SetInvoiceLogoURL(*s)
	}
	return icuo
}

// ClearInvoiceLogoURL clears the value of the "invoice_logo_url" field.
func (icuo *InvoiceControlUpdateOne) ClearInvoiceLogoURL() *InvoiceControlUpdateOne {
	icuo.mutation.ClearInvoiceLogoURL()
	return icuo
}

// SetInvoiceDateFormat sets the "invoice_date_format" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceDateFormat(idf invoicecontrol.InvoiceDateFormat) *InvoiceControlUpdateOne {
	icuo.mutation.SetInvoiceDateFormat(idf)
	return icuo
}

// SetNillableInvoiceDateFormat sets the "invoice_date_format" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceDateFormat(idf *invoicecontrol.InvoiceDateFormat) *InvoiceControlUpdateOne {
	if idf != nil {
		icuo.SetInvoiceDateFormat(*idf)
	}
	return icuo
}

// SetInvoiceDueAfterDays sets the "invoice_due_after_days" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceDueAfterDays(u uint8) *InvoiceControlUpdateOne {
	icuo.mutation.ResetInvoiceDueAfterDays()
	icuo.mutation.SetInvoiceDueAfterDays(u)
	return icuo
}

// SetNillableInvoiceDueAfterDays sets the "invoice_due_after_days" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceDueAfterDays(u *uint8) *InvoiceControlUpdateOne {
	if u != nil {
		icuo.SetInvoiceDueAfterDays(*u)
	}
	return icuo
}

// AddInvoiceDueAfterDays adds u to the "invoice_due_after_days" field.
func (icuo *InvoiceControlUpdateOne) AddInvoiceDueAfterDays(u int8) *InvoiceControlUpdateOne {
	icuo.mutation.AddInvoiceDueAfterDays(u)
	return icuo
}

// SetInvoiceLogoWidth sets the "invoice_logo_width" field.
func (icuo *InvoiceControlUpdateOne) SetInvoiceLogoWidth(u uint16) *InvoiceControlUpdateOne {
	icuo.mutation.ResetInvoiceLogoWidth()
	icuo.mutation.SetInvoiceLogoWidth(u)
	return icuo
}

// SetNillableInvoiceLogoWidth sets the "invoice_logo_width" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableInvoiceLogoWidth(u *uint16) *InvoiceControlUpdateOne {
	if u != nil {
		icuo.SetInvoiceLogoWidth(*u)
	}
	return icuo
}

// AddInvoiceLogoWidth adds u to the "invoice_logo_width" field.
func (icuo *InvoiceControlUpdateOne) AddInvoiceLogoWidth(u int16) *InvoiceControlUpdateOne {
	icuo.mutation.AddInvoiceLogoWidth(u)
	return icuo
}

// SetShowAmountDue sets the "show_amount_due" field.
func (icuo *InvoiceControlUpdateOne) SetShowAmountDue(b bool) *InvoiceControlUpdateOne {
	icuo.mutation.SetShowAmountDue(b)
	return icuo
}

// SetNillableShowAmountDue sets the "show_amount_due" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableShowAmountDue(b *bool) *InvoiceControlUpdateOne {
	if b != nil {
		icuo.SetShowAmountDue(*b)
	}
	return icuo
}

// SetAttachPdf sets the "attach_pdf" field.
func (icuo *InvoiceControlUpdateOne) SetAttachPdf(b bool) *InvoiceControlUpdateOne {
	icuo.mutation.SetAttachPdf(b)
	return icuo
}

// SetNillableAttachPdf sets the "attach_pdf" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableAttachPdf(b *bool) *InvoiceControlUpdateOne {
	if b != nil {
		icuo.SetAttachPdf(*b)
	}
	return icuo
}

// SetShowInvoiceDueDate sets the "show_invoice_due_date" field.
func (icuo *InvoiceControlUpdateOne) SetShowInvoiceDueDate(b bool) *InvoiceControlUpdateOne {
	icuo.mutation.SetShowInvoiceDueDate(b)
	return icuo
}

// SetNillableShowInvoiceDueDate sets the "show_invoice_due_date" field if the given value is not nil.
func (icuo *InvoiceControlUpdateOne) SetNillableShowInvoiceDueDate(b *bool) *InvoiceControlUpdateOne {
	if b != nil {
		icuo.SetShowInvoiceDueDate(*b)
	}
	return icuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (icuo *InvoiceControlUpdateOne) SetOrganizationID(id uuid.UUID) *InvoiceControlUpdateOne {
	icuo.mutation.SetOrganizationID(id)
	return icuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (icuo *InvoiceControlUpdateOne) SetOrganization(o *Organization) *InvoiceControlUpdateOne {
	return icuo.SetOrganizationID(o.ID)
}

// SetBusinessUnitID sets the "business_unit" edge to the BusinessUnit entity by ID.
func (icuo *InvoiceControlUpdateOne) SetBusinessUnitID(id uuid.UUID) *InvoiceControlUpdateOne {
	icuo.mutation.SetBusinessUnitID(id)
	return icuo
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (icuo *InvoiceControlUpdateOne) SetBusinessUnit(b *BusinessUnit) *InvoiceControlUpdateOne {
	return icuo.SetBusinessUnitID(b.ID)
}

// Mutation returns the InvoiceControlMutation object of the builder.
func (icuo *InvoiceControlUpdateOne) Mutation() *InvoiceControlMutation {
	return icuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (icuo *InvoiceControlUpdateOne) ClearOrganization() *InvoiceControlUpdateOne {
	icuo.mutation.ClearOrganization()
	return icuo
}

// ClearBusinessUnit clears the "business_unit" edge to the BusinessUnit entity.
func (icuo *InvoiceControlUpdateOne) ClearBusinessUnit() *InvoiceControlUpdateOne {
	icuo.mutation.ClearBusinessUnit()
	return icuo
}

// Where appends a list predicates to the InvoiceControlUpdate builder.
func (icuo *InvoiceControlUpdateOne) Where(ps ...predicate.InvoiceControl) *InvoiceControlUpdateOne {
	icuo.mutation.Where(ps...)
	return icuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *InvoiceControlUpdateOne) Select(field string, fields ...string) *InvoiceControlUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated InvoiceControl entity.
func (icuo *InvoiceControlUpdateOne) Save(ctx context.Context) (*InvoiceControl, error) {
	icuo.defaults()
	return withHooks(ctx, icuo.sqlSave, icuo.mutation, icuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *InvoiceControlUpdateOne) SaveX(ctx context.Context) *InvoiceControl {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *InvoiceControlUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *InvoiceControlUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (icuo *InvoiceControlUpdateOne) defaults() {
	if _, ok := icuo.mutation.UpdatedAt(); !ok {
		v := invoicecontrol.UpdateDefaultUpdatedAt()
		icuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (icuo *InvoiceControlUpdateOne) check() error {
	if v, ok := icuo.mutation.InvoiceNumberPrefix(); ok {
		if err := invoicecontrol.InvoiceNumberPrefixValidator(v); err != nil {
			return &ValidationError{Name: "invoice_number_prefix", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_number_prefix": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.CreditMemoNumberPrefix(); ok {
		if err := invoicecontrol.CreditMemoNumberPrefixValidator(v); err != nil {
			return &ValidationError{Name: "credit_memo_number_prefix", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.credit_memo_number_prefix": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.InvoiceLogoURL(); ok {
		if err := invoicecontrol.InvoiceLogoURLValidator(v); err != nil {
			return &ValidationError{Name: "invoice_logo_url", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_logo_url": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.InvoiceDateFormat(); ok {
		if err := invoicecontrol.InvoiceDateFormatValidator(v); err != nil {
			return &ValidationError{Name: "invoice_date_format", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_date_format": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.InvoiceDueAfterDays(); ok {
		if err := invoicecontrol.InvoiceDueAfterDaysValidator(v); err != nil {
			return &ValidationError{Name: "invoice_due_after_days", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_due_after_days": %w`, err)}
		}
	}
	if v, ok := icuo.mutation.InvoiceLogoWidth(); ok {
		if err := invoicecontrol.InvoiceLogoWidthValidator(v); err != nil {
			return &ValidationError{Name: "invoice_logo_width", err: fmt.Errorf(`ent: validator failed for field "InvoiceControl.invoice_logo_width": %w`, err)}
		}
	}
	if _, ok := icuo.mutation.OrganizationID(); icuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceControl.organization"`)
	}
	if _, ok := icuo.mutation.BusinessUnitID(); icuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "InvoiceControl.business_unit"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (icuo *InvoiceControlUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InvoiceControlUpdateOne {
	icuo.modifiers = append(icuo.modifiers, modifiers...)
	return icuo
}

func (icuo *InvoiceControlUpdateOne) sqlSave(ctx context.Context) (_node *InvoiceControl, err error) {
	if err := icuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invoicecontrol.Table, invoicecontrol.Columns, sqlgraph.NewFieldSpec(invoicecontrol.FieldID, field.TypeUUID))
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InvoiceControl.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoicecontrol.FieldID)
		for _, f := range fields {
			if !invoicecontrol.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invoicecontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.UpdatedAt(); ok {
		_spec.SetField(invoicecontrol.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := icuo.mutation.InvoiceNumberPrefix(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceNumberPrefix, field.TypeString, value)
	}
	if value, ok := icuo.mutation.CreditMemoNumberPrefix(); ok {
		_spec.SetField(invoicecontrol.FieldCreditMemoNumberPrefix, field.TypeString, value)
	}
	if value, ok := icuo.mutation.InvoiceTerms(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceTerms, field.TypeString, value)
	}
	if icuo.mutation.InvoiceTermsCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceTerms, field.TypeString)
	}
	if value, ok := icuo.mutation.InvoiceFooter(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceFooter, field.TypeString, value)
	}
	if icuo.mutation.InvoiceFooterCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceFooter, field.TypeString)
	}
	if value, ok := icuo.mutation.InvoiceLogoURL(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceLogoURL, field.TypeString, value)
	}
	if icuo.mutation.InvoiceLogoURLCleared() {
		_spec.ClearField(invoicecontrol.FieldInvoiceLogoURL, field.TypeString)
	}
	if value, ok := icuo.mutation.InvoiceDateFormat(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceDateFormat, field.TypeEnum, value)
	}
	if value, ok := icuo.mutation.InvoiceDueAfterDays(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceDueAfterDays, field.TypeUint8, value)
	}
	if value, ok := icuo.mutation.AddedInvoiceDueAfterDays(); ok {
		_spec.AddField(invoicecontrol.FieldInvoiceDueAfterDays, field.TypeUint8, value)
	}
	if value, ok := icuo.mutation.InvoiceLogoWidth(); ok {
		_spec.SetField(invoicecontrol.FieldInvoiceLogoWidth, field.TypeUint16, value)
	}
	if value, ok := icuo.mutation.AddedInvoiceLogoWidth(); ok {
		_spec.AddField(invoicecontrol.FieldInvoiceLogoWidth, field.TypeUint16, value)
	}
	if value, ok := icuo.mutation.ShowAmountDue(); ok {
		_spec.SetField(invoicecontrol.FieldShowAmountDue, field.TypeBool, value)
	}
	if value, ok := icuo.mutation.AttachPdf(); ok {
		_spec.SetField(invoicecontrol.FieldAttachPdf, field.TypeBool, value)
	}
	if value, ok := icuo.mutation.ShowInvoiceDueDate(); ok {
		_spec.SetField(invoicecontrol.FieldShowInvoiceDueDate, field.TypeBool, value)
	}
	if icuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   invoicecontrol.OrganizationTable,
			Columns: []string{invoicecontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   invoicecontrol.OrganizationTable,
			Columns: []string{invoicecontrol.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if icuo.mutation.BusinessUnitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicecontrol.BusinessUnitTable,
			Columns: []string{invoicecontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   invoicecontrol.BusinessUnitTable,
			Columns: []string{invoicecontrol.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(icuo.modifiers...)
	_node = &InvoiceControl{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invoicecontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	icuo.mutation.done = true
	return _node, nil
}
