package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.term entity.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetQueryParameters the [set] in which the term is created.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetQueryParameters
}

// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder {
	m := &ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children/{term%2Did}/children/{term%2Did1}/set{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder instantiates a new SetRequestBuilder and sets the default values.
func NewItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the [set] in which the term is created.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
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

// ToGetRequestInformation the [set] in which the term is created.
func (m *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoreSetsItemParentGroupSetsItemChildrenItemChildrenItemSetRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
