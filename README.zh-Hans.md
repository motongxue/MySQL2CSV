# MySQL2CSV

<div align="center">
<strong>
<samp>

[English](README.md) · [简体中文](README.zh-Hans.md)

</samp>
</strong>
</div>

## 概述

MySQL2CSV 是一个开源工具，旨在帮助用户将 MySQL 数据库中的数据以并发方式导出成 CSV
格式的文件。它支持字段选择，具有并发性能优化，可以有效地降低导出时间，并支持批量操作。

## 功能特点

- 并发导出：利用 Go 语言的并发特性，可以同时处理多个数据块，提高导出性能。
- 字段选择：允许用户选择要导出的特定字段，以满足个性化需求。
- 批量操作：支持批量导出多个表或查询结果，提高效率。
- 高性能：通过并发和优化的方法，显著降低导出数据所需的时间。
- 易于使用：简单的配置和命令行界面使得项目易于使用和集成。

## 使用方法

1. **获取代码**：使用以下命令获取 MySQL2CSV 代码库：

   ```sh
   git clone https://github.com/motongxue/MySQL2CSV.git
   ```
2. 配置参数：在 config.toml 中配置 MySQL 数据库连接信息、要导出的表、字段选择、导出文件路劲、文件名、是否保留临时表等。
3. 获取相关依赖
   ```sh
   go mod tidy
   ```
4. 运行导出：在终端中运行以下命令开始导出数据：

   ```sh
   go run main.go
   ```
   或者您想自定义配置文件的路径，则可以
   ```sh
   go run main.go -f "config.toml"
   ```

## 配置示例

```toml
# 配置应用，默认可以不更改
[app]
name = "MySQL2CSV"
thread_num = 12                         # 线程数
batch_size = 10000                      # 批量大小
output_dir = "./output/"                # 输出文件目录
output_file_name = "output_file_name"   # 输出文件名
save_tmp_file = "true"                  # 是否保存临时文件

# 配置mysql
[mysql]
host = "127.0.0.1"              # 主机
port = "3306"                   # 端口
database = "db_test"            # 数据库
username = "root"               # 用户名
password = "root" # 密码
table = "users"                 # 表名
columns = "name,age,email"      # 列名
```

## 项目结构
```
MySQL2CSV/
├── cmd/                
│   ├── init.go             # 初始化
├── conf/                   # 配置文件
│   ├── config.go           # 配置文件解析
│   ├── load.go             # 加载配置文件
├── models/                 # 数据模型
│   └── models.go           # 表定义
├── output/                 # 输出文件
│   └── output.csv          # 数据模型定义
├── utils/                  # 工具包
│   ├── csv.go              # CSV文件处理
│   ├── hash.go             # CSV文件处理
├── config.toml             # 配置文件，包含数据库连接和导出设置
├── main.go                 # 主应用程序入口
├── LICENSE                 # 许可证文件
├── README.md               # 项目说明文档
└── README.zh-Hans.md       # 项目说明文档
```

## 系统逻辑
1. **数据连接池的建立：**
在系统启动时，我们会首先建立一个数据连接池，以确保我们能够高效地访问所需的数据。

2. **工单数量的查询：**
我们会根据特定的查询条件，从数据库中获取符合条件的工单数量，将这个数量记作 total，为后续的操作做好准备。

3. **工单的分割与分配：**
基于总工单数量 total 以及设定的线程数，我们会将工单进行分割，实现均匀的分配给每个线程。这里的线程数量由 thread_num 决定。

4. **每线程数据量的计算：**
为了确保每个线程都能有效地处理数据，我们会计算出每个线程需要处理的数据量，这个数量被称为 avg_num，计算方式为 total/thread_num。

5. **每次处理的数量设定：**
我们会定义每个线程在每次执行任务时需要处理的数据数量，这个设定称为 batch_size。它决定了每次处理的工作量。

6. **Goroutine 的创建：**
在准备就绪后，我们会创建 thread_num 个 goroutine，为后续的任务执行做好准备。

7. **CSV 导出任务的执行：**
每个 goroutine 都会执行一个单独的 CSV 导出命令任务。这些导出任务需要分页处理，确保数据的有序导出。

8. **分页页码计算：**
每个导出命令的起始页码（pageNo）是根据所在 goroutine 的编号计算得出的。例如，第三个 goroutine 的初始 pageNo 为 (3-1)*avg_num。


## 贡献

欢迎提出问题和建议，以及参与到项目的贡献中来。在提交 Pull Request 之前，请确保您已经阅读了贡献指南。

## 许可证

这个项目基于 MIT 许可证 进行分发和使用。

---

希望 MySQL2CSV 能够帮助您高效地导出 MySQL 数据库中的数据并降低时间成本。如果您有任何问题、建议或贡献，请随时联系我们。

作者：motongxue

联系方式：https://github.com/motongxue