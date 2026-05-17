package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telebot",
	Short: "SerialSnackBot — Telegram-бот з рецептами перекусів для перегляду серіалів 🍕",
	Long: `SerialSnackBot — це Telegram-бот, який пропонує швидкі рецепти
перекусів для перегляду серіалів та фільмів.

Доступні перекуси: піца, нагетси, гуакамоле, картопля фрі,
курячі крильця та мікс на компанію.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
