package commands

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

type ICommand interface {
	GetCommand() *cobra.Command
	WithDatabase(db *gorm.DB) ICommand
	getMainCommand() *cobra.Command
}
