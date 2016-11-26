window.addEventListener("load", function() {
  if (!isNaN(location.pathname.substring(location.pathname.lastIndexOf('/') + 1)) && location.pathname != "/") {
  var socket = new WebSocket("ws://" + location.host + "/sockets" + location.pathname);

$('#createComment').on ('submit', function (event) {

event.preventDefault ();
event.stopImmediatePropagation ();

    var message = {
    comment: this.comment.value,
    author: "",
    date: ""
    };
    socket.send(JSON.stringify(message));
    this.reset();
});

socket.onmessage = function(event) {
console.log(event.data);
  var incomingMessage = JSON.parse(event.data);
  showMessage(incomingMessage);
};

function showMessage(message) {
    var comment = '<div class="card card-block"> <h4 class="card-title">' + message.author + '</h4>' +
                       message.date + '<p class="card-text">' + message.comment +'</p></div>';
    var newComment = document.createElement('div');
      newComment.innerHTML = comment;
      subscribe.insertBefore(newComment, subscribe.children[0]);
}

  }

});