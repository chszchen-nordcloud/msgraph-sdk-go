package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
type ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetQueryParameters retrieve a user's single FIDO2 Security Key Authentication Method object.
type ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetQueryParameters
}

// NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal instantiates a new Fido2AuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder {
	m := &ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/fido2Methods/{fido2AuthenticationMethod%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder instantiates a new Fido2AuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes a user's FIDO2 Security Key Authentication Method object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/fido2authenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderDeleteRequestConfiguration) error {
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

// Get retrieve a user's single FIDO2 Security Key Authentication Method object.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/fido2authenticationmethod-get?view=graph-rest-1.0
func (m *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Fido2AuthenticationMethodable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFido2AuthenticationMethodFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Fido2AuthenticationMethodable), nil
}

// ToDeleteRequestInformation deletes a user's FIDO2 Security Key Authentication Method object.
func (m *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation retrieve a user's single FIDO2 Security Key Authentication Method object.
func (m *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
