package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOutlookSupportedTimeZonesRequestBuilder provides operations to call the supportedTimeZones method.
type ItemOutlookSupportedTimeZonesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemOutlookSupportedTimeZonesRequestBuilderGetQueryParameters invoke function supportedTimeZones
type ItemOutlookSupportedTimeZonesRequestBuilderGetQueryParameters struct {
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

// ItemOutlookSupportedTimeZonesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookSupportedTimeZonesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemOutlookSupportedTimeZonesRequestBuilderGetQueryParameters
}

// NewItemOutlookSupportedTimeZonesRequestBuilderInternal instantiates a new SupportedTimeZonesRequestBuilder and sets the default values.
func NewItemOutlookSupportedTimeZonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOutlookSupportedTimeZonesRequestBuilder {
	m := &ItemOutlookSupportedTimeZonesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/supportedTimeZones(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
	}
	return m
}

// NewItemOutlookSupportedTimeZonesRequestBuilder instantiates a new SupportedTimeZonesRequestBuilder and sets the default values.
func NewItemOutlookSupportedTimeZonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOutlookSupportedTimeZonesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemOutlookSupportedTimeZonesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function supportedTimeZones
func (m *ItemOutlookSupportedTimeZonesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookSupportedTimeZonesRequestBuilderGetRequestConfiguration) (ItemOutlookSupportedTimeZonesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookSupportedTimeZonesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemOutlookSupportedTimeZonesResponseable), nil
}

// ToGetRequestInformation invoke function supportedTimeZones
func (m *ItemOutlookSupportedTimeZonesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookSupportedTimeZonesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
