package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAgreementAcceptancesRequestBuilder provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
type ItemAgreementAcceptancesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemAgreementAcceptancesRequestBuilderGetQueryParameters retrieve the signed-in user's agreementAcceptance objects.
type ItemAgreementAcceptancesRequestBuilderGetQueryParameters struct {
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

// ItemAgreementAcceptancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAgreementAcceptancesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemAgreementAcceptancesRequestBuilderGetQueryParameters
}

// ByAgreementAcceptanceId provides operations to manage the agreementAcceptances property of the microsoft.graph.user entity.
func (m *ItemAgreementAcceptancesRequestBuilder) ByAgreementAcceptanceId(agreementAcceptanceId string) *ItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if agreementAcceptanceId != "" {
		urlTplParams["agreementAcceptance%2Did"] = agreementAcceptanceId
	}
	return NewItemAgreementAcceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemAgreementAcceptancesRequestBuilderInternal instantiates a new AgreementAcceptancesRequestBuilder and sets the default values.
func NewItemAgreementAcceptancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAgreementAcceptancesRequestBuilder {
	m := &ItemAgreementAcceptancesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/agreementAcceptances{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemAgreementAcceptancesRequestBuilder instantiates a new AgreementAcceptancesRequestBuilder and sets the default values.
func NewItemAgreementAcceptancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAgreementAcceptancesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemAgreementAcceptancesRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemAgreementAcceptancesRequestBuilder) Count() *ItemAgreementAcceptancesCountRequestBuilder {
	return NewItemAgreementAcceptancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get retrieve the signed-in user's agreementAcceptance objects.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/user-list-agreementacceptances?view=graph-rest-1.0
func (m *ItemAgreementAcceptancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAgreementAcceptancesRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementAcceptanceCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceCollectionResponseable), nil
}

// ToGetRequestInformation retrieve the signed-in user's agreementAcceptance objects.
func (m *ItemAgreementAcceptancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAgreementAcceptancesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
