# MySQL2CSV
<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

## Overview

MySQL2CSV is an open-source tool designed to help users export data from a MySQL database to CSV format files with concurrency support. It allows field selection, leverages concurrency for performance improvement, reduces export time significantly, and supports batch operations.


## Features

- Concurrent Export: Utilizes Go language's concurrency features to process multiple data chunks simultaneously, enhancing export performance.
- Field Selection: Allows users to choose specific fields to export, catering to personalized needs.
- Batch Operations: Supports exporting multiple tables or query results in batches for enhanced efficiency.
- High Performance: By employing concurrency and optimization techniques, reduces the time required for data export.
- User-Friendly: Simple configuration and command-line interface make the project easy to use and integrate.

## Usage

1. Get the Code: Use the following command to clone the MySQL2CSV code repository:
   ```sh
   git clone https://github.com/motongxue/MySQL2CSV.git
   ```

2. Configure Parameters: Configure MySQL database connection information, tables to export, field selection, export file paths, filenames, and whether to retain temporary tables in `config.toml`.

   ```sh
   go get github.com/yourusername/MySQL2CSV
   ```
3. Get Dependencies:
   ```sh
   go mod tidy
   ```
4. Run Export: Execute the following command in your terminal to start exporting data:

   ```sh
   go run main.go
   ```
   Alternatively, if you want to customize the path of the configuration file, you can
   ```sh
   go run main.go -f "config.toml"
   ```

## Configuration Example
```toml
# Application configuration, can be left unchanged
[app]
name = "MySQL2CSV"
thread_num = 12                         # Number of threads
batch_size = 10000                      # Batch size
output_dir = "./output/"                # Output file directory
output_file_name = "output_file_name"   # Output file name
save_tmp_file = "true"                  # Whether to save temporary files

# MySQL configuration
[mysql]
host = "127.0.0.1"              # Host
port = "3306"                   # Port
database = "db_test"            # Database
username = "root"               # Username
password = "root" # Password
table = "users"                 # Table name
columns = "name,age,email"      # Column names
```

## Project Structure
```
MySQL2CSV/
├── cmd/                
│   ├── init.go             # Initialization
├── conf/                   # Configuration files
│   ├── config.go           # Configuration file parsing
│   ├── load.go             # Load configuration file
├── models/                 # Data models
│   └── models.go           # Table definitions
├── output/                 # Output files
│   └── output.csv          # Data model definitions
├── utils/                  # Utility packages
│   ├── csv.go              # CSV file handling
│   ├── hash.go             # Hashing functions
├── config.toml             # Configuration file containing database connection and export settings
├── main.go                 # Main application entry point
├── LICENSE                 # License file
├── README.md               # Project documentation
└── README.zh-Hans.md       # Simplified Chinese version of the project documentation
```

## Project Logic
1. **Establishment of Data Connection Pool:**
At the system's initialization, we first establish a data connection pool to ensure efficient access to the required data.

2. **Querying Work Order Count:**
Based on specific query conditions, we retrieve the count of work orders that meet the criteria from the database, referred to as "total," to prepare for subsequent operations.

3. **Work Order Segmentation and Allocation:**
Based on the total count of work orders, "total," and the specified number of threads, we segment the work orders and evenly distribute them among each thread. The number of threads is determined by "thread_num."

4. **Calculation of Data Volume per Thread:**
To ensure effective data processing by each thread, we calculate the amount of data each thread needs to handle, referred to as "avg_num." This is calculated as "total/thread_num."

5. **Setting Quantity Processed per Iteration:**
We define the quantity of data each thread needs to process during each task execution. This setting is referred to as "batch_size." It determines the workload processed per iteration.

6. **Creation of Goroutines:**
Once prepared, we create "thread_num" goroutines to set the stage for subsequent task execution.

7. **Execution of CSV Export Tasks:**
Each goroutine executes an individual CSV export command task. These export tasks need to be handled in a paginated manner to ensure the orderly export of data.

8. **Calculation of Pagination Page Numbers:**
The starting page number ("pageNo") for each export command is calculated based on the goroutine's assigned number. For example, the starting pageNo for the third goroutine is calculated as "(3-1)*avg_num."

## Contribution
Feel free to raise issues, suggest enhancements, and contribute to the project. Make sure to read the contribution guidelines before submitting a Pull Request.

## License
This project is distributed under the MIT License.

---

We hope MySQL2CSV helps you efficiently export data from your MySQL database and saves you time. For any questions, suggestions, or contributions, please contact us.

Author: motongxue

Contact: https://github.com/motongxue