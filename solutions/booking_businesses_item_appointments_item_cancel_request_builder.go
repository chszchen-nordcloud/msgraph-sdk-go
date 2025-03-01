package solutions

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BookingBusinessesItemAppointmentsItemCancelRequestBuilder provides operations to call the cancel method.
type BookingBusinessesItemAppointmentsItemCancelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal instantiates a new CancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemAppointmentsItemCancelRequestBuilder {
	m := &BookingBusinessesItemAppointmentsItemCancelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/appointments/{bookingAppointment%2Did}/cancel", pathParameters),
	}
	return m
}

// NewBookingBusinessesItemAppointmentsItemCancelRequestBuilder instantiates a new CancelRequestBuilder and sets the default values.
func NewBookingBusinessesItemAppointmentsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemAppointmentsItemCancelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewBookingBusinessesItemAppointmentsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}

// Post cancel the specified bookingAppointment in the specified bookingBusiness and send a message to the involved customer and staff members.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingappointment-cancel?view=graph-rest-1.0
func (m *BookingBusinessesItemAppointmentsItemCancelRequestBuilder) Post(ctx context.Context, body BookingBusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation cancel the specified bookingAppointment in the specified bookingBusiness and send a message to the involved customer and staff members.
func (m *BookingBusinessesItemAppointmentsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, body BookingBusinessesItemAppointmentsItemCancelPostRequestBodyable, requestConfiguration *BookingBusinessesItemAppointmentsItemCancelRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
