# Zypto

A simple Golang app to zip and encrypt your files

## Getting Started

These instructions will give you a copy of the project up and running on
your local machine for development and testing purposes.

### Prerequisites

Please make sure to have Go installed on your machine

### Installing

Install our command to the `$GOPATH/bin` directory:

    $ go install

Run `zypto`

``` 
$ zypto -h

NAME:
   zypto - zip and encrypt your files

USAGE:
   zypto [global options] command [command options] [arguments...]

COMMANDS:
   encrypt, e  encrypt a file or dir
   decrypt, d  decrypt a file or dir
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --password value, -p value  A password to use for encrypting
   --name value, -n value      name of the file to encrypt
   --help, -h                  show help (default: false)
 ```


## License

This project is licensed under the [CC0 1.0 Universal](LICENSE.md)
Creative Commons License - see the [LICENSE.md](LICENSE.md) file for
details