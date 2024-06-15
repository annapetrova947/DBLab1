package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
)

type Item struct {
    ID             int
    BorrowerID     int
    FirstName      string
    LastName       string
    Surname        string
    Passport       string
    INN            string
    SNILS          string
    DriverLicense  string
    AdditionalDocs string
    Comment        string
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/add", addHandler)
    http.HandleFunc("/edit", editHandler)
    http.HandleFunc("/delete", deleteHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("sqlite3", "./database/database.db")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, borrowerID, firstName, lastName, surname, passport, inn, snils, driverLicense, additionalDocs, comment FROM my_table")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.BorrowerID, &item.FirstName, &item.LastName, &item.Surname, &item.Passport, &item.INN, &item.SNILS, &item.DriverLicense, &item.AdditionalDocs, &item.Comment); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        items = append(items, item)
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, items)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        borrowerID := r.FormValue("borrowerID")
        firstName := r.FormValue("firstName")
        lastName := r.FormValue("lastName")
        surname := r.FormValue("surname")
        passport := r.FormValue("passport")
        inn := r.FormValue("inn")
        snils := r.FormValue("snils")
        driverLicense := r.FormValue("driverLicense")
        additionalDocs := r.FormValue("additionalDocs")
        comment := r.FormValue("comment")

        db, err := sql.Open("sqlite3", "./database/database.db")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer db.Close()

        _, err = db.Exec("INSERT INTO my_table (borrowerID, firstName, lastName, surname, passport, inn, snils, driverLicense, additionalDocs, comment) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
            borrowerID, firstName, lastName, surname, passport, inn, snils, driverLicense, additionalDocs, comment)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        id := r.FormValue("id")
        borrowerID := r.FormValue("borrowerID")
        firstName := r.FormValue("firstName")
        lastName := r.FormValue("lastName")
        surname := r.FormValue("surname")
        passport := r.FormValue("passport")
        inn := r.FormValue("inn")
        snils := r.FormValue("snils")
        driverLicense := r.FormValue("driverLicense")
        additionalDocs := r.FormValue("additionalDocs")
        comment := r.FormValue("comment")

        db, err := sql.Open("sqlite3", "./database/database.db")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer db.Close()

        _, err = db.Exec("UPDATE my_table SET borrowerID=?, firstName=?, lastName=?, surname=?, passport=?, inn=?, snils=?, driverLicense=?, additionalDocs=?, comment=? WHERE id=?",
            borrowerID, firstName, lastName, surname, passport, inn, snils, driverLicense, additionalDocs, comment, id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    db, err := sql.Open("sqlite3", "./database/database.db")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM my_table WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
