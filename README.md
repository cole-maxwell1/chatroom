# Chatroom Project

This project is a chatroom application that leverages [Websockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) for real-time communication and a ring buffer data structure to store messages efficiently.

## Concepts Explored

- **Real-time Communication**: Utilizes s *Websocket* to push messages and total connected users to clients in real-time.
- **Efficient Message Storage**: Implements a ring buffer data structure to manage and store chat messages, preventing high memory usage.

## Technologies Used

- **Go**: Backend server implementation.
- **HTMX**: Frontend client-side scripting.
- **Makefile**: Build automation.

## Setup Instructions

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/cole-maxwell1/chatroom.git && cd chatroom
    ```

1. **Install Dependencies**:
    ```sh
    npm install
    ```

1. **Build the Project**:
    ```sh
    make
    ```

1. **Run the Server**:
    ```sh
    make run
    ```

1. **Open the Client**:
    Open `index.html` in your web browser to start the chatroom client.

## Usage

- Open the chatroom client in multiple browser tabs or different browsers. By default http://localhost:8080
- Start typing messages and see them appear in real-time across all clients.