package education

import (
	"context"
	iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/chszchen-nordcloud/msgraph-sdk-go/models"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder provides operations to call the unsubmit method.
type ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal instantiates a new UnsubmitRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder {
	m := &ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/unsubmit", pathParameters),
	}
	return m
}

// NewClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder instantiates a new UnsubmitRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal(urlParams, requestAdapter)
}

// Post indicate that a student wants to work on the submission of the assignment after it was turned in. Only teachers, students, and applications with application permissions can perform this operation. This method changes the status of the submission from `submitted` to `working`. During the submit process, all the resources are copied from **submittedResources** to  **workingResources**. The teacher will be looking at the working resources list for grading. A teacher can also unsubmit a student's assignment on their behalf.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-unsubmit?view=graph-rest-1.0
func (m *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) Post(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration) (iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
		"5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable), nil
}

// ToPostRequestInformation indicate that a student wants to work on the submission of the assignment after it was turned in. Only teachers, students, and applications with application permissions can perform this operation. This method changes the status of the submission from `submitted` to `working`. During the submit process, all the resources are copied from **submittedResources** to  **workingResources**. The teacher will be looking at the working resources list for grading. A teacher can also unsubmit a student's assignment on their behalf.
func (m *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
	requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
	requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
	requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
	requestInfo.Headers.Add("Accept", "application/json")
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	return requestInfo, nil
}
