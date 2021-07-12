# Backend for a Restaurant Visualizer Application

BackEnd of an application for restaurant customers and transactions viewers.

## Used Technologies

- Language: Go.
- Database: Dgraph.
- API Router: ```chi```.
- Interface: Android(Kotlin).

## Endpoints

### Load data from an External API to a Local database:

- Allows loading data from an endpoint based on a specific date.
- Download, process and store in a local database. 

##### Data to be load:

1. List of products.
2. List of buyers.
3. List of transactions.

### List of All Buyers:

- List all people that have bought on the platform.
- Use a local database with loaded data, day does not matter.

### Buyers by ID:

- Buyer Information.
- Shopping history.
- Other buyers using same IP.
- Some product recommendations base on products that others people buy with the products bought on the buyer transaction.

