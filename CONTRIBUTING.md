# Contributing to Ruuvi

Thank you for your interest in contributing to the Ruuvi Go module! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Code Style](#code-style)
- [Submitting Changes](#submitting-changes)

## Code of Conduct

Please be respectful and professional in all interactions. We are committed to providing a welcoming and inclusive environment for all contributors.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- golangci-lint (for linting)

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/ruuvi.git
   cd ruuvi
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/marcgeld/ruuvi.git
   ```
4. Install dependencies:
   ```bash
   go mod download
   ```

## Development Workflow

### Building

Build the entire project:
```bash
make build
# or
go build ./...
```

Build the CLI tool:
```bash
cd cmd/ruuvi
go build .
```

### Testing

Run all tests:
```bash
make test
# or
go test -v -race ./...
```

Run tests with coverage:
```bash
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
```

### Linting

Run linters:
```bash
make lint
# or
golangci-lint run --timeout=5m
```

### Code Formatting

Format your code:
```bash
go fmt ./...
goimports -w -local github.com/marcgeld/ruuvi .
```

## Testing

- Write tests for all new functionality
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Follow existing test patterns in the codebase
- Test both success and error cases
- Include tests based on the official Ruuvi test vectors when applicable

Example test structure:
```go
func TestDecodeFormat5(t *testing.T) {
    tests := []struct {
        name    string
        input   []byte
        want    *Format5Data
        wantErr bool
    }{
        // test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := DecodeFormat5(tt.input)
            // assertions...
        })
    }
}
```

## Code Style

### Go Standards

- Follow standard Go conventions and idioms
- Run `gofmt` and `goimports` before committing
- Use meaningful variable and function names
- Write clear, concise comments for exported functions and types
- Keep functions focused and reasonably sized

### Project-Specific Guidelines

- **Strict parsing**: Always validate data format and length
- **Type safety**: Use type-safe structs for sensor readings
- **Nil handling**: Use pointers for optional fields; `nil` indicates invalid/unavailable data
- **Error handling**: Return descriptive errors with context using `fmt.Errorf()`
- **Binary data**: Use `encoding/binary` with `binary.BigEndian` for RuuviTag data

### Package Organization

```
ruuvi/
├── common/          # Shared types and utilities
└── tag/             # RuuviTag format decoders/encoders
```

## Submitting Changes

### Pull Request Process

1. **Create a branch**: Create a feature branch from `main`
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. **Make your changes**: Implement your feature or fix
   - Write clear, focused commits
   - Include tests for new functionality
   - Update documentation as needed

3. **Test your changes**:
   ```bash
   make test
   make lint
   go mod tidy
   ```

4. **Commit your changes**:
   ```bash
   git add .
   git commit -m "Add feature: description of changes"
   ```
   
   Use clear, descriptive commit messages:
   - Use present tense ("Add feature" not "Added feature")
   - Be specific about what changed
   - Reference issue numbers if applicable

5. **Push to your fork**:
   ```bash
   git push origin feature/my-new-feature
   ```

6. **Open a Pull Request**:
   - Go to the original repository on GitHub
   - Click "New Pull Request"
   - Select your fork and branch
   - Fill out the PR template with:
     - Description of changes
     - Related issue numbers
     - Testing performed
     - Breaking changes (if any)

### PR Requirements

- All tests must pass
- Code must pass linting (`golangci-lint`)
- `go mod tidy` must not produce any changes
- Maintain or improve code coverage
- Update documentation if needed
- Add entries to CHANGELOG.md for notable changes

## Reporting Issues

When reporting issues, please include:

- Go version (`go version`)
- Operating system and architecture
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Relevant code snippets or error messages

## Questions?

If you have questions about contributing, feel free to:

- Open an issue with the "question" label
- Check existing issues and pull requests
- Review the README.md for project details

## License

By contributing to this project, you agree that your contributions will be licensed under the same MIT License that covers the project. See the [LICENSE](LICENSE) file for details.

Thank you for contributing to Ruuvi! 🎉
