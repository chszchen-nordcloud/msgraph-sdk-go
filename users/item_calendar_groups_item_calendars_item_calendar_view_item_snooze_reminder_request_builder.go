package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder provides operations to call the snoozeReminder method.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderInternal instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder {
	m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/snoozeReminder", pathParameters),
	}
	return m
}

// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder instantiates a new SnoozeReminderRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderInternal(urlParams, requestAdapter)
}

// Post postpone a reminder for an event in a user calendar until a new time.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/event-snoozereminder?view=graph-rest-1.0
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder) Post(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation postpone a reminder for an event in a user calendar until a new time.
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemSnoozeReminderRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
