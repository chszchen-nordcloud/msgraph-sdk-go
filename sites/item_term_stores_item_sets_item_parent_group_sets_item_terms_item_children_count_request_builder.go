package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder provides operations to count the resources in the collection.
type ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
}

// ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetQueryParameters
}

// NewItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder {
	m := &ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/children/$count{?%24search,%24filter}", pathParameters),
	}
	return m
}

// NewItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the number of the resource
func (m *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration) (*int32, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*int32), nil
}

// ToGetRequestInformation get the number of the resource
func (m *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoresItemSetsItemParentGroupSetsItemTermsItemChildrenCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	requestInfo.Headers.Add("Accept", "text/plain")
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
