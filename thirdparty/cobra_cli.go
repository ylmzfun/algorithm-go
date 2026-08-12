package thirdparty

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewRootCmd 创建 CLI 根命令
// 思路：根命令作为命令树的入口，挂载 greet 与 calc 两个子命令
// 作用：返回可执行命令树，调用方可用 SetArgs + Execute 运行，也可在测试中注入参数
// 业务场景：命令行工具的主入口，如运维脚本、数据迁移工具
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cli-demo",
		Short: "cobra 命令行演示工具",
		Long: `cli-demo 演示 cobra 框架的基本用法，
包含 greet（问候）与 calc（计算）两个子命令。`,
	}
	rootCmd.AddCommand(NewGreetCmd())
	rootCmd.AddCommand(NewCalcCmd())
	return rootCmd
}

// NewGreetCmd 创建 greet 子命令
// 作用：解析 --name 与 --times 两个 flag 并输出问候语，--name 为空时返回错误
// 复杂度：O(n)，n 为 --times 指定的次数
func NewGreetCmd() *cobra.Command {
	var name string
	var times int

	cmd := &cobra.Command{
		Use:   "greet",
		Short: "输出问候语",
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" {
				return errors.New("--name 不能为空")
			}
			if times < 1 {
				times = 1
			}
			for i := 0; i < times; i++ {
				fmt.Fprintf(cmd.OutOrStdout(), "Hello, %s!\n", name)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "要问候的名字")
	cmd.Flags().IntVar(&times, "times", 1, "问候次数")
	return cmd
}

// NewCalcCmd 创建 calc 子命令
// 作用：解析 --a/--b/--op 三个 flag 并输出四则运算结果，
// 除数为 0 或运算符不支持时返回错误
// 复杂度：O(1)
func NewCalcCmd() *cobra.Command {
	var a, b int
	var op string

	cmd := &cobra.Command{
		Use:   "calc",
		Short: "四则运算计算器",
		RunE: func(cmd *cobra.Command, args []string) error {
			var result int
			switch op {
			case "add":
				result = a + b
			case "sub":
				result = a - b
			case "mul":
				result = a * b
			case "div":
				if b == 0 {
					return errors.New("除数不能为 0")
				}
				result = a / b
			default:
				return fmt.Errorf("不支持的运算符: %s（可选 add/sub/mul/div）", op)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%d %s %d = %d\n", a, op, b, result)
			return nil
		},
	}
	cmd.Flags().IntVar(&a, "a", 0, "第一个操作数")
	cmd.Flags().IntVar(&b, "b", 0, "第二个操作数")
	cmd.Flags().StringVar(&op, "op", "add", "运算符: add/sub/mul/div")
	return cmd
}
