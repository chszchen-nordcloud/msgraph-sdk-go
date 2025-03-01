package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder provides operations to call the accept method.
type ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderInternal instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder {
	m := &ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/accept", pathParameters),
	}
	return m
}

// NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder instantiates a new AcceptRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderInternal(urlParams, requestAdapter)
}

// Post accept the specified event in a user calendar.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/event-accept?view=graph-rest-1.0
func (m *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder) Post(ctx context.Context, body ItemCalendarCalendarViewItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToPostRequestInformation accept the specified event in a user calendar.
func (m *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarCalendarViewItemInstancesItemAcceptPostRequestBodyable, requestConfiguration *ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
