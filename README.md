# Chatroom Project

This project is a chatroom application I threw together to learn more about Websockets, Go channels, and the hot Javascript library of the moment [HTMX](https://htmx.org/). The chatroom allows users to send messages to each other in real-time and see how many other users are connected to the chatroom. The last 100 messages are stored in a ring buffer and displayed to new users when they join the chatroom.

## Concepts Explored

- **Websockets**: Utilizes a *Websocket* to push messages and total connected users to clients in real-time.
- **Efficient Message Storage**: Implements a [ring buffer](https://en.wikipedia.org/wiki/Circular_buffer) data structure to manage and store chat messages, preventing high memory usage. The application stores the last 100 messages.
- **Injection Attacks**: Uses the `templ` library to prevent injection attacks in the chatroom.
- **Javascript Internationalization API**: Utilizes the [Intl API](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl) to format dates and times in the chatroom.
- **Go Channels**: This server makes heavy use of Go channels to manage the communication between the server and clients. All sends and receives are done through channels and a central broker to prevent blocks and slow delivery to other connected clients.

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

## Injection Attacks
One of the big reasons for using the `templ` library is to prevent injection attacks in the chatroom. The `templ` library has some nice built checks to prevent this. For example, the `templ` library will not allow the use of the `<script>` tag in the chatroom. Don't believe me? Try it out for yourself!
```html
<script> 
    body.addEventListener('htmx:wsBeforeMessage'
    function (event) { console.log('Nasty code was run'); document.getElementById('status').innerText = 'Nasty code was run!!!'; }); 
</script>
```

## HTMX Takeaways
I really enjoyed getting a little closer to the DOM and using HTMX. The web sockets extension was really easy to work with. I'd like to seem the HTMX team add some additional documentation on what [websocket close codes](https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent/code) the automatic reconnect feature will work on, for the record it's only codes `1006, 1012, 1013` that reconnect automatically. There was some learning curve with the client side date formatting, but I was able to get it working using the `htmx:oobBeforeSwap` event, which wasn't well documented. The nice thing about using a small library like HTMX is that it's actually easy to read the source code and figure out what it's doing and how to use it.