package drives

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder provides operations to call the clear method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderInternal instantiates a new ClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder {
	m := &ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/axes/seriesAxis/format/line/clear", pathParameters),
	}
	return m
}

// NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder instantiates a new ClearRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderInternal(urlParams, requestAdapter)
}

// Post clear the line format of a chart element.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/chartlineformat-clear?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation clear the line format of a chart element.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisFormatLineClearRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
