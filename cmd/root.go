package cmd

import (
	"fmt"
	"net/url"
	"os"

	"gurl/methods"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "A http request client, written in Go!",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        rUrl := args[0]
        
        pUrl, err := url.Parse(rUrl)
        if err != nil {
            fmt.Print(err)
        }
        
        if !pUrl.IsAbs() {
            rUrl = "https://" + pUrl.String()
        }
        
		switch method {
		case "GET":
            body, err := methods.Get(rUrl)
            if err != nil {
                fmt.Print(err)
            }
            
            fmt.Print(body) 
		case "POST":
            body, err := methods.Post(rUrl, headers, data)        
            if err != nil {
                fmt.Print(err)
            }

           fmt.Print(body) 
		case "DELETE":
			fmt.Println("TODO")	
		case "PUT":
			fmt.Println("TODO")
		case "HEAD":
			fmt.Println("TODO")
		default:
			fmt.Println("Invalid method")
		}
	},
}

var method string
var headers []string
var data string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&method, "method", "X", "GET", "http request method")
    rootCmd.Flags().StringArrayVarP(&headers, "headers", "H", []string{}, "request headers")
    rootCmd.Flags().StringVarP(&data, "data", "d", "", "request data")
}


