package users

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder provides operations to call the provisionEmail method.
type ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder {
	m := &ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/provisionEmail", pathParameters),
	}
	return m
}

// NewItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder instantiates a new ProvisionEmailRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderInternal(urlParams, requestAdapter)
}

// Post provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a **channel** by default. To have Teams provision an email address, you can call **provisionEmail**, or through the Teams user interface, select **Get email address**, which triggers Teams to generate an email address if it has not already provisioned one. To remove the email address of a **channel**, use the removeEmail method.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/channel-provisionemail?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProvisionChannelEmailResultable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateProvisionChannelEmailResultFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ProvisionChannelEmailResultable), nil
}

// ToPostRequestInformation provision an email address for a channel. Microsoft Teams doesn't automatically provision an email address for a **channel** by default. To have Teams provision an email address, you can call **provisionEmail**, or through the Teams user interface, select **Get email address**, which triggers Teams to generate an email address if it has not already provisioned one. To remove the email address of a **channel**, use the removeEmail method.
func (m *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedTeamsItemPrimaryChannelProvisionEmailRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
