# Proto Reference Guide

## Quick Reference

### User Service
```bash
# Port: 9091
# Database: PostgreSQL (user_db)
```

**RPCs:**
- `Register(email, password, phone, name)` → User + token
- `Login(email, password)` → token + User
- `GetProfile(user_id)` → User
- `UpdateProfile(user_id, name, phone)` → User
- `ValidateToken(token)` → user_id + role

### Event Service
```bash
# Port: 9092
# Database: MongoDB (event_db)
```

**RPCs:**
- `CreateEvent(name, description, category, image_url)` → Event
- `GetEvent(event_id)` → Event
- `ListEvents(city, category, page, page_size)` → Event[]
- `CreateVenue(name, city, address, total_seats)` → Venue
- `CreateShow(event_id, venue_id, show_time, base_price)` → Show
- `ListShows(event_id)` → Show[]
- `GetSeatLayout(show_id)` → Seat[]
- `UpdateSeatStatus(show_id, seat_ids[], status)` → success

### Booking Service
```bash
# Port: 9093
# Database: MongoDB (booking_db)
```

**RPCs:**
- `ReserveSeats(user_id, show_id, seat_ids[])` → reservation_id + expires_at
- `ConfirmBooking(reservation_id, user_id, payment_id)` → Booking
- `CancelBooking(booking_id, user_id)` → success
- `GetBooking(booking_id, user_id)` → Booking
- `ListBookings(user_id, page, page_size)` → Booking[]
- `ReleaseSeats(reservation_id, show_id, seat_ids[])` → success

### Payment Service
```bash
# Port: 9094
# Database: PostgreSQL (payment_db)
```

**RPCs:**
- `InitiatePayment(booking_id, user_id, amount)` → payment_id + order_id
- `VerifyPayment(payment_id, transaction_id, signature)` → Payment
- `ProcessRefund(payment_id, amount, reason)` → Refund
- `GetPaymentStatus(payment_id)` → Payment

### Notification Service
```bash
# Port: 9095
# Database: MongoDB
```

**RPCs:**
- `SendEmail(to, subject, body, template, data)` → success
- `SendSMS(to, message)` → success
- `SendBookingConfirmation(booking_id, user_email, ...)` → success

## Status Enums

```go
// Seat Status
Available = "available"
Reserved = "reserved"
Booked = "booked"
Blocked = "blocked"

// Booking Status
Pending = "pending"
Confirmed = "confirmed"
Cancelled = "cancelled"
Expired = "expired"

// Payment Status
Pending = "pending"
Success = "success"
Failed = "failed"
Refunded = "refunded"
```

## Usage Examples

### Import Proto
```go
import (
    userpb "ticket-booking-platform/proto/user"
    eventpb "ticket-booking-platform/proto/event"
    bookingpb "ticket-booking-platform/proto/booking"
    paymentpb "ticket-booking-platform/proto/payment"
    notificationpb "ticket-booking-platform/proto/notification"
)
```

### Create Client
```go
conn, _ := grpc.Dial("localhost:9091", grpc.WithInsecure())
defer conn.Close()
client := userpb.NewUserServiceClient(conn)
```

### Make RPC Call
```go
resp, err := client.Register(ctx, &userpb.RegisterRequest{
    Email:    "user@example.com",
    Password: "password123",
    Phone:    "+1234567890",
    Name:     "John Doe",
})
```
