package tenantrelationships

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder provides operations to count the resources in the collection.
type DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetQueryParameters get the number of the resource
type DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
}

// DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetQueryParameters
}

// NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder {
	m := &DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}/serviceManagementDetails/$count{?%24search,%24filter}", pathParameters),
	}
	return m
}

// NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the number of the resource
func (m *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetRequestConfiguration) (*int32, error) {
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
func (m *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedAdminCustomersItemServiceManagementDetailsCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
