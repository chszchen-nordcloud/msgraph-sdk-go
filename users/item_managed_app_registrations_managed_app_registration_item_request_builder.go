package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
type ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters zero or more managed app registrations that belong to the user.
type ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters
}

// NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder {
	m := &ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedAppRegistrations/{managedAppRegistration%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder instantiates a new ManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get zero or more managed app registrations that belong to the user.
func (m *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}

// ToGetRequestInformation zero or more managed app registrations that belong to the user.
func (m *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
