package teams

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder provides operations to call the softDelete method.
type ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderInternal instantiates a new SoftDeleteRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder {
	m := &ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/softDelete", pathParameters),
	}
	return m
}

// NewItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder instantiates a new SoftDeleteRequestBuilder and sets the default values.
func NewItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}

// Post delete a single chatMessage or a chat message reply in a channel or a chat.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/chatmessage-softdelete?view=graph-rest-1.0
func (m *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration) error {
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
func (m *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemChannelsItemMessagesItemRepliesItemSoftDeleteRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
