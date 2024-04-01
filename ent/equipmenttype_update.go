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
	"github.com/emoss08/trenova/ent/equipmenttype"
	"github.com/emoss08/trenova/ent/predicate"
)

// EquipmentTypeUpdate is the builder for updating EquipmentType entities.
type EquipmentTypeUpdate struct {
	config
	hooks     []Hook
	mutation  *EquipmentTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EquipmentTypeUpdate builder.
func (etu *EquipmentTypeUpdate) Where(ps ...predicate.EquipmentType) *EquipmentTypeUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetUpdatedAt sets the "updated_at" field.
func (etu *EquipmentTypeUpdate) SetUpdatedAt(t time.Time) *EquipmentTypeUpdate {
	etu.mutation.SetUpdatedAt(t)
	return etu
}

// SetVersion sets the "version" field.
func (etu *EquipmentTypeUpdate) SetVersion(i int) *EquipmentTypeUpdate {
	etu.mutation.ResetVersion()
	etu.mutation.SetVersion(i)
	return etu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableVersion(i *int) *EquipmentTypeUpdate {
	if i != nil {
		etu.SetVersion(*i)
	}
	return etu
}

// AddVersion adds i to the "version" field.
func (etu *EquipmentTypeUpdate) AddVersion(i int) *EquipmentTypeUpdate {
	etu.mutation.AddVersion(i)
	return etu
}

// SetStatus sets the "status" field.
func (etu *EquipmentTypeUpdate) SetStatus(e equipmenttype.Status) *EquipmentTypeUpdate {
	etu.mutation.SetStatus(e)
	return etu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableStatus(e *equipmenttype.Status) *EquipmentTypeUpdate {
	if e != nil {
		etu.SetStatus(*e)
	}
	return etu
}

// SetCode sets the "code" field.
func (etu *EquipmentTypeUpdate) SetCode(s string) *EquipmentTypeUpdate {
	etu.mutation.SetCode(s)
	return etu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableCode(s *string) *EquipmentTypeUpdate {
	if s != nil {
		etu.SetCode(*s)
	}
	return etu
}

// SetDescription sets the "description" field.
func (etu *EquipmentTypeUpdate) SetDescription(s string) *EquipmentTypeUpdate {
	etu.mutation.SetDescription(s)
	return etu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableDescription(s *string) *EquipmentTypeUpdate {
	if s != nil {
		etu.SetDescription(*s)
	}
	return etu
}

// ClearDescription clears the value of the "description" field.
func (etu *EquipmentTypeUpdate) ClearDescription() *EquipmentTypeUpdate {
	etu.mutation.ClearDescription()
	return etu
}

// SetCostPerMile sets the "cost_per_mile" field.
func (etu *EquipmentTypeUpdate) SetCostPerMile(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetCostPerMile()
	etu.mutation.SetCostPerMile(f)
	return etu
}

// SetNillableCostPerMile sets the "cost_per_mile" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableCostPerMile(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetCostPerMile(*f)
	}
	return etu
}

// AddCostPerMile adds f to the "cost_per_mile" field.
func (etu *EquipmentTypeUpdate) AddCostPerMile(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddCostPerMile(f)
	return etu
}

// ClearCostPerMile clears the value of the "cost_per_mile" field.
func (etu *EquipmentTypeUpdate) ClearCostPerMile() *EquipmentTypeUpdate {
	etu.mutation.ClearCostPerMile()
	return etu
}

// SetEquipmentClass sets the "equipment_class" field.
func (etu *EquipmentTypeUpdate) SetEquipmentClass(ec equipmenttype.EquipmentClass) *EquipmentTypeUpdate {
	etu.mutation.SetEquipmentClass(ec)
	return etu
}

// SetNillableEquipmentClass sets the "equipment_class" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableEquipmentClass(ec *equipmenttype.EquipmentClass) *EquipmentTypeUpdate {
	if ec != nil {
		etu.SetEquipmentClass(*ec)
	}
	return etu
}

// SetFixedCost sets the "fixed_cost" field.
func (etu *EquipmentTypeUpdate) SetFixedCost(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetFixedCost()
	etu.mutation.SetFixedCost(f)
	return etu
}

// SetNillableFixedCost sets the "fixed_cost" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableFixedCost(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetFixedCost(*f)
	}
	return etu
}

