package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetQueryParameters get the first level children of a [set] or [term] resource using the children navigation property.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetQueryParameters
}

// ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByTermId provides operations to manage the children property of the microsoft.graph.termStore.set entity.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) ByTermId(termId string) *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if termId != "" {
		urlTplParams["term%2Did"] = termId
	}
	return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderInternal instantiates a new ChildrenRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder {
	m := &ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStores/{store%2Did}/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/children{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder instantiates a new ChildrenRequestBuilder and sets the default values.
func NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) Count() *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenCountRequestBuilder {
	return NewItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the first level children of a [set] or [term] resource using the children navigation property.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/termstore-term-list-children?view=graph-rest-1.0
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.TermCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.TermCollectionResponseable), nil
}

// Post create a new term object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/termstore-term-post?view=graph-rest-1.0
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) Post(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderPostRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation get the first level children of a [set] or [term] resource using the children navigation property.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a new term object.
func (m *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, requestConfiguration *ItemSitesItemTermStoresItemSetsItemParentGroupSetsItemChildrenRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
