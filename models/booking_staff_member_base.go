package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingStaffMemberBase 
type BookingStaffMemberBase struct {
    Entity
}
// NewBookingStaffMemberBase instantiates a new bookingStaffMemberBase and sets the default values.
func NewBookingStaffMemberBase()(*BookingStaffMemberBase) {
    m := &BookingStaffMemberBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingStaffMemberBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingStaffMemberBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingStaffMemberBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingStaffMemberBase) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *BookingStaffMemberBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
