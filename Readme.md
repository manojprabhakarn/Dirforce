# Directory Bruteforce Tool
This tool is used to bruteforce directories on a web server.
## Usage

To use the tool, run the following command:

` go run dirforceV1 --rhost http://example.com --wordlist /path/to/dictionary.txt `

The **--rhost** flag specifies the base URL of the server to bruteforce, and the **--wordlist** flag specifies the path to a dictionary file containing a list of directories to try.

The tool will make HTTP GET requests to each of the directories in the dictionary file and print a message if the directory was found (HTTP status code 200) or not found (any other status code).
## help
if yout have any doubt run the file with **--help** flag

` go run dirforceV1 --help `

## Options

    --rhost: The base URL of the server to bruteforce. Required.
    --wordlist: The path to a dictionary file containing a list of directories to try. Required.

## Examples

Bruteforce the directories on http://example.com using the dictionary file /tmp/dictionary.txt:

` go run dirforceV1 --rhost http://example.com --wordlist /tmp/dictionary.txt `
