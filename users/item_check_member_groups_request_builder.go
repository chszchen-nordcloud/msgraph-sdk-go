package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemCheckMemberGroupsRequestBuilder provides operations to call the checkMemberGroups method.
type ItemCheckMemberGroupsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemCheckMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCheckMemberGroupsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemCheckMemberGroupsRequestBuilderInternal instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewItemCheckMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckMemberGroupsRequestBuilder {
	m := &ItemCheckMemberGroupsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/checkMemberGroups", pathParameters),
	}
	return m
}

// NewItemCheckMemberGroupsRequestBuilder instantiates a new CheckMemberGroupsRequestBuilder and sets the default values.
func NewItemCheckMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemCheckMemberGroupsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemCheckMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Azure AD. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-checkmembergroups?view=graph-rest-1.0
func (m *ItemCheckMemberGroupsRequestBuilder) Post(ctx context.Context, body ItemCheckMemberGroupsPostRequestBodyable, requestConfiguration *ItemCheckMemberGroupsRequestBuilderPostRequestConfiguration) (ItemCheckMemberGroupsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckMemberGroupsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ItemCheckMemberGroupsResponseable), nil
}

// ToPostRequestInformation check for membership in a specified list of group IDs, and return from that list those groups (identified by IDs) of which the specified user, group, service principal, organizational contact, device, or directory object is a member. This function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups provisioned in Azure AD. Because Microsoft 365 groups cannot contain other groups, membership in a Microsoft 365 group is always direct.
func (m *ItemCheckMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCheckMemberGroupsPostRequestBodyable, requestConfiguration *ItemCheckMemberGroupsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
