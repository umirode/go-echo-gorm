package Command

import (
	"github.com/spf13/cobra"
)

type ICommand interface {
	GetCommand() *cobra.Command
	getMainCommand() *cobra.Command
}
