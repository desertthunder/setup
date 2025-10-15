package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type sysCheck struct {
	name string
	fn   func() error
}

// CheckArchLinux verifies the system is running Arch Linux by checking for arch-release file.
func CheckArchLinux() error {
	if _, err := os.Stat("/etc/arch-release"); os.IsNotExist(err) {
		return fmt.Errorf("not running on Arch Linux - /etc/arch-release not found")
	}
	return nil
}

// CheckSudoPrivileges verifies the user has sudo access without requiring a password prompt.
func CheckSudoPrivileges() error {
	cmd := exec.Command("sudo", "-n", "true")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sudo privileges required - run 'sudo -v' first or configure NOPASSWD in sudoers")
	}
	return nil
}

// CheckInternetConnectivity verifies internet connection by attempting to reach archlinux.org.
func CheckInternetConnectivity() error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://archlinux.org")
	if err != nil {
		return fmt.Errorf("no internet connectivity: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("internet connectivity check failed with status: %d", resp.StatusCode)
	}
	return nil
}

// CheckDiskSpace verifies at least 5GB of free disk space is available on the root partition.
func CheckDiskSpace() error {
	var stat syscall.Statfs_t
	if err := syscall.Statfs("/", &stat); err != nil {
		return fmt.Errorf("failed to check disk space: %w", err)
	}

	availableGB := float64(stat.Bavail*uint64(stat.Bsize)) / (1024 * 1024 * 1024)

	if availableGB < 5.0 {
		return fmt.Errorf("insufficient disk space: %.2f GB available, need at least 5 GB", availableGB)
	}
	return nil
}

// RunSystemChecks executes all system checks and reports results.
func RunSystemChecks() error {
	checks := []sysCheck{
		{"Arch Linux", CheckArchLinux},
		{"Sudo Privileges", CheckSudoPrivileges},
		{"Internet Connectivity", CheckInternetConnectivity},
		{"Disk Space", CheckDiskSpace},
	}

	Print.NewLns(StyleInfoC, "Running system checks...")
	for _, check := range checks {
		fmt.Printf("  %s ", check.name)
		if err := check.fn(); err != nil {
			Print.Err("✗")
			return fmt.Errorf("%s check failed: %w", check.name, err)
		} else {
			Print.Success("✓")
		}
	}

	Print.Beforeln(StyleSuccess, "All system checks passed!")
	return nil
}

// GetAvailableSpace returns the available disk space in GB as a formatted string.
func GetAvailableSpace() (string, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs("/", &stat); err != nil {
		return "", err
	}

	availableGB := float64(stat.Bavail*uint64(stat.Bsize)) / (1024 * 1024 * 1024)
	return strconv.FormatFloat(availableGB, 'f', 2, 64) + " GB", nil
}

// CheckCommandExists verifies if a command is available in PATH.
func CheckCommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// GetPackageManager detects which AUR helper is installed (yay or paru).
func GetPackageManager() (string, error) {
	if CheckCommandExists("yay") {
		return "yay", nil
	}
	if CheckCommandExists("paru") {
		return "paru", nil
	}
	return "", fmt.Errorf("no AUR helper found - install yay or paru first")
}

// GetInstalledPackages returns a list of explicitly installed packages.
func GetInstalledPackages() ([]string, error) {
	cmd := exec.Command("pacman", "-Qe")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get installed packages: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	packages := make([]string, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			packages = append(packages, fields[0])
		}
	}

	return packages, nil
}
