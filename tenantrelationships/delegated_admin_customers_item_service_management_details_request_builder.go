package tenantrelationships

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
type DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetQueryParameters get a list of the delegatedAdminServiceManagementDetail objects and their properties.
type DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetQueryParameters struct {
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

// DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetQueryParameters
}

// DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByDelegatedAdminServiceManagementDetailId provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) ByDelegatedAdminServiceManagementDetailId(delegatedAdminServiceManagementDetailId string) *DelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if delegatedAdminServiceManagementDetailId != "" {
		urlTplParams["delegatedAdminServiceManagementDetail%2Did"] = delegatedAdminServiceManagementDetailId
	}
	return NewDelegatedAdminCustomersItemServiceManagementDetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderInternal instantiates a new ServiceManagementDetailsRequestBuilder and sets the default values.
func NewDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder {
	m := &DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}/serviceManagementDetails{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
	}
	return m
}

// NewDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder instantiates a new ServiceManagementDetailsRequestBuilder and sets the default values.
func NewDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) Count() *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder {
	return NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get a list of the delegatedAdminServiceManagementDetail objects and their properties.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/delegatedadmincustomer-list-servicemanagementdetails?view=graph-rest-1.0
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminServiceManagementDetailCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailCollectionResponseable), nil
}

// Post create new navigation property to serviceManagementDetails for tenantRelationships
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable), nil
}

// ToGetRequestInformation get a list of the delegatedAdminServiceManagementDetail objects and their properties.
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create new navigation property to serviceManagementDetails for tenantRelationships
func (m *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
