package groups

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder provides operations to manage the parentSectionGroup property of the microsoft.graph.onenoteSection entity.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetQueryParameters the section group that contains the section.  Read-only.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetQueryParameters
}

// NewItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderInternal instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder {
	m := &ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sections/{onenoteSection%2Did}/parentSectionGroup{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder instantiates a new ParentSectionGroupRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the section group that contains the section.  Read-only.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSectionGroupFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SectionGroupable), nil
}

// ToGetRequestInformation the section group that contains the section.  Read-only.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionsItemParentSectionGroupRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
