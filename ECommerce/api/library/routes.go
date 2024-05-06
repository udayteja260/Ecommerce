package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/add", AddHandler(db))
	http.HandleFunc("/update", UpdateHandler(db))
	http.HandleFunc("/delete", DeleteHandler(db))
	http.HandleFunc("/list", ListHandler(db))

}
