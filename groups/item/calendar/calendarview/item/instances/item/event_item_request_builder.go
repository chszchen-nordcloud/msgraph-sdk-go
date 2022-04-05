package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i03c20b721830723f54831b527e543a48a7f2b60c82d93c7553533911faae7e29 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/snoozereminder"
    i18c9dcdb1481cc835ace44e085f40089356487a1978d3da7ed6906ee57eee913 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/multivalueextendedproperties"
    i1e1d28a4880b32cd9d34077dc20d66555410bcfcabf9311cec9f18123f1588ab "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/accept"
    i2f0049ed9e28dc8a66f0ab5d8707aff93da3fc33ffb842844bcae2f4d6ca6f55 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/forward"
    i2f3e2c7bce3122084bdbd424284896f3f64c302860793a13f5739a38e581d314 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/tentativelyaccept"
    i63bc1490c8416695403ed8eafd481600f0e6ce32e1eae3b1f3e951e850218e72 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/dismissreminder"
    i78b66b67438fcc6f00ccc5167cb0015b98ef225118bc6af78c2d3319c87dec44 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/extensions"
    i808301d66b9dde5df1b7969261fd7bf2b2a03cfabf17e0668ee80f17c80d607a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/cancel"
    i864ee5009750acc8b6512348fcc02020df2716cd4135ea453d4901fe578b0880 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/decline"
    i8b2f8359a9e5162b4446a90363486ea7de973da5bbc11fbe0f454460453cbbdd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/attachments"
    ia41bcf0865c237f73307ee55fdb00972cce2636f2ce130d2871bc29bdafaa5a8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties"
    id3bf01a4d4097fb20a6fa3544147adcab3a8adc2236bfb8e28cb47ffbbd05cfc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/calendar"
    i3221dc5ca868a1e6cd7c5e5c456cb6c4d9660bb2d5e06821ddc4eaf488434bea "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i922bc548c1702452cd7dd5783841a3fa2c1670c543d1754231bcda50aba3204d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/multivalueextendedproperties/item"
    i9396377bc9fce46fc50efd38fa73033a302e1e421dd2589a42457f8a18851e93 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/extensions/item"
    iac37f278ba9367baf4a79f85e98dfb558391faf98c8982226b2ef47bc4e051b5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item/instances/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i1e1d28a4880b32cd9d34077dc20d66555410bcfcabf9311cec9f18123f1588ab.AcceptRequestBuilder) {
    return i1e1d28a4880b32cd9d34077dc20d66555410bcfcabf9311cec9f18123f1588ab.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i8b2f8359a9e5162b4446a90363486ea7de973da5bbc11fbe0f454460453cbbdd.AttachmentsRequestBuilder) {
    return i8b2f8359a9e5162b4446a90363486ea7de973da5bbc11fbe0f454460453cbbdd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iac37f278ba9367baf4a79f85e98dfb558391faf98c8982226b2ef47bc4e051b5.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return iac37f278ba9367baf4a79f85e98dfb558391faf98c8982226b2ef47bc4e051b5.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id3bf01a4d4097fb20a6fa3544147adcab3a8adc2236bfb8e28cb47ffbbd05cfc.CalendarRequestBuilder) {
    return id3bf01a4d4097fb20a6fa3544147adcab3a8adc2236bfb8e28cb47ffbbd05cfc.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i808301d66b9dde5df1b7969261fd7bf2b2a03cfabf17e0668ee80f17c80d607a.CancelRequestBuilder) {
    return i808301d66b9dde5df1b7969261fd7bf2b2a03cfabf17e0668ee80f17c80d607a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/calendarView/{event_id}/instances/{event_id1}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i864ee5009750acc8b6512348fcc02020df2716cd4135ea453d4901fe578b0880.DeclineRequestBuilder) {
    return i864ee5009750acc8b6512348fcc02020df2716cd4135ea453d4901fe578b0880.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i63bc1490c8416695403ed8eafd481600f0e6ce32e1eae3b1f3e951e850218e72.DismissReminderRequestBuilder) {
    return i63bc1490c8416695403ed8eafd481600f0e6ce32e1eae3b1f3e951e850218e72.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i78b66b67438fcc6f00ccc5167cb0015b98ef225118bc6af78c2d3319c87dec44.ExtensionsRequestBuilder) {
    return i78b66b67438fcc6f00ccc5167cb0015b98ef225118bc6af78c2d3319c87dec44.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9396377bc9fce46fc50efd38fa73033a302e1e421dd2589a42457f8a18851e93.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i9396377bc9fce46fc50efd38fa73033a302e1e421dd2589a42457f8a18851e93.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i2f0049ed9e28dc8a66f0ab5d8707aff93da3fc33ffb842844bcae2f4d6ca6f55.ForwardRequestBuilder) {
    return i2f0049ed9e28dc8a66f0ab5d8707aff93da3fc33ffb842844bcae2f4d6ca6f55.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i18c9dcdb1481cc835ace44e085f40089356487a1978d3da7ed6906ee57eee913.MultiValueExtendedPropertiesRequestBuilder) {
    return i18c9dcdb1481cc835ace44e085f40089356487a1978d3da7ed6906ee57eee913.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i922bc548c1702452cd7dd5783841a3fa2c1670c543d1754231bcda50aba3204d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i922bc548c1702452cd7dd5783841a3fa2c1670c543d1754231bcda50aba3204d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia41bcf0865c237f73307ee55fdb00972cce2636f2ce130d2871bc29bdafaa5a8.SingleValueExtendedPropertiesRequestBuilder) {
    return ia41bcf0865c237f73307ee55fdb00972cce2636f2ce130d2871bc29bdafaa5a8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3221dc5ca868a1e6cd7c5e5c456cb6c4d9660bb2d5e06821ddc4eaf488434bea.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3221dc5ca868a1e6cd7c5e5c456cb6c4d9660bb2d5e06821ddc4eaf488434bea.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i03c20b721830723f54831b527e543a48a7f2b60c82d93c7553533911faae7e29.SnoozeReminderRequestBuilder) {
    return i03c20b721830723f54831b527e543a48a7f2b60c82d93c7553533911faae7e29.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2f3e2c7bce3122084bdbd424284896f3f64c302860793a13f5739a38e581d314.TentativelyAcceptRequestBuilder) {
    return i2f3e2c7bce3122084bdbd424284896f3f64c302860793a13f5739a38e581d314.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
