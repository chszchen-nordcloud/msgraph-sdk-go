package identity

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters get userFlowIdentityProviders from identity
type B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters struct {
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

// B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetQueryParameters
}

// ByIdentityProviderBaseId gets an item from the github.com/chszchen-nordcloud/msgraph-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) ByIdentityProviderBaseId(identityProviderBaseId string) *B2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if identityProviderBaseId != "" {
		urlTplParams["identityProviderBase%2Did"] = identityProviderBaseId
	}
	return NewB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder {
	m := &B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userFlowIdentityProviders{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder instantiates a new UserFlowIdentityProvidersRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewB2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) Count() *B2xUserFlowsItemUserFlowIdentityProvidersCountRequestBuilder {
	return NewB2xUserFlowsItemUserFlowIdentityProvidersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get userFlowIdentityProviders from identity
func (m *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityProviderBaseCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityProviderBaseCollectionResponseable), nil
}

// Ref provides operations to manage the collection of identityContainer entities.
func (m *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) Ref() *B2xUserFlowsItemUserFlowIdentityProvidersRefRequestBuilder {
	return NewB2xUserFlowsItemUserFlowIdentityProvidersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation get userFlowIdentityProviders from identity
func (m *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserFlowIdentityProvidersRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
