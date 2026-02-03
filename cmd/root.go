/*
Copyright © 2026 Mirac Bedir <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

var (
	Length  int
	Symbols int
	Numbers int
	Name    string
	ID      int
	NewPass string
)

var DoSave bool = false
var ListAll bool = false
var DoDelete bool = false
var DoUpdate bool = false
var AutoGenerateNewPass bool = false

var (
	DB     *sql.DB
	DbPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "./pwgen",
	Short: "A simple password generator",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		saveFlag, _ := cmd.Flags().GetBool("save")
		listFlag, _ := cmd.Flags().GetBool("vault")
		deleteFlag, _ := cmd.Flags().GetBool("remove")
		updateFlag, _ := cmd.Flags().GetBool("update")
		if saveFlag || listFlag || deleteFlag || updateFlag {
			setupDatabase()
		}

	},
	Run: func(cmd *cobra.Command, args []string) {

		err := validateInputs()
		if err != nil {
			fmt.Printf("Validation failed: %v\n", err)
			os.Exit(1)
		}

		Password := PassGen()

		if DoSave && ListAll {
			fmt.Println("You cannot save and list at the same time!")
			os.Exit(1)
		}

		if DoSave {
			data := PasswordData{
				Appname:  Name,
				Password: Password,
			}
			Save(data)
			fmt.Printf("Password: %v is saved\n", Password)
		} else if ListAll {
			List()
		} else if DoDelete {
			if ID < 1 {
				fmt.Println("You didn't enter a valid id")
				os.Exit(1)
			}
			Delete(ID)

		} else if DoUpdate {
			if ID < 1 {
				fmt.Println("You didn't enter a valid id")
				os.Exit(1)
			}
			if AutoGenerateNewPass {
				Update(ID, Password)
				fmt.Printf("Your new password is %v for id %d.\n", Password, ID)
			} else {
				if NewPass == " " {
					fmt.Println("You probably didn't enter a new password. If you did, password cannot be am empty string")
					os.Exit(1)
				}
				Update(ID, NewPass)
			}

		} else {
			fmt.Println(Password)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if DB != nil {
			DB.Close()
			// fmt.Println("Connection closed.") // uncomment for debug
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var currentTime = time.Now().Format("2006-01-02 15:04:05")

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passwordGenerator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntVarP(&Length, "length", "l", 16, "Length of the generated password")
	rootCmd.Flags().IntVarP(&Symbols, "symbol", "c", 5, "Include special characters (e.g. @, #, $)")
	rootCmd.Flags().IntVarP(&Numbers, "digit", "d", 5, "Include numbers (0-9)")
	rootCmd.Flags().BoolVarP(&DoSave, "save", "s", false, "Would you like to save the password created")
	rootCmd.Flags().StringVarP(&Name, "name", "n", currentTime, "Enter the name of the password belongs to")
	rootCmd.Flags().BoolVarP(&ListAll, "vault", "v", false, "Would you like to see your password list")
	rootCmd.Flags().BoolVarP(&DoDelete, "remove", "r", false, "Would you like to delete a password")
	rootCmd.Flags().IntVarP(&ID, "id", "i", 0, "ID for passwords which is used to delete or update the passwords with -d, and -u flags ")
	rootCmd.Flags().BoolVarP(&DoUpdate, "update", "u", false, "Would you like to update a password")
	rootCmd.Flags().BoolVarP(&AutoGenerateNewPass, "auto", "a", false, "Would you like to update a password with auto generated password")
	rootCmd.Flags().StringVarP(&NewPass, "new-pass", "p", " ", "Your new password (cannot be an emptystring)")

}

func setupDatabase() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error locating home directory:", err)
		return
	}

	// Define the application config directory path
	// This points to ~/.config/password-gen
	appDir := filepath.Join(homeDir, ".config", "password-gen")

	// Create the directory if it doesn't exist
	// os.MkdirAll creates the folder and any missing parents
	// 0755: Owner can read/write/execute, others can read/execute
	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	DbPath = filepath.Join(appDir, "data.db")

	DB, err = sql.Open("sqlite", DbPath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS passwords (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    appname TEXT NOT NULL,
	    password TEXT NOT NULL
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}
