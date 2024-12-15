# URL Shortener Frontend

A React-based frontend application for a URL shortener service. This project provides a sleek and user-friendly interface for generating and managing short URLs.

## Features

- Input long URLs to generate short ones.
- View previously shortened URLs.
- Copy short URLs to the clipboard.
- Responsive design for mobile and desktop users.

## Technologies Used

- **React**: JavaScript library for building user interfaces.
- **Vite**: Fast development server and build tool.
- **ESLint**: Linting and code style enforcement.
- **Node.js**: Dependency management with npm.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Node.js](https://nodejs.org/) (v14 or newer)
- [npm](https://www.npmjs.com/) (usually bundled with Node.js)

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd url-shortener-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Configure environment variables:
   - Create a `.env` file in the root directory (or edit the existing one).
   - Set the backend API URL or other configuration options as required:
     ```env
     VITE_API_URL=https://your-backend-url.com
     ```

### Running the Application

1. Start the development server:
   ```bash
   npm run dev
   ```

2. Open your browser and navigate to:
   ```
   http://localhost:3000
   ```

### Building for Production

1. Build the project:
   ```bash
   npm run build
   ```

2. The production-ready files will be in the `dist` directory.

### Linting and Formatting

1. Run ESLint:
   ```bash
   npm run lint
   ```

2. Fix linting issues:
   ```bash
   npm run lint:fix
   ```

## Directory Structure

```plaintext
url-shortener-frontend/
├── public/          # Static assets
├── src/             # Application source code
│   ├── components/  # Reusable UI components
│   ├── pages/       # Page-level components
│   ├── services/    # API interaction logic
│   ├── styles/      # Global styles
│   └── main.jsx     # Application entry point
├── .env             # Environment variables
├── vite.config.js   # Vite configuration
└── README.md        # Project documentation
```

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Description of changes"
   ```
4. Push to your fork and create a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to customize this README further as needed!

