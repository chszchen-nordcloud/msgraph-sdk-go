package deviceappmanagement

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
type MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters the list of categories for this app.
type MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters
}

// NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderInternal instantiates a new MobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder {
	m := &MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/categories/{mobileAppCategory%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder instantiates a new MobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewMobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the list of categories for this app.
func (m *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppCategoryable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMobileAppCategoryFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MobileAppCategoryable), nil
}

// ToGetRequestInformation the list of categories for this app.
func (m *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
