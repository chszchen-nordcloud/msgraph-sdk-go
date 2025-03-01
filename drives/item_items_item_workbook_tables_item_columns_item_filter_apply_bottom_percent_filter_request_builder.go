package drives

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder provides operations to call the applyBottomPercentFilter method.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderInternal instantiates a new ApplyBottomPercentFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder {
	m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyBottomPercentFilter", pathParameters),
	}
	return m
}

// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder instantiates a new ApplyBottomPercentFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderInternal(urlParams, requestAdapter)
}

// Post invoke action applyBottomPercentFilter
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToPostRequestInformation invoke action applyBottomPercentFilter
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyBottomPercentFilterRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
