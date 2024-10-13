
# IOT Smart Measure KWH Backend

This repository contains the backend for the **IoT Smart Measure KWH** project. It is designed to manage hotels, sections (floors), rooms, and IoT devices for monitoring electricity usage in real time. The backend provides APIs for user authentication, hotel management, and KWH monitoring.

## Table of Contents
- [Features](#features)
- [API Endpoints](#api-endpoints)
  - [Authentication APIs](#authentication-apis)
  - [Hotel Management APIs](#hotel-management-apis)
  - [Section Management APIs](#section-management-apis)
  - [Room Management APIs](#room-management-apis)
  - [Device Management APIs](#device-management-apis)
  - [KWH Data & Monitoring APIs](#kwh-data--monitoring-apis)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Features
- User Authentication (Login, Register, Logout, Forgot Password)
- Hotel, Section (Floor), and Room management
- IoT Device management for monitoring KWH usage
- Real-time and historical KWH data tracking

## API Endpoints

### Authentication APIs
1. **POST** `/auth/login`  
   Authenticate a user and return a JWT token.  
   **Input:** `username`, `password`  
   **Output:** JWT Token, user info

2. **POST** `/auth/register`  
   Register a new user (admin/hotel manager only).  
   **Input:** `username`, `email`, `password`, `hotel_id`  
   **Output:** Success message, user details

3. **POST** `/auth/logout`  
   Invalidate the current session.

4. **POST** `/auth/forgot-password`  
   Handle password recovery.

### Hotel Management APIs
1. **GET** `/hotels`  
   Retrieve a list of all hotels.  
   **Output:** List of hotels (`hotel_name`, `address`)

2. **POST** `/hotels`  
   Create a new hotel (admin-only).  
   **Input:** `hotel_name`, `address`  
   **Output:** Hotel details

3. **PUT** `/hotels/{hotel_id}`  
   Update hotel details (admin-only).  
   **Input:** `hotel_name`, `address`

4. **DELETE** `/hotels/{hotel_id}`  
   Delete a hotel (admin-only).  
   **Output:** Success or failure message

### Section Management APIs (Floor Management)
1. **GET** `/hotels/{hotel_id}/sections`  
   Retrieve all sections (floors) for a specific hotel.  
   **Output:** List of sections (`section_name`, `section_id`)

2. **POST** `/hotels/{hotel_id}/sections`  
   Add a section to a hotel.  
   **Input:** `section_name`  
   **Output:** Section details

3. **PUT** `/sections/{section_id}`  
   Update section details.  
   **Input:** `section_name`

4. **DELETE** `/sections/{section_id}`  
   Delete a section.  
   **Output:** Success or failure message

### Room Management APIs
1. **GET** `/sections/{section_id}/rooms`  
   Retrieve all rooms within a section.  
   **Output:** List of rooms (`room_number`, `occupancy_status`, `room_id`)

2. **POST** `/sections/{section_id}/rooms`  
   Add a room to a section.  
   **Input:** `room_number`, `occupancy_status`  
   **Output:** Room details

3. **PUT** `/rooms/{room_id}`  
   Update room details.  
   **Input:** `room_number`, `occupancy_status`

4. **DELETE** `/rooms/{room_id}`  
   Delete a room.  
   **Output:** Success or failure message

### Device Management APIs
1. **GET** `/rooms/{room_id}/devices`  
   Retrieve all devices in a room.  
   **Output:** List of devices (`device_name`, `device_id`, `status`)

2. **POST** `/rooms/{room_id}/devices`  
   Register a new device in a room.  
   **Input:** `device_name`, `ssid`, `password`  
   **Output:** Device details

3. **PUT** `/devices/{device_id}`  
   Update device details (e.g., SSID/password).  
   **Input:** `device_name`, `ssid`, `password`

4. **DELETE** `/devices/{device_id}`  
   Remove a device from the system.  
   **Output:** Success or failure message

### KWH Data & Monitoring APIs
1. **GET** `/devices/{device_id}/readings?start={start_time}&end={end_time}`  
   Fetch KWH readings for a device within a specified time range.  
   **Output:** List of readings (`kwh_value`, `timestamp`)

2. **GET** `/rooms/{room_id}/real-time-usage`  
   Get real-time KWH usage for a room.  
   **Output:** Real-time KWH data (`current_kwh`, `timestamp`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/iot-smart-measure-kwh-backend.git
   cd iot-smart-measure-kwh-backend
   ```

2. Install dependencies:


3. Set up environment variables by creating a `.env` file in the root of the project and configuring the following:


4. Run the application:
  

## Usage

- Use Postman or any HTTP client to interact with the API endpoints.
- Ensure that you include the JWT token in the `Authorization` header for authenticated routes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
