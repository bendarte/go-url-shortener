# Modern Web Tool Template

A lightweight, fast web application template built with Go and SQLite. Features a beautiful, responsive UI that can be adapted for various tools and utilities.

## Features

- Modern, responsive interface with gradient accents
- Clean and intuitive user experience
- SQLite database for persistent storage
- Minimal external dependencies
- Built-in security features
- Mobile-friendly design

## UI Components

- Gradient-accented headings
- Smooth animations and transitions
- Loading states and success indicators
- Copy-to-clipboard functionality
- Error handling with visual feedback
- Responsive layout for all screen sizes

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd web-tool-template
```

2. Install the SQLite driver:
```bash
go mod download
```

## Running the Service

Start the server with:
```bash
go run main.go
```

The server will start on port 8081. Access the web interface at:
```
http://localhost:8081
```

## API Structure

### 1. Process Input
- **Endpoint**: `/app/process`
- **Method**: POST
- **Form Parameters**: 
  - Customizable based on tool requirements
- **Response**: Processed result as plain text
- **Example**:
  ```bash
  curl -X POST -d "input=your-data-here" http://localhost:8081/app/process
  ```

## Database Schema

The service uses SQLite with a customizable schema:

```sql
CREATE TABLE IF NOT EXISTS entries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    input TEXT NOT NULL,
    output TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Project Structure

```
.
├── main.go           # Main server code
├── static/           # Static web files
│   └── index.html    # Web interface with modern UI
├── go.mod           # Go module file
├── go.sum           # Go module checksums
└── data.db          # SQLite database (created on first run)
```

## UI Features

The interface includes:
- Gradient backgrounds and accents
- Smooth animations
- Loading states
- Success/error notifications
- Mobile-responsive design
- Copy-to-clipboard functionality
- Form validation
- Interactive buttons and inputs

## Error Handling

The service includes proper error handling for:
- Invalid inputs
- Missing required fields
- Database errors
- Server errors
- Network issues

## Security Features

- Input validation
- SQL injection prevention
- Security headers:
  - X-Frame-Options
  - X-Content-Type-Options
  - X-XSS-Protection

## Customization

The template can be easily adapted for various tools such as:
- Text processors
- Data converters
- Formatters
- Generators
- Calculators
- Encoders/Decoders

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details.

