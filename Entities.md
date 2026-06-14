# DB Entities

---

## `users`

| Column        | Type          | Notes                 |
|---------------|---------------|-----------------------|
| `id`          | UUID / SERIAL | Primary key           |
| `first_name`  | TEXT          | Not null              |
| `second_name` | TEXT          | Not null              |
| `email`       | TEXT          | Unique, not null      |
| `password`    | TEXT          | bcrypt hash           |
<!-- | `role`        | TEXT          | `customer` \| `admin` | -->
| `created_at`  | TIMESTAMPTZ   | Default NOW()         |
<!-- | `updated_at`  | TIMESTAMPTZ   |                       | -->

---

## `products`

| Column        | Type          | Notes         |
|---------------|---------------|---------------|
| `id`          | UUID / SERIAL | Primary key   |
| `name`        | TEXT          | Not null      |
| `description` | TEXT          |               |
| `price`       | NUMERIC(10,2) | Not null      |
| `stock`       | INTEGER       | Default 0     |
<!-- | `category`    | TEXT          |               | -->
| `image_url`   | TEXT          |               |
| `created_at`  | TIMESTAMPTZ   | Default NOW() |
<!-- | `updated_at`  | TIMESTAMPTZ   |               | -->

---

## `orders`

| Column       | Type          | Notes                                                   |
|--------------|---------------|---------------------------------------------------------|
| `id`         | UUID / SERIAL | Primary key                                             |
| `user_id`    | UUID / INT    | FK → `users.id`                                         |
| `status`     | TEXT          | `pending` \| `processing` \| `completed` \| `cancelled` |
| `total`      | NUMERIC(10,2) | Price snapshot at order time                            |
| `created_at` | TIMESTAMPTZ   | Default NOW()                                           |
<!-- | `updated_at` | TIMESTAMPTZ   |                                                         | -->

---

## `order_items`

| Column       | Type          | Notes                        |
|--------------|---------------|------------------------------|
| `id`         | UUID / SERIAL | Primary key                  |
| `order_id`   | UUID / INT    | FK → `orders.id`             |
| `product_id` | UUID / INT    | FK → `products.id`           |
| `quantity`   | INTEGER       | Not null                     |
| `unit_price` | NUMERIC(10,2) | Price snapshot at order time |
| `created_at` | TIMESTAMPTZ   | Default NOW()                |
