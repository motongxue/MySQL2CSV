package utils

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/motongxue/MySQL2CSV/conf"
	"github.com/motongxue/MySQL2CSV/models"
	"os"
	"strings"
	"sync"
)

func MergeToFile(table *models.Table, hashStr string, threadNum int, outputDir string) (string, error) {
	headers := table.ColName
	// 合并到一个大文件
	outputFile := fmt.Sprintf("%s%s_%s_%s_output.csv", outputDir, table.TableName, table.Cols, hashStr)
	outFile, err := os.Create(outputFile)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()
	writer.WriteString(fmt.Sprintf("%s\n", strings.Join(headers, ",")))

	for i := 0; i < threadNum; i++ {
		// 按照顺序依次打开小的csv
		file, err := os.Open(fmt.Sprintf("%s%s_%s_%d_%s_tmp.csv", outputDir, table.TableName, table.Cols, i, hashStr))
		if err != nil {
			return "", err
		}
		defer file.Close()

		// 用扫描器开始读取文件
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			record := scanner.Text()
			if len(record) == 0 {
				continue
			}
			writer.WriteString(fmt.Sprintf("%s\n", record))
		}
	}
	return outputFile, nil
}

func ExportToCSVConcurrently(appConf *conf.App, table *models.Table, db *sql.DB, blocks [][2]int, hashStr string, errChan chan<- error) error {
	var wg sync.WaitGroup
	for i := 0; i < appConf.ThreadNum; i++ {
		wg.Add(1)
		go func(block [2]int, filename string) error {
			defer wg.Done()
			var err error // Declare an error variable
			// 开始导出csv
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()

			writer := bufio.NewWriter(file)
			defer writer.Flush()

			for start := block[0]; start < block[1]; start += appConf.BatchSize {
				end := start + appConf.BatchSize
				if end > block[1] {
					end = block[1]
				}
				querySQL := fmt.Sprintf("SELECT %s FROM users LIMIT ?, ?", table.Cols)
				rows, err := db.Query(querySQL, start, end-start)
				if err != nil {
					return err
				}
				// 处理查询结果
				for rows.Next() {
					colLen := len(table.ColName)
					record := make([]string, colLen)
					scanArgs := make([]interface{}, colLen)
					for i := range scanArgs {
						scanArgs[i] = &record[i]
					}
					if err := rows.Scan(scanArgs...); err != nil {
						return err
					}
					// TODO 可在这里对数据行进行所需要的处理
					// 采用的是带缓冲的写入
					writer.WriteString(fmt.Sprintf("%s\n", strings.Join(record, ",")))
				}
				if err := rows.Err(); err != nil {
					return err
				}
			}
			if err != nil {
				errChan <- err // Send the error to the channel
			}
			// 要加上前缀，用来区分不同的文件
			// 对table.TableName+table.Cols取md5
			return nil
		}(blocks[i], fmt.Sprintf("%s%s_%s_%d_%s_tmp.csv", appConf.OutputDir, table.TableName, table.Cols, i, hashStr))
	}
	wg.Wait()
	close(errChan)
	return nil
}
