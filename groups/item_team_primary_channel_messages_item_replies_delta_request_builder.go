package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder provides operations to call the delta method.
type ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters invoke function delta
type ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters struct {
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

// ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetQueryParameters
}

// NewItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder {
	m := &ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/messages/{chatMessage%2Did}/replies/delta(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
	}
	return m
}

// NewItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder instantiates a new DeltaRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderInternal(urlParams, requestAdapter)
}

// Get invoke function delta
func (m *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration) (ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemTeamPrimaryChannelMessagesItemRepliesDeltaResponseable), nil
}

// ToGetRequestInformation invoke function delta
func (m *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelMessagesItemRepliesDeltaRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
