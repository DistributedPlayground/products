# Products
The Products service is a [REST](https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) API containerized with [Docker](https://www.docker.com/). 
It allows stores to manage their products and product collections.

- [Architecture](#architecture)
- [Endpoint Description](#endpoints)
- [Running the Service](#running)
- [Testing the Service](#testing)

## Architecture
I chose a REST architecture for this service because the stores will only be doing CRUD operations on their own products and collections. More complex queries are not necessary for this use case, so GraphQL would not be beneficial. 

### Endpoints:
| Method | Endpoint        | Description                                              | Request Body       | Path Parameters | 
|--------|-----------------|----------------------------------------------------------|--------------------|-----------------|
| GET    | `/heartbeat`    | Checks the service's heartbeat.                          | None               | None            |
| POST   | `/collection`   | Creates a new collection.                                | JSON (Collection)  | None            |
| PUT    | `/collection/:id` | Updates an existing collection by ID.                   | JSON (Collection)  | `id`: Collection ID |
| POST   | `/product`      | Creates a new product.                                   | JSON (Product)     | None            |
| PUT    | `/product/:id`  | Updates an existing product by ID.                        | JSON (Product)     | `id`: Product ID    |


#### Request Body Examples:

- JSON (Collection):
```json
{
  "name": "Spring Collection",
  "description": "New arrivals for the Spring season."
}
```

- 
```json
{
  "name": "Floral Dress",
  "description": "A beautiful floral dress.",
  "price": 29.99
}
```

# Distributed Playground
The purpose of this repo is to practice the development of distributed systems
## Project 1 - Ecommerce Platform
For my first project, I'll create an ecommerce platform. This platform will allow users to purchase goods that are maintained by us. This is *not* a two sided marketplace with our users being both companies and consumers -- at least for now. To simplify, we'll centrally define the products.

### Architecture

#### Functional Requirements
- Users should be able to make purchases using card payments
- Users should be able to see available products, including their remaining stock
- Users should be able to sort and filter products

#### Quality Attributes
- Should be designed to handle >10M users per day
- Should have 99.9% uptime
- Should allow for concurrent development from a large distributed team

#### Constraints
- The core services must be completed by 1 engineer (me) in < 1 month. 

#### Design
- We will use Golang as the primary backend language.

- We will provide an internal CP db optimized for writes in postgres
    - This db will be exposed to our systems through a Products API
- Implementing CQRS (Command Query Responsibility Segregation), we will create a AP db optimized for reads to be used for queries (Cassandra)
    - This will be updated through a service that reads from a Kafka queue

#### Services
- **API Gateway**
- **Products**: REST for future compatibility with 2 sided marketplace. postgres db, writes to kafka on upsert
- **Search**: GraphQL, cassandra db, updates from kafka
- **Inventory**: gRPC for internal use, redis k/v store
- **Orders**: Orchestrator for *Payments*, *Fufillment*, and *Notifications*, state management with kafka
- **Order Recovery**