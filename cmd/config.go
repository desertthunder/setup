package cmd

import (
	"fmt"
	"os"
	"strings"
)

var (
	// NvimConfig represents the neovim configuration.
	NvimConfig = &ConfigType{
		Name:       "neovim",
		RepoPath:   "config/nvim",
		SystemPath: "~/.config/nvim",
		IsFile:     false,
		Excludes:   []string{},
	}

	// ZshConfig represents the zsh configuration.
	ZshConfig = &ConfigType{
		Name:       "zsh",
		RepoPath:   "config/zshrc",
		SystemPath: "~/.zshrc",
		IsFile:     true,
		Excludes:   []string{},
	}

	// AsdfConfig represents the asdf tool versions configuration.
	AsdfConfig = &ConfigType{
		Name:       "asdf",
		RepoPath:   "config/tool-versions",
		SystemPath: "~/.tool-versions",
		IsFile:     true,
		Excludes:   []string{},
	}

	// AlacrittyConfig represents the alacritty terminal configuration.
	AlacrittyConfig = &ConfigType{
		Name:       "alacritty",
		RepoPath:   "config/alacritty",
		SystemPath: "~/.config/alacritty",
		IsFile:     false,
		Excludes:   []string{".DS_Store"},
	}

	// OhMyPoshConfig represents the oh-my-posh prompt configuration.
	OhMyPoshConfig = &ConfigType{
		Name:       "oh-my-posh",
		RepoPath:   "config/omp.json",
		SystemPath: "~/.omp.json",
		IsFile:     true,
		Excludes:   []string{},
	}

	// ZshSecretsConfig represents the zsh secrets template (for reference only).
	// This is not deployed automatically to prevent overwriting user secrets.
	ZshSecretsConfig = &ConfigType{
		Name:       "zsh-secrets",
		RepoPath:   "config/zsh_secrets.templ",
		SystemPath: "~/.zsh_secrets",
		IsFile:     true,
		Excludes:   []string{},
	}

	// AllConfigs contains all available configurations.
	AllConfigs = []*ConfigType{
		NvimConfig,
		ZshConfig,
		AsdfConfig,
		AlacrittyConfig,
		OhMyPoshConfig,
	}

	// SecretConfigs contains secret configurations that require manual setup.
	SecretConfigs = []*ConfigType{
		ZshSecretsConfig,
	}
)

// GetConfigByName returns a config by its name (case-insensitive).
func GetConfigByName(name string) (*ConfigType, error) {
	nameLower := strings.ToLower(name)
	for _, config := range AllConfigs {
		if strings.ToLower(config.Name) == nameLower {
			return config, nil
		}
	}
	return nil, fmt.Errorf("unknown config: %s", name)
}

// DeployConfig deploys a config from repo to system.
func DeployConfig(configName string) error {
	config, err := GetConfigByName(configName)
	if err != nil {
		return err
	}
	return SyncConfig(config, true)
}

// DeployAllConfigs deploys all configs from repo to system.
func DeployAllConfigs() error {
	Print.NewLns(StyleInfoC, "Deploying all configurations...")
	for _, config := range AllConfigs {
		if err := SyncConfig(config, true); err != nil {
			Print.Err(fmt.Sprintf("Failed to deploy %s: %v", config.Name, err))
			return err
		}
		Print.Info()
	}
	Print.Success("All configurations deployed successfully!")
	return nil
}

// BackupConfigByName backs up a config from system to repo.
func BackupConfigByName(configName string) error {
	config, err := GetConfigByName(configName)
	if err != nil {
		return err
	}
	return SyncConfig(config, false)
}

// BackupAllConfigs backs up all configs from system to repo.
func BackupAllConfigs() error {
	Print.NewLns(StyleInfoC, "Backing up all configurations...")

	for _, config := range AllConfigs {
		if err := SyncConfig(config, false); err != nil {
			msg := fmt.Sprintf("Warning: Failed to backup %s: %v", config.Name, err)
			Print.Warn(msg)
			// Continue with other configs even if one fails
		}
		Print.Info()
	}

	Print.Success("Configuration backup completed!")
	return nil
}

// ListConfigs displays all available configurations.
func ListConfigs() error {
	Print.NewLns(StyleInfoC, "Available Configurations:")

	for _, config := range AllConfigs {
		kind := "directory"
		if config.IsFile {
			kind = "file"
		}
		msg := fmt.Sprintf("  %s %s\n", BoldMagenta(config.Name), Dim(fmt.Sprintf("(%s)", kind)))
		msg += fmt.Sprintf("    %s %s\n", Dim("Repo:"), config.RepoPath)
		msg += fmt.Sprintf("    %s %s\n", Dim("System:"), config.SystemPath)
		Print.NewLns(StyleInfo, msg)
	}
	return nil
}

// ValidateConfigs checks that all config files exist in the repo.
func ValidateConfigs() error {
	Print.NewLns(StyleInfoC, "Validating configurations...")

	isOk := true
	for _, config := range AllConfigs {
		repoPath, err := config.GetConfigPath(true)
		if err != nil {
			return err
		}

		fmt.Printf("  %s ", config.Name)
		if _, err := os.Stat(repoPath); err != nil {
			Print.Err("✗ (missing)")
			isOk = false
		} else {
			Print.Success("✓")
		}
	}

	if isOk {
		Print.Beforeln(StyleSuccess, "All configurations are present!")
		return nil
	}
	Print.Info()
	return fmt.Errorf("some configurations are missing from the repo")
}
