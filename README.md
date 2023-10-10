# go-survey-app
Basic survey web application using go and tailwind.

# Project Structure
survey-app/
│
├── cmd/
│   └── server/          # Entry point for the application
│       └── main.go
│
├── api/
│   ├── routes/          # Route definitions mapping to controller logic
│   ├── controllers/     # Business logic for each route (handling user actions, serving data, etc.)
│   └── middleware/      # Utility logic (authentication, logging, etc.)
│
├── models/              # GORM database models
│
├── db/                  
│   ├── migrations/      # Database migration scripts
│   └── db.go            # Database connection and initialization
│
├── static/              # Static assets
│   ├── css/             # TailwindCSS and other styles
│   ├── js/              # JavaScript files
│   └── images/          # Images, if any
│
├── templates/           # HTML templates
│
├── utils/               # Interface for shared utilities, exposing middleware functions
│   └── utils.go         # Exposes middleware functions via a Middleware global variable
│
├── vendor/              # Go dependencies (if you're vendoring)
│
├── go.mod               # Go module file
├── go.sum               # Go module checksums
│
├── .gitignore           # Ignore file for Git
├── README.md            # Project documentation
│
└── Dockerfile           # Optional: For containerization

Further Context:

cmd/server: It's a common pattern in Go projects to have the cmd directory for the application's entry point. This makes it clear where the application starts.

api: This is the core of your backend logic. It's separated into three distinct parts:

routes: Define the HTTP methods and paths and map them to specific controllers.
controllers: Handle the logic for each route, from processing input to returning responses.
middleware: Utility functions that act on requests before they reach the controllers.
models: These represent your database tables and relationships. With GORM, these will also define how Go structures map to database tables.

db: Any logic directly related to setting up and managing your database connection is here.

static: Static assets like CSS and JS files that are sent as-is to the client.

templates: HTML files with dynamic content placeholders, which your application will fill in with actual data when rendering.

utils: By introducing the global Middleware variable, any part of your application can easily access shared utility functions. This makes using shared functionality like authentication checks or logging very intuitive.

vendor: If you opt to vendor your Go dependencies (i.e., include them directly in your project rather than downloading them fresh on each build), they'll reside here.

The design aims to keep related files close to each other while offering a clear separation of concerns. It provides an intuitive layout for both developers familiar with Go's idiomatic project structures and those who might be coming from other web development backgrounds.

Generated by chatgpt lol