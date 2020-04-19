# Cab Booking with Gin Framework (GOLANG)

An basic cab booking REST apis, using gin framework and mysql for database. This can be used to create a new user, get all the cabs near by user location, book a cab from point A to point B and also get details for all past bookings for particular user.

### 1. Create user.
##### POST
```
http://localhost:8080/api/v1/booking/createUser
```
##### Request
```
{
	"name":"Mark Jons",
	"email":"mark@yopmail.com",
	"password":"mark@123",
	"mobile_number":"7894561230"
}
```
##### Response
```
{
    "user_Id": 1,
    "name": "Mark Jons",
    "email": "mark@yopmail.com",
    "mobile_number": "7894561230",
    "message": "User created successfully!",
    "status": 201
}
```
### 2. Get Near by cabs.

##### GET
```
http://localhost:8080/api/v1/booking/cabs?location=katraj
```
##### Request
```
- Used query parameter "cabs?location=katraj"
- Get cabs list whichever is available in that location
- No request body
```
##### Response
```
{
    "data": [
        {
            "id": 1,
            "driver_name": "Peter",
            "driver_number": "9874561111"
        },
        {
            "id": 2,
            "driver_name": "Michal",
            "driver_number": "9874562222"
        }
    ],
    "status": 200
}
```

### 3. Book a cab from location A to location B.

##### POST
```
http://localhost:8080/api/v1/booking/book
```
##### Request
```
{
	"user_id":1,
	"cab_id":1,
	"from_loc":"katraj",
	"to_loc":"wakad"
}
```
##### Response
```
{
    "bookingId": 10,
    "cab_number": "MHV 1209",
    "driver_name": "John",
    "driver_number": "9874563333",
    "message": "Cab booked successfully!",
    "status": 201
}
```
### 4. Get past bookings.

##### GET
```
http://localhost:8080/api/v1/booking/getBookings/2
```
##### Request
```
- Used path parameter by providing user id
- Get past bookings list of the user
- No request body
```
##### Response
```
{
    "data": [
        {
            "id": 3,
            "from_loc": "katraj",
            "to_loc": "wakad",
            "booked_at": "2020-04-18T21:07:29+05:30"
        },
        {
            "id": 10,
            "from_loc": "katraj",
            "to_loc": "wakad",
            "booked_at": "2020-04-19T17:41:12+05:30"
        }
    ],
    "status": 200
}
```

## Test driven development.

* Running `go test` in the same directory as this file cabs_booking_test.go will run the tests.


