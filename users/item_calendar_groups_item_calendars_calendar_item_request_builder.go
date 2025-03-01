package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.calendarGroup entity.
type ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters the calendars in the calendar group. Navigation property. Read-only. Nullable.
type ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters struct {
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetQueryParameters
}

// ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string) *ItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserRequestBuilder {
	return NewItemCalendarGroupsItemCalendarsItemAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, user)
}

// CalendarPermissions provides operations to manage the calendarPermissions property of the microsoft.graph.calendar entity.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarPermissions() *ItemCalendarGroupsItemCalendarsItemCalendarPermissionsRequestBuilder {
	return NewItemCalendarGroupsItemCalendarsItemCalendarPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// CalendarView provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) CalendarView() *ItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilder {
	return NewItemCalendarGroupsItemCalendarsItemCalendarViewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder {
	m := &ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}{?%24select}", pathParameters),
	}
	return m
}

// NewItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property calendars for users
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration) error {
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

// Events provides operations to manage the events property of the microsoft.graph.calendar entity.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) Events() *ItemCalendarGroupsItemCalendarsItemEventsRequestBuilder {
	return NewItemCalendarGroupsItemCalendarsItemEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}

// GetSchedule provides operations to call the getSchedule method.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) GetSchedule() *ItemCalendarGroupsItemCalendarsItemGetScheduleRequestBuilder {
	return NewItemCalendarGroupsItemCalendarsItemGetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch update the navigation property calendars in users
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}

// ToDeleteRequestInformation delete navigation property calendars for users
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToGetRequestInformation the calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property calendars in users
func (m *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, requestConfiguration *ItemCalendarGroupsItemCalendarsCalendarItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
