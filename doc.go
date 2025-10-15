// Package main implements thunderize.
//
// # Overview
//
// Thunderize automates the setup and synchronization of development environments,
// managing configuration files, package installations, and system setup across
// different platforms by providing a unified interface for deploying dotfiles,
// installing development tools, and maintaining consistent configurations.
//
// # Installation
//
// Build the tool locally:
//
//	go build -o tmp/thunderize
//	./tmp/thunderize --help
//
// Or install it globally:
//
//	go install
//	thunderize --help
//
// # Architecture
//
// The tool is organized into several key components:
//
//   - Config Management: Synchronizes dotfiles between repository and system locations
//   - Package Management: Handles package installation via pacman, AUR, and asdf
//   - Secrets Management: Securely manages API keys and credentials
//   - System Checks: Validates system requirements and tool availability
//
// # Commands
//
// ## Configuration Management
//
// Deploy configurations from repository to system:
//
//	thunderize config deploy <name>    # Deploy specific config
//	thunderize config deploy all       # Deploy all configs
//
// Backup configurations from system to repository:
//
//	thunderize config backup <name>    # Backup specific config
//	thunderize config backup all       # Backup all configs
//
// List and validate configurations:
//
//	thunderize config list             # Show all available configs
//	thunderize config validate         # Verify configs exist in repo
//
// Available configurations:
//   - neovim:     Neovim editor configuration (~/.config/nvim)
//   - zsh:        Zsh shell configuration (~/.zshrc)
//   - asdf:       asdf version manager tool versions (~/.tool-versions)
//   - alacritty:  Alacritty terminal emulator config (~/.config/alacritty)
//   - oh-my-posh: oh-my-posh prompt theme (~/.omp.json)
//
// ## Package Installation
//
// Install packages on Arch Linux:
//
//	thunderize install pacman          # Install from official repositories
//	thunderize install aur             # Install from AUR
//	thunderize install dev             # Install development tools via asdf
//	thunderize install all             # Install everything
//
// Package lists are maintained in the packages/ directory:
//   - packages/pacman.txt: Official repository packages
//   - packages/aur.txt:    AUR packages and asdf plugins
//   - packages/dev.txt:    Language-specific dev tools (pip, cargo, npm, etc.)
//
// Supported languages and tools:
//   - Python:  black, mypy, ruff, pytest, pylsp
//   - Go:      golangci-lint, air, delve, gopls
//   - Rust:    rust-analyzer, cargo-watch, cargo-edit, cargo-audit
//   - OCaml:   merlin, dune, ocamlformat, utop
//   - .NET:    fsautocomplete, dotnet-script
//   - LSPs:    bash, yaml, json, markdown, lua, typescript, html, css
//
// ## Secrets Management
//
// Initialize secrets file from template:
//
//	thunderize secrets init
//
// This copies config/zsh_secrets.templ to ~/.zsh_secrets with secure permissions.
// The template contains placeholders for:
//   - OPENROUTER_API_KEY: AI service API key
//   - REPOFLOW_PASSWORD:  Application password
//   - Custom secrets:     Add any additional environment variables
//
// Check secrets file status:
//
//	thunderize secrets check
//
// Verifies file exists and has secure 0600 permissions (owner read/write only).
//
// Edit secrets file:
//
//	thunderize secrets edit
//
// Opens ~/.zsh_secrets in your default $EDITOR.
//
// Security best practices:
//   - Never commit ~/.zsh_secrets to version control (gitignored)
//   - Keep only the template (config/zsh_secrets.templ) in git
//   - Use 0600 permissions for secrets file
//   - Rotate credentials regularly
//
// ## System Setup
//
// Run complete system setup:
//
//	thunderize setup
//
// This executes:
//  1. System checks (validates required tools)
//  2. Package installation (all packages)
//  3. Config deployment (all configurations)
//
// Run system checks only:
//
//	thunderize check
//
// Verifies availability of required system tools.
//
// # Cross-Platform Configuration
//
// The zshrc configuration (config/zshrc) includes platform detection and
// conditional logic to work seamlessly across macOS and Arch Linux:
//
// Platform detection:
//   - Automatically detects OS via $OSTYPE
//   - Sets PLATFORM variable (mac/linux)
//   - Applies platform-specific configurations
//
// Conditional tool loading:
//   - Checks for tool existence before sourcing
//   - Prevents errors from missing tools
//   - Gracefully handles different install locations
//
// Package manager support:
//   - macOS: Homebrew (/opt/homebrew or /usr/local)
//   - Arch:  pacman, AUR helpers (yay, paru)
//
// Path management:
//   - asdf shims for version-managed tools
//   - Cargo/Rust binaries (~/.cargo/bin)
//   - Platform-specific SDK paths (Flutter, Android)
//   - Language-specific paths (opam, ghcup, dune)
//
// # Project Structure
//
//	.
//	├── main.go                  # CLI entry point and command definitions
//	├── doc.go                   # This documentation file
//	├── cmd/
//	│   ├── checks.go           # System validation checks
//	│   ├── config.go           # Configuration management
//	│   ├── packages.go         # Package installation logic
//	│   ├── printer.go          # Terminal output styling
//	│   ├── secrets.go          # Secrets management
//	│   ├── sync.go             # File synchronization (rsync)
//	│   └── utils.go            # Helper utilities
//	├── config/
//	│   ├── nvim/               # Neovim configuration
//	│   ├── alacritty/          # Alacritty terminal config
//	│   ├── zshrc               # Zsh shell configuration
//	│   ├── omp.json            # oh-my-posh prompt theme
//	│   ├── tool-versions       # asdf tool versions
//	│   └── zsh_secrets.templ # Secrets template
//	├── packages/
//	│   ├── pacman.txt          # Official repo packages
//	│   ├── aur.txt             # AUR packages
//	│   └── dev.txt             # Development tools
//	└── doc/                    # Additional documentation
//
// # Configuration System
//
// The ConfigType struct defines how configurations are managed:
//
//	type ConfigType struct {
//	    Name       string      // Display name (e.g., "neovim")
//	    RepoPath   string      // Path in repository
//	    SystemPath string      // Path on system (~ expanded)
//	    IsFile     bool        // Single file vs directory
//	    Excludes   []string    // rsync exclude patterns
//	}
//
// All configs are synchronized using rsync with:
//   - Archive mode (-a): Preserves permissions and timestamps
//   - Delete flag: Removes files not in source (directories only)
//   - Default excludes: .git, *.swp, *.swo
//   - Custom excludes: Per-config patterns
//
// # Package Installation
//
// Package installation follows this workflow:
//
//  1. Read package list from embedded filesystem
//  2. Filter out comments and empty lines
//  3. Validate package manager availability
//  4. Install packages with appropriate command:
//     - pacman: sudo pacman -S --needed
//     - AUR:    yay -S --needed (or paru)
//     - asdf:   asdf plugin add && asdf install
//     - dev:    Tool-specific installers (pip, cargo, npm, etc.)
//
// The --needed flag prevents reinstallation of up-to-date packages.
//
// # Error Handling
//
// All errors are surfaced to the user with:
//   - Colored output (red for errors, yellow for warnings)
//   - Context about what failed
//   - Suggestions for resolution
//   - Non-zero exit codes for script integration
//
// # Development
//
// Build for development:
//
//	go build -o tmp/thunderize
//
// The tmp/ directory is gitignored for development builds.
//
// Run tests:
//
//	go test ./...
//
// Code organization:
//   - Keep functions focused and single-purpose
//   - Use early returns to reduce nesting
//   - Preserve error chains with %w formatting
//   - Document all exported functions
//
// # Examples
//
// Initial system setup on new machine:
//
//	# Clone repository
//	git clone <repo-url> ~/dotfiles
//	cd ~/dotfiles
//
//	# Build tool
//	go build -o tmp/thunderize
//
//	# Run system setup
//	./tmp/thunderize setup
//
//	# Initialize secrets
//	./tmp/thunderize secrets init
//	$EDITOR ~/.zsh_secrets
//
//	# Deploy configurations
//	./tmp/thunderize config deploy all
//
// Update configurations after editing:
//
//	# Edit configs in repository
//	vim config/nvim/init.lua
//
//	# Deploy changes to system
//	./tmp/thunderize config deploy neovim
//
// Backup system configurations:
//
//	# Make changes in system location
//	vim ~/.config/nvim/lua/plugins/lsp.lua
//
//	# Backup to repository
//	./tmp/thunderize config backup neovim
//
//	# Commit changes
//	git add config/nvim
//	git commit -m "Update LSP configuration"
//
// Install new development tools:
//
//	# Add packages to appropriate list
//	echo "ripgrep" >> packages/pacman.txt
//	echo "rust-analyzer" >> packages/dev.txt
//
//	# Install new packages
//	./tmp/thunderize install pacman
//	./tmp/thunderize install dev
//
// # Environment Variables
//
// The tool respects these environment variables:
//
//   - EDITOR:     		Used for 'secrets edit' command
//   - HOME:       		User home directory (standard)
//   - ASDF_DATA_DIR: 	Custom asdf data directory (optional)
//
// # Platform-Specific Notes
//
// macOS:
//   - Homebrew must be installed (see https://brew.sh)
//   - Xcode Command Line Tools required
//   - Supports both Intel (/usr/local) and Apple Silicon (/opt/homebrew)
//
// Arch Linux:
//   - Base system with pacman required
//   - AUR helper (yay or paru) recommended
//   - sudo access needed for package installation
//
// # See Also
//
// Related documentation:
//   - doc/GUIDE.md:        Detailed setup guide
//   - doc/README.md:       Additional documentation
//   - doc/plugin-notes.md: Plugin-specific notes
//
// Configuration formats:
//   - Neovim:    	Lua configuration
//   - Alacritty: 	TOML configuration
//   - asdf:      	Simple VERSION files
//   - oh-my-posh: 	JSON theme configuration
package main
