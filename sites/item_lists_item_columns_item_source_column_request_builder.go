package sites

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemListsItemColumnsItemSourceColumnRequestBuilder provides operations to manage the sourceColumn property of the microsoft.graph.columnDefinition entity.
type ItemListsItemColumnsItemSourceColumnRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemListsItemColumnsItemSourceColumnRequestBuilderGetQueryParameters the source column for the content type column.
type ItemListsItemColumnsItemSourceColumnRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemListsItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemListsItemColumnsItemSourceColumnRequestBuilderGetQueryParameters
}

// NewItemListsItemColumnsItemSourceColumnRequestBuilderInternal instantiates a new SourceColumnRequestBuilder and sets the default values.
func NewItemListsItemColumnsItemSourceColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemColumnsItemSourceColumnRequestBuilder {
	m := &ItemListsItemColumnsItemSourceColumnRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/columns/{columnDefinition%2Did}/sourceColumn{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemListsItemColumnsItemSourceColumnRequestBuilder instantiates a new SourceColumnRequestBuilder and sets the default values.
func NewItemListsItemColumnsItemSourceColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemListsItemColumnsItemSourceColumnRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemListsItemColumnsItemSourceColumnRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the source column for the content type column.
func (m *ItemListsItemColumnsItemSourceColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnDefinitionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnDefinitionable), nil
}

// ToGetRequestInformation the source column for the content type column.
func (m *ItemListsItemColumnsItemSourceColumnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemColumnsItemSourceColumnRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
