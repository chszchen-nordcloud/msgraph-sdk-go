package communications

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CallsItemUpdateRecordingStatusPostRequestBody
type CallsItemUpdateRecordingStatusPostRequestBody struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewCallsItemUpdateRecordingStatusPostRequestBody instantiates a new CallsItemUpdateRecordingStatusPostRequestBody and sets the default values.
func NewCallsItemUpdateRecordingStatusPostRequestBody() *CallsItemUpdateRecordingStatusPostRequestBody {
	m := &CallsItemUpdateRecordingStatusPostRequestBody{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCallsItemUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemUpdateRecordingStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCallsItemUpdateRecordingStatusPostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetAdditionalData() map[string]any {
	val, err := m.backingStore.Get("additionalData")
	if err != nil {
		panic(err)
	}
	if val == nil {
		var value = make(map[string]any)
		m.SetAdditionalData(value)
	}
	return val.(map[string]any)
}

// GetBackingStore gets the backingStore property value. Stores model information.
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetClientContext gets the clientContext property value. The clientContext property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetClientContext() *string {
	val, err := m.GetBackingStore().Get("clientContext")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["clientContext"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetClientContext(val)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseRecordingStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus))
		}
		return nil
	}
	return res
}

// GetStatus gets the status property value. The status property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) GetStatus() *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus {
	val, err := m.GetBackingStore().Get("status")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus)
	}
	return nil
}

// Serialize serializes information the current object
func (m *CallsItemUpdateRecordingStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("clientContext", m.GetClientContext())
		if err != nil {
			return err
		}
	}
	if m.GetStatus() != nil {
		cast := (*m.GetStatus()).String()
		err := writer.WriteStringValue("status", &cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetAdditionalData(value map[string]any) {
	err := m.GetBackingStore().Set("additionalData", value)
	if err != nil {
		panic(err)
	}
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetClientContext sets the clientContext property value. The clientContext property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetClientContext(value *string) {
	err := m.GetBackingStore().Set("clientContext", value)
	if err != nil {
		panic(err)
	}
}

// SetStatus sets the status property value. The status property
func (m *CallsItemUpdateRecordingStatusPostRequestBody) SetStatus(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus) {
	err := m.GetBackingStore().Set("status", value)
	if err != nil {
		panic(err)
	}
}

// CallsItemUpdateRecordingStatusPostRequestBodyable
type CallsItemUpdateRecordingStatusPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetClientContext() *string
	GetStatus() *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetClientContext(value *string)
	SetStatus(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecordingStatus)
}
