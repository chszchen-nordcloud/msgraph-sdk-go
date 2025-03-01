package identitygovernance

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder {
	m := &EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/getAvailableExtensionProperties", pathParameters),
	}
	return m
}

// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}

// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration) (EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesResponseable), nil
}

// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:+ user+ group+ administrativeUnit+ application+ device+ organization
func (m *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *EntitlementManagementConnectedOrganizationsItemExternalSponsorsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
