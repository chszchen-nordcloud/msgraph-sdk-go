package devicemanagement

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetEffectivePermissionsWithScopeRequestBuilder provides operations to call the getEffectivePermissions method.
type GetEffectivePermissionsWithScopeRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// GetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters retrieves the effective permissions of the currently authenticated user
type GetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// GetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *GetEffectivePermissionsWithScopeRequestBuilderGetQueryParameters
}

// NewGetEffectivePermissionsWithScopeRequestBuilderInternal instantiates a new GetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGetEffectivePermissionsWithScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, scope *string) *GetEffectivePermissionsWithScopeRequestBuilder {
	m := &GetEffectivePermissionsWithScopeRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/getEffectivePermissions(scope='{scope}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
	}
	if scope != nil {
		m.BaseRequestBuilder.PathParameters["scope"] = *scope
	}
	return m
}

// NewGetEffectivePermissionsWithScopeRequestBuilder instantiates a new GetEffectivePermissionsWithScopeRequestBuilder and sets the default values.
func NewGetEffectivePermissionsWithScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *GetEffectivePermissionsWithScopeRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewGetEffectivePermissionsWithScopeRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Get retrieves the effective permissions of the currently authenticated user
func (m *GetEffectivePermissionsWithScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *GetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration) (GetEffectivePermissionsWithScopeResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetEffectivePermissionsWithScopeResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(GetEffectivePermissionsWithScopeResponseable), nil
}

// ToGetRequestInformation retrieves the effective permissions of the currently authenticated user
func (m *GetEffectivePermissionsWithScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetEffectivePermissionsWithScopeRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
