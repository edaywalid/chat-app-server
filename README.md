# Go (Gin) Chat App Server

This project is a real-time chat server built using the **Go** programming language and the **Gin** web framework. It features user authentication, email verification, password reset functionalities, and supports various chat types (1-1, group, and broadcast chats). Below is a high-level overview of the features and technologies used.

## Features

1. **User Authentication**: JWT-based authentication is used to secure the API. Users must sign up and confirm their email addresses before gaining access.
2. **Email Confirmation**: Users receive a confirmation link in their email to verify their account.
3. **Password Reset**: A password reset flow is implemented, allowing users to reset their password via email.
4. **PostgreSQL for User Management**: User information (such as email, password, etc.) is stored securely in a PostgreSQL database.
5. **1-1 Chat**: Users can engage in one-on-one conversations.
6. **Group Chat**: Multiple users can join a group and chat in real-time.
7. **Broadcast Chat**: Send messages to a large number of users in a broadcast fashion.
8. **MongoDB for Chat History**: All chat messages (1-1, group, and broadcast) are stored in MongoDB for easy access and scalability.
9. **Redis Message Broker**: Redis is used as a message broker for handling real-time messaging efficiently. Go channels and concurrency help to manage message flow and background tasks.
10. **WebSockets for Real-time Communication**: WebSockets are used for real-time bi-directional communication between the client and server, ensuring seamless chat experiences.
11. **Concurrency with Go Routines**: Concurrency is managed with Go routines and channels to handle multiple users and chats efficiently.

## Technologies

- **Go (Gin)**: Web framework for building the REST API.
- **JWT**: Token-based authentication for secure API access.
- **PostgreSQL**: Storing user data.
- **MongoDB**: Storing chat messages.
- **Redis**: Message broker for real-time communication.
- **Go Channels & Goroutines**: Concurrency management for efficient message handling.
- **WebSockets**: Real-time communication for chat functionalities.

## Getting Started

### Prerequisites

- Golang >= 1.22.5
- Docker >= 27.0.3
- Docker Compose >= 2.28.1
- Make >= 4.4.1

### Installation

1. Clone the repository:

```bash
    git clone www.github.com/edaywalid/chat-app-server
    cd chat-app-server
```

2. Create a `.env` file in the root directory and add the following environment variables:

```bash
    POSTGRES_URL=postgres://root:password@localhost:5432
    MONGO_URI=mongodb://root:password@localhost:27017
    JWT_SECRET=
    JWT_RESET_KEY=
    SMTP_USER=
    SMTP_PASS=
    SMTP_HOST=
    SMTP_PORT=
    REDIS_ADDR=
    REDIS_URL=
```

3. Build and run the project using Docker Compose:

```bash
    docker-compose up -d
    make run
```

### Notes

- The project is still under development and may contain bugs or incomplete features.
