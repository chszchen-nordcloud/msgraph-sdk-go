package identitygovernance

import (
	"context"
	ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/chszchen-nordcloud/msgraph-sdk-go/models/odataerrors"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder provides operations to call the stop method.
type AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AccessReviewsDefinitionsItemInstancesItemStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStopRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAccessReviewsDefinitionsItemInstancesItemStopRequestBuilderInternal instantiates a new StopRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder {
	m := &AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stop", pathParameters),
	}
	return m
}

// NewAccessReviewsDefinitionsItemInstancesItemStopRequestBuilder instantiates a new StopRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAccessReviewsDefinitionsItemInstancesItemStopRequestBuilderInternal(urlParams, requestAdapter)
}

// Post stop a currently active accessReviewInstance. After the access review instance stops, the instance status is marked as `Completed`, the reviewers can no longer give input, and the access review decisions are applied. Stopping an instance will not stop future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
// [Find more info here]
//
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-stop?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilderPostRequestConfiguration) error {
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

// ToPostRequestInformation stop a currently active accessReviewInstance. After the access review instance stops, the instance status is marked as `Completed`, the reviewers can no longer give input, and the access review decisions are applied. Stopping an instance will not stop future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
func (m *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStopRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
