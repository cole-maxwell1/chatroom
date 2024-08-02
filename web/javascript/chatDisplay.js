// listen for click event on the Set Username button
document.getElementById('set-username').addEventListener('click', function (event) {
    //event.preventDefault();
    const username = document.getElementById('input-username').value;
    if (username) {
        document.getElementById('input-username').disabled = true;
        document.querySelector('textarea[name="message"]').disabled = false;
        document.querySelector('button[name="send-message"]').disabled = false;
        document.getElementById('username').value = username;
        document.querySelector('textarea[name="message"]').focus();
        // hide the username configuration
        document.getElementById('config-username').style.display = 'none';
        // show the chat form
        document.getElementById('chat-form').style.display = 'flex';
        document.getElementById('chat-label').innerText = `Send a new message as ${username}`;
    }
});

// listen for htmx:beforeSend event on the chat form and clear the textarea
document.getElementById('chat-form').addEventListener('htmx:afterRequest', function (event) {
    document.querySelector('textarea[name="message"]').value = '';
});
