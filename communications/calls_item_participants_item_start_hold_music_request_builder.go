package communications

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CallsItemParticipantsItemStartHoldMusicRequestBuilder provides operations to call the startHoldMusic method.
type CallsItemParticipantsItemStartHoldMusicRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CallsItemParticipantsItemStartHoldMusicRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemParticipantsItemStartHoldMusicRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCallsItemParticipantsItemStartHoldMusicRequestBuilderInternal instantiates a new StartHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemStartHoldMusicRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CallsItemParticipantsItemStartHoldMusicRequestBuilder {
	m := &CallsItemParticipantsItemStartHoldMusicRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/calls/{call%2Did}/participants/{participant%2Did}/startHoldMusic", pathParameters),
	}
	return m
}

// NewCallsItemParticipantsItemStartHoldMusicRequestBuilder instantiates a new StartHoldMusicRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemStartHoldMusicRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CallsItemParticipantsItemStartHoldMusicRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCallsItemParticipantsItemStartHoldMusicRequestBuilderInternal(urlParams, requestAdapter)
}

// Post put a participant on hold and play music in the background.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/participant-startholdmusic?view=graph-rest-1.0
func (m *CallsItemParticipantsItemStartHoldMusicRequestBuilder) Post(ctx context.Context, body CallsItemParticipantsItemStartHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemStartHoldMusicRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StartHoldMusicOperationable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStartHoldMusicOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StartHoldMusicOperationable), nil
}

// ToPostRequestInformation put a participant on hold and play music in the background.
func (m *CallsItemParticipantsItemStartHoldMusicRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemParticipantsItemStartHoldMusicPostRequestBodyable, requestConfiguration *CallsItemParticipantsItemStartHoldMusicRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
