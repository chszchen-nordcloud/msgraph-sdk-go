package policies

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder provides operations to manage the authenticationMethodConfigurations property of the microsoft.graph.authenticationMethodsPolicy entity.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetQueryParameters
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal instantiates a new AuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder {
	m := &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/{authenticationMethodConfiguration%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder instantiates a new AuthenticationMethodConfigurationItemRequestBuilder and sets the default values.
func NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property authenticationMethodConfigurations for policies
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable), nil
}

// Patch update the navigation property authenticationMethodConfigurations in policies
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodConfigurationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable), nil
}

// ToDeleteRequestInformation delete navigation property authenticationMethodConfigurations for policies
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation represents the settings for each authentication method. Automatically expanded on GET /policies/authenticationMethodsPolicy.
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property authenticationMethodConfigurations in policies
func (m *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodConfigurationable, requestConfiguration *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsAuthenticationMethodConfigurationItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
