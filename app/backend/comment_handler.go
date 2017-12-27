package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/CzarSimon/util"
)

// Comment Datastructure to hold comments
type Comment struct {
	Body string `json:"body"`
}

// IsEmpty Checks if a comment is empty
func (comment Comment) IsEmpty() bool {
	return comment.Body == ""
}

// HandleCommentRequests Handles request about comments
func (env *Env) HandleCommentRequests(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		env.getComments(res, req)
	case http.MethodPost:
		env.postComment(res, req)
	default:
		util.SendErrStatus(
			res, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

// getComments Retrieves and sends comments to the requestor
func (env *Env) getComments(res http.ResponseWriter, req *http.Request) {
	comments, err := getCommentsFromDB(env.DB)
	if err != nil {
		log.Println(err.Error())
		util.SendErrStatus(
			res, errors.New("Failed to get comments"), http.StatusInternalServerError)
		return
	}
	jsonBody, err := json.Marshal(comments)
	if err != nil {
		log.Println(err.Error())
		util.SendErrStatus(
			res, errors.New("Failed to get comments"), http.StatusInternalServerError)
		return
	}
	util.SendJSONRes(res, jsonBody)
}

// getCommentsFromDB Retrieves all comments from the database ordererd
// in reverse cronological order
func getCommentsFromDB(db *sql.DB) ([]Comment, error) {
	comments := make([]Comment, 0)
	rows, err := db.Query("SELECT BODY FROM COMMENTS ORDER BY DATE_INSERTED DESC")
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	return createCommentSlice(rows)
}

// createCommentSlice Create a slice of comments from a supplied set of rows
func createCommentSlice(rows *sql.Rows) ([]Comment, error) {
	comments := make([]Comment, 0)
	var comment Comment
	for rows.Next() {
		err := rows.Scan(&comment.Body)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// postComment Stores a sent comment to the database
func (env *Env) postComment(res http.ResponseWriter, req *http.Request) {
	var comment Comment
	err := util.DecodeJSON(req.Body, &comment)
	if err != nil || comment.IsEmpty() {
		util.SendErrStatus(
			res, errors.New("Failed to parse supplied comment"), http.StatusBadRequest)
		return
	}
	err = insertComment(comment, env.DB)
	if err != nil {
		log.Println(err.Error())
		util.SendErrStatus(
			res, errors.New("Failed to save comment"), http.StatusInternalServerError)
		return
	}
	util.SendOK(res)
}

// insertComment Inserts a supplied comment in the database
func insertComment(comment Comment, db *sql.DB) error {
	stmt, err := db.Prepare(
		"INSERT INTO COMMENTS (BODY, DATE_INSERTED) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(comment.Body, getCurrentTime())
	return err
}

// getCurrentTime Gets the current UTC timestamp
func getCurrentTime() time.Time {
	return time.Now().UTC()
}
