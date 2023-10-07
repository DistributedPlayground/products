# Products
The Products service is a [REST](https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm) API written in Go and containerized with [Docker](https://www.docker.com/). It's purpose is to facilitate seller management of their products and product collections.
It is a fundamental part of our E-commerce platform within the [DistributedPlayground](https://github.com/DistributedPlayground) project. See the [project description](https://github.com/DistributedPlayground/project-description) for more details.


It allows stores to manage their products and product collections.

- [Service Architecture](#service-architecture)
- [Endpoint Description](#endpoint-description)
- [Running the Service](#running-the-service)

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
Follow the instructions in the [DistributedPlayground project description](https://github.com/DistributedPlayground/project-description).