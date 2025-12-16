# Task: In-Memory Analytics (more-or-less)

## Objective
Build an in-memory analytics component that efficiently ingests request events and answers top-N queries.

## Requirements

### 1. Data Ingestion
pImplement a method of a component that records a request event:

`record_request(user_id: string, project_id: string, model_id: string)`

Each call represents one inference request. pThe component must store this data in a structure optimized for the query patterns below.

### 2. Query Interface
pImplement methods of a component to retrieve the top N entries by request count:

* `get_top_models(n: int) -> List[(model_id, count)]`
* `get_top_projects(n: int) -> List[(project_id, count)]`
* `get_top_users(n: int) -> List[(user_id, count)]`

**(Optional if there is time left) Scoped queries (customer view)**
* `get_top_models_in_project(project_id: string, n: int) -> List[(model_id, count)]`
* `get_top_users_in_project(project_id: string, n: int) -> List[(user_id, count)]`

Results must be sorted by count in descending order.

### 3. Constraints
* All data must be held in memory (no external databases)
* Optimize for query performance - assume queries are frequent
* N will not exceed 1,000
* Choose appropriate data structures and justify your choices

## Stretch Goals (if time permits)
* Expose the component via a REST API
* Add concurrency safety for parallel ingestion
* Discuss how your design would change if N could be arbitrarily large
