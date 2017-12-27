window.onload = function() {
  document.getElementById('post-button')
    .addEventListener('click', postComment)
  document.getElementById('comment-form')
    .addEventListener('onsubmit', postComment)
  fillComments();
  getCommentInputElement().focus();
}

// postComment Posts a new comment to the backend
var postComment = function(event) {
  event.preventDefault();
  var message = getMessage();
  postRequest('/api/comment', {body: message})
    .then(function() {
      clearMessageForm();
      fillComments();
    })
    .catch(console.log);
}

// getMessage Gets the entered input comment
var getMessage = function() {
  return getCommentInputElement().value
}

// clearMessageForm Clears the comment input form
var clearMessageForm = function() {
  getCommentInputElement().value = '';
}

// getCommentInputElement Gets the comment input form dom element
var getCommentInputElement = function() {
  return document.getElementById('comment-input');
}

// fillComments Fills the published comments in the html document
var fillComments = function() {
  getComments()
    .then(function(comments) {
      populatePageWithComments(comments);
    })
    .catch(function(err) {
      console.log('Failed to get comments:', err);
    });
}

// getComments Retrieves comments from the backend
var getComments = function() {
  return getRequest('/api/comment')
          .then(checkReponse)
}

// populatePageWithComments
var populatePageWithComments = function(comments) {
  var commentsHTML = mapCommentsToHtml(comments);
  document.getElementById('comments').innerHTML = commentsHTML;
}

// mapCommentsToHtml Creates a string of comment divs from a array of comments
var mapCommentsToHtml = function(comments) {
  return comments.map(function(comment) {
    return createCommentPost(comment.body);
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

/* ----- HTTP API functions ----- */
var BACKEND_HOST = '192.168.99.100:30360';

// toURL Creates a http route by prepending protocol and backend host and port
var toURL = function(route) {
  return 'http://' + BACKEND_HOST + route;
}

// getRequest Creates and executes a get request
var getRequest = function(route) {
  return fetch(toURL(route))
          .then(checkReponse)
          .then(function(res) {
            return res.json();
          });
}

// postRequest Creates and executes a post request
var postRequest = function(route, body) {
  return fetch(toURL(route), postRequestObject(body))
          .then(checkReponse);
}

// checkReponse Checks if a response has a 200 status
var checkReponse = function(response) {
  if (response.status != 500) {
    return response;
  } else {
    var error = new Error(response.statusText);
    error.response = response;
    throw error;
  }
}

// postRequestObject Creates a object to to do a post request
var postRequestObject = function(body) {
  return makeRequestObject('POST', body);
}

// makeRequestObject Creates a request object to be used by the FETCH api
var makeRequestObject = function(method, body) {
  return {
    method,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
    mode: 'no-cors',
    body: JSON.stringify(body)
  };
}
