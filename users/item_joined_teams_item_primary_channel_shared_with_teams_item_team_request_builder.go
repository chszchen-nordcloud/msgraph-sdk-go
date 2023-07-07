package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder provides operations to manage the team property of the microsoft.graph.teamInfo entity.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetQueryParameters get team from users
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetQueryParameters
}

// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder {
	m := &ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/team{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get team from users
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}

// ToGetRequestInformation get team from users
func (m *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelSharedWithTeamsItemTeamRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
