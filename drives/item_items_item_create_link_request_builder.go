package drives

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemsItemCreateLinkRequestBuilder provides operations to call the createLink method.
type ItemItemsItemCreateLinkRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemItemsItemCreateLinkRequestBuilderInternal instantiates a new CreateLinkRequestBuilder and sets the default values.
func NewItemItemsItemCreateLinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemCreateLinkRequestBuilder {
	m := &ItemItemsItemCreateLinkRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/createLink", pathParameters),
	}
	return m
}

// NewItemItemsItemCreateLinkRequestBuilder instantiates a new CreateLinkRequestBuilder and sets the default values.
func NewItemItemsItemCreateLinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemItemsItemCreateLinkRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemItemsItemCreateLinkRequestBuilderInternal(urlParams, requestAdapter)
}

// Post you can use **createLink** action to share a DriveItem via a sharing link. The **createLink** action will create a new sharing link if the specified link type doesn't already exist for the calling application.If a sharing link of the specified type already exists for the app, the existing sharing link will be returned. DriveItem resources inherit sharing permissions from their ancestors.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-createlink?view=graph-rest-1.0
func (m *ItemItemsItemCreateLinkRequestBuilder) Post(ctx context.Context, body ItemItemsItemCreateLinkPostRequestBodyable, requestConfiguration *ItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable), nil
}

// ToPostRequestInformation you can use **createLink** action to share a DriveItem via a sharing link. The **createLink** action will create a new sharing link if the specified link type doesn't already exist for the calling application.If a sharing link of the specified type already exists for the app, the existing sharing link will be returned. DriveItem resources inherit sharing permissions from their ancestors.
func (m *ItemItemsItemCreateLinkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemCreateLinkPostRequestBodyable, requestConfiguration *ItemItemsItemCreateLinkRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
