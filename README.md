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