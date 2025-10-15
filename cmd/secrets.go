package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InitSecrets copies the secrets template to ~/.zsh_secrets if it doesn't exist.
// If the file exists, it prompts the user before overwriting.
func InitSecrets() error {
	repoPath, err := ZshSecretsConfig.GetConfigPath(true)
	if err != nil {
		return err
	}

	sysPath, err := ZshSecretsConfig.GetConfigPath(false)
	if err != nil {
		return err
	}

	if _, err := os.Stat(sysPath); err == nil {
		Print.Warn(fmt.Sprintf("Warning: Secrets file already exists at: %s", sysPath))
		fmt.Print("Overwrite existing file? (y/N): ")
		reader := bufio.NewReader(os.Stdin)
		resp, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read user input: %w", err)
		}
		resp = strings.ToLower(strings.TrimSpace(resp))
		if resp != "y" && resp != "yes" {
			Print.InfoC("Secrets initialization cancelled.")
			return nil
		}
	}

	if err := RunRsync(repoPath, sysPath, "zsh-secrets", "Initializing", true, []string{}); err != nil {
		return err
	}
	if err := os.Chmod(sysPath, 0600); err != nil {
		return fmt.Errorf("failed to set secure permissions: %w", err)
	}

	Print.NewLns(StyleSuccess, "Secrets file initialized successfully!")
	Print.NewLns(StyleInfoC, "Next steps:")
	Print.Info(fmt.Sprintf("  1. Edit the file: %s", sysPath))
	Print.Info("  2. Replace placeholder values with your actual credentials")
	Print.NewLns(StyleInfo, "  3. The file is protected with 600 permissions (owner read/write only)")
	Print.Warn("Security reminder:")
	Print.Info("  - NEVER commit ~/.zsh_secrets to version control")
	Print.Info("  - Keep only the template (config/zsh_secrets.templ) in git")

	return nil
}

// CheckSecrets verifies if the secrets file exists and has secure permissions.
func CheckSecrets() error {
	systemPath, err := ZshSecretsConfig.GetConfigPath(false)
	if err != nil {
		return err
	}

	Print.NewLns(StyleInfoC, "Checking secrets file...")

	info, err := os.Stat(systemPath)
	if os.IsNotExist(err) {
		Print.NewLns(StyleWarn, fmt.Sprintf("Secrets file not found: %s", systemPath))
		Print.Info(fmt.Sprintf("Run %s to create it from template", BoldCyan("thunderize secrets init")))
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check secrets file: %w", err)
	}

	mode := info.Mode().Perm()
	msg := fmt.Sprintf("  %s %s\n", BoldMagenta("Location:"), systemPath)
	msg += fmt.Sprintf("  %s %04o\n", BoldMagenta("Permissions:"), mode)
	Print.NewLns(StyleInfo, msg)

	if mode != 0600 {
		Print.Warn("Warning: Secrets file has insecure permissions!")
		Print.NewLns(StyleInfo, "Recommended permissions: 0600 (owner read/write only)")
		fmt.Print("Fix permissions now? (Y/n): ")

		reader := bufio.NewReader(os.Stdin)
		resp, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read user input: %w", err)
		}

		resp = strings.ToLower(strings.TrimSpace(resp))
		if resp == "" || resp == "y" || resp == "yes" {
			if err := os.Chmod(systemPath, 0600); err != nil {
				return fmt.Errorf("failed to set permissions: %w", err)
			}
			Print.Success("Permissions updated to 0600")
		}
	} else {
		Print.Success("Secrets file exists and has secure permissions!")
	}

	return nil
}

// EditSecrets opens the secrets file in the user's default editor.
func EditSecrets() error {
	sysPath, err := ZshSecretsConfig.GetConfigPath(false)
	if err != nil {
		return err
	}

	if _, err := os.Stat(sysPath); os.IsNotExist(err) {
		Print.NewLns(StyleWarn, fmt.Sprintf("Secrets file not found: %s", sysPath))
		fmt.Print("Initialize secrets file now? (Y/n): ")

		reader := bufio.NewReader(os.Stdin)
		resp, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read user input: %w", err)
		}

		resp = strings.ToLower(strings.TrimSpace(resp))
		if resp == "" || resp == "y" || resp == "yes" {
			if err := InitSecrets(); err != nil {
				return err
			}
		} else {
			return nil
		}
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	Print.NewLns(StyleInfoC, fmt.Sprintf("Opening secrets file in %s", editor))
	cmd := fmt.Sprintf("%s %s", editor, sysPath)
	return RunShellCommand(cmd)
}