// AddFixedCost adds f to the "fixed_cost" field.
func (etu *EquipmentTypeUpdate) AddFixedCost(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddFixedCost(f)
	return etu
}

// ClearFixedCost clears the value of the "fixed_cost" field.
func (etu *EquipmentTypeUpdate) ClearFixedCost() *EquipmentTypeUpdate {
	etu.mutation.ClearFixedCost()
	return etu
}

// SetVariableCost sets the "variable_cost" field.
func (etu *EquipmentTypeUpdate) SetVariableCost(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetVariableCost()
	etu.mutation.SetVariableCost(f)
	return etu
}

// SetNillableVariableCost sets the "variable_cost" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableVariableCost(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetVariableCost(*f)
	}
	return etu
}

// AddVariableCost adds f to the "variable_cost" field.
func (etu *EquipmentTypeUpdate) AddVariableCost(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddVariableCost(f)
	return etu
}

// ClearVariableCost clears the value of the "variable_cost" field.
func (etu *EquipmentTypeUpdate) ClearVariableCost() *EquipmentTypeUpdate {
	etu.mutation.ClearVariableCost()
	return etu
}

// SetHeight sets the "height" field.
func (etu *EquipmentTypeUpdate) SetHeight(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetHeight()
	etu.mutation.SetHeight(f)
	return etu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableHeight(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetHeight(*f)
	}
	return etu
}

// AddHeight adds f to the "height" field.
func (etu *EquipmentTypeUpdate) AddHeight(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddHeight(f)
	return etu
}

// ClearHeight clears the value of the "height" field.
func (etu *EquipmentTypeUpdate) ClearHeight() *EquipmentTypeUpdate {
	etu.mutation.ClearHeight()
	return etu
}

// SetLength sets the "length" field.
func (etu *EquipmentTypeUpdate) SetLength(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetLength()
	etu.mutation.SetLength(f)
	return etu
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableLength(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetLength(*f)
	}
	return etu
}

// AddLength adds f to the "length" field.
func (etu *EquipmentTypeUpdate) AddLength(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddLength(f)
	return etu
}

// ClearLength clears the value of the "length" field.
func (etu *EquipmentTypeUpdate) ClearLength() *EquipmentTypeUpdate {
	etu.mutation.ClearLength()
	return etu
}

// SetWidth sets the "width" field.
func (etu *EquipmentTypeUpdate) SetWidth(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetWidth()
	etu.mutation.SetWidth(f)
	return etu
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableWidth(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetWidth(*f)
	}
	return etu
}

// AddWidth adds f to the "width" field.
func (etu *EquipmentTypeUpdate) AddWidth(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddWidth(f)
	return etu
}

// ClearWidth clears the value of the "width" field.
func (etu *EquipmentTypeUpdate) ClearWidth() *EquipmentTypeUpdate {
	etu.mutation.ClearWidth()
	return etu
}

// SetWeight sets the "weight" field.
func (etu *EquipmentTypeUpdate) SetWeight(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetWeight()
	etu.mutation.SetWeight(f)
	return etu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableWeight(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetWeight(*f)
	}
	return etu
}

// AddWeight adds f to the "weight" field.
func (etu *EquipmentTypeUpdate) AddWeight(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddWeight(f)
	return etu
}

// ClearWeight clears the value of the "weight" field.
func (etu *EquipmentTypeUpdate) ClearWeight() *EquipmentTypeUpdate {
	etu.mutation.ClearWeight()
	return etu
}

// SetIdlingFuelUsage sets the "idling_fuel_usage" field.
func (etu *EquipmentTypeUpdate) SetIdlingFuelUsage(f float64) *EquipmentTypeUpdate {
	etu.mutation.ResetIdlingFuelUsage()
	etu.mutation.SetIdlingFuelUsage(f)
	return etu
}

