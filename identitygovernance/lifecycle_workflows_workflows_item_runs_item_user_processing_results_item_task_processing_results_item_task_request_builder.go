package identitygovernance

import (
	"context"
	ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/identitygovernance"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
type LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters the related workflow task
type LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetQueryParameters
}

// NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderInternal instantiates a new TaskRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder {
	m := &LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/task{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder instantiates a new TaskRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the related workflow task
func (m *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration) (ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable), nil
}

// ToGetRequestInformation the related workflow task
func (m *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsItemTaskProcessingResultsItemTaskRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
