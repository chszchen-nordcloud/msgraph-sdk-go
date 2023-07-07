package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters to indicate which terms are related to the current term as either pinned or reused.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetQueryParameters
}

// ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderInternal instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder {
	m := &ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/children/{term%2Did1}/relations/{relation%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder instantiates a new RelationItemRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property relations for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// FromTerm provides operations to manage the fromTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) FromTerm() *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemFromTermRequestBuilder {
	return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemFromTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable), nil
}

// Patch update the navigation property relations in sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateRelationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable), nil
}

// Set provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) Set() *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilder {
	return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemSetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation delete navigation property relations for sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation to indicate which terms are related to the current term as either pinned or reused.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToPatchRequestInformation update the navigation property relations in sites
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Relationable, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
	requestInfo.Headers.Add("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToTerm provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
func (m *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsRelationItemRequestBuilder) ToTerm() *ItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilder {
	return NewItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemRelationsItemToTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
