package users

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemCloneRequestBuilder provides operations to call the clone method.
type ItemJoinedTeamsItemCloneRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemCloneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemCloneRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemJoinedTeamsItemCloneRequestBuilderInternal instantiates a new CloneRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemCloneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemCloneRequestBuilder {
	m := &ItemJoinedTeamsItemCloneRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/clone", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemCloneRequestBuilder instantiates a new CloneRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemCloneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemCloneRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemCloneRequestBuilderInternal(urlParams, requestAdapter)
}

// Post create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they are put into an unconfigured state -- they are displayed on the tab bar in Microsoft Teams, and the first time you open them, you'll go through the configuration screen. (If the person opening the tab does not have permission to configure apps, they will see a message explaining that the tab hasn't been configured.) Cloning is a long-running operation.After the POST clone returns, you need to GET the operation to see if it's 'running' or 'succeeded' or 'failed'. You should continue to GET until the status is not 'running'. The recommended delay between GETs is 5 seconds.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/team-clone?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemCloneRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemClonePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemCloneRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
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

// ToPostRequestInformation create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they are put into an unconfigured state -- they are displayed on the tab bar in Microsoft Teams, and the first time you open them, you'll go through the configuration screen. (If the person opening the tab does not have permission to configure apps, they will see a message explaining that the tab hasn't been configured.) Cloning is a long-running operation.After the POST clone returns, you need to GET the operation to see if it's 'running' or 'succeeded' or 'failed'. You should continue to GET until the status is not 'running'. The recommended delay between GETs is 5 seconds.
func (m *ItemJoinedTeamsItemCloneRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemClonePostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemCloneRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
