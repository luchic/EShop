# AI-Powered E-Shop Template

Hackathon Project -- buildTHEshop

## 0. Document Purpose

This document defines the product vision, core scope, and domain concepts.

It is intentionally high-level and stable.
Detailed architecture, API design, database schema, and AI implementation
will be documented in separate files later.

This file is the vision baseline and should not change frequently.

## 1. Overview

The goal of this project is to build a modern e-commerce web application
for a weekly challenge.

Unlike traditional shop templates that rely only on category navigation
or keyword search, this project introduces an AI-powered deep product
search feature. Users can describe what they are looking for in natural
language, and the AI system suggests relevant products based on intent
rather than simple keyword matching.

The system is designed as a scalable and modular web platform, allowing
easy customization and extension.

### Main Goals

-   Provide a reusable e-commerce template
-   Demonstrate AI integration in product discovery
-   Implement clean architecture suitable for scaling
-   Provide clear developer documentation and diagrams

------------------------------------------------------------------------

# 2. Actors and External Systems

## Primary Actors

-   **Guest**: Browses products and discovers items
-   **Customer**: Manages cart, completes checkout, views order history
-   **Admin**: Manages catalog, inventory, and orders

## External Systems

-   **Payment Provider**: Authorizes and confirms payments
-   **AI Service**: Interprets natural-language search and ranking
-   **Image Storage/CDN**: Serves product media assets
-   **Email Service**: Sends order confirmations and account emails

------------------------------------------------------------------------

# 3. Scope and Priorities

## MVP (Hackathon Scope)

-   Product browsing by category and details
-   Cart management and order creation
-   Checkout with mock
-   AI search over store catalog data
-   Basic admin operations for product and order management

------------------------------------------------------------------------

# 4. Features

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
-   Payment integration
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
3.  Searches product catalog data
4.  Returns ranked suggestions with short explanations

------------------------------------------------------------------------

# 5. Key User Flows

-   **Browse and Buy**: Browse catalog -> view product -> add to cart -> checkout -> order confirmation
-   **AI Discovery**: Enter natural-language query -> view ranked suggestions -> open product details -> add to cart
-   **Admin Update**: Update product or inventory -> changes reflected in catalog/search results

------------------------------------------------------------------------

# 6. Non-Functional Requirements (Vision-Level)

-   Security: Authentication required for customer account and admin operations
-   Reliability: Core purchase flow should remain available even if AI search is degraded
-   Performance: Standard catalog browsing must remain responsive for typical MVP load
-   Maintainability: Modular architecture and clear domain boundaries
-   Delivery Goal: Demo-ready MVP first, production hardening later


------------------------------------------------------------------------

## Administrative Capabilities

-   Add / edit products
-   Manage inventory
-   View orders
-   Manage categories

------------------------------------------------------------------------

# 7. Core Domain Model

The system uses these main entities:

-   **User**: A registered customer who can browse products, maintain a cart, and place orders
-   **Product**: An item available for purchase, organized within a category
-   **Category**: An organizational structure for grouping and discovering products
-   **Cart**: A temporary container for a user's selected items before checkout
-   **Order**: A completed purchase record containing ordered items and pricing

Supporting concepts expected in detailed design:

-   **Cart Item** and **Order Item** for line-level purchase details
-   **Payment** for checkout transaction outcomes
-   **Inventory Movement** for stock changes
-   **Product Attributes/Tags** to improve filtering and AI matching

Detailed entity attributes, database schema, and data types will be defined in the Architecture document.

------------------------------------------------------------------------

# 8. AI Search Concepts

The AI-powered search feature requires these data concepts:

-   **Search Query**: Captures user natural-language input and context
-   **AI Recommendation**: Stores AI-generated product suggestions with relevance scoring and explanation

Assumptions and rules (vision-level):

-   AI search uses the store catalog as its primary knowledge source
-   Results should include short user-facing explanations
-   Users may combine natural-language intent with standard filters
-   If AI processing fails, the system falls back to keyword/category search

Implementation details such as storage, indexing, and ranking algorithms will be defined in the AI Search Design document.

------------------------------------------------------------------------

# 9. Future Architecture (Planned)

This document will later be used to design:

-   System architecture diagram
-   Database ER diagram
-   API structure
-   Service interactions
-   AI search pipeline
