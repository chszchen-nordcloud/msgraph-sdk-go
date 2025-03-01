package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetQueryParameters
}

// ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Children provides operations to manage the children property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Children() *ItemSitesItemTermStoresItemGroupsItemSetsItemChildrenRequestBuilder {
	return NewItemSitesItemTermStoresItemGroupsItemSetsItemChildrenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder {
	m := &ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did1}/sets/{set%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property sets for groups
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get all sets under the group in a term [store].
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}

// ParentGroup provides operations to manage the parentGroup property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) ParentGroup() *ItemSitesItemTermStoresItemGroupsItemSetsItemParentGroupRequestBuilder {
	return NewItemSitesItemTermStoresItemGroupsItemSetsItemParentGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch update the navigation property sets in groups
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Patch(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}

// Relations provides operations to manage the relations property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Relations() *ItemSitesItemTermStoresItemGroupsItemSetsItemRelationsRequestBuilder {
	return NewItemSitesItemTermStoresItemGroupsItemSetsItemRelationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Terms provides operations to manage the terms property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) Terms() *ItemSitesItemTermStoresItemGroupsItemSetsItemTermsRequestBuilder {
	return NewItemSitesItemTermStoresItemGroupsItemSetsItemTermsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation delete navigation property sets for groups
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation all sets under the group in a term [store].
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property sets in groups
func (m *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, requestConfiguration *ItemSitesItemTermStoresItemGroupsItemSetsSetItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
