# Gotic

A modern static site generator written in Go.

## Overview

Gotic is a command-line tool for creating and managing static websites. It provides a simple interface to scaffold new site projects with a well-organized directory structure.

## Features

- **Project Scaffolding**: Quickly create new static site projects with organized directory structure
- **CLI Interface**: Clean command-line interface built with Cobra and BubbleTea
- **Directory Management**: Automatically creates essential directories (pages, styles, assets)
- **Go-powered**: Fast and efficient, built with Go

## Installation

### From Source

```bash
git clone https://github.com/hammadmajid/gotic.git
cd gotic
go build -o gotic ./cmd/gotic
```

### Using Go Install

```bash
go install github.com/hammadmajid/gotic/cmd/gotic@latest
```

## Usage

### Create a New Site

Create a new static site project:

```bash
# Create a site with a specific name
gotic new my-awesome-site

# Create a site with default name (my-site)
gotic new
```

This will create a new directory with the following structure:

```txt
my-awesome-site/
├── pages/      # Markdown pages and content
├── styles/     # CSS stylesheets
└── assets/     # Images, fonts, and other static assets
```

### Help

Get help for any command:

```bash
gotic --help
gotic new --help
```

## Project Structure

The generated project follows a conventional structure:

- **`pages/`** - Contains your markdown pages and content files
- **`styles/`** - CSS stylesheets for styling your site
- **`assets/`** - Static assets like images, fonts, JavaScript files

## Development

### Prerequisites

- Go 1.23.4 or later

### Building from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/hammadmajid/gotic.git
   cd gotic
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the project:

   ```bash
   go build -o gotic ./cmd/gotic
   ```

4. Run the binary:

   ```bash
   ./gotic --help
   ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Roadmap

- [ ] Add boilerplate code generation
- [ ] Interactive site creation with TUI (Bubble Tea)
- [ ] Template system for different site types
- [ ] Build and serve commands
- [ ] Theme support
- [ ] Content management features

## License

This project is open source. Please check the LICENSE file for details.

## Author

**Hammad Majid** - [@hammadmajid](https://github.com/hammadmajid)

---

*Gotic - Simple static site generation, powered by Go* 🚀
