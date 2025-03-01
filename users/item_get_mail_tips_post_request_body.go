package users

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemGetMailTipsPostRequestBody
type ItemGetMailTipsPostRequestBody struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewItemGetMailTipsPostRequestBody instantiates a new ItemGetMailTipsPostRequestBody and sets the default values.
func NewItemGetMailTipsPostRequestBody() *ItemGetMailTipsPostRequestBody {
	m := &ItemGetMailTipsPostRequestBody{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateItemGetMailTipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetMailTipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemGetMailTipsPostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemGetMailTipsPostRequestBody) GetAdditionalData() map[string]any {
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
func (m *ItemGetMailTipsPostRequestBody) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetEmailAddresses gets the emailAddresses property value. The EmailAddresses property
func (m *ItemGetMailTipsPostRequestBody) GetEmailAddresses() []string {
	val, err := m.GetBackingStore().Get("emailAddresses")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]string)
	}
	return nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemGetMailTipsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["emailAddresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetEmailAddresses(res)
		}
		return nil
	}
	res["mailTipsOptions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseMailTipsType)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMailTipsOptions(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType))
		}
		return nil
	}
	return res
}

// GetMailTipsOptions gets the mailTipsOptions property value. The MailTipsOptions property
func (m *ItemGetMailTipsPostRequestBody) GetMailTipsOptions() *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType {
	val, err := m.GetBackingStore().Get("mailTipsOptions")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType)
	}
	return nil
}

// Serialize serializes information the current object
func (m *ItemGetMailTipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetEmailAddresses() != nil {
		err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
		if err != nil {
			return err
		}
	}
	if m.GetMailTipsOptions() != nil {
		cast := (*m.GetMailTipsOptions()).String()
		err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *ItemGetMailTipsPostRequestBody) SetAdditionalData(value map[string]any) {
	err := m.GetBackingStore().Set("additionalData", value)
	if err != nil {
		panic(err)
	}
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemGetMailTipsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetEmailAddresses sets the emailAddresses property value. The EmailAddresses property
func (m *ItemGetMailTipsPostRequestBody) SetEmailAddresses(value []string) {
	err := m.GetBackingStore().Set("emailAddresses", value)
	if err != nil {
		panic(err)
	}
}

// SetMailTipsOptions sets the mailTipsOptions property value. The MailTipsOptions property
func (m *ItemGetMailTipsPostRequestBody) SetMailTipsOptions(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType) {
	err := m.GetBackingStore().Set("mailTipsOptions", value)
	if err != nil {
		panic(err)
	}
}

// ItemGetMailTipsPostRequestBodyable
type ItemGetMailTipsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetEmailAddresses() []string
	GetMailTipsOptions() *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetEmailAddresses(value []string)
	SetMailTipsOptions(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType)
}
