package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/termstore"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.term entity.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetQueryParameters the [set] in which the term is created.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetQueryParameters
}

// NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderInternal instantiates a new SetRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder {
	m := &ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/parentGroup/sets/{set%2Did1}/terms/{term%2Did}/children/{term%2Did1}/set{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder instantiates a new SetRequestBuilder and sets the default values.
func NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the [set] in which the term is created.
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetRequestConfiguration) (ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
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
func (m *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreSetsItemParentGroupSetsItemTermsItemChildrenItemSetRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
