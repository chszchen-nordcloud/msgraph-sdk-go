package sites

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder provides operations to call the restoreVersion method.
type ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderInternal instantiates a new RestoreVersionRequestBuilder and sets the default values.
func NewItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder {
	m := &ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/versions/{listItemVersion%2Did}/restoreVersion", pathParameters),
	}
	return m
}

// NewItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder instantiates a new RestoreVersionRequestBuilder and sets the default values.
func NewItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderInternal(urlParams, requestAdapter)
}

// Post restore a previous version of a ListItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the item.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/listitemversion-restore?view=graph-rest-1.0
func (m *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation restore a previous version of a ListItem to be the current version. This will create a new version with the contents of the previous version, but preserves all existing versions of the item.
func (m *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsItemVersionsItemRestoreVersionRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
