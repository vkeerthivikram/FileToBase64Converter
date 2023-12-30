# Installation
This tool allows for converting any file into a Base64 encoded format.
# Key Features
The tool, when given a file path as a command-line argument, reads the file and encodes it using base64 encoding.
It then writes the encoded file to the same directory with the filename "base64_filename_timestamp.txt", where 'filename' represents the original file name and 'timestamp' is the Unix timestamp at the time of encoding.
Any errors encountered during the process are logged and result in program termination.
# Installation
To install `FileToBase64Converter` run the below command in your terminal.
```bash
go install github.com/vkeerthivikram/FileToBase64Converter@latest
```
# Contributing
Feel free to fork the repository and submit pull requests.
For major changes, please open an issue first to discuss what you would like to change or add.
