package security

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder provides operations to call the asHierarchy method.
type CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters invoke function asHierarchy
type CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
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

// CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetQueryParameters
}

// NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal instantiates a new MicrosoftGraphSecurityAsHierarchyRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder {
	m := &CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/tags/microsoft.graph.security.asHierarchy(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
	}
	return m
}

// NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder instantiates a new MicrosoftGraphSecurityAsHierarchyRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function asHierarchy
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration) (CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyAsHierarchyResponseable), nil
}

// ToGetRequestInformation invoke function asHierarchy
func (m *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemTagsMicrosoftGraphSecurityAsHierarchyRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
