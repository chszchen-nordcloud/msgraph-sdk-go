package connections

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc "github.com/chszchen-nordcloud/msgraph-sdk-go/models/externalconnectors"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse struct {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}

// NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse instantiates a new ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse and sets the default values.
func NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse() *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse {
	m := &ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse{
		BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.CreateExternalActivityResultFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable)
				}
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) GetValue() []i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable {
	val, err := m.GetBackingStore().Get("value")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) SetValue(value []i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable) {
	err := m.GetBackingStore().Set("value", value)
	if err != nil {
		panic(err)
	}
}

// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable interface {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable
	SetValue(value []i648e92ed22999203da3c8fad3bc63deefe974fd0d511e7f830d70ea0aff57ffc.ExternalActivityResultable)
}
