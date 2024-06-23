# CafeManagerAPI

CafeManagerAPI is a RESTful API for managing various aspects of a cafe, including menu items, customer orders, customer details, and reservations. This project demonstrates CRUD operations, basic authentication, and API design principles using Go and the Gorilla Mux package.

## Features

- **Menu Item Management**: Add, update, delete, and retrieve menu items.
- **Order Management**: Create, update, delete, and retrieve customer orders.
- **Customer Management**: Manage customer details, including loyalty points and order history.
- **Reservation Management**: Handle table reservations, including availability checks.

## Endpoints

### Menu Items

- **GET /menu**: Retrieve all menu items.
- **POST /menu**: Create a new menu item.

### Orders

- **GET /order/{id}**: Retrieve a specific order by ID.
- **POST /order**: Create a new order.

### Customers

- **GET /customer/{id}**: Retrieve a specific customer by ID.
- **POST /customer**: Create a new customer.

### Reservations

- **GET /reservation/{id}**: Retrieve a specific reservation by ID.
- **POST /reservation**: Create a new reservation.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16+ recommended)
