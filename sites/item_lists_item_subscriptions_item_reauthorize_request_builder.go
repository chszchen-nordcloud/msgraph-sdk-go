package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemListsItemSubscriptionsItemReauthorizeRequestBuilder provides operations to call the reauthorize method.
type ItemListsItemSubscriptionsItemReauthorizeRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemSubscriptionsItemReauthorizeRequestBuilder {
	m := &ItemListsItemSubscriptionsItemReauthorizeRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/subscriptions/{subscription%2Did}/reauthorize", pathParameters),
	}
	return m
}

// NewItemListsItemSubscriptionsItemReauthorizeRequestBuilder instantiates a new ReauthorizeRequestBuilder and sets the default values.
func NewItemListsItemSubscriptionsItemReauthorizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemSubscriptionsItemReauthorizeRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemListsItemSubscriptionsItemReauthorizeRequestBuilderInternal(urlParams, requestAdapter)
}

// Post reauthorize a subscription when you receive a **reauthorizationRequired** challenge.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/subscription-reauthorize?view=graph-rest-1.0
func (m *ItemListsItemSubscriptionsItemReauthorizeRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation reauthorize a subscription when you receive a **reauthorizationRequired** challenge.
func (m *ItemListsItemSubscriptionsItemReauthorizeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemSubscriptionsItemReauthorizeRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
