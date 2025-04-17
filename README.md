# GOTH Stack Starter Template

This is a starter template for building web applications using the GOTH stack:

- **G**o - A fast, statically typed programming language
- **O**rganized Structure - Clean architecture with separation of concerns
- **T**empl - Typed HTML templating for Go
- **H**TMX - HTML-based AJAX for dynamic interfaces without writing JavaScript

## Features

- ✅ Fast and light Go-based server
- ✅ Chi router with middleware support
- ✅ Templ for type-safe HTML templating
- ✅ HTMX for interactive UI without JavaScript complexity
- ✅ Environment variable configuration
- ✅ Development and production build modes
- ✅ Embedded static file serving for production
- ✅ Error handling middleware
- 🔄 Tailwind CSS (coming soon)

## Project Structure

```
.
├── handlers/        # HTTP request handlers
├── middlewares/     # Custom middleware functions
├── public/          # Static assets (CSS, JS, images)
├── utils/           # Utility functions
├── views/           # Templ templates
│   ├── home/        # Home page templates
│   └── layout/      # Layout templates
├── main.go          # Application entry point
├── static_dev.go    # Static file handling for development
├── static_prod.go   # Static file handling for production
└── go.mod           # Go module definition
```

## How It Works

### Server Setup (main.go)

The main.go file initializes the server:
- Loads environment variables
- Sets up the Chi router
- Configures static file serving
- Registers HTTP routes
- Starts the HTTP server

### Templating (views/)

The application uses Templ for rendering HTML:
- Layout templates provide the common structure
- Content templates define page-specific content
- Templates are type-safe and compiled to Go code

### Static Files

The application handles static files differently based on the build mode:
- In development (`static_dev.go`): Serves files directly from the filesystem
- In production (`static_prod.go`): Serves files embedded in the binary

### Handlers

Handlers process HTTP requests and return responses:
- Uses Chi router for routing
- Middleware for common functionality (error handling, etc.)
- Renders HTML using Templ templates

### Frontend

The frontend utilizes:
- HTMX for dynamic content without writing JavaScript
- CSS for styling (Tailwind CSS will be integrated soon)

## Getting Started

### Prerequisites

- Go 1.20 or later
- Make sure to set up environment variables (see `.env.example`)

### Development

1. Clone the repository
   ```
   git clone https://github.com/yourusername/goth-stack-starter.git
   cd goth-stack-starter
   ```

2. Install dependencies
   ```
   go mod download
   ```

3. Create a `.env` file with the required configuration
   ```
   LISTEN_ADDR=:8080
   ```

4. Run the development server
   ```
   go run -tags dev .
   ```

5. Visit `http://localhost:8080` in your browser

### Building for Production

```
go build -o app
```

### Using HTMX

HTMX is already included in the layout. You can use HTMX attributes in your HTML to create dynamic interfaces:

```html
<button hx-post="/api/example" hx-swap="outerHTML">
  Click Me
</button>
```

## Upcoming Features

- Tailwind CSS integration
- More examples of HTMX usage
- Authentication support
- Database integration examples

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
