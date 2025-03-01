package drives

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder provides operations to call the applyTopItemsFilter method.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderInternal instantiates a new ApplyTopItemsFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder {
	m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyTopItemsFilter", pathParameters),
	}
	return m
}

// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder instantiates a new ApplyTopItemsFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderInternal(urlParams, requestAdapter)
}

// Post invoke action applyTopItemsFilter
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation invoke action applyTopItemsFilter
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyTopItemsFilterRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
