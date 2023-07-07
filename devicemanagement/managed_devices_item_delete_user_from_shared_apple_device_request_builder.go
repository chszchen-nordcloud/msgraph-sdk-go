package devicemanagement

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder provides operations to call the deleteUserFromSharedAppleDevice method.
type ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder {
	m := &ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deleteUserFromSharedAppleDevice", pathParameters),
	}
	return m
}

// NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder instantiates a new DeleteUserFromSharedAppleDeviceRequestBuilder and sets the default values.
func NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderInternal(urlParams, requestAdapter)
}

// Post delete user from shared Apple device
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-devices-manageddevice-deleteuserfromsharedappledevice?view=graph-rest-1.0
func (m *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) Post(ctx context.Context, body ManagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyable, requestConfiguration *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation delete user from shared Apple device
func (m *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesItemDeleteUserFromSharedAppleDevicePostRequestBodyable, requestConfiguration *ManagedDevicesItemDeleteUserFromSharedAppleDeviceRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
