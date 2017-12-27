window.onload = function() {
  document.getElementById('post-button')
    .addEventListener('click', postComment)
  document.getElementById('comment-form')
    .addEventListener('onsubmit', postComment)
  fillComments();
}

// postComment Posts a new comment to the backend
var postComment = function(event) {
  event.preventDefault();
  var message = getMessage();
  console.log(message);
}

// getMessage Gets the entered input comment
var getMessage = function() {
  return document.getElementById('comment-input').value
}

// fillComments Fills the published comments in the html document
var fillComments = function() {
  var comments = getComments()
  document.getElementById('comments').innerHTML = mapCommentsToHtml(comments);
}

// getComments Retrieves comments from the backend
var getComments = function() {
  return [
    'Trying out kube, its cool',
    'Hayek is da bomb'
  ];
}

// mapCommentsToHtml Create
var mapCommentsToHtml = function(comments) {
  return comments.map(function(comment) {
    return createCommentPost(comment);
  }).join('');
}

// createCommentPost Creates a populated comment html div
var createCommentPost = function(message) {
  var commentDiv = [
    '<div class="comment card">',
    '<p>', message, '</p>',
    '</div>'
  ];
  return commentDiv.join('');
}
