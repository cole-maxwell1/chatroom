import ws from "k6/ws";
import { check } from "k6";

export let options = {
  stages: [
    { duration: "30s", target: 50 }, // simulate ramp-up of traffic from 1 to 10 users over 30 seconds.
    { duration: "1m", target: 50 }, // stay at 10 users for 1 minute
    { duration: "10s", target: 0 }, // ramp-down to 0 users
  ],
};

export default function () {
  const url = "ws://localhost:8080/ws"; // replace with your WebSocket URL
  const params = { tags: { my_tag: "websocket_test" } };

  const response = ws.connect(url, params, function (socket) {
    socket.on("open", function () {
      console.log("Connected");

      const messageOptions = [
        { message: "Hello", username: "Bobby" },
        { message: "How are you?", username: "Jimmy" },
        { message: "Goodbye", username: "Donny" },
        { message: "I'm doing well", username: "Lenny" },
        { message: "I'm not doing well", username: "Jenny" },
        { message: "I'm doing okay", username: "Kenny" },
        {
          message:
            "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
          username: "RaNd0XXxxXX",
        },
        {
          message: "🤔",
          username: "🤔",
        },
        {
          message: "This is fucking awesome!!!",
          username: "Secret Admirer",
        },
      ];

      // Select a random message
      const payload =
        messageOptions[Math.floor(Math.random() * messageOptions.length)];

      socket.send(JSON.stringify(payload));

      socket.setTimeout(function () {
        socket.close();
      }, 1000);
    });

    socket.on("message", function (data) {
      console.log("Received message: ", data);
    });

    socket.on("close", function () {
      console.log("Disconnected");
    });

    socket.on("error", function (e) {
      if (e.error() != "websocket: close sent") {
        console.log("An unexpected error occurred: ", e.error());
      }
    });
  });

  check(response, {
    "Connected successfully": (res) => res && res.status === 101,
  });
}
