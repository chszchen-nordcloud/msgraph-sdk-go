package identitygovernance

import (
	ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/identitygovernance"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody
type LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody struct {
	// Stores model information.
	backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}

// NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody instantiates a new LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody() *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody {
	m := &LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody{}
	m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance()
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetAdditionalData() map[string]any {
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore {
	return m.backingStore
}

// GetData gets the data property value. The data property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetData() ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable {
	val, err := m.GetBackingStore().Get("data")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable)
	}
	return nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateCustomTaskExtensionCallbackDataFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetData(val.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable))
		}
		return nil
	}
	res["source"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSource(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetType(val)
		}
		return nil
	}
	return res
}

// GetSource gets the source property value. The source property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetSource() *string {
	val, err := m.GetBackingStore().Get("source")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// GetType gets the type property value. The type property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) GetType() *string {
	val, err := m.GetBackingStore().Get("typeEscaped")
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val.(*string)
	}
	return nil
}

// Serialize serializes information the current object
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("data", m.GetData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("source", m.GetSource())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetType())
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetAdditionalData(value map[string]any) {
	err := m.GetBackingStore().Set("additionalData", value)
	if err != nil {
		panic(err)
	}
}

// SetBackingStore sets the backingStore property value. Stores model information.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
	m.backingStore = value
}

// SetData sets the data property value. The data property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetData(value ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable) {
	err := m.GetBackingStore().Set("data", value)
	if err != nil {
		panic(err)
	}
}

// SetSource sets the source property value. The source property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetSource(value *string) {
	err := m.GetBackingStore().Set("source", value)
	if err != nil {
		panic(err)
	}
}

// SetType sets the type property value. The type property
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBody) SetType(value *string) {
	err := m.GetBackingStore().Set("typeEscaped", value)
	if err != nil {
		panic(err)
	}
}

// LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyable
type LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeResumePostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBackingStore() ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
	GetData() ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable
	GetSource() *string
	GetType() *string
	SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
	SetData(value ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCallbackDataable)
	SetSource(value *string)
	SetType(value *string)
}
