package solutions

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder provides operations to manage the staffMembers property of the microsoft.graph.bookingBusiness entity.
type BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
type BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters struct {
	// Expand related entities
	Expand []string `uriparametername:"%24expand"`
	// Select properties to be returned
	Select []string `uriparametername:"%24select"`
}

// BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetQueryParameters
}

// BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderInternal instantiates a new BookingStaffMemberBaseItemRequestBuilder and sets the default values.
func NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder {
	m := &BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/bookingBusinesses/{bookingBusiness%2Did}/staffMembers/{bookingStaffMemberBase%2Did}{?%24select,%24expand}", pathParameters),
	}
	return m
}

// NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder instantiates a new BookingStaffMemberBaseItemRequestBuilder and sets the default values.
func NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewBookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a bookingStaffMember in the specified bookingBusiness.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingstaffmember-delete?view=graph-rest-1.0
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
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

// Get get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingstaffmember-get?view=graph-rest-1.0
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingStaffMemberBaseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable), nil
}

// Patch update the properties of a bookingStaffMember in the specified bookingBusiness.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/bookingstaffmember-update?view=graph-rest-1.0
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBookingStaffMemberBaseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable), nil
}

// ToDeleteRequestInformation delete a bookingStaffMember in the specified bookingBusiness.
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}

// ToGetRequestInformation get the properties and relationships of a bookingStaffMember in the specified bookingBusiness.
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation update the properties of a bookingStaffMember in the specified bookingBusiness.
func (m *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BookingStaffMemberBaseable, requestConfiguration *BookingBusinessesItemStaffMembersBookingStaffMemberBaseItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
