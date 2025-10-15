package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GetRepoDir returns the repository directory containing nvim config.
// It resolves the path to <repo>/config/nvim where this binary is located.
func GetRepoDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	// Get the directory containing the executable
	exeDir := filepath.Dir(exe)

	// Go up to the repo root (assuming exe is in tmp/ or similar)
	repoRoot := filepath.Dir(exeDir)

	// Build path to config/nvim
	nvimConfigPath := filepath.Join(repoRoot, "config", "nvim")

	return nvimConfigPath, nil
}

// GetNvimConfig returns the nvim config directory path (~/.config/nvim).
func GetNvimConfig() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "nvim"), nil
}

// RunRsync executes rsync to synchronize config directories.
// It ensures the target directory exists and runs rsync with sensible flags.
func RunRsync(source, target, operation string) error {
	// Ensure target directory exists
	if err := os.MkdirAll(target, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	fmt.Println(BoldCyan(fmt.Sprintf("%s neovim config...", operation)))
	fmt.Printf("%s %s\n", Dim("Source:"), source)
	fmt.Printf("%s %s\n", Dim("Target:"), target)

	cmd := exec.Command("rsync",
		"-av",
		"--delete",
		"--exclude=.git",
		"--exclude=*.swp",
		"--exclude=*.swo",
		source+"/",
		target+"/",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync failed: %w\nOutput: %s", err, string(output))
	}

	fmt.Println(BoldGreen(fmt.Sprintf("Config %s successfully", operation)))
	return nil
}

// UpdateConfig syncs nvim config from repo to ~/.config/nvim.
func UpdateConfig() error {
	repoDir, err := GetRepoDir()
	if err != nil {
		return err
	}

	nvimConfig, err := GetNvimConfig()
	if err != nil {
		return err
	}

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		return fmt.Errorf("%s Repository config not found at %s", BoldRed("Error:"), repoDir)
	}

	return RunRsync(repoDir, nvimConfig, "Updating")
}

// BackupConfig syncs currently installed nvim config from ~/.config/nvim to repo.
func BackupConfig() error {
	repoDir, err := GetRepoDir()
	if err != nil {
		return err
	}

	nvimConfig, err := GetNvimConfig()
	if err != nil {
		return err
	}

	if _, err := os.Stat(nvimConfig); os.IsNotExist(err) {
		return fmt.Errorf("%s Nvim config not found at %s", BoldRed("Error:"), nvimConfig)
	}

	return RunRsync(nvimConfig, repoDir, "Backing up")
}
