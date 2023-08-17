package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/motongxue/MySQL2CSV/cmd"
	"github.com/motongxue/MySQL2CSV/conf"
	"github.com/motongxue/MySQL2CSV/models"
	"github.com/motongxue/MySQL2CSV/utils"
	"log"
	"time"
)

var (
	appConf        *conf.App
	table          *models.Table
	err            error
	hashStr        string
	dataSourceName string
)

func main() {
	// ========初始化阶段========
	err := initial()
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()

	// ========连接阶段========
	// 连接数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	// 获取数据总量
	var total int
	countSql := fmt.Sprintf("SELECT COUNT(*) FROM %s", table.TableName)
	if err := db.QueryRow(countSql).Scan(&total); err != nil {
		log.Fatal(err)
	}
	// ========计算每个携程处理的范围========
	log.Printf("total size of table %s is %d\n", table.TableName, total)
	blockSize := total / appConf.ThreadNum
	blocks := make([][2]int, appConf.ThreadNum)
	for b := 0; b < appConf.ThreadNum; b++ {
		blocks[b][0] = b * blockSize
		blocks[b][1] = b*blockSize + blockSize
		if b == appConf.ThreadNum-1 {
			blocks[b][1] = total
		}
	}

	// ========开启携程========
	errChan := make(chan error, appConf.ThreadNum)
	err = utils.ExportToCSVConcurrently(appConf, table, db, blocks, hashStr, errChan)
	if err != nil {
		log.Fatal(err)
	}

	// ========将所有文件合并========
	outputFile, err := utils.MergeToFile(appConf, table, hashStr, appConf.ThreadNum, appConf.OutputDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("final output file is %s\n", outputFile)

	end := time.Now()
	log.Printf("Time elapsed: %v\n", end.Sub(start))
}

func initial() error {
	cmd.Execute()
	mysqlConf := conf.C().MySQL
	appConf = conf.C().App
	table, err = models.NewTable()
	if err != nil {
		return err
	}
	// 取哈希防止文件名相同，并将hashStr截取前6位
	hashStr = utils.Hash(table.TableName + table.Cols)[:6]
	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlConf.UserName, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
	return nil
}
