# Game Store Application

A modern web application for an online game store built with React, TypeScript, and Go.

## Project Structure

The project is organized into two main directories:

- **Frontend** - React TypeScript application
- **Backend** - Go RESTful API server

## Features

- User authentication (login, register)
- Game browsing and search
- Shopping cart functionality
- User profiles
- Game reviews
- Admin panel for managing games, categories, and orders

## Prerequisites

To run this project, you need to have the following installed:

- **Node.js** (v14+)
- **npm** (v6+) or **yarn**
- **Go** (v1.16+)
- **PostgreSQL** (v12+)

## Setup and Installation

### Database Setup

1. Create a PostgreSQL database
2. Configure database connection in `.env` file (see Configuration section)

### Backend Setup

1. Navigate to the Backend directory
```
cd Backend
```

2. Install Go dependencies
```
go mod tidy
```

3. Build the backend
```
go build -o gamestore-backend cmd/app/main.go
```

4. Run the backend server
```
./gamestore-backend
```

### Frontend Setup

1. Navigate to the Frontend directory
```
cd Frontend
```

2. Install dependencies
```
npm install
```

3. Start the development server
```
npm start
```

## Building for Production

A build script is provided to build both frontend and backend for production:

```
./build.ps1
```

This will:
- Build the Go backend into an executable
- Build the React frontend into optimized static files
- Combine them into a single build directory

## Configuration

Create a `.env` file in the Backend directory with the following variables:

```
# Server Configuration
PORT=your_port
IP=your_ip

# Database Configuration
DB_HOST=your_host
DB_PORT=your_port
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db

# JWT Configuration
JWT_SECRET_KEY=your_secret_key
APP_ENV=development

# Admin Configuration
ADMIN_PASSWORD=your_password
```