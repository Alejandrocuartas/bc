\# API Documentation

This document provides an overview of the endpoints available in the API.
The API has 2 models: cafeterias and products. They have one to many relationship.

## Base URL

The base URL for all API endpoints is: https://bctest.onrender.com

## Authentication

Authentication is not required for accessing the API endpoints.

## Endpoints

### Create a Cafeteria

- **URL:** \`/api/cafeteria\`
- **Method:** \`POST\`
- **Request Body:**
    - \`name\` (string): The name of the cafeteria.

#### Example

\`\`\`json
{
    "name": "Cafeteria Example"
}
\`\`\`

### Get All Cafeterias

- **URL:** \`/api/cafeteria\`
- **Method:** \`GET\`

### Get All Products

- **URL:** \`/api/product\`
- **Method:** \`GET\`

### Create a Product

- **URL:** \`/api/product\`
- **Method:** \`POST\`
- **Request Body:**
    - \`name\` (string): The name of the product.
    - \`price\` (int): The price of the product.
    - \`categoria_id\` (int): The category ID of the product.

#### Example

\`\`\`json
{
    "name": "Product Example",
    "price": 100,
    "categoria_id": 1
}
\`\`\`

### Get Populated Cafeterias

Gets each cafeteria with all its products

- **URL:** \`/api/cafeteria/populate\`
- **Method:** \`GET\`

## Middleware

The API includes the following middlewares:

- \`ShowMiddleware()\`: Middleware for displaying information about the request.
