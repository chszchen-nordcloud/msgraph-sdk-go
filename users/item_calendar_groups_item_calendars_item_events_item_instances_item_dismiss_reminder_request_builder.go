package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderInternal instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder {
	m := &ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/dismissReminder", pathParameters),
	}
	return m
}

// NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}

// Post dismiss a reminder that has been triggered for an event in a user calendar.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/event-dismissreminder?view=graph-rest-1.0
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
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

// ToPostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
func (m *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemEventsItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
