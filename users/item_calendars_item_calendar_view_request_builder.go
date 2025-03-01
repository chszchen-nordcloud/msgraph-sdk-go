package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarsItemCalendarViewRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type ItemCalendarsItemCalendarViewRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type ItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
	EndDateTime *string
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
	StartDateTime *string
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemCalendarsItemCalendarViewRequestBuilderGetQueryParameters
}

// ByEventId provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarsItemCalendarViewRequestBuilder) ByEventId(eventId string) *ItemCalendarsItemCalendarViewEventItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if eventId != "" {
		urlTplParams["event%2Did"] = eventId
	}
	return NewItemCalendarsItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemCalendarsItemCalendarViewRequestBuilderInternal instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarsItemCalendarViewRequestBuilder {
	m := &ItemCalendarsItemCalendarViewRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView{?startDateTime*,endDateTime*,%24top,%24skip,%24filter,%24count,%24orderby,%24select}", pathParameters),
	}
	return m
}

// NewItemCalendarsItemCalendarViewRequestBuilder instantiates a new CalendarViewRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarsItemCalendarViewRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarsItemCalendarViewRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemCalendarsItemCalendarViewRequestBuilder) Count() *ItemCalendarsItemCalendarViewCountRequestBuilder {
	return NewItemCalendarsItemCalendarViewCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Delta provides operations to call the delta method.
func (m *ItemCalendarsItemCalendarViewRequestBuilder) Delta() *ItemCalendarsItemCalendarViewDeltaRequestBuilder {
	return NewItemCalendarsItemCalendarViewDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get the calendar view for the calendar. Navigation property. Read-only.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/calendar-list-calendarview?view=graph-rest-1.0
func (m *ItemCalendarsItemCalendarViewRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable), nil
}

// ToGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *ItemCalendarsItemCalendarViewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
