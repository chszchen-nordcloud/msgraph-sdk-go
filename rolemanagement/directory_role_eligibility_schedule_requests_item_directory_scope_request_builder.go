package rolemanagement

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
type DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetQueryParameters the directory object that is the scope of the role eligibility. Read-only. Supports $expand.
type DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetQueryParameters
}

// NewDirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderInternal instantiates a new DirectoryScopeRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder {
	m := &DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/directoryScope{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewDirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder instantiates a new DirectoryScopeRequestBuilder and sets the default values.
func NewDirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewDirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the directory object that is the scope of the role eligibility. Read-only. Supports $expand.
func (m *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
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

// ToGetRequestInformation the directory object that is the scope of the role eligibility. Read-only. Supports $expand.
func (m *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleEligibilityScheduleRequestsItemDirectoryScopeRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
