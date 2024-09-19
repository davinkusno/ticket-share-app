# Ticket Share App - Backend Documentation

This document provides a detailed overview of the backend API endpoints and the database structure for the Ticket Share application.

## Table of Contents
- [Installation Guide](#installation-guide)
- [API Endpoints](#api-endpoints)
  - [User Authentication](#user-authentication)
  - [Events](#events)
  - [Cart](#cart)
  - [Health Check](#health-check)
- [Database Structure](#database-structure)

## How to Download and Run the Backend

Follow these steps to download and run the backend locally.

### 1. Clone the Repository

Start by cloning the repository to your local machine:

```bash
git clone https://github.com/davinkusno/ticket-share-app.git
cd ticket-share-backend

### 2. Install Dependencies
Once inside the project directory, run the following command to install the necessary dependencies:
go mod tidy

### 3. Setup Your Database
Make sure you have PostgreSQL installed and running on your local machine. Create a PostgreSQL database and update the database credentials in main.go to match your local environment:

dbCredential := database.Credential{
    Host:         "localhost",   // Your database host
    Username:     "myuser",      // Your database username
    Password:     "mypassword",  // Your database password
    DatabaseName: "postgres",    // Your database name
    Port:         5432,          // Your database port
}

### 4. Create .env
create .env file aligned with main.go, and insert your DB credentials and jwt secret key just like the .env.example

### 5. Run the Application
go run main.go

## API Endpoints

### User Authentication

#### **Register User**
- **Endpoint**: `POST /register`
- **Description**: Registers a new user.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response**:
  ```json
  {
    "message": "User registered successfully!"
  }
  ```

#### **Login User**
- **Endpoint**: `POST /login`
- **Description**: Logs in a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response**:
  ```json
  {
    "token": "jwt_token_here"
  }
  ```

#### **Get User Profile**
- **Endpoint**: `GET /profile`
- **Description**: Retrieves the profile of the authenticated user.
- **Headers**: `Authorization: Bearer {jwt_token}`
- **Response**:
  ```json
  {
    "Name": "John Doe",
    "Email": "john@example.com"
  }
  ```

### Events

#### **Get All Events**
- **Endpoint**: `GET /events`
- **Description**: Fetches all available events.
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Music Concert",
      "description": "A live music concert featuring top artists.",
      "date": "2024-10-01",
      "price": 50
    },
    {
      "id": 2,
      "name": "Tech Conference",
      "description": "A conference about the latest in tech.",
      "date": "2024-10-08",
      "price": 100
    }
  ]
  ```

#### **Get Event by ID**
- **Endpoint**: `GET /events/:id`
- **Description**: Fetches an event by its ID.
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Music Concert",
    "description": "A live music concert featuring top artists.",
    "date": "2024-10-01",
    "price": 50
  }
  ```

#### **Create Event**
- **Endpoint**: `POST /events`
- **Description**: Creates a new event.
- **Request Body**:
  ```json
  {
    "name": "New Event",
    "description": "A new event description",
    "date": "2024-11-15",
    "price": 75
  }
  ```
- **Response**:
  ```json
  {
    "id": 5,
    "name": "New Event",
    "description": "A new event description",
    "date": "2024-11-15",
    "price": 75
  }
  ```

#### **Update Event**
- **Endpoint**: `PUT /events/:id`
- **Description**: Updates an existing event.
- **Request Body**:
  ```json
  {
    "name": "Updated Event",
    "description": "Updated description",
    "date": "2024-12-01",
    "price": 80
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Updated Event",
    "description": "Updated description",
    "date": "2024-12-01",
    "price": 80
  }
  ```

#### **Delete Event**
- **Endpoint**: `DELETE /events/:id`
- **Description**: Deletes an event by ID.
- **Response**:
  ```json
  {
    "message": "Event deleted successfully"
  }
  ```

### Cart

#### **Get All Cart Items**
- **Endpoint**: `GET /cart/:user_id`
- **Description**: Fetches all items in the user's cart.
- **Response**:
  ```json
  {
    "cartItems": [
      {
        "id": 1,
        "event": {
          "id": 1,
          "name": "Music Concert",
          "price": 50
        },
        "quantity": 2
      }
    ],
    "totalPrice": 100
  }
  ```

#### **Add to Cart**
- **Endpoint**: `POST /cart`
- **Description**: Adds an item to the user's cart.
- **Request Body**:
  ```json
  {
    "user_id": 1,
    "event_id": 1,
    "quantity": 2
  }
  ```
- **Response**:
  ```json
  {
    "message": "Added to cart successfully",
    "cart": {
      "id": 1,
      "event_id": 1,
      "user_id": 1,
      "quantity": 2
    }
  }
  ```

#### **Update Cart Item Quantity**
- **Endpoint**: `PUT /cart/:id`
- **Description**: Updates the quantity of an item in the cart.
- **Request Body**:
  ```json
  {
    "quantity": 3
  }
  ```
- **Response**:
  ```json
  {
    "message": "Cart item updated successfully",
    "cart": {
      "id": 1,
      "quantity": 3
    }
  }
  ```

#### **Delete Cart Item**
- **Endpoint**: `DELETE /cart/:id`
- **Description**: Deletes an item from the user's cart.
- **Response**:
  ```json
  {
    "message": "Cart item deleted successfully"
  }
  ```

### Health Check

#### **Health Check**
- **Endpoint**: `GET /health`
- **Description**: A health check endpoint to verify the server status.
- **Response**:
  ```json
  {
    "status": "OK"
  }
  ```

## Database Structure

The following is the database structure used by the Ticket Share App:

| Table   | Fields                                                       |
|---------|---------------------------------------------------------------|
| `users` | `id`, `name`, `email`, `password`                             |
| `events`| `id`, `name`, `description`, `date`, `price`                  |
| `cart`  | `id`, `user_id`, `event_id`, `quantity`                       |

### `users` Table

| Field     | Type         | Description                  |
|-----------|--------------|------------------------------|
| `id`      | INT          | Primary key, auto-increment   |
| `name`    | VARCHAR(255) | Name of the user              |
| `email`   | VARCHAR(255) | Email address of the user     |
| `password`| VARCHAR(255) | Hashed password               |

### `events` Table

| Field        | Type         | Description                        |
|--------------|--------------|------------------------------------|
| `id`         | INT          | Primary key, auto-increment         |
| `name`       | VARCHAR(255) | Name of the event                   |
| `description`| TEXT         | Description of the event            |
| `date`       | DATE         | Date of the event                   |
| `price`      | DECIMAL      | Ticket price                        |

### `cart` Table

| Field      | Type         | Description                            |
|------------|--------------|----------------------------------------|
| `id`       | INT          | Primary key, auto-increment             |
| `user_id`  | INT          | Foreign key referencing `users`         |
| `event_id` | INT          | Foreign key referencing `events`        |
| `quantity` | INT          | Number of tickets in the cart           |
