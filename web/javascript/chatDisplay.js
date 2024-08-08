// listen for htmx:wsConnecting event
document.body.addEventListener("htmx:wsConnecting", function (event) {
  console.log("WebSocket connecting...");

  setStatusIndicator("yellow");
  updateStatus("connecting...");
});

// listen for htmx:wsOpen event
document.body.addEventListener("htmx:wsOpen", function (event) {
  console.log("WebSocket connection opened");

  // Update status-indicator class
  setStatusIndicator("green");
  updateStatus("Connected 📡");
});

// listen for htmx:wsClose event
document.body.addEventListener("htmx:wsClose", function (event) {
  // @ts-ignore
  const closeCode = event.detail.event.code;
  console.log(`WebSocket connection closed with code ${closeCode}`);

  if (willReconnect(closeCode)) {
    setStatusIndicator("yellow");
    updateStatus("Connection was lost, retrying...");
  } else {
    setStatusIndicator("red");
    updateStatus(`Connection closed, please refresh the page to reconnect.
    If the problem persists, please contact an administrator. CODE: ${closeCode}`);
  }
});

// listen for htmx:wsError event
document.body.addEventListener("htmx:wsError", function (event) {
  // @ts-ignore
  console.error("WebSocket error:", event.detail);
  setStatusIndicator("red");
  updateStatus("Connection error, retrying...");
});

// When a new message is received, scroll to the bottom of the chat
document.body.addEventListener("htmx:wsAfterMessage", (event) => {
  scrollMessages();
});

// listen for htmx:wsAfterSend event on the chat form and clear the textarea
document.body.addEventListener("htmx:wsAfterSend", function (event) {
  console.log(event);
  clearChatTextarea();
  countMessageCharacters();
});

const messageTextarea = document.getElementById("message");
if (messageTextarea) {
  messageTextarea.addEventListener("input", function (event) {
    countMessageCharacters();
  });
}

const setUserNameButton = document.getElementById("config-username");
// listen for click event on the Set Username button
if (setUserNameButton) {
  setUserNameButton.addEventListener("submit", function (event) {
    // prevent the form from submitting
    event.preventDefault();

    const username = document.getElementById("input-username").value;
    // check if the username is not empty
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

/**
 * Sets the status indicator color
 * @param { "green" | "red" | "yellow"} color
 */
function setStatusIndicator(color) {
  cleanStatusIndicator();
  const statusIndicator = document.getElementById("status-indicator");
  const subStatusIndicator = document.getElementById("sub-status-indicator");

  if (!statusIndicator || !subStatusIndicator) {
    return;
  }

  switch (color) {
    case "green":
      statusIndicator.classList.add("bg-green-500/20");
      subStatusIndicator.classList.add("bg-green-500");
      break;
    case "red":
      statusIndicator.classList.add("bg-red-500/20");
      subStatusIndicator.classList.add("bg-red-500");
      break;
    case "yellow":
      statusIndicator.classList.add("bg-yellow-500/20");
      subStatusIndicator.classList.add("bg-yellow-500");
      break;
  }
}

/**
 * Helper function to clean the status indicator classes
 * @returns {void}
 */
function cleanStatusIndicator() {
  const statusIndicator = document.getElementById("status-indicator");
  const subStatusIndicator = document.getElementById("sub-status-indicator");

  if (!statusIndicator || !subStatusIndicator) {
    return;
  }

  const possibleStatusIndicatorClass = [
    "bg-red-500/20",
    "bg-yellow-500/20",
    "bg-green-500/20",
  ];

  const possibleSubStatusIndicatorClass = [
    "bg-red-500",
    "bg-yellow-500",
    "bg-green-500",
  ];

  for (let i = 0; i < possibleStatusIndicatorClass.length; i++) {
    statusIndicator.classList.remove(possibleStatusIndicatorClass[i]);
  }

  for (let i = 0; i < possibleSubStatusIndicatorClass.length; i++) {
    subStatusIndicator.classList.remove(possibleSubStatusIndicatorClass[i]);
  }
}

/**
 * Checks if the Websocket code will reconnect
 * @param {number} socketCode
 * @returns {boolean}
 */
function willReconnect(socketCode) {
  // HTMX will automatically reconnect on these codes
  const reconnectCodes = [1006, 1012, 1013];

  if (reconnectCodes.includes(socketCode)) {
    return true;
  }

  return false;
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
    messages.scrollTop = messages.scrollHeight;
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
