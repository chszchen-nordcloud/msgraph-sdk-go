package communications

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallsItemParticipantsItemStopHoldMusicRequestBuilder provides operations to call the stopHoldMusic method.
type CallsItemParticipantsItemStopHoldMusicRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CallsItemParticipantsItemStopHoldMusicRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemParticipantsItemStopHoldMusicRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCallsItemParticipantsItemStopHoldMusicRequestBuilderInternal instantiates a new StopHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemStopHoldMusicRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CallsItemParticipantsItemStopHoldMusicRequestBuilder {
	m := &CallsItemParticipantsItemStopHoldMusicRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/participants/{participant%2Did}/stopHoldMusic", pathParameters),
	}
	return m
}

// NewCallsItemParticipantsItemStopHoldMusicRequestBuilder instantiates a new StopHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemStopHoldMusicRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CallsItemParticipantsItemStopHoldMusicRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCallsItemParticipantsItemStopHoldMusicRequestBuilderInternal(urlParams, requestAdapter)
}

// Post reincorporate a participant previously put on hold to the call.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/participant-stopholdmusic?view=graph-rest-1.0
func (m *CallsItemParticipantsItemStopHoldMusicRequestBuilder) Post(ctx context.Context, body CallsItemParticipantsItemStopHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemStopHoldMusicRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StopHoldMusicOperationable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStopHoldMusicOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StopHoldMusicOperationable), nil
}

// ToPostRequestInformation reincorporate a participant previously put on hold to the call.
func (m *CallsItemParticipantsItemStopHoldMusicRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemParticipantsItemStopHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemStopHoldMusicRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
