package groups

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemOwnersItemRefRequestBuilder provides operations to manage the collection of group entities.
type ItemOwnersItemRefRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemOwnersItemRefRequestBuilderDeleteQueryParameters remove an owner from a Microsoft 365 group or a security group through the **owners** navigation property. Once owners are assigned to a group, the last owner (a user object) of the group cannot be removed.
type ItemOwnersItemRefRequestBuilderDeleteQueryParameters struct {
	// Delete Uri
	Id *string `uriparametername:"%40id"`
}

// ItemOwnersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOwnersItemRefRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemOwnersItemRefRequestBuilderDeleteQueryParameters
}

// NewItemOwnersItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewItemOwnersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOwnersItemRefRequestBuilder {
	m := &ItemOwnersItemRefRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}/$ref{?%40id*}", pathParameters),
	}
	return m
}

// NewItemOwnersItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewItemOwnersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemOwnersItemRefRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemOwnersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete remove an owner from a Microsoft 365 group or a security group through the **owners** navigation property. Once owners are assigned to a group, the last owner (a user object) of the group cannot be removed.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/group-delete-owners?view=graph-rest-1.0
func (m *ItemOwnersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOwnersItemRefRequestBuilderDeleteRequestConfiguration) error {
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

// ToDeleteRequestInformation remove an owner from a Microsoft 365 group or a security group through the **owners** navigation property. Once owners are assigned to a group, the last owner (a user object) of the group cannot be removed.
func (m *ItemOwnersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOwnersItemRefRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
