package identitygovernance

import (
	"context"
	ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/chszchen-nordcloud/msgraph-sdk-go/models/identitygovernance"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.task entity.
type LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters the result of processing the task.
type LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters
}

// NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder {
	m := &LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}/tasks/{task%2Did}/taskProcessingResults/{taskProcessingResult%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the result of processing the task.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration) (ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskProcessingResultable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskProcessingResultFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskProcessingResultable), nil
}

// MicrosoftGraphIdentityGovernanceResume provides operations to call the resume method.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) MicrosoftGraphIdentityGovernanceResume() *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilder {
	return NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Subject() *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemSubjectRequestBuilder {
	return NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Task() *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemTaskRequestBuilder {
	return NewLifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsItemTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation the result of processing the task.
func (m *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowTemplatesItemTasksItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
