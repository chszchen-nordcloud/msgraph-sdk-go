package security

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntelligenceProfileIndicatorCollectionResponse
type IntelligenceProfileIndicatorCollectionResponse struct {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}

// NewIntelligenceProfileIndicatorCollectionResponse instantiates a new IntelligenceProfileIndicatorCollectionResponse and sets the default values.
func NewIntelligenceProfileIndicatorCollectionResponse() *IntelligenceProfileIndicatorCollectionResponse {
	m := &IntelligenceProfileIndicatorCollectionResponse{
		BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateIntelligenceProfileIndicatorCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntelligenceProfileIndicatorCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewIntelligenceProfileIndicatorCollectionResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *IntelligenceProfileIndicatorCollectionResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateIntelligenceProfileIndicatorFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]IntelligenceProfileIndicatorable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(IntelligenceProfileIndicatorable)
				}
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *IntelligenceProfileIndicatorCollectionResponse) GetValue() []IntelligenceProfileIndicatorable {
	val, err := m.GetBackingStore().Get("value")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]IntelligenceProfileIndicatorable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *IntelligenceProfileIndicatorCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
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
func (m *IntelligenceProfileIndicatorCollectionResponse) SetValue(value []IntelligenceProfileIndicatorable) {
	err := m.GetBackingStore().Set("value", value)
	if err != nil {
		panic(err)
	}
}

// IntelligenceProfileIndicatorCollectionResponseable
type IntelligenceProfileIndicatorCollectionResponseable interface {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []IntelligenceProfileIndicatorable
	SetValue(value []IntelligenceProfileIndicatorable)
}
