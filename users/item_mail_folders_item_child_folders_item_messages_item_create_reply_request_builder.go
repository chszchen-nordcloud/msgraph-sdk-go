package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder provides operations to call the createReply method.
type ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal instantiates a new CreateReplyRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder {
	m := &ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/createReply", pathParameters),
	}
	return m
}

// NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder instantiates a new CreateReplyRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderInternal(urlParams, requestAdapter)
}

// Post create a draft to reply to the sender of a message in either JSON or MIME format. When using JSON format:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- If **replyTo** is specified in the original message, per Internet Message Format (RFC 2822), you should send the reply to the recipients in **replyTo**, and not the recipients in **from**.- You can update the draft later to add reply content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, reply to a message in a single operation.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/message-createreply?view=graph-rest-1.0
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyPostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable), nil
}

// ToPostRequestInformation create a draft to reply to the sender of a message in either JSON or MIME format. When using JSON format:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- If **replyTo** is specified in the original message, per Internet Message Format (RFC 2822), you should send the reply to the recipients in **replyTo**, and not the recipients in **from**.- You can update the draft later to add reply content to the **body** or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. Send the draft message in a subsequent operation. Alternatively, reply to a message in a single operation.
func (m *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyPostRequestBodyable, requestConfiguration *ItemMailFoldersItemChildFoldersItemMessagesItemCreateReplyRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
