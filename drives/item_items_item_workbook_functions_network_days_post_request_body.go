package drives

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody
type ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody instantiates a new ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody and sets the default values.
func NewItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody() *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody {
	m := &ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetAdditionalData() map[string]any {
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
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetEndDate gets the endDate property value. The endDate property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetEndDate() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable {
	val, err := m.GetBackingStore().Get("endDate")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
	}
	return nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["endDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEndDate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
		}
		return nil
	}
	res["holidays"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHolidays(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
		}
		return nil
	}
	res["startDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateJsonFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartDate(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable))
		}
		return nil
	}
	return res
}

// GetHolidays gets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetHolidays() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable {
	val, err := m.GetBackingStore().Get("holidays")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
	}
	return nil
}

// GetStartDate gets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) GetStartDate() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable {
	val, err := m.GetBackingStore().Get("startDate")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("endDate", m.GetEndDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("holidays", m.GetHolidays())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("startDate", m.GetStartDate())
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
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) SetAdditionalData(value map[string]any) {
	err := m.GetBackingStore().Set("additionalData", value)
	if err != nil {
		panic(err)
	}
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetEndDate sets the endDate property value. The endDate property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) SetEndDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
	err := m.GetBackingStore().Set("endDate", value)
	if err != nil {
		panic(err)
	}
}

// SetHolidays sets the holidays property value. The holidays property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) SetHolidays(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
	err := m.GetBackingStore().Set("holidays", value)
	if err != nil {
		panic(err)
	}
}

// SetStartDate sets the startDate property value. The startDate property
func (m *ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBody) SetStartDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable) {
	err := m.GetBackingStore().Set("startDate", value)
	if err != nil {
		panic(err)
	}
}

// ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBodyable
type ItemItemsItemWorkbookFunctionsNetworkDaysPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetEndDate() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
	GetHolidays() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
	GetStartDate() iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetEndDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
	SetHolidays(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
	SetStartDate(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Jsonable)
}
