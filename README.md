# URL Shortener

A modern, full-stack URL shortening application with click tracking and analytics.

## Features

- **URL Shortening**: Convert long URLs into short, easy-to-share links
- **Click Tracking**: Monitor how many times your shortened links have been clicked
- **Modern UI**: Clean, responsive interface with animations and visual feedback
- **Secure**: All links are validated and sanitized

## Tech Stack

### Backend
- Go (Golang)
- SQLite database
- RESTful API

### Frontend
- React
- TypeScript
- Tailwind CSS
- Vite

## Getting Started

### Prerequisites

- Go 1.16+
- Node.js 14+
- npm or yarn

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/urlshortener.git
   cd urlshortener
   ```

2. Install backend dependencies:
   ```bash
   go mod download
   ```

3. Install frontend dependencies:
   ```bash
   cd frontend
   npm install
   ```

### Running the Application

1. Start the backend server (from the project root):
   ```bash
   go run main.go
   ```
   The server will start on http://localhost:8000

2. Start the frontend development server (from the frontend directory):
   ```bash
   cd frontend
   npm run dev
   ```
   The frontend will be available at http://localhost:5173

## API Endpoints

### Shorten URL
- **URL**: `/shorten`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "url": "https://example.com/very/long/url/that/needs/shortening"
  }
  ```
- **Response**:
  ```json
  {
    "short_url": "http://localhost:8000/abc123",
    "clicks": 0,
    "created_at": "2023-01-01T12:00:00Z"
  }
  ```

### Redirect to Original URL
- **URL**: `/:shortCode`
- **Method**: `GET`
- **Response**: Redirects to the original URL

### Health Check
- **URL**: `/health`
- **Method**: `GET`
- **Response**: Returns "OK" if the server is running

## Database Schema

The application uses SQLite with the following schema:

```sql
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    long_url TEXT NOT NULL,
    short_code TEXT NOT NULL UNIQUE,
    clicks INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Troubleshooting

### Port Already in Use

If you see an error like `listen tcp :8000: bind: address already in use`, it means another process is using port 8000. You can find and stop the process with:

```bash
# Find the process
lsof -i :8000

# Kill the process (replace PID with the actual process ID)
kill PID
```

### Frontend Can't Connect to Backend

If the frontend shows "Backend server is offline", make sure:
1. The backend server is running
2. CORS is properly configured (already set up in the code)
3. You're using the correct URL (http://localhost:8000)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Go SQLite3 Driver](https://github.com/mattn/go-sqlite3)
- [React](https://reactjs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Vite](https://vitejs.dev/) 