// SetNillableIdlingFuelUsage sets the "idling_fuel_usage" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableIdlingFuelUsage(f *float64) *EquipmentTypeUpdate {
	if f != nil {
		etu.SetIdlingFuelUsage(*f)
	}
	return etu
}

// AddIdlingFuelUsage adds f to the "idling_fuel_usage" field.
func (etu *EquipmentTypeUpdate) AddIdlingFuelUsage(f float64) *EquipmentTypeUpdate {
	etu.mutation.AddIdlingFuelUsage(f)
	return etu
}

// ClearIdlingFuelUsage clears the value of the "idling_fuel_usage" field.
func (etu *EquipmentTypeUpdate) ClearIdlingFuelUsage() *EquipmentTypeUpdate {
	etu.mutation.ClearIdlingFuelUsage()
	return etu
}

// SetExemptFromTolls sets the "exempt_from_tolls" field.
func (etu *EquipmentTypeUpdate) SetExemptFromTolls(b bool) *EquipmentTypeUpdate {
	etu.mutation.SetExemptFromTolls(b)
	return etu
}

// SetNillableExemptFromTolls sets the "exempt_from_tolls" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableExemptFromTolls(b *bool) *EquipmentTypeUpdate {
	if b != nil {
		etu.SetExemptFromTolls(*b)
	}
	return etu
}

