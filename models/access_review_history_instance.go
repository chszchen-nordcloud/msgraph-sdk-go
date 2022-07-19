package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewHistoryInstance provides operations to manage the admin singleton.
type AccessReviewHistoryInstance struct {
    Entity
    // Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated. Required.
    downloadUri *string
    // Timestamp when this instance and associated data expires and the history is deleted. Required.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp when all of the available data for this instance was collected. This will be set after this instance's status is set to done. Required.
    fulfilledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp, reviews ending on or before this date will be included in the fetched history data.
    reviewHistoryPeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp, reviews starting on or after this date will be included in the fetched history data.
    reviewHistoryPeriodStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timestamp when the instance's history data is scheduled to be generated.
    runDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue. Once the status has been marked as done, a link can be generated to retrieve the instance's data by calling generateDownloadUri method.
    status *AccessReviewHistoryStatus
}
// NewAccessReviewHistoryInstance instantiates a new accessReviewHistoryInstance and sets the default values.
func NewAccessReviewHistoryInstance()(*AccessReviewHistoryInstance) {
    m := &AccessReviewHistoryInstance{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewHistoryInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessReviewHistoryInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewHistoryInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewHistoryInstance(), nil
}
// GetDownloadUri gets the downloadUri property value. Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated. Required.
func (m *AccessReviewHistoryInstance) GetDownloadUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.downloadUri
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. Timestamp when this instance and associated data expires and the history is deleted. Required.
func (m *AccessReviewHistoryInstance) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewHistoryInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["downloadUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUri(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["fulfilledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFulfilledDateTime(val)
        }
        return nil
    }
    res["reviewHistoryPeriodEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodEndDateTime(val)
        }
        return nil
    }
    res["reviewHistoryPeriodStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodStartDateTime(val)
        }
        return nil
    }
    res["runDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewHistoryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AccessReviewHistoryStatus))
        }
        return nil
    }
    return res
}
// GetFulfilledDateTime gets the fulfilledDateTime property value. Timestamp when all of the available data for this instance was collected. This will be set after this instance's status is set to done. Required.
func (m *AccessReviewHistoryInstance) GetFulfilledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.fulfilledDateTime
    }
}
// GetReviewHistoryPeriodEndDateTime gets the reviewHistoryPeriodEndDateTime property value. Timestamp, reviews ending on or before this date will be included in the fetched history data.
func (m *AccessReviewHistoryInstance) GetReviewHistoryPeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodEndDateTime
    }
}
// GetReviewHistoryPeriodStartDateTime gets the reviewHistoryPeriodStartDateTime property value. Timestamp, reviews starting on or after this date will be included in the fetched history data.
func (m *AccessReviewHistoryInstance) GetReviewHistoryPeriodStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewHistoryPeriodStartDateTime
    }
}
// GetRunDateTime gets the runDateTime property value. Timestamp when the instance's history data is scheduled to be generated.
func (m *AccessReviewHistoryInstance) GetRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.runDateTime
    }
}
// GetStatus gets the status property value. Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue. Once the status has been marked as done, a link can be generated to retrieve the instance's data by calling generateDownloadUri method.
func (m *AccessReviewHistoryInstance) GetStatus()(*AccessReviewHistoryStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *AccessReviewHistoryInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("downloadUri", m.GetDownloadUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("fulfilledDateTime", m.GetFulfilledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodEndDateTime", m.GetReviewHistoryPeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodStartDateTime", m.GetReviewHistoryPeriodStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("runDateTime", m.GetRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDownloadUri sets the downloadUri property value. Uri which can be used to retrieve review history data. This URI will be active for 24 hours after being generated. Required.
func (m *AccessReviewHistoryInstance) SetDownloadUri(value *string)() {
    if m != nil {
        m.downloadUri = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Timestamp when this instance and associated data expires and the history is deleted. Required.
func (m *AccessReviewHistoryInstance) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetFulfilledDateTime sets the fulfilledDateTime property value. Timestamp when all of the available data for this instance was collected. This will be set after this instance's status is set to done. Required.
func (m *AccessReviewHistoryInstance) SetFulfilledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.fulfilledDateTime = value
    }
}
// SetReviewHistoryPeriodEndDateTime sets the reviewHistoryPeriodEndDateTime property value. Timestamp, reviews ending on or before this date will be included in the fetched history data.
func (m *AccessReviewHistoryInstance) SetReviewHistoryPeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewHistoryPeriodEndDateTime = value
    }
}
// SetReviewHistoryPeriodStartDateTime sets the reviewHistoryPeriodStartDateTime property value. Timestamp, reviews starting on or after this date will be included in the fetched history data.
func (m *AccessReviewHistoryInstance) SetReviewHistoryPeriodStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewHistoryPeriodStartDateTime = value
    }
}
// SetRunDateTime sets the runDateTime property value. Timestamp when the instance's history data is scheduled to be generated.
func (m *AccessReviewHistoryInstance) SetRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.runDateTime = value
    }
}
// SetStatus sets the status property value. Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue. Once the status has been marked as done, a link can be generated to retrieve the instance's data by calling generateDownloadUri method.
func (m *AccessReviewHistoryInstance) SetStatus(value *AccessReviewHistoryStatus)() {
    if m != nil {
        m.status = value
    }
}
