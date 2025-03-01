package rolemanagement

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
type EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters the schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
type EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetQueryParameters
}

// NewEntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderInternal instantiates a new TargetScheduleRequestBuilder and sets the default values.
func NewEntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder {
	m := &EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/targetSchedule{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewEntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder instantiates a new TargetScheduleRequestBuilder and sets the default values.
func NewEntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
func (m *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable), nil
}

// ToGetRequestInformation the schedule for a role eligibility that is referenced through the targetScheduleId property. Supports $expand.
func (m *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementRoleEligibilityScheduleRequestsItemTargetScheduleRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
