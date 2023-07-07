package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
type ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters the to [term] of the relation. The term to which the relationship is defined.
type ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetQueryParameters
}

// NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderInternal instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder {
	m := &ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/children/{term%2Did}/relations/{relation%2Did}/toTerm{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder instantiates a new ToTermRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}

// ToGetRequestInformation the to [term] of the relation. The term to which the relationship is defined.
func (m *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemChildrenItemRelationsItemToTermRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
