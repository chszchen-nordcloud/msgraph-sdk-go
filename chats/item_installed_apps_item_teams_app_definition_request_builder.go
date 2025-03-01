package chats

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
type ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetQueryParameters the details of this version of the app.
type ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetQueryParameters
}

// NewItemInstalledAppsItemTeamsAppDefinitionRequestBuilderInternal instantiates a new TeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledAppsItemTeamsAppDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder {
	m := &ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/installedApps/{teamsAppInstallation%2Did}/teamsAppDefinition{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewItemInstalledAppsItemTeamsAppDefinitionRequestBuilder instantiates a new TeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledAppsItemTeamsAppDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewItemInstalledAppsItemTeamsAppDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the details of this version of the app.
func (m *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppDefinitionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable), nil
}

// ToGetRequestInformation the details of this version of the app.
func (m *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstalledAppsItemTeamsAppDefinitionRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
