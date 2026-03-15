# AI-Powered E-Shop Template

Hackathon Project -- buildTHEshop

## 1. Overview

The goal of this project is to build a modern e-commerce web application
for weekend challange.

Unlike traditional shop templates that rely only on category navigation
or keyword search, this project introduces an AI-powered deep product
search feature. Users can describe what they are looking for in natural
language, and the AI system suggests relevant products based on intent
rather than simple keyword matching.

Example:

User input: "I need a lightweight laptop for programming and travel
under €1500."

AI response: - Suggests products matching weight, performance, and price
criteria - Explains why the products match the request

The system is designed as a scalable and modular web platform, allowing
easy customization and extension.

### Main Goals

-   Provide a reusable e-commerce template
-   Demonstrate AI integration in product discovery
-   Implement clean architecture suitable for scaling
-   Provide clear developer documentation and diagrams

------------------------------------------------------------------------

# 2. Features

## Core E-Commerce Features

### User Management

-   User registration
-   Login / authentication
-   User profile
-   Order history

### Product Catalog

-   Product listing
-   Product categories
-   Product details page
-   Product images and descriptions
-   Product availability

### Shopping Cart

-   Add product to cart
-   Remove product
-   Update quantity
-   Cart persistence

### Checkout

-   Order creation
-   Order summary
-   Payment integration (mock or real)
-   Order confirmation

------------------------------------------------------------------------

## AI Features

### Deep Search (AI Product Discovery)

Users can search using natural language queries instead of traditional
filters.

Example queries:

-   "Comfortable running shoes for winter"
-   "Cheap mechanical keyboard for programming"
-   "Camera for beginner vloggers"

AI processes the query and:

1.  Understands intent
2.  Extracts product requirements
3.  Searches product database
4.  Returns ranked suggestions

Possible implementation:

-   LLM API (OpenAI / Claude / local model)
-   Embedding-based semantic search
-   Hybrid search (filters + vector search)

------------------------------------------------------------------------

## Admin Features (Optional for Hackathon)

-   Add / edit products
-   Manage inventory
-   View orders
-   Manage categories

------------------------------------------------------------------------

# 3. Core Entities

These represent the main domain objects in the system.

## User

Represents a registered customer.

Attributes:

-   id
-   name
-   email
-   password_hash
-   created_at

Relationships:

-   has many Orders
-   has one Cart

------------------------------------------------------------------------

## Product

Represents an item available for purchase.

Attributes:

-   id
-   name
-   description
-   price
-   category_id
-   stock_quantity
-   created_at

Relationships:

-   belongs to Category

------------------------------------------------------------------------

## Category

Used to organize products.

Attributes:

-   id
-   name
-   description

Relationships:

-   has many Products

------------------------------------------------------------------------

## Cart

Represents a user's shopping cart.

Attributes:

-   id
-   user_id
-   created_at
-   updated_at

Relationships:

-   belongs to User
-   contains CartItems

------------------------------------------------------------------------

## CartItem

Represents a product inside a cart.

Attributes:

-   id
-   cart_id
-   product_id
-   quantity

Relationships:

-   belongs to Cart
-   references Product

------------------------------------------------------------------------

## Order

Represents a completed purchase.

Attributes:

-   id
-   user_id
-   total_price
-   status
-   created_at

Relationships:

-   belongs to User
-   contains OrderItems

------------------------------------------------------------------------

## OrderItem

Represents a purchased product inside an order.

Attributes:

-   id
-   order_id
-   product_id
-   quantity
-   price

Relationships:

-   belongs to Order
-   references Product

------------------------------------------------------------------------

# 4. AI Search Components

Additional components required for the deep search feature.

## Search Query

Stores user search input.

Attributes:

-   id
-   user_id
-   query_text
-   created_at

------------------------------------------------------------------------

## AI Recommendation Result

Stores AI-generated product suggestions.

Attributes:

-   id
-   query_id
-   product_id
-   score
-   explanation

------------------------------------------------------------------------

# 5. Future Architecture (Planned)

This document will later be used to design:

-   System architecture diagram
-   Database ER diagram
-   API structure
-   Service interactions
-   AI search pipeline
