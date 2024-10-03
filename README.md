# go-osint
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Open Source](https://img.shields.io/badge/Open_Source-green?style=for-the-badge&logo=opensource&logoColor=white)
![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

A lightweight, efficient OSINT tool built in Go for username and keyword searches across multiple platforms. Inspired by Sherlock but optimized for simplicity and speed using Go routines.

## 🌟 Features
- **Username Search**: Find digital footprints across various platforms
- **Keyword Discovery**: Uncover URLs related to specific search terms
- **Domain Search**: Identify platforms associated with a domain
- **Concurrent Processing**: Leverages Go routines for fast, efficient searches
- **File Output**: Save search results for later analysis
- **Platform Specific**: Option to target searches to individual platforms

## 🚀 Quick Start
1. Clone and build:
```bash
git clone https://github.com/vehbiu/go-osint.git
cd go-osint
go build
```

2. Run a search:
```bash
go-osint username johndoe
```

## 🛠️ Usage Examples

### 📋 Command Structure
```bash
go-osint [type] [term] (platform)
```

### 🔍 Search Types
1. **All-Platform Username Search**:
   ```bash
   go-osint username johndoe
   ```

2. **Platform-Specific Search**:
   ```bash
   go-osint username johndoe github
   ```

3. **Keyword Search**:
   ```bash
   go-osint keywords johndoe.22
   ```

4. **Domain Search**:
   ```bash
    go-osint domain example.com
    ```

5. **Save Results to File**:
   ```bash
   go-osint username johndoe --output results.txt
   ```

## 📁 Project Structure
```
go-osint/
├── main.go           # Entry point
├── username/         # Username-specific search functions
├── keywords/         # Keyword-specific search functions, (currently only google seearches)
├── domain/           # Domain-specific search functions
└── README.md         # Documentation
```

## ⚙️ Technical Details
- Built in Go for maximum efficiency
- Uses Go routines for concurrent searches
- Platform-specific functions for unique handling

## 📝 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing
Contributions are welcome! Feel free to:
- Open issues
- Submit pull requests
- Suggest new platforms to add

## 👤 Author
**Vehbi**
- GitHub: [@vehbiu](https://github.com/vehbiu)

## 🔒 Responsible Use
This tool is for educational purposes only. Users are responsible for ensuring their usage complies with applicable laws and respects privacy.

## 🙏 Acknowledgments
- Inspired by [Sherlock](https://github.com/sherlock-project/sherlock)

## 📊 Stats
![GitHub stars](https://img.shields.io/github/stars/vehbiu/go-osint?style=social)
![GitHub forks](https://img.shields.io/github/forks/vehbiu/go-osint?style=social)

## 💭 Developer's Note
While this project shares similarities with Sherlock, it's different. Rather than copying sherlock, the methods are implemented manually (which can be seen from the abstraction), although not the greatest, it works for such a project.

---
Made with ❤️ by [@vehbiu](https://github.com/vehbiu)
