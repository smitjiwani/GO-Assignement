## Prerequisites

- Docker
- Docker Compose
- Git

## Setup Instructions

### Option 1: Using Docker (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd GO-Assignment
```

2. Start the application using Docker Compose:
```bash
docker-compose up --build
```

3. Access the application:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

### Option 2: Manual Setup

#### Prerequisites
- Go 1.19 or higher
- Node.js 16 or higher
- npm or yarn

#### Backend Setup
1. Navigate to the server directory:
```bash
cd server
```

2. Install Go dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```
The server will start on http://localhost:8080

#### Frontend Setup
1. Open a new terminal and navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
# or
yarn
```

3. Update the API URL in `vite.config.js`:
```js
export default defineConfig({
  server: {
    proxy: {
      '/api': 'http://localhost:8080'
    }
  }
})
```

4. Start the development server:
```bash
npm run dev
# or
yarn dev
```
The frontend will be available at http://localhost:5173

## API Endpoints

- GET `/api/numbers` - Retrieve the current list of numbers
- POST `/api/numbers` - Add a new number
  - Request body: `{ "number": <integer> }`

## Testing

### Backend Tests
1. Navigate to the server directory:
```bash
cd server
```

2. Run the tests:
```bash
go test -v
```

### Integration Tests
1. Make sure both frontend and backend are running
2. Test the API endpoints using curl:
```bash
# Get current numbers
curl http://localhost:8080/api/numbers

# Add a new number
curl -X POST -H "Content-Type: application/json" \
     -d '{"number": 42}' \
     http://localhost:8080/api/numbers
```
