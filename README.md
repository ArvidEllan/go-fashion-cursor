# Virtual Try-On Application

A web-based application that allows women to virtually try on clothes using their uploaded photos. The application uses AI/ML to superimpose clothing items onto user photos, providing a realistic preview of how the clothes would look.

## Features

- User authentication and profile management
- Virtual try-on using uploaded photos
- Product browsing and filtering
- Shopping cart functionality
- Try-on history
- Responsive design

## Tech Stack

### Backend
- Go (Golang)
- PostgreSQL
- Gin Web Framework
- JWT Authentication
- GORM ORM

### Frontend
- TypeScript
- React
- Material-UI
- Redux Toolkit
- Axios

## Prerequisites

- Go 1.21 or higher
- Node.js 16 or higher
- PostgreSQL 12 or higher
- Docker (optional)

## Setup

### Backend Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd virtual-tryon
```

2. Install Go dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```
PORT=8080
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=virtual_tryon
DB_PORT=5432
JWT_SECRET=your-secret-key-here
JWT_EXPIRATION=24h
```

4. Create the database:
```bash
createdb virtual_tryon
```

5. Run the backend server:
```bash
go run main.go
```

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Create a `.env` file in the frontend directory:
```
REACT_APP_API_URL=http://localhost:8080
```

4. Start the development server:
```bash
npm start
```

## Usage

1. Register a new account or log in to an existing one
2. Upload a photo of yourself
3. Browse available clothing items
4. Select an item to try on
5. View the virtual try-on result
6. Add items to cart and proceed to checkout

## API Endpoints

### Authentication
- POST /api/auth/register - Register a new user
- POST /api/auth/login - Login user

### User
- GET /api/user/profile - Get user profile
- PUT /api/user/profile - Update user profile

### Try-On
- POST /api/try-on/upload - Upload photo for try-on
- POST /api/try-on/process - Process try-on request
- GET /api/try-on/history - Get try-on history
- DELETE /api/try-on/history/:id - Delete try-on record

### Products
- GET /api/products - Get all products
- GET /api/products/:id - Get product details

### Cart
- POST /api/cart - Add item to cart
- GET /api/cart - Get cart items

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Material-UI for the frontend components
- Gin framework for the backend API
- GORM for database operations 