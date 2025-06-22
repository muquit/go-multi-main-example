## Introduction

This project demonstrates the enhanced @GO_XBUILD_GO@ **v1.0.5+** 
functionality with support for multiple main packages and JSON 
configuration.

## Project Structure

```
go-multi-main-example/
├── cmd/
│   ├── cli/main.go          # CLI application
│   └── server/main.go       # HTTP server application
├── build-config.json        # Multi-target build configuration
├── platforms.txt            # Supported platforms
├── VERSION                  # Version file
├── go.mod                   # Go module
└── README.md                # This file
```

## Applications

### CLI Tool (`cmd/cli`)
A command-line interface tool with the following features:
- Version information display
- Configuration file support
- Multiple commands (process, status, help)
- Verbose logging option

**Usage:**
```bash
./example-cli --version
./example-cli --verbose process
./example-cli status
```

### Server (`cmd/server`)
An HTTP server application with:
- Health check endpoint (`/health`)
- API information endpoint (`/api/info`)
- Graceful shutdown handling
- Configurable host and port

**Usage:**
```bash
./example-server --port 8080 --host localhost
./example-server --version
```

**Endpoints:**
- `GET /` - Welcome page
- `GET /health` - Health check (JSON)
- `GET /api/info` - Service information (JSON)

## Building manually
```
go build -o example-cli cmd/cli/main.go
go build -o example-server cmd/cli/main.go
```

## Building with go-xbuild-go

### Prerequisites
Download pre-build binary from @GO_XBUILD_GO@ project

### Single Binary (Legacy Mode)
Build individual binaries the traditional way:
```bash
# Traditional single binary build (uses main.go in project root)
go-xbuild-go

# With additional files
go-xbuild-go -additional-files "foo.txt,bar.txt"
```

**Note:**
 The following files are automatically included if they exist:
 README.md, LICENSE.txt, LICENSE, platforms.txt, <project>.1
 Do not specify these files in -additional-files as they will conflict.

### Multi-Binary Configuration Mode (This project)
Use the JSON configuration for unified builds:

```bash
# List available targets
go-xbuild-go -config build-config.json -list-targets

# Build all targets for all platforms
go-xbuild-go -config build-config.json

..

```

### Build Output

The build process creates:
```
bin/
├── example-cli-v1.0.1-checksums.txt
├── example-cli-v1.0.1-linux-amd64.d.tar.gz
├── example-cli-v1.0.1-raspberry-pi-jessie.d.tar.gz
├── example-cli-v1.0.1-raspberry-pi.d.tar.gz
├── example-cli-v1.0.1-windows-amd64.d.zip
├── example-server-v1.0.1-checksums.txt
├── example-server-v1.0.1-linux-amd64.d.tar.gz
├── example-server-v1.0.1-raspberry-pi-jessie.d.tar.gz
├── example-server-v1.0.1-raspberry-pi.d.tar.gz
└── example-server-v1.0.1-windows-amd64.d.zip
```

Each archive contains:
- The compiled binary
- Additional files (if specified)

Example:
```
# unzip -l bin/example-cli-v1.0.1-windows-amd64.d.zip
Archive:  bin/example-cli-v1.0.1-windows-amd64.d.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
     1076  06-22-2025 17:11   example-cli-v1.0.1-windows-amd64.d/LICENSE
     7289  06-22-2025 17:11   example-cli-v1.0.1-windows-amd64.d/README.md
  1715200  06-22-2025 17:11   example-cli-v1.0.1-windows-amd64.d/example-cli-v1.0.1-windows-amd64.exe
      908  06-22-2025 17:11   example-cli-v1.0.1-windows-amd64.d/platforms.txt
---------                     -------
  1724473                     4 files
```

## Configuration Features Demonstrated

### JSON Configuration (`build-config.json`)
- **Multiple targets**: CLI and server applications with different paths
- **Output naming**: Custom binary names per target
- **Shared settings**: Default ldflags and build flags for all targets
- **File references**: Points to VERSION and platforms.txt files

### Build Variables
The configuration demonstrates variable substitution in ldflags:
- `{{.Version}}` - From VERSION file
- `{{.Commit}}` - From git commit hash  
- `{{.Date}}` - Build timestamp

Look at `cmd/cli/main.go`
```
./example-cli -version
Example CLI Tool
Version: dev
Commit: xyz
Built: Sun Jun 22 17:16:23 EDT 2025
```

### Legacy Compatibility
- **platforms.txt**: Standard platform definitions file
- **VERSION file**: Simple version management
- **Additional files**: Can be specified via command line in legacy mode

## Development

### Local Development
```bash
# Run CLI locally
go run ./cmd/cli --help

# Run server locally
go run ./cmd/server --port 8080

# Test server endpoints
curl http://localhost:8080/health
curl http://localhost:8080/api/info
```

### Testing the Build
```bash
# Quick build test
go-xbuild-go -config build-config.json

# List what targets would be built
go-xbuild-go -config build-config.json -list-targets

# Extract and test
cd dist
tar -xzf example-cli-1.0.0-linux-amd64.tar.gz
./example-cli --version

tar -xzf example-server-1.0.0-linux-amd64.tar.gz  
./example-server --version
```

## Migration Guide

### From Single Binary to Multi-Binary

If you have an existing project using go-xbuild-go v1.0.4:

1. **Create `cmd/` structure**:
   ```bash
   mkdir -p cmd/myapp
   mv main.go cmd/myapp/
   ```

2. **Create `build-config.json`**:
   ```json
   {
     "project_name": "myapp",
     "version_file": "VERSION",
     "platforms_file": "platforms.txt", 
     "targets": [{
       "name": "myapp",
       "path": "./cmd/myapp",
       "output_name": "myapp"
     }]
   }
   ```

3. **Update build commands**:
   ```bash
   # Old way (still works)
   go-xbuild-go
   
   # New way
   go-xbuild-go -config build-config.json
   ```

### Backward Compatibility
go-xbuild-go v1.0.5 maintains full backward compatibility:
```bash
# This still works exactly as before
go-xbuild-go

# With additional files
go-xbuild-go -additional-files "new_file.txt,foo.txt"

# Create GitHub release
go-xbuild-go -release -release-note "New version"
```

**Note:**
 The following files are automatically included if they exist:
 README.md, LICENSE.txt, LICENSE, platforms.txt, <project>.1
 Do not specify these files in -additional-files as they will conflict.

## Best Practices

### Project Organization
- Use `cmd/` directory for multiple main packages
- Keep shared code in internal packages
- Include version information in all binaries

### Configuration Management
- Use descriptive target names
- Include comprehensive metadata
- Document additional files and their purposes

### Cross-Platform Considerations
- Test key platforms during development
- Consider platform-specific build flags
- Include platform-appropriate additional files

## License

MIT License - see LICENSE file for details.

## Authors

Developed with @CLAUDE@, working under my guidance and instructions.
