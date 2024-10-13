# IOTSmartMeasureKWH
This Is Backend For IOT Smart Measure KWH Project

Features
User Authentication (Login, Register, Logout, Forgot Password)
Hotel, Section (Floor), and Room management
IoT Device management for monitoring KWH usage
Real-time and historical KWH data tracking
API Endpoints
Authentication APIs
POST /auth/login
Authenticate a user and return a JWT token.
Input: username, password
Output: JWT Token, user info

POST /auth/register
Register a new user (admin/hotel manager only).
Input: username, email, password, hotel_id
Output: Success message, user details

POST /auth/logout
Invalidate the current session.

POST /auth/forgot-password
Handle password recovery.

Hotel Management APIs
GET /hotels
Retrieve a list of all hotels.
Output: List of hotels (hotel_name, address)

POST /hotels
Create a new hotel (admin-only).
Input: hotel_name, address
Output: Hotel details

PUT /hotels/{hotel_id}
Update hotel details (admin-only).
Input: hotel_name, address

DELETE /hotels/{hotel_id}
Delete a hotel (admin-only).
Output: Success or failure message

Section Management APIs (Floor Management)
GET /hotels/{hotel_id}/sections
Retrieve all sections (floors) for a specific hotel.
Output: List of sections (section_name, section_id)

POST /hotels/{hotel_id}/sections
Add a section to a hotel.
Input: section_name
Output: Section details

PUT /sections/{section_id}
Update section details.
Input: section_name

DELETE /sections/{section_id}
Delete a section.
Output: Success or failure message

Room Management APIs
GET /sections/{section_id}/rooms
Retrieve all rooms within a section.
Output: List of rooms (room_number, occupancy_status, room_id)

POST /sections/{section_id}/rooms
Add a room to a section.
Input: room_number, occupancy_status
Output: Room details

PUT /rooms/{room_id}
Update room details.
Input: room_number, occupancy_status

DELETE /rooms/{room_id}
Delete a room.
Output: Success or failure message

Device Management APIs
GET /rooms/{room_id}/devices
Retrieve all devices in a room.
Output: List of devices (device_name, device_id, status)

POST /rooms/{room_id}/devices
Register a new device in a room.
Input: device_name, ssid, password
Output: Device details

PUT /devices/{device_id}
Update device details (e.g., SSID/password).
Input: device_name, ssid, password

DELETE /devices/{device_id}
Remove a device from the system.
Output: Success or failure message

KWH Data & Monitoring APIs
GET /devices/{device_id}/readings?start={start_time}&end={end_time}
Fetch KWH readings for a device within a specified time range.
Output: List of readings (kwh_value, timestamp)

GET /rooms/{room_id}/real-time-usage
Get real-time KWH usage for a room.
Output: Real-time KWH data (current_kwh, timestamp)
