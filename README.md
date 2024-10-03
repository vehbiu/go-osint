# go-osint
`go-osint` is a lightweight, easy-to-use tool designed for Open Source Intelligence (OSINT) tasks, specifically for username and keyword searches across various platforms. Built in Go and leveraging Go routines for efficiency, this tool serves as a simplified alternative to [Sherlock](https://github.com/sherlock-project/sherlock), making it accessible for users of all skill levels.

## Features

- **Username Search**: Find usernames across multiple platforms.
- **Keyword Search**: Discover URLs related to specific keywords.
- **Output to File**: Easily save search results to a specified output file.

## Installation
To get started, clone the repository and build the project:

```bash
git clone https://github.com/vehbiu/go-osint.git
cd go-osint
go build
```

## Usage

### Quick Start

To run `go-osint`, use the following command structure:

```bash
go-osint [type] [term] (platform)
```

- **`[type]`**: Specify the type of search:
    - `"username"` for username searches
    - `"keywords"` for keyword searches
- **`[term]`**: The username, domain, or email you want to search for.
- **`(platform)`**: (Optional) Specify the platform to search for the username on.

### Examples

1. **Search for a username across all platforms**:
   ```bash
   go-osint username johndoe
   ```

2. **Search for a username on a specific platform**:
   ```bash
   go-osint username johndoe github
   ```

3. **Search for keywords**:
   ```bash
   go-osint keywords johndoe.22
   ```

4. **Save results to a file**:
   ```bash
   go-osint username johndoe --output results.txt
   ```

### Output
Search results will be displayed in the terminal, and if an output file is specified, the results will also be saved there.

## Contribution
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
`go-osint` is heavily inspired by [Sherlock](https://github.com/sherlock-project/sherlock). Special thanks to the original authors for their work in making OSINT accessible to everyone.

The authors or contributors of this project are not responsible for any misuse of the tool. Use responsibly and respect the privacy of others. Do not use this tool for illegal activities.


## Side-note / Rant
I did not want to be a 1-1 copy of Sherlock, hence, that's why all the platforms have their own functions. Although this is not the greatest feature for abstraction, it was to add a bit of uniqueness to the project. However, it should not matter much due to the nature and size of the project.