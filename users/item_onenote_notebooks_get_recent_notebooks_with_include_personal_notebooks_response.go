package users

import (
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse struct {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}

// NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse instantiates a new ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse and sets the default values.
func NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse() *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse {
	m := &ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse{
		BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
	}
	return m
}

// CreateItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRecentNotebookFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable)
				}
			}
			m.SetValue(res)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable {
	val, err := m.GetBackingStore().Get("value")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable)
	}
	return nil
}

// Serialize serializes information the current object
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable) {
	err := m.GetBackingStore().Set("value", value)
	if err != nil {
		panic(err)
	}
}

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable interface {
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable
	SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RecentNotebookable)
}
