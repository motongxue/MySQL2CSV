package cmd

import (
	"github.com/motongxue/MySQL2CSV/conf"
	"github.com/spf13/cobra"
)

var (
	confFilePath string
)
var RootCmd = &cobra.Command{
	Use:   "start", // 根据Use指定的start跟cmd中的start命令捆绑在一起，从而找到对应的函数
	Short: "MySQL2CSV服务",
	Long:  "MySQL2CSV服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := conf.LoadConfigFromToml(confFilePath); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confFilePath, "config-file", "f", "config.toml", "the service config from file")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
