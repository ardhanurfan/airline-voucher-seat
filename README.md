# Airline Voucher Seat Management App

An application for generate airline voucher seat, with separate backend and frontend services. This guide will help you set up and run the application in a development environment.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Environment Variables](#environment-variables)
- [Folder Structure](#folder-structure)
  - [Backend Structure (Golang)](#backend-structure-golang)
  - [Frontend Structure (Next.js)](#frontend-structure-nextjs)
- [Local Setup](#local-setup)
  - [Backend (Golang)](#backend-setup-airline-voucher-seat-be)
  - [Frontend (Next.js)](#frontend-setup-airline-voucher-seat-fe)
- [Setup with Docker](#setup-with-docker)
- [Development Workflow](#development-workflow)
- [API Documentation](#api-documentation)

---

## Prerequisites

Ensure the following tools are installed on your system:

- **Golang** (version 1.24 or newer)
- **Node.js** (version 18 or newer)
- **npm** or **yarn**
- **Git**
- **Make** (to run the backend)
- **Docker** & **Docker Compose** (optional, for containerized setup)

---

## Environment Variables

Before running the application (either locally or with Docker), you need to create a `.env` file.

1.  Inside the `airline-voucher-seat-be` and `airline-voucher-seat-fe` directories, copy the `.env.example` file to `.env`.

    ```sh
    # Inside the backend folder
    cp .env.example .env

    # Inside the frontend folder
    cp .env.example .env
    ```

2.  Adjust the variables inside the newly created `.env` files according to your configuration.

---

## Folder Structure

This project is divided into two main directories:

```
airline-voucher-seat-app/
├── airline-voucher-seat-be/  # Backend Application (Golang)
└── airline-voucher-seat-fe/  # Frontend Application (Next.js)
```

### Backend Structure (Golang)

The backend follows a _Clean Architecture_ approach to separate business logic from implementation details.

```
airline-voucher-seat-be/
├── cmd/main/              # Main application entry point
├── config/                # Database and server configuration
├── data/                  # Directory for the SQLite database file (if used)
├── internal/              # Core application logic not intended for export
│   ├── controllers/       # (Handler) Receives HTTP requests and calls use cases
│   ├── models/            # Data structures (structs) for entities
│   ├── repositories/      # Logic for accessing and manipulating data in the database
│   ├── routes/            # API route definitions and binding to controllers
│   └── usecase/           # Main business logic of the application
├── middlewares/           # Middleware for HTTP requests (e.g., body validator)
├── Makefile               # Commands for build, run, test, etc.
└── Dockerfile             # Configuration for building the Docker image
```

### Frontend Structure (Next.js)

The frontend uses the standard Next.js project structure.

```
airline-voucher-seat-fe/
├── .next/                 # Build output from Next.js
├── lib/                   # Helper or utility functions
├── public/                # Static assets like images and fonts
├── src/                   # Main source code of the frontend application
├── types/                 # TypeScript type definitions
└── Dockerfile             # Configuration for building the Docker image
```

---

## Local Setup

### Backend Setup (airline-voucher-seat-be)

1.  Open a terminal and navigate to the backend directory:

    ```sh
    cd airline-voucher-seat-be
    ```

2.  Download all required dependencies:

    ```sh
    go mod tidy
    ```

3.  Run the backend server using Make:
    ```sh
    make run
    ```

The backend server will run at `http://localhost:8080`.

### Frontend Setup (airline-voucher-seat-fe)

1.  Open a new terminal and navigate to the frontend directory:

    ```sh
    cd airline-voucher-seat-fe
    ```

2.  Install all required dependencies:

    ```sh
    npm install
    # or
    yarn install
    ```

3.  Run the frontend development server:
    ```sh
    npm run dev
    # or
    yarn dev
    ```

The frontend application will be accessible at `http://localhost:3000`.

---

## Setup with Docker

Alternatively, you can run the entire application using Docker Compose.

1.  Ensure you are in the project's root directory (the one containing the `docker-compose.yml` file).

2.  Build and run the containers:
    ```sh
    docker-compose up --build
    ```

This command will build the images and run the containers for both the backend and frontend simultaneously.

- The frontend will be available at `http://localhost:3000`
- The backend will be available at `http://localhost:8080`

---

## Development Workflow

1.  Ensure all environment variables are set correctly.
2.  Start the backend server first.
3.  Start the frontend server.
4.  Access the application via a browser at `http://localhost:3000`.

---

## API Documentation

The following is a list of available API endpoints in the backend.

### 1. Generate Voucher

Creates a new voucher and generates random seat numbers.

- **Endpoint**: `POST /api/generate`
- **Request Body**:
  ```json
  {
    "name": "Sarah",
    "id": "98123",
    "flightNumber": "GA102",
    "date": "2025-07-12",
    "aircraft": "Airbus 320"
  }
  ```
- **Success Response**:
  ```json
  {
    "success": true,
    "message": "Voucher generated successfully",
    "data": {
      "seats": ["6A", "14C", "11D"]
    }
  }
  ```

### 2. Check Voucher

Checks if a voucher for a specific flight number and date already exists.

- **Endpoint**: `POST /api/check`
- **Request Body**:
  ```json
  {
    "flightNumber": "GA102",
    "date": "2025-07-12"
  }
  ```
- **Success Response (if exists)**:
  ```json
  {
    "success": true,
    "message": "Voucher checked successfully",
    "data": {
      "exists": true
    }
  }
  ```
- **Success Response (if not exists)**:
  ```json
  {
    "success": true,
    "message": "Voucher checked successfully",
    "data": {
      "exists": false
    }
  }
  ```

### 3. Get Aircraft Types

Retrieves a list of all available aircraft types.

- **Endpoint**: `GET /api/aircrafts`
- **Success Response**:
  ```json
  {
    "success": true,
    "message": "Aircraft types fetched successfully",
    "data": ["ATR", "Airbus 320", "Boeing 737 Max"]
  }
  ```

### 4. Get All Vouchers

Retrieves a list of all vouchers that have been created.

- **Endpoint**: `GET /api/vouchers`
- **Success Response**:
  ```json
  {
    "success": true,
    "message": "Vouchers fetched successfully",
    "data": [
      {
        "id": 1,
        "crew_name": "Sarah",
        "crew_id": "98123",
        "flight_number": "GA102",
        "flight_date": "2025-07-12",
        "aircraft_type": "Airbus 320",
        "seat1": "24C",
        "seat2": "18A",
        "seat3": "22B",
        "created_at": "2025-07-23T21:25:21.809075+07:00"
      }
    ]
  }
  ```

---

<p align="center">
  ardhanurfan Copyrights @2025 for BookCabin Project Test
</p>
