package directoryroletemplates

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRestoreRequestBuilder provides operations to call the restore method.
type ItemRestoreRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRestoreRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemRestoreRequestBuilderInternal instantiates a new RestoreRequestBuilder and sets the default values.
func NewItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemRestoreRequestBuilder {
	m := &ItemRestoreRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directoryRoleTemplates/{directoryRoleTemplate%2Did}/restore", pathParameters),
	}
	return m
}

// NewItemRestoreRequestBuilder instantiates a new RestoreRequestBuilder and sets the default values.
func NewItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemRestoreRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}

// Post restore a recently deleted application, group, servicePrincipal, administrative unit, or user object from deleted items. If an item was accidentally deleted, you can fully restore the item. This is not applicable to security groups, which are deleted permanently. A recently deleted item will remain available for up to 30 days. After 30 days, the item is permanently deleted.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/directory-deleteditems-restore?view=graph-rest-1.0
func (m *ItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemRestoreRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}

// ToPostRequestInformation restore a recently deleted application, group, servicePrincipal, administrative unit, or user object from deleted items. If an item was accidentally deleted, you can fully restore the item. This is not applicable to security groups, which are deleted permanently. A recently deleted item will remain available for up to 30 days. After 30 days, the item is permanently deleted.
func (m *ItemRestoreRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemRestoreRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
