package teamwork

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetQueryParameters
}

// NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder {
	m := &DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder instantiates a new ConversationMemberItemRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get a collection of team members who have access to the shared channel.
func (m *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable), nil
}

// ToGetRequestInformation a collection of team members who have access to the shared channel.
func (m *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersConversationMemberItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
