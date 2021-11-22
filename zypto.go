package main

import (
	"fmt"
	"log"
	"os"

	"github.com/playsthisgame/zypto/commons"
	"github.com/urfave/cli/v2"
)

func main() {
	var dir string
	var password string
	var name string
	app := &cli.App{
		Name:  "zypto",
		Usage: "zip and encrypt your files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "A password to use for encrypting",
				Destination: &password,
			},
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "name of the file to encrypt",
				Destination: &name,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "encrypt a file or dir",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						dir = c.Args().First()
					}
					if dir == "" {
						fmt.Println("Include an argument")
					}
					if name == "" {
						fmt.Println("Include a name for the decrypted file")
					}
					if password == "" {
						fmt.Println("Include a password")
					}

					// Zip the dir
					commons.ZipSource(dir, name)

					// Read the zipped dir
					f, err := os.ReadFile(name + ".zip")
					if err != nil {
						fmt.Println(err)
					}

					// Create the encrypted file
					ef, err := os.Create(name)
					defer ef.Close()

					// Write contents to encrypted file
					ef.Write(commons.Encrypt(f, password))

					// Remove the zip file
					e := os.Remove(name + ".zip")
					if e != nil {
						log.Fatal(e)
					}

					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "decrypt a file or dir",
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						dir = c.Args().First()
					}
					if dir == "" {
						fmt.Println("Include an argument")
					}
					if name == "" {
						fmt.Println("Include a name for the decrypted file")
					}
					if password == "" {
						fmt.Println("Include a password")
					}

					// Read the file
					f, err := os.ReadFile(dir)
					if err != nil {
						fmt.Println(err)
					}

					// Decrypt the file
					var decrypted = commons.Decrypt(f, password)
					if decrypted != nil {
						// Create the decrypted zip file
						ef, err := os.Create(name + ".zip")
						if err != nil {
							fmt.Println(err)
						}
						defer ef.Close()
						// Write to the decrypted file
						ef.Write(decrypted)

						// Unzip file
						commons.UnzipSource(name+".zip", name)

						// Remove zip file
						e := os.Remove(name + ".zip")
						if e != nil {
							log.Fatal(e)
						}

					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
