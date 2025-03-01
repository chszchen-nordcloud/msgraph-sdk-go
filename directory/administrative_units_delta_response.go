package directory

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnitsDeltaResponse
type AdministrativeUnitsDeltaResponse struct {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseDeltaFunctionResponse
}

// NewAdministrativeUnitsDeltaResponse instantiates a new AdministrativeUnitsDeltaResponse and sets the default values.
func NewAdministrativeUnitsDeltaResponse() *AdministrativeUnitsDeltaResponse {
	m := &AdministrativeUnitsDeltaResponse{
		BaseDeltaFunctionResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseDeltaFunctionResponse(),
	}
	return m
}

// CreateAdministrativeUnitsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAdministrativeUnitsDeltaResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *AdministrativeUnitsDeltaResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAdministrativeUnitFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable)
				}
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *AdministrativeUnitsDeltaResponse) GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable {
	val, err := m.GetBackingStore().Get("value")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *AdministrativeUnitsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.BaseDeltaFunctionResponse.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetValue() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
		for i, v := range m.GetValue() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err = writer.WriteCollectionOfObjectValues("value", cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValue sets the value property value. The value property
func (m *AdministrativeUnitsDeltaResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable) {
	err := m.GetBackingStore().Set("value", value)
	if err != nil {
		panic(err)
	}
}

// AdministrativeUnitsDeltaResponseable
type AdministrativeUnitsDeltaResponseable interface {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseDeltaFunctionResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable
	SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AdministrativeUnitable)
}
