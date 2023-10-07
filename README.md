# Products
The Products service is a [REST](https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) API written in Go and containerized with [Docker](https://www.docker.com/). It's purpose is to facilitate seller management of their products and product collections. Is a fundamental part of our E-commerce platform within the Distributed Playground project.

It allows stores to manage their products and product collections.

- [Service Architecture](#service-architecture)
- [Endpoint Description](#endpoint-description)
- [Running the Service](#running-the-service)
- [Distributed Playground Description](#distributed-playground)

## Service Architecture
I chose a REST architecture for this service because the stores will only be doing CRUD operations on their own products and collections. More complex queries are not necessary for this use case, so GraphQL would not be beneficial. 

### Endpoint Description
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

- JSON (Product):
```json
{
  "name": "Floral Dress",
  "description": "A beautiful floral dress.",
  "price": 29.99,
  "inventory": 100,
  "collectionId": "f4059396-09fe-4bda-a620-64f14aca646d"
}
```

## Running the Service
1. Ensure that you have cloned each repository within the DistributedPlayground organization. Each repository should be cloned into the same directory so that they are parallel, forming the complete project structure.
2. Follow the instructions detailed in the [infra](https://github.com/DistributedPlayground/infra) repository to set up and run the service.

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
    - This db will be exposed to our systems through a REST API
- Implementing a psudo CQRS (Command Query Responsibility Segregation)
    - "Psudo" here because we really just want to segregrate the most read heavy users (customers) to a separate db and api optimized for their needs
    - This customer read db will be updated through a service that reads from a Kafka queue

#### Services
- **API Gateway**: *Not Implemented*
- [Products](https://github.com/DistributedPlayground/products): A REST API with a postgres db. It allows sellers to manage their products and product collections. This service publishes writes to a kafka queue.
- [Product Search](https://github.com/DistributedPlayground/product-search): A GraphQL API with mongodb. It allows customers to query the current products and collections. It reads updates from kafka to update the mongodb database.
- [Inventory](https://github.com/DistributedPlayground/inventory): A gRPC API with redis. It is for internal use only, and is intended to maintain the most up-to-date state of product inventory. It reads updates to product inventory made by sellers from kafka, and also writes updates to the product inventory made by customers through purchases.
- **Orders**: *Not Implemented* Orchestrator for *Payments*, *Fufillment*, and *Notifications*, state management with kafka
- **Order Recovery**: *Not Implemented*