package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder provides operations to call the softDelete method.
type ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderInternal instantiates a new SoftDeleteRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder {
	m := &ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/softDelete", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder instantiates a new SoftDeleteRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}

// Post delete a single chatMessage or a chat message reply in a channel or a chat.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/chatmessage-softdelete?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation delete a single chatMessage or a chat message reply in a channel or a chat.
func (m *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
