package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder provides operations to count the resources in the collection.
type ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetQueryParameters struct {
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
}

// ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetQueryParameters
}

// NewItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder {
	m := &ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/extensions/$count{?%24filter}", pathParameters),
	}
	return m
}

// NewItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the number of the resource
func (m *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetRequestConfiguration) (*int32, error) {
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
func (m *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemInstancesItemExtensionsCountRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
