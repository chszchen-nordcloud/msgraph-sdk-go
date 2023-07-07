package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder provides operations to call the getAllMessages method.
type ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetQueryParameters invoke function getAllMessages
type ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// The payment model for the API
	Model *string
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

// ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetQueryParameters
}

// NewItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderInternal instantiates a new GetAllMessagesRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder {
	m := &ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/getAllMessages(){?model*,%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder instantiates a new GetAllMessagesRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function getAllMessages
func (m *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetRequestConfiguration) (ItemJoinedTeamsItemChannelsGetAllMessagesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemChannelsGetAllMessagesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemJoinedTeamsItemChannelsGetAllMessagesResponseable), nil
}

// ToGetRequestInformation invoke function getAllMessages
func (m *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemChannelsGetAllMessagesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
