package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ConfigType represents a configuration that can be synced.
type ConfigType struct {
	Name       string   // Display name (e.g., "neovim", "zsh")
	RepoPath   string   // Path in repo (e.g., "config/nvim")
	SystemPath string   // Path on system (e.g., "~/.config/nvim")
	IsFile     bool     // true if config is a single file, false if directory
	Excludes   []string // rsync exclude patterns
}

// GetRepoRoot returns the repository root directory where the binary is located.
func GetRepoRoot() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	exeDir := filepath.Dir(exe)
	repoRoot := filepath.Dir(exeDir)
	return repoRoot, nil
}

// GetHomeDir returns the user's home directory.
func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return homeDir, nil
}

// ExpandPath expands ~ to home directory in paths.
func ExpandPath(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	if len(path) == 1 {
		return homeDir, nil
	}

	return filepath.Join(homeDir, path[1:]), nil
}

// GetConfigPath returns the full path for a config (repo or system).
func (c *ConfigType) GetConfigPath(isRepo bool) (string, error) {
	if isRepo {
		repoRoot, err := GetRepoRoot()
		if err != nil {
			return "", err
		}
		return filepath.Join(repoRoot, c.RepoPath), nil
	}
	return ExpandPath(c.SystemPath)
}

// RunRsync executes rsync to synchronize config directories or files.
func RunRsync(source, target, configName, operation string, isFile bool, excludes []string) error {
	targetPath := target
	if isFile {
		targetPath = filepath.Dir(target)
	}

	if err := os.MkdirAll(targetPath, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	Print.InfoC(fmt.Sprintf("%s %s config...", operation, configName))
	fmt.Printf("%s %s\n", Dim("Source:"), source)
	fmt.Printf("%s %s\n", Dim("Target:"), target)

	args := []string{"-av"}

	if !isFile {
		args = append(args, "--delete")
	}

	args = append(args, "--exclude=.git", "--exclude=*.swp", "--exclude=*.swo")

	for _, exclude := range excludes {
		args = append(args, "--exclude="+exclude)
	}

	if isFile {
		args = append(args, source, target)
	} else {
		args = append(args, source+"/", target+"/")
	}

	cmd := exec.Command("rsync", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync failed: %w\nOutput: %s", err, string(output))
	}

	Print.Success(fmt.Sprintf("%s config %s successfully", configName, operation))
	return nil
}

// SyncConfig synchronizes a config between repo and system.
func SyncConfig(config *ConfigType, toSystem bool) error {
	repoPath, err := config.GetConfigPath(true)
	if err != nil {
		return err
	}

	systemPath, err := config.GetConfigPath(false)
	if err != nil {
		return err
	}

	var source, target, operation string
	if toSystem {
		source = repoPath
		target = systemPath
		operation = "Deploying"
	} else {
		source = systemPath
		target = repoPath
		operation = "Backing up"
	}

	if _, err := os.Stat(source); os.IsNotExist(err) {
		return fmt.Errorf("%s config not found at %s", config.Name, source)
	}

	return RunRsync(source, target, config.Name, operation, config.IsFile, config.Excludes)
}
