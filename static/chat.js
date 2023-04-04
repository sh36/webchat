$(document).ready(function() {
    var chatBox = $("#chat-box");
    var chatForm = $("#chat-form");
    var messageInput = $("#message-input");

    chatForm.submit(function(event) {
        event.preventDefault();

        var message = messageInput.val();
        if (message) {
            //chatBox.append("<p class='message-sent'>" + message + "</p>");

            $.ajax({
                url: "/send-message",
                type: "POST",
                data: {
                    message: message
                },
                success: function(response) {
                    chatBox.append("<p class='message-received'>" + response + "</p>");
                },
                error: function() {
                    chatBox.append("<p class='error'>Error sending message.</p>");
                }
            });
        }

        messageInput.val("");
    });
});
