package drives

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemPermissionsItemGrantRequestBuilder provides operations to call the grant method.
type ItemItemsItemPermissionsItemGrantRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemPermissionsItemGrantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemPermissionsItemGrantRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemPermissionsItemGrantRequestBuilderInternal instantiates a new GrantRequestBuilder and sets the default values.
func NewItemItemsItemPermissionsItemGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemPermissionsItemGrantRequestBuilder {
	m := &ItemItemsItemPermissionsItemGrantRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/permissions/{permission%2Did}/grant", pathParameters),
	}
	return m
}

// NewItemItemsItemPermissionsItemGrantRequestBuilder instantiates a new GrantRequestBuilder and sets the default values.
func NewItemItemsItemPermissionsItemGrantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemPermissionsItemGrantRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemPermissionsItemGrantRequestBuilderInternal(urlParams, requestAdapter)
}

// Post grant users access to a link represented by a [permission][].
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/permission-grant?view=graph-rest-1.0
func (m *ItemItemsItemPermissionsItemGrantRequestBuilder) Post(ctx context.Context, body ItemItemsItemPermissionsItemGrantPostRequestBodyable, requestConfiguration *ItemItemsItemPermissionsItemGrantRequestBuilderPostRequestConfiguration) (ItemItemsItemPermissionsItemGrantResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemPermissionsItemGrantResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemItemsItemPermissionsItemGrantResponseable), nil
}

// ToPostRequestInformation grant users access to a link represented by a [permission][].
func (m *ItemItemsItemPermissionsItemGrantRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemPermissionsItemGrantPostRequestBodyable, requestConfiguration *ItemItemsItemPermissionsItemGrantRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
