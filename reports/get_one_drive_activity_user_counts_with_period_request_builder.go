package reports

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetOneDriveActivityUserCountsWithPeriodRequestBuilder provides operations to call the getOneDriveActivityUserCounts method.
type GetOneDriveActivityUserCountsWithPeriodRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// GetOneDriveActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOneDriveActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal instantiates a new GetOneDriveActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string) *GetOneDriveActivityUserCountsWithPeriodRequestBuilder {
	m := &GetOneDriveActivityUserCountsWithPeriodRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOneDriveActivityUserCounts(period='{period}')", pathParameters),
	}
	if period != nil {
		m.BaseRequestBuilder.PathParameters["period"] = *period
	}
	return m
}

// NewGetOneDriveActivityUserCountsWithPeriodRequestBuilder instantiates a new GetOneDriveActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOneDriveActivityUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *GetOneDriveActivityUserCountsWithPeriodRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Get invoke function getOneDriveActivityUserCounts
func (m *GetOneDriveActivityUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetOneDriveActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.([]byte), nil
}

// ToGetRequestInformation invoke function getOneDriveActivityUserCounts
func (m *GetOneDriveActivityUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetOneDriveActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
