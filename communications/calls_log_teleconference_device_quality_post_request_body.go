package communications

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CallsLogTeleconferenceDeviceQualityPostRequestBody
type CallsLogTeleconferenceDeviceQualityPostRequestBody struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewCallsLogTeleconferenceDeviceQualityPostRequestBody instantiates a new CallsLogTeleconferenceDeviceQualityPostRequestBody and sets the default values.
func NewCallsLogTeleconferenceDeviceQualityPostRequestBody() *CallsLogTeleconferenceDeviceQualityPostRequestBody {
	m := &CallsLogTeleconferenceDeviceQualityPostRequestBody{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCallsLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCallsLogTeleconferenceDeviceQualityPostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) GetAdditionalData() map[string]any {
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
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetFieldDeserializers the deserialization information for the current model
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["quality"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeleconferenceDeviceQualityFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetQuality(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable))
		}
		return nil
	}
	return res
}

// GetQuality gets the quality property value. The quality property
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) GetQuality() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable {
	val, err := m.GetBackingStore().Get("quality")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("quality", m.GetQuality())
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
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) SetAdditionalData(value map[string]any) {
	err := m.GetBackingStore().Set("additionalData", value)
	if err != nil {
		panic(err)
	}
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetQuality sets the quality property value. The quality property
func (m *CallsLogTeleconferenceDeviceQualityPostRequestBody) SetQuality(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable) {
	err := m.GetBackingStore().Set("quality", value)
	if err != nil {
		panic(err)
	}
}

// CallsLogTeleconferenceDeviceQualityPostRequestBodyable
type CallsLogTeleconferenceDeviceQualityPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetQuality() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetQuality(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable)
}
