package drives

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder provides operations to manage the format property of the microsoft.graph.workbookChartAxisTitle entity.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetQueryParameters represents the formatting of chart axis title. Read-only.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetQueryParameters
}

// ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderInternal instantiates a new FormatRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder {
	m := &ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/axes/seriesAxis/title/format{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder instantiates a new FormatRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete navigation property format for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
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

// Font provides operations to manage the font property of the microsoft.graph.workbookChartAxisTitleFormat entity.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) Font() *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatFontRequestBuilder {
	return NewItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatFontRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get represents the formatting of chart axis title. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable), nil
}

// Patch update the navigation property format in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable), nil
}

// ToDeleteRequestInformation delete navigation property format for drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation represents the formatting of chart axis title. Read-only.
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the navigation property format in drives
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookChartAxisTitleFormatable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesAxisTitleFormatRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
