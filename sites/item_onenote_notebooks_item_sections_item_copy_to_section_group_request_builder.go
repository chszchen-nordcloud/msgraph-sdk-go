package sites

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder provides operations to call the copyToSectionGroup method.
type ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderInternal instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder {
	m := &ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sections/{onenoteSection%2Did}/copyToSectionGroup", pathParameters),
	}
	return m
}

// NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder instantiates a new CopyToSectionGroupRequestBuilder and sets the default values.
func NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}

// Post for Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/section-copytosectiongroup?view=graph-rest-1.0
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder) Post(ctx context.Context, body ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteOperationable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenoteOperationFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenoteOperationable), nil
}

// ToPostRequestInformation for Copy operations, you follow an asynchronous calling pattern:  First call the Copy action, and then poll the operation endpoint for the result.
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyable, requestConfiguration *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
