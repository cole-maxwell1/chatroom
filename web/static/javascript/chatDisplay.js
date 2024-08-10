//////////////////////////////////////
// Initial page setup
//////////////////////////////////////
document.addEventListener("DOMContentLoaded", () => {
  const chat = document.getElementById("chat-messages");
  if (chat) {
    formatMessageTimestamps(chat);
    scrollMessages();
  }
});

//////////////////////////////////////
// constants
//////////////////////////////////////
const STATUS_COLORS = {
  green: ["bg-green-500/20", "bg-green-500"],
  red: ["bg-red-500/20", "bg-red-500"],
  yellow: ["bg-yellow-500/20", "bg-yellow-500"],
};

//////////////////////////////////////
// Event listeners
//////////////////////////////////////

/* Websocket connection */
document.body.addEventListener("htmx:wsConnecting", () => {
  console.log("WebSocket connecting...");
  setStatusIndicator("yellow");
  updateStatus("connecting...");
});

document.body.addEventListener("htmx:wsOpen", () => {
  console.log("WebSocket connection opened");
  setStatusIndicator("green");
  updateStatus("Connected 📡");
});

document.body.addEventListener("htmx:wsClose", (event) => {
  const closeCode = event.detail.event.code;
  console.log(`WebSocket connection closed with code ${closeCode}`);

  if (willReconnect(closeCode)) {
    setStatusIndicator("yellow");
    updateStatus("Connection was lost, retrying...");
  } else {
    setStatusIndicator("red");
    updateStatus(
      `Connection closed, please refresh the page to reconnect. If the problem persists, please contact an administrator. CODE: ${closeCode}`
    );
  }
});

document.body.addEventListener("htmx:wsError", (event) => {
  console.error("WebSocket error:", event.detail);
  setStatusIndicator("red");
  updateStatus("Connection error, retrying...");
});

/* Websocket messages */
document.body.addEventListener("htmx:wsAfterMessage", scrollMessages);

// If the websocket swap is a new message, format the timestamp
document.body.addEventListener("htmx:oobBeforeSwap", (event) => {
  formatMessageTimestamps(event.detail.fragment);
});

document.body.addEventListener("htmx:wsAfterSend", () => {
  clearChatTextarea();
  countMessageCharacters();
});

/* 
Handles the form state. A user must set a username before sending a message.
The username is stored in a hidden input field and is sent with the message.
*/
const messageTextarea = document.getElementById("message");
if (messageTextarea) {
  messageTextarea.addEventListener("input", countMessageCharacters);
}

const setUserNameButton = document.getElementById("config-username");
if (setUserNameButton) {
  setUserNameButton.addEventListener("submit", (event) => {
    event.preventDefault();
    const username = document.getElementById("input-username").value;
    if (username) {
      document.getElementById("input-username").disabled = true;
      document.querySelector('textarea[name="message"]').disabled = false;
      document.getElementById("send-message").disabled = false;
      document.getElementById("username").value = username;
      document.querySelector('textarea[name="message"]').focus();
      document.getElementById("config-username").style.display = "none";
      document.getElementById("chat-form").style.display = "flex";
      document.getElementById(
        "chat-label"
      ).innerText = `Send a new message as "${username}"`;
    }
  });
}

//////////////////////////////////////
// Helper functions
//////////////////////////////////////

/**
 * Sets the status indicator color
 * @param { "green" | "red" | "yellow"} color
 */
function setStatusIndicator(color) {
  cleanStatusIndicator();
  const statusIndicator = document.getElementById("status-indicator");
  const subStatusIndicator = document.getElementById("sub-status-indicator");

  if (!statusIndicator || !subStatusIndicator) return;

  const [bgColor, subBgColor] = STATUS_COLORS[color];
  statusIndicator.classList.add(bgColor);
  subStatusIndicator.classList.add(subBgColor);
}

/**
 * Helper function to clean the status indicator classes
 * @returns {void}
 */
function cleanStatusIndicator() {
  const statusIndicator = document.getElementById("status-indicator");
  const subStatusIndicator = document.getElementById("sub-status-indicator");

  if (!statusIndicator || !subStatusIndicator) return;

  Object.values(STATUS_COLORS)
    .flat()
    .forEach((colorClass) => {
      statusIndicator.classList.remove(colorClass);
      subStatusIndicator.classList.remove(colorClass);
    });
}

/**
 * Checks if the WebSocket code will reconnect
 * @param {number} socketCode
 * @returns {boolean}
 */
function willReconnect(socketCode) {
  const reconnectCodes = [1006, 1012, 1013];
  return reconnectCodes.includes(socketCode);
}

/**
 * Update the status element with the new status text
 * @param {string} status
 */
function updateStatus(status) {
  const statusElement = document.getElementById("status");
  if (statusElement) {
    statusElement.innerText = status;
  }
}

function scrollMessages() {
  const messages = document.getElementById("chat-messages");
  if (messages) {
    requestAnimationFrame(() => {
      messages.scrollTop = messages.scrollHeight;
    });
  }
}

function clearChatTextarea() {
  const message = document.getElementById("message");
  if (message) {
    message.value = "";
  }
}

function countMessageCharacters() {
  const message = document.getElementById("message");
  const messageLength = message.maxLength;
  if (message && messageLength) {
    const messageCounter = document.getElementById("message-counter");
    if (messageCounter) {
      messageCounter.innerText = `${message.value.length}/${messageLength}`;
    }
  }
}

/**
 * Format the date when a new message is received
 * @param {HTMLElement} element
 */
function formatMessageTimestamps(element) {
  const messageTimestamps = element.querySelectorAll("#message-timestamp");
  if (messageTimestamps) {
    messageTimestamps.forEach((timestamp) => {
      const date = new Date(timestamp.innerText);
      timestamp.innerText = formatDate(date);
    });
  }
}

/**
 * Format date to a readable string
 * @example "August 8th, 2024 at 6:11 PM"
 * @param {Date} date
 * @returns {string}
 */
function formatDate(date) {
  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

  const daySuffix = (day) => {
    if (day > 3 && day < 21) return "th";
    switch (day % 10) {
      case 1:
        return "st";
      case 2:
        return "nd";
      case 3:
        return "rd";
      default:
        return "th";
    }
  };

  const hours = date.getHours();
  const minutes = date.getMinutes();
  const ampm = hours >= 12 ? "PM" : "AM";
  const formattedHours = hours % 12 || 12;
  const formattedMinutes = minutes < 10 ? "0" + minutes : minutes;

  const formattedDate = `${
    monthNames[date.getMonth()]
  } ${date.getDate()}${daySuffix(
    date.getDate()
  )}, ${date.getFullYear()} at ${formattedHours}:${formattedMinutes} ${ampm}`;

  return formattedDate;
}