// SetColor sets the "color" field.
func (etu *EquipmentTypeUpdate) SetColor(s string) *EquipmentTypeUpdate {
	etu.mutation.SetColor(s)
	return etu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (etu *EquipmentTypeUpdate) SetNillableColor(s *string) *EquipmentTypeUpdate {
	if s != nil {
		etu.SetColor(*s)
	}
	return etu
}

// ClearColor clears the value of the "color" field.
func (etu *EquipmentTypeUpdate) ClearColor() *EquipmentTypeUpdate {
	etu.mutation.ClearColor()
	return etu
}

// Mutation returns the EquipmentTypeMutation object of the builder.
func (etu *EquipmentTypeUpdate) Mutation() *EquipmentTypeMutation {
	return etu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EquipmentTypeUpdate) Save(ctx context.Context) (int, error) {
	etu.defaults()
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EquipmentTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EquipmentTypeUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EquipmentTypeUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etu *EquipmentTypeUpdate) defaults() {
	if _, ok := etu.mutation.UpdatedAt(); !ok {
		v := equipmenttype.UpdateDefaultUpdatedAt()
		etu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etu *EquipmentTypeUpdate) check() error {
	if v, ok := etu.mutation.Status(); ok {
		if err := equipmenttype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.status": %w`, err)}
		}
	}
	if v, ok := etu.mutation.Code(); ok {
		if err := equipmenttype.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.code": %w`, err)}
		}
	}
	if v, ok := etu.mutation.EquipmentClass(); ok {
		if err := equipmenttype.EquipmentClassValidator(v); err != nil {
			return &ValidationError{Name: "equipment_class", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.equipment_class": %w`, err)}
		}
	}
	if _, ok := etu.mutation.BusinessUnitID(); etu.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentType.business_unit"`)
	}
	if _, ok := etu.mutation.OrganizationID(); etu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentType.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etu *EquipmentTypeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EquipmentTypeUpdate {
	etu.modifiers = append(etu.modifiers, modifiers...)
	return etu
}

func (etu *EquipmentTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := etu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipmenttype.Table, equipmenttype.Columns, sqlgraph.NewFieldSpec(equipmenttype.FieldID, field.TypeUUID))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmenttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := etu.mutation.Version(); ok {
		_spec.SetField(equipmenttype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := etu.mutation.AddedVersion(); ok {
		_spec.AddField(equipmenttype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := etu.mutation.Status(); ok {
		_spec.SetField(equipmenttype.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := etu.mutation.Code(); ok {
		_spec.SetField(equipmenttype.FieldCode, field.TypeString, value)
	}
	if value, ok := etu.mutation.Description(); ok {
		_spec.SetField(equipmenttype.FieldDescription, field.TypeString, value)
	}
	if etu.mutation.DescriptionCleared() {
		_spec.ClearField(equipmenttype.FieldDescription, field.TypeString)
	}
	if value, ok := etu.mutation.CostPerMile(); ok {
		_spec.SetField(equipmenttype.FieldCostPerMile, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedCostPerMile(); ok {
		_spec.AddField(equipmenttype.FieldCostPerMile, field.TypeFloat64, value)
	}
	if etu.mutation.CostPerMileCleared() {
		_spec.ClearField(equipmenttype.FieldCostPerMile, field.TypeFloat64)
	}
	if value, ok := etu.mutation.EquipmentClass(); ok {
		_spec.SetField(equipmenttype.FieldEquipmentClass, field.TypeEnum, value)
	}
	if value, ok := etu.mutation.FixedCost(); ok {
		_spec.SetField(equipmenttype.FieldFixedCost, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedFixedCost(); ok {
		_spec.AddField(equipmenttype.FieldFixedCost, field.TypeFloat64, value)
	}
	if etu.mutation.FixedCostCleared() {
		_spec.ClearField(equipmenttype.FieldFixedCost, field.TypeFloat64)
	}
	if value, ok := etu.mutation.VariableCost(); ok {
		_spec.SetField(equipmenttype.FieldVariableCost, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedVariableCost(); ok {
		_spec.AddField(equipmenttype.FieldVariableCost, field.TypeFloat64, value)
	}
	if etu.mutation.VariableCostCleared() {
		_spec.ClearField(equipmenttype.FieldVariableCost, field.TypeFloat64)
	}
	if value, ok := etu.mutation.Height(); ok {
		_spec.SetField(equipmenttype.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedHeight(); ok {
		_spec.AddField(equipmenttype.FieldHeight, field.TypeFloat64, value)
	}
	if etu.mutation.HeightCleared() {
		_spec.ClearField(equipmenttype.FieldHeight, field.TypeFloat64)
	}
	if value, ok := etu.mutation.Length(); ok {
		_spec.SetField(equipmenttype.FieldLength, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedLength(); ok {
		_spec.AddField(equipmenttype.FieldLength, field.TypeFloat64, value)
	}
	if etu.mutation.LengthCleared() {
		_spec.ClearField(equipmenttype.FieldLength, field.TypeFloat64)
	}
	if value, ok := etu.mutation.Width(); ok {
		_spec.SetField(equipmenttype.FieldWidth, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedWidth(); ok {
		_spec.AddField(equipmenttype.FieldWidth, field.TypeFloat64, value)
	}
	if etu.mutation.WidthCleared() {
		_spec.ClearField(equipmenttype.FieldWidth, field.TypeFloat64)
	}
	if value, ok := etu.mutation.Weight(); ok {
		_spec.SetField(equipmenttype.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedWeight(); ok {
		_spec.AddField(equipmenttype.FieldWeight, field.TypeFloat64, value)
	}
	if etu.mutation.WeightCleared() {
		_spec.ClearField(equipmenttype.FieldWeight, field.TypeFloat64)
	}
	if value, ok := etu.mutation.IdlingFuelUsage(); ok {
		_spec.SetField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64, value)
	}
	if value, ok := etu.mutation.AddedIdlingFuelUsage(); ok {
		_spec.AddField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64, value)
	}
	if etu.mutation.IdlingFuelUsageCleared() {
		_spec.ClearField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64)
	}
	if value, ok := etu.mutation.ExemptFromTolls(); ok {
		_spec.SetField(equipmenttype.FieldExemptFromTolls, field.TypeBool, value)
	}
	if value, ok := etu.mutation.Color(); ok {
		_spec.SetField(equipmenttype.FieldColor, field.TypeString, value)
	}
	if etu.mutation.ColorCleared() {
		_spec.ClearField(equipmenttype.FieldColor, field.TypeString)
	}
	_spec.AddModifiers(etu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// EquipmentTypeUpdateOne is the builder for updating a single EquipmentType entity.
type EquipmentTypeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EquipmentTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (etuo *EquipmentTypeUpdateOne) SetUpdatedAt(t time.Time) *EquipmentTypeUpdateOne {
	etuo.mutation.SetUpdatedAt(t)
	return etuo
}

// SetVersion sets the "version" field.
func (etuo *EquipmentTypeUpdateOne) SetVersion(i int) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetVersion()
	etuo.mutation.SetVersion(i)
	return etuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableVersion(i *int) *EquipmentTypeUpdateOne {
	if i != nil {
		etuo.SetVersion(*i)
	}
	return etuo
}

// AddVersion adds i to the "version" field.
func (etuo *EquipmentTypeUpdateOne) AddVersion(i int) *EquipmentTypeUpdateOne {
	etuo.mutation.AddVersion(i)
	return etuo
}

// SetStatus sets the "status" field.
func (etuo *EquipmentTypeUpdateOne) SetStatus(e equipmenttype.Status) *EquipmentTypeUpdateOne {
	etuo.mutation.SetStatus(e)
	return etuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableStatus(e *equipmenttype.Status) *EquipmentTypeUpdateOne {
	if e != nil {
		etuo.SetStatus(*e)
	}
	return etuo
}

// SetCode sets the "code" field.
func (etuo *EquipmentTypeUpdateOne) SetCode(s string) *EquipmentTypeUpdateOne {
	etuo.mutation.SetCode(s)
	return etuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableCode(s *string) *EquipmentTypeUpdateOne {
	if s != nil {
		etuo.SetCode(*s)
	}
	return etuo
}

// SetDescription sets the "description" field.
func (etuo *EquipmentTypeUpdateOne) SetDescription(s string) *EquipmentTypeUpdateOne {
	etuo.mutation.SetDescription(s)
	return etuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableDescription(s *string) *EquipmentTypeUpdateOne {
	if s != nil {
		etuo.SetDescription(*s)
	}
	return etuo
}

// ClearDescription clears the value of the "description" field.
func (etuo *EquipmentTypeUpdateOne) ClearDescription() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearDescription()
	return etuo
}

// SetCostPerMile sets the "cost_per_mile" field.
func (etuo *EquipmentTypeUpdateOne) SetCostPerMile(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetCostPerMile()
	etuo.mutation.SetCostPerMile(f)
	return etuo
}

// SetNillableCostPerMile sets the "cost_per_mile" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableCostPerMile(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetCostPerMile(*f)
	}
	return etuo
}

// AddCostPerMile adds f to the "cost_per_mile" field.
func (etuo *EquipmentTypeUpdateOne) AddCostPerMile(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddCostPerMile(f)
	return etuo
}

// ClearCostPerMile clears the value of the "cost_per_mile" field.
func (etuo *EquipmentTypeUpdateOne) ClearCostPerMile() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearCostPerMile()
	return etuo
}

// SetEquipmentClass sets the "equipment_class" field.
func (etuo *EquipmentTypeUpdateOne) SetEquipmentClass(ec equipmenttype.EquipmentClass) *EquipmentTypeUpdateOne {
	etuo.mutation.SetEquipmentClass(ec)
	return etuo
}

// SetNillableEquipmentClass sets the "equipment_class" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableEquipmentClass(ec *equipmenttype.EquipmentClass) *EquipmentTypeUpdateOne {
	if ec != nil {
		etuo.SetEquipmentClass(*ec)
	}
	return etuo
}

// SetFixedCost sets the "fixed_cost" field.
func (etuo *EquipmentTypeUpdateOne) SetFixedCost(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetFixedCost()
	etuo.mutation.SetFixedCost(f)
	return etuo
}

// SetNillableFixedCost sets the "fixed_cost" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableFixedCost(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetFixedCost(*f)
	}
	return etuo
}

// AddFixedCost adds f to the "fixed_cost" field.
func (etuo *EquipmentTypeUpdateOne) AddFixedCost(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddFixedCost(f)
	return etuo
}

// ClearFixedCost clears the value of the "fixed_cost" field.
func (etuo *EquipmentTypeUpdateOne) ClearFixedCost() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearFixedCost()
	return etuo
}

// SetVariableCost sets the "variable_cost" field.
func (etuo *EquipmentTypeUpdateOne) SetVariableCost(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetVariableCost()
	etuo.mutation.SetVariableCost(f)
	return etuo
}

// SetNillableVariableCost sets the "variable_cost" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableVariableCost(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetVariableCost(*f)
	}
	return etuo
}

// AddVariableCost adds f to the "variable_cost" field.
func (etuo *EquipmentTypeUpdateOne) AddVariableCost(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddVariableCost(f)
	return etuo
}

// ClearVariableCost clears the value of the "variable_cost" field.
func (etuo *EquipmentTypeUpdateOne) ClearVariableCost() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearVariableCost()
	return etuo
}

// SetHeight sets the "height" field.
func (etuo *EquipmentTypeUpdateOne) SetHeight(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetHeight()
	etuo.mutation.SetHeight(f)
	return etuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableHeight(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetHeight(*f)
	}
	return etuo
}

// AddHeight adds f to the "height" field.
func (etuo *EquipmentTypeUpdateOne) AddHeight(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddHeight(f)
	return etuo
}

// ClearHeight clears the value of the "height" field.
func (etuo *EquipmentTypeUpdateOne) ClearHeight() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearHeight()
	return etuo
}

// SetLength sets the "length" field.
func (etuo *EquipmentTypeUpdateOne) SetLength(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetLength()
	etuo.mutation.SetLength(f)
	return etuo
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableLength(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetLength(*f)
	}
	return etuo
}

// AddLength adds f to the "length" field.
func (etuo *EquipmentTypeUpdateOne) AddLength(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddLength(f)
	return etuo
}

// ClearLength clears the value of the "length" field.
func (etuo *EquipmentTypeUpdateOne) ClearLength() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearLength()
	return etuo
}

// SetWidth sets the "width" field.
func (etuo *EquipmentTypeUpdateOne) SetWidth(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetWidth()
	etuo.mutation.SetWidth(f)
	return etuo
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableWidth(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetWidth(*f)
	}
	return etuo
}

// AddWidth adds f to the "width" field.
func (etuo *EquipmentTypeUpdateOne) AddWidth(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddWidth(f)
	return etuo
}

// ClearWidth clears the value of the "width" field.
func (etuo *EquipmentTypeUpdateOne) ClearWidth() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearWidth()
	return etuo
}

// SetWeight sets the "weight" field.
func (etuo *EquipmentTypeUpdateOne) SetWeight(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetWeight()
	etuo.mutation.SetWeight(f)
	return etuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableWeight(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetWeight(*f)
	}
	return etuo
}

// AddWeight adds f to the "weight" field.
func (etuo *EquipmentTypeUpdateOne) AddWeight(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddWeight(f)
	return etuo
}

// ClearWeight clears the value of the "weight" field.
func (etuo *EquipmentTypeUpdateOne) ClearWeight() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearWeight()
	return etuo
}

// SetIdlingFuelUsage sets the "idling_fuel_usage" field.
func (etuo *EquipmentTypeUpdateOne) SetIdlingFuelUsage(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.ResetIdlingFuelUsage()
	etuo.mutation.SetIdlingFuelUsage(f)
	return etuo
}

// SetNillableIdlingFuelUsage sets the "idling_fuel_usage" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableIdlingFuelUsage(f *float64) *EquipmentTypeUpdateOne {
	if f != nil {
		etuo.SetIdlingFuelUsage(*f)
	}
	return etuo
}

// AddIdlingFuelUsage adds f to the "idling_fuel_usage" field.
func (etuo *EquipmentTypeUpdateOne) AddIdlingFuelUsage(f float64) *EquipmentTypeUpdateOne {
	etuo.mutation.AddIdlingFuelUsage(f)
	return etuo
}

// ClearIdlingFuelUsage clears the value of the "idling_fuel_usage" field.
func (etuo *EquipmentTypeUpdateOne) ClearIdlingFuelUsage() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearIdlingFuelUsage()
	return etuo
}

// SetExemptFromTolls sets the "exempt_from_tolls" field.
func (etuo *EquipmentTypeUpdateOne) SetExemptFromTolls(b bool) *EquipmentTypeUpdateOne {
	etuo.mutation.SetExemptFromTolls(b)
	return etuo
}

// SetNillableExemptFromTolls sets the "exempt_from_tolls" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableExemptFromTolls(b *bool) *EquipmentTypeUpdateOne {
	if b != nil {
		etuo.SetExemptFromTolls(*b)
	}
	return etuo
}

// SetColor sets the "color" field.
func (etuo *EquipmentTypeUpdateOne) SetColor(s string) *EquipmentTypeUpdateOne {
	etuo.mutation.SetColor(s)
	return etuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (etuo *EquipmentTypeUpdateOne) SetNillableColor(s *string) *EquipmentTypeUpdateOne {
	if s != nil {
		etuo.SetColor(*s)
	}
	return etuo
}

// ClearColor clears the value of the "color" field.
func (etuo *EquipmentTypeUpdateOne) ClearColor() *EquipmentTypeUpdateOne {
	etuo.mutation.ClearColor()
	return etuo
}

// Mutation returns the EquipmentTypeMutation object of the builder.
func (etuo *EquipmentTypeUpdateOne) Mutation() *EquipmentTypeMutation {
	return etuo.mutation
}

// Where appends a list predicates to the EquipmentTypeUpdate builder.
func (etuo *EquipmentTypeUpdateOne) Where(ps ...predicate.EquipmentType) *EquipmentTypeUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EquipmentTypeUpdateOne) Select(field string, fields ...string) *EquipmentTypeUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EquipmentType entity.
func (etuo *EquipmentTypeUpdateOne) Save(ctx context.Context) (*EquipmentType, error) {
	etuo.defaults()
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EquipmentTypeUpdateOne) SaveX(ctx context.Context) *EquipmentType {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EquipmentTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EquipmentTypeUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etuo *EquipmentTypeUpdateOne) defaults() {
	if _, ok := etuo.mutation.UpdatedAt(); !ok {
		v := equipmenttype.UpdateDefaultUpdatedAt()
		etuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etuo *EquipmentTypeUpdateOne) check() error {
	if v, ok := etuo.mutation.Status(); ok {
		if err := equipmenttype.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.status": %w`, err)}
		}
	}
	if v, ok := etuo.mutation.Code(); ok {
		if err := equipmenttype.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.code": %w`, err)}
		}
	}
	if v, ok := etuo.mutation.EquipmentClass(); ok {
		if err := equipmenttype.EquipmentClassValidator(v); err != nil {
			return &ValidationError{Name: "equipment_class", err: fmt.Errorf(`ent: validator failed for field "EquipmentType.equipment_class": %w`, err)}
		}
	}
	if _, ok := etuo.mutation.BusinessUnitID(); etuo.mutation.BusinessUnitCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentType.business_unit"`)
	}
	if _, ok := etuo.mutation.OrganizationID(); etuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EquipmentType.organization"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (etuo *EquipmentTypeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EquipmentTypeUpdateOne {
	etuo.modifiers = append(etuo.modifiers, modifiers...)
	return etuo
}

func (etuo *EquipmentTypeUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentType, err error) {
	if err := etuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipmenttype.Table, equipmenttype.Columns, sqlgraph.NewFieldSpec(equipmenttype.FieldID, field.TypeUUID))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmenttype.FieldID)
		for _, f := range fields {
			if !equipmenttype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmenttype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := etuo.mutation.Version(); ok {
		_spec.SetField(equipmenttype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := etuo.mutation.AddedVersion(); ok {
		_spec.AddField(equipmenttype.FieldVersion, field.TypeInt, value)
	}
	if value, ok := etuo.mutation.Status(); ok {
		_spec.SetField(equipmenttype.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := etuo.mutation.Code(); ok {
		_spec.SetField(equipmenttype.FieldCode, field.TypeString, value)
	}
	if value, ok := etuo.mutation.Description(); ok {
		_spec.SetField(equipmenttype.FieldDescription, field.TypeString, value)
	}
	if etuo.mutation.DescriptionCleared() {
		_spec.ClearField(equipmenttype.FieldDescription, field.TypeString)
	}
	if value, ok := etuo.mutation.CostPerMile(); ok {
		_spec.SetField(equipmenttype.FieldCostPerMile, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedCostPerMile(); ok {
		_spec.AddField(equipmenttype.FieldCostPerMile, field.TypeFloat64, value)
	}
	if etuo.mutation.CostPerMileCleared() {
		_spec.ClearField(equipmenttype.FieldCostPerMile, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.EquipmentClass(); ok {
		_spec.SetField(equipmenttype.FieldEquipmentClass, field.TypeEnum, value)
	}
	if value, ok := etuo.mutation.FixedCost(); ok {
		_spec.SetField(equipmenttype.FieldFixedCost, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedFixedCost(); ok {
		_spec.AddField(equipmenttype.FieldFixedCost, field.TypeFloat64, value)
	}
	if etuo.mutation.FixedCostCleared() {
		_spec.ClearField(equipmenttype.FieldFixedCost, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.VariableCost(); ok {
		_spec.SetField(equipmenttype.FieldVariableCost, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedVariableCost(); ok {
		_spec.AddField(equipmenttype.FieldVariableCost, field.TypeFloat64, value)
	}
	if etuo.mutation.VariableCostCleared() {
		_spec.ClearField(equipmenttype.FieldVariableCost, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.Height(); ok {
		_spec.SetField(equipmenttype.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedHeight(); ok {
		_spec.AddField(equipmenttype.FieldHeight, field.TypeFloat64, value)
	}
	if etuo.mutation.HeightCleared() {
		_spec.ClearField(equipmenttype.FieldHeight, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.Length(); ok {
		_spec.SetField(equipmenttype.FieldLength, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedLength(); ok {
		_spec.AddField(equipmenttype.FieldLength, field.TypeFloat64, value)
	}
	if etuo.mutation.LengthCleared() {
		_spec.ClearField(equipmenttype.FieldLength, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.Width(); ok {
		_spec.SetField(equipmenttype.FieldWidth, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedWidth(); ok {
		_spec.AddField(equipmenttype.FieldWidth, field.TypeFloat64, value)
	}
	if etuo.mutation.WidthCleared() {
		_spec.ClearField(equipmenttype.FieldWidth, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.Weight(); ok {
		_spec.SetField(equipmenttype.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedWeight(); ok {
		_spec.AddField(equipmenttype.FieldWeight, field.TypeFloat64, value)
	}
	if etuo.mutation.WeightCleared() {
		_spec.ClearField(equipmenttype.FieldWeight, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.IdlingFuelUsage(); ok {
		_spec.SetField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64, value)
	}
	if value, ok := etuo.mutation.AddedIdlingFuelUsage(); ok {
		_spec.AddField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64, value)
	}
	if etuo.mutation.IdlingFuelUsageCleared() {
		_spec.ClearField(equipmenttype.FieldIdlingFuelUsage, field.TypeFloat64)
	}
	if value, ok := etuo.mutation.ExemptFromTolls(); ok {
		_spec.SetField(equipmenttype.FieldExemptFromTolls, field.TypeBool, value)
	}
	if value, ok := etuo.mutation.Color(); ok {
		_spec.SetField(equipmenttype.FieldColor, field.TypeString, value)
	}
	if etuo.mutation.ColorCleared() {
		_spec.ClearField(equipmenttype.FieldColor, field.TypeString)
	}
	_spec.AddModifiers(etuo.modifiers...)
	_node = &EquipmentType{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmenttype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}
