package solutions

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BookingBusinessesItemUnpublishRequestBuilder provides operations to call the unpublish method.
type BookingBusinessesItemUnpublishRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewBookingBusinessesItemUnpublishRequestBuilderInternal instantiates a new UnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemUnpublishRequestBuilder {
	m := &BookingBusinessesItemUnpublishRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/unpublish", pathParameters),
	}
	return m
}

// NewBookingBusinessesItemUnpublishRequestBuilder instantiates a new UnpublishRequestBuilder and sets the default values.
func NewBookingBusinessesItemUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemUnpublishRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewBookingBusinessesItemUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}

// Post make the scheduling page of this business not available to external customers. Set the **isPublished** property to `false`, and the **publicUrl** property to `null`.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingbusiness-unpublish?view=graph-rest-1.0
func (m *BookingBusinessesItemUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
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

// ToPostRequestInformation make the scheduling page of this business not available to external customers. Set the **isPublished** property to `false`, and the **publicUrl** property to `null`.
func (m *BookingBusinessesItemUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemUnpublishRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
