# Chatroom Project

This project is a chatroom application I threw together to learn more about Websockets, Go channels, and the hot Javascript library of the moment [HTMX](https://htmx.org/). The chatroom allows users to send messages to each other in real-time and see how many other users are connected to the chatroom. The last 100 messages are stored in a ring buffer and displayed to new users when they join the chatroom.

## Concepts Explored

- **Websockets**: Utilizes a *Websocket* to push messages and total connected users to clients in real-time.
- **Efficient Message Storage**: Implements a [ring buffer](https://en.wikipedia.org/wiki/Circular_buffer) data structure to manage and store chat messages, preventing high memory usage. The application stores the last 100 messages.
- **Injection Attacks**: Uses the `templ` library to prevent injection attacks in the chatroom.
- **Javascript Internationalization API**: Utilizes the [Intl API](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl) to format dates and times in the chatroom.
- **Go Channels**: This server makes heavy use of Go channels to manage the communication between the server and clients. All sends and receives are done through channels and a central broker to prevent blocks and slow delivery to other connected clients.

## Usage

 1. You can pull down the published docker image and give it a try yourself:
    ```sh
    docker run -d -p 8080:8080 ghcr.io/cole-maxwell1/chatroom:latest
    ```
 1. Or view live at [chat.colemaxwell.dev](https://chat.colemaxwell.dev/)

## K6 Load Testing

I ran some load testing on the chatroom using the [K6](https://grafana.com/docs/k6/latest/get-started/) load testing tool. The test ramps to 50 users over 30 seconds and sustains 50 for one minute. Each connection sends a message after connecting. The test was run on a Hetzner VPS with 2 vCPUs and 2GB of RAM. `loadtest.js` is configured to run on a local instance. You can run the test yourself with the following command:
```sh
k6 run loadtest.js
```

Results:
```sh
✓ Connected successfully

     checks................: 100.00% ✓ 3316        ✗ 0
     data_received.........: 165 MB  1.6 MB/s
     data_sent.............: 2.6 MB  26 kB/s
     iteration_duration....: avg=1.21s    min=1.14s    med=1.16s    max=2.55s p(90)=1.35s    p(95)=1.45s
     iterations............: 3316    32.915225/s
     vus...................: 4       min=2         max=50
     vus_max...............: 50      min=50        max=50
     ws_connecting.........: avg=209.99ms min=139.68ms med=162.28ms max=1.53s p(90)=348.65ms p(95)=451.28ms
     ws_msgs_received......: 343532  3409.961751/s
     ws_msgs_sent..........: 3316    32.915225/s
     ws_session_duration...: avg=1.21s    min=1.14s    med=1.16s    max=2.55s p(90)=1.35s    p(95)=1.45s
     ws_sessions...........: 3316    32.915225/s


running (1m40.7s), 00/50 VUs, 3316 complete and 0 interrupted iterations
default ✓ [======================================] 00/50 VUs  1m40s
```

## Development

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
    make build
    ```

1. **Run the Server**:
    ```sh
    make run
    ```

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
