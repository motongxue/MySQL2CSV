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

## Contribution
Feel free to raise issues, suggest enhancements, and contribute to the project. Make sure to read the contribution guidelines before submitting a Pull Request.

## License
This project is distributed under the MIT License.

---

We hope MySQL2CSV helps you efficiently export data from your MySQL database and saves you time. For any questions, suggestions, or contributions, please contact us.

Author: motongxue

Contact: https://github.com/motongxue