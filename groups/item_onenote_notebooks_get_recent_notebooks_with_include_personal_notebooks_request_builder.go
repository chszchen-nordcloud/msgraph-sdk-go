package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder provides operations to call the getRecentNotebooks method.
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters invoke function getRecentNotebooks
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters struct {
	// Include count of items
	Count *bool `uriparametername:"%24count"`
	// Filter items by property values
	Filter *string `uriparametername:"%24filter"`
	// Search items by search phrases
	Search *string `uriparametername:"%24search"`
	// Skip the first n items
	Skip *int32 `uriparametername:"%24skip"`
	// Show only the first n items
	Top *int32 `uriparametername:"%24top"`
}

// ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetQueryParameters
}

// NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal instantiates a new GetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, includePersonalNotebooks *bool) *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder {
	m := &ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/onenote/notebooks/getRecentNotebooks(includePersonalNotebooks={includePersonalNotebooks}){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
	}
	if includePersonalNotebooks != nil {
		m.BaseRequestBuilder.PathParameters["includePersonalNotebooks"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatBool(*includePersonalNotebooks)
	}
	return m
}

// NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder instantiates a new GetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Get invoke function getRecentNotebooks
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration) (ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable), nil
}

// ToGetRequestInformation invoke function getRecentNotebooks
func (m *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
