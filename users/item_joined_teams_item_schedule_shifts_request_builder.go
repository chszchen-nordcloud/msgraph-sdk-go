package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemScheduleShiftsRequestBuilder provides operations to manage the shifts property of the microsoft.graph.schedule entity.
type ItemJoinedTeamsItemScheduleShiftsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetQueryParameters get the list of shift instances in a schedule.
type ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Order items by property values
	Orderby []string `uriparametername:"%24orderby"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetQueryParameters
}

// ItemJoinedTeamsItemScheduleShiftsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemScheduleShiftsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByShiftId provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) ByShiftId(shiftId string) *ItemJoinedTeamsItemScheduleShiftsShiftItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if shiftId != "" {
		urlTplParams["shift%2Did"] = shiftId
	}
	return NewItemJoinedTeamsItemScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewItemJoinedTeamsItemScheduleShiftsRequestBuilderInternal instantiates a new ShiftsRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleShiftsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemScheduleShiftsRequestBuilder {
	m := &ItemJoinedTeamsItemScheduleShiftsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/schedule/shifts{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select}", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemScheduleShiftsRequestBuilder instantiates a new ShiftsRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemScheduleShiftsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemScheduleShiftsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemScheduleShiftsRequestBuilderInternal(urlParams, requestAdapter)
}

// Count provides operations to count the resources in the collection.
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) Count() *ItemJoinedTeamsItemScheduleShiftsCountRequestBuilder {
	return NewItemJoinedTeamsItemScheduleShiftsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the list of shift instances in a schedule.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-list-shifts?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ShiftCollectionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateShiftCollectionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ShiftCollectionResponseable), nil
}

// Post create a new shift instance in a schedule. The duration of a shift cannot be less than 1 minute or longer than 24 hours.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/schedule-post-shifts?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Shiftable, requestConfiguration *ItemJoinedTeamsItemScheduleShiftsRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Shiftable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateShiftFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Shiftable), nil
}

// ToGetRequestInformation get the list of shift instances in a schedule.
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemScheduleShiftsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a new shift instance in a schedule. The duration of a shift cannot be less than 1 minute or longer than 24 hours.
func (m *ItemJoinedTeamsItemScheduleShiftsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Shiftable, requestConfiguration *ItemJoinedTeamsItemScheduleShiftsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
