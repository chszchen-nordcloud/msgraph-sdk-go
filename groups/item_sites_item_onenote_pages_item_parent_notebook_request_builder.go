package groups

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
type ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetQueryParameters the notebook that contains the page.  Read-only.
type ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetQueryParameters
}

// NewItemSitesItemOnenotePagesItemParentNotebookRequestBuilderInternal instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewItemSitesItemOnenotePagesItemParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder {
	m := &ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/pages/{onenotePage%2Did}/parentNotebook{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemSitesItemOnenotePagesItemParentNotebookRequestBuilder instantiates a new ParentNotebookRequestBuilder and sets the default values.
func NewItemSitesItemOnenotePagesItemParentNotebookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemSitesItemOnenotePagesItemParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the notebook that contains the page.  Read-only.
func (m *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNotebookFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Notebookable), nil
}

// ToGetRequestInformation the notebook that contains the page.  Read-only.
func (m *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenotePagesItemParentNotebookRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
