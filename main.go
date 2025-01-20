package main

import (
    "database/sql"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "net/url"
    "strings"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type URLEntry struct {
    ID        int64  `json:"id"`
    LongURL   string `json:"long_url"`
    ShortCode string `json:"short_code"`
}

var db *sql.DB

func main() {
    // Initialize SQLite database
    var err error
    db, err = sql.Open("sqlite3", "./urlshortener.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create table if it doesn't exist
    createTable()

    // Serve static files (for index.html)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    // Handle URL shortening
    http.HandleFunc("/app/shorten", shortenHandler)
    
    // Handle redirects for short URLs
    http.HandleFunc("/", redirectHandler)

    log.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable() {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        long_url TEXT NOT NULL,
        short_code TEXT NOT NULL UNIQUE
    );`

    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }
}

func generateShortCode() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    const length = 6
    
    // Create a byte slice to store the result
    result := make([]byte, length)
    
    // Generate random bytes
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    
    return string(result)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    // Skip favicon.ico requests
    if r.URL.Path == "/favicon.ico" {
        return
    }

    // Skip the static directory
    if strings.HasPrefix(r.URL.Path, "/static/") {
        return
    }

    // Get the short code from the URL path
    shortCode := strings.TrimPrefix(r.URL.Path, "/")
    
    // If it's the root path, serve the index.html
    if shortCode == "" {
        http.ServeFile(w, r, "static/index.html")
        return
    }

    // Query the database for the long URL
    var longURL string
    err := db.QueryRow("SELECT long_url FROM urls WHERE short_code = ?", shortCode).Scan(&longURL)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Short URL not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    // Redirect to the long URL
    http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    longURL := r.FormValue("url")
    if longURL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    // Check if URL is valid
    _, err := url.Parse(longURL)
    if err != nil {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
        return
    }

    // Generate a unique short code
    var shortCode string
    for {
        shortCode = generateShortCode()
        // Check if code already exists
        var exists bool
        err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE short_code = ?)", shortCode).Scan(&exists)
        if err != nil {
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        if !exists {
            break
        }
    }

    // Insert into database
    _, err = db.Exec("INSERT INTO urls (long_url, short_code) VALUES (?, ?)", longURL, shortCode)
    if err != nil {
        http.Error(w, "Error creating short URL", http.StatusInternalServerError)
        return
    }

    // Return the shortened URL
    shortURL := fmt.Sprintf("http://%s/%s", r.Host, shortCode)
    w.Write([]byte(shortURL))
} 