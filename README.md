# MySQL Benchmark Test for Unread Messages

This Go application benchmarks the performance of fetching the count of unread messages from a MySQL database.

## Overview

The code connects to a MySQL database and measures the time it takes to fetch the count of unread messages for a specific user. The benchmarking process considers different indexing strategies to optimize query performance.

### Indexing Performance

- Without Index: approximately 200ms
- Index on `touser`: approximately 150ms
- Index on `touser` and `timestamp`: approximately 120ms
- Index on `touser`, `timestamp`, and `fromuser`: approximately 100ms

## Setup and Usage

### Prerequisites

- Go 1.16+
- MySQL server
- `github.com/go-sql-driver/mysql` package

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Aman123at/benchmark-unread-messages.git
   cd benchmark-unread-messages
   ```
2. Install the MySQL driver:
  ```bash
   go get -u github.com/go-sql-driver/mysql
  ```
3. Update the MySQL connection string in the `NewConn` function:
   ```bash
   db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/indicator")
   ```

### Running the Benchmark

1. Run the application to benchmark the unread message count query:
   ```bash
   go run main.go
   ```
2. The output will display the time taken to execute the query with the current indexing strategy.

### Bulk Insert (Optional)

To insert 1 million dummy messages into the database for testing:
1. Uncomment the `bulkInsertMessagesIntoSQL()` function call in the `main` function.
2. Run the application:
   ```bash
   go run main.go
   ```

### Code Structure

- **`main.go`**: The entry point of the application.

  - **`main()`**: 
    - Starts the benchmark test by calling `fetchUnreadMessageCount()`.
    - Contains a commented-out line to trigger the bulk insertion of messages.

  - **`NewConn()`**:
    - Establishes a connection to the MySQL database.
    - Returns a pointer to the `sql.DB` object.

  - **`fetchUnreadMessageCount()`**:
    - Connects to the database.
    - Executes a query to count the number of distinct unread messages for a specific user.
    - Logs the time taken to execute the query.
    - Compares performance with and without indexing on the relevant columns.

  - **`bulkInsertMessagesIntoSQL()`** *(Commented)*:
    - Inserts 1 million rows of dummy messages into the `messages` table.
    - Uses a prepared statement to optimize the insertion process.
    - Contains helper functions:
      - **`getRandomFromAndToUser()`**: Generates random `fromuser` and `touser` combinations.
      - **`getRandomMsg()`**: Selects a random message from a predefined list.
      - **`getCurrentTimeStamp()`**: Returns the current timestamp in `YYYY-MM-DD HH:MM:SS` format.

- **Helper Functions** *(Commented)*:
  - **`getRandomFromAndToUser()`**: Generates random users to simulate message exchanges.
  - **`getRandomMsg()`**: Returns a random message from a list of dummy messages.
  - **`getCurrentTimeStamp()`**: Fetches the current timestamp for inserting into the database.


