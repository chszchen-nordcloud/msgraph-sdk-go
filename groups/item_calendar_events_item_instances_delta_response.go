package groups

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarEventsItemInstancesDeltaResponse
type ItemCalendarEventsItemInstancesDeltaResponse struct {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseDeltaFunctionResponse
}

// NewItemCalendarEventsItemInstancesDeltaResponse instantiates a new ItemCalendarEventsItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarEventsItemInstancesDeltaResponse() *ItemCalendarEventsItemInstancesDeltaResponse {
	m := &ItemCalendarEventsItemInstancesDeltaResponse{
		BaseDeltaFunctionResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseDeltaFunctionResponse(),
	}
	return m
}

// CreateItemCalendarEventsItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemCalendarEventsItemInstancesDeltaResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemInstancesDeltaResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)
				}
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *ItemCalendarEventsItemInstancesDeltaResponse) GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable {
	val, err := m.GetBackingStore().Get("value")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *ItemCalendarEventsItemInstancesDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ItemCalendarEventsItemInstancesDeltaResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable) {
	err := m.GetBackingStore().Set("value", value)
	if err != nil {
		panic(err)
	}
}

// ItemCalendarEventsItemInstancesDeltaResponseable
type ItemCalendarEventsItemInstancesDeltaResponseable interface {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseDeltaFunctionResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable
	SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable)
}
