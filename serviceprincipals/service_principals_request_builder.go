package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i280dc50e69b7ff09cee61df281de2a8da0149cc215fb3c8f75da14e6dcbdc96c "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/count"
    i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/delta"
    i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/validateproperties"
    i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/getavailableextensionproperties"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/getbyids"
)

// ServicePrincipalsRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ServicePrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ServicePrincipalsRequestBuilderGetOptions options for Get
type ServicePrincipalsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ServicePrincipalsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ServicePrincipalsRequestBuilderGetQueryParameters get entities from servicePrincipals
type ServicePrincipalsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ServicePrincipalsRequestBuilderPostOptions options for Post
type ServicePrincipalsRequestBuilderPostOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewServicePrincipalsRequestBuilderInternal instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
func NewServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    m := &ServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsRequestBuilder instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
func NewServicePrincipalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ServicePrincipalsRequestBuilder) Count()(*i280dc50e69b7ff09cee61df281de2a8da0149cc215fb3c8f75da14e6dcbdc96c.CountRequestBuilder) {
    return i280dc50e69b7ff09cee61df281de2a8da0149cc215fb3c8f75da14e6dcbdc96c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from servicePrincipals
func (m *ServicePrincipalsRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to servicePrincipals
func (m *ServicePrincipalsRequestBuilder) CreatePostRequestInformation(options *ServicePrincipalsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Delta provides operations to call the delta method.
func (m *ServicePrincipalsRequestBuilder) Delta()(*i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080.DeltaRequestBuilder) {
    return i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from servicePrincipals
func (m *ServicePrincipalsRequestBuilder) Get(options *ServicePrincipalsRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalCollectionResponseable), nil
}
// GetAvailableExtensionProperties the getAvailableExtensionProperties property
func (m *ServicePrincipalsRequestBuilder) GetAvailableExtensionProperties()(*i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc.GetAvailableExtensionPropertiesRequestBuilder) {
    return i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetByIds the getByIds property
func (m *ServicePrincipalsRequestBuilder) GetByIds()(*idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5.GetByIdsRequestBuilder) {
    return idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to servicePrincipals
func (m *ServicePrincipalsRequestBuilder) Post(options *ServicePrincipalsRequestBuilderPostOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServicePrincipalFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServicePrincipalable), nil
}
// ValidateProperties the validateProperties property
func (m *ServicePrincipalsRequestBuilder) ValidateProperties()(*i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd.ValidatePropertiesRequestBuilder) {
    return i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
