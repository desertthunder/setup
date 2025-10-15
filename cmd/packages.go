package cmd

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ReadPackageList reads a package list file and returns non-empty, non-comment lines.
func ReadPackageList(fs embed.FS, filename string) ([]string, error) {
	data, err := fs.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}

	var packages []string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		packages = append(packages, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading package list: %w", err)
	}

	return packages, nil
}

// FilterInstalledPackages removes already-installed packages from the list.
func FilterInstalledPackages(packages []string) ([]string, error) {
	installed, err := GetInstalledPackages()
	if err != nil {
		return nil, err
	}

	installedMap := make(map[string]bool)
	for _, pkg := range installed {
		installedMap[pkg] = true
	}

	var toInstall []string
	for _, pkg := range packages {
		if !installedMap[pkg] {
			toInstall = append(toInstall, pkg)
		}
	}

	return toInstall, nil
}

// InstallPacmanPackages installs packages using pacman.
func InstallPacmanPackages(fs embed.FS) error {
	Print.NewLns(StyleInfoC, "Installing pacman packages...")

	packages, err := ReadPackageList(fs, "packages/pacman.txt")
	if err != nil {
		return err
	}

	fmt.Printf("%s Found %d packages in list\n", Dim("→"), len(packages))

	toInstall, err := FilterInstalledPackages(packages)
	if err != nil {
		return err
	}

	if len(toInstall) == 0 {
		Print.Success("All pacman packages already installed!")
		return nil
	}

	fmt.Printf("%s Installing %d new packages...\n\n", Dim("→"), len(toInstall))

	args := append([]string{"-S", "--needed", "--noconfirm"}, toInstall...)
	cmd := exec.Command("sudo", append([]string{"pacman"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pacman installation failed: %w", err)
	}

	Print.NewLns(StyleSuccess, "Pacman packages installed successfully!")
	return nil
}

// InstallAURHelper installs yay or paru if not already installed.
func InstallAURHelper() error {
	if CheckCommandExists("yay") || CheckCommandExists("paru") {
		return nil
	}

	Print.NewLns(StyleInfoC, "Installing AUR helper (yay)...")

	prereqs := []string{"base-devel", "git"}
	cmd := exec.Command("sudo", "pacman", "-S", "--needed", "--noconfirm", prereqs[0], prereqs[1])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install prerequisites: %w", err)
	}

	tmpDir := "/tmp/yay-install"
	if err := os.RemoveAll(tmpDir); err != nil {
		return fmt.Errorf("failed to clean temp directory: %w", err)
	}

	cloneCmd := exec.Command("git", "clone", "https://aur.archlinux.org/yay.git", tmpDir)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		return fmt.Errorf("failed to clone yay: %w", err)
	}

	buildCmd := exec.Command("makepkg", "-si", "--noconfirm")
	buildCmd.Dir = tmpDir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build yay: %w", err)
	}

	Print.NewLns(StyleSuccess, "AUR helper installed successfully!")
	return nil
}

// InstallAURPackages installs packages from AUR using yay or paru.
func InstallAURPackages(fs embed.FS) error {
	Print.NewLns(StyleInfoC, "Installing AUR packages...")
	if err := InstallAURHelper(); err != nil {
		return err
	}

	aurHelper, err := GetPackageManager()
	if err != nil {
		return err
	}

	fmt.Printf("%s Using %s as AUR helper\n", Dim("→"), aurHelper)

	packages, err := ReadPackageList(fs, "packages/aur.txt")
	if err != nil {
		return err
	}

	fmt.Printf("%s Found %d packages in list\n", Dim("→"), len(packages))

	toInstall, err := FilterInstalledPackages(packages)
	if err != nil {
		return err
	}

	if len(toInstall) == 0 {
		Print.Success("All AUR packages already installed!")
		return nil
	}

	fmt.Printf("%s Installing %d new packages...\n\n", Dim("→"), len(toInstall))

	args := append([]string{"-S", "--needed", "--noconfirm"}, toInstall...)
	cmd := exec.Command(aurHelper, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("AUR installation failed: %w", err)
	}

	Print.NewLns(StyleSuccess, "AUR packages installed successfully!")
	return nil
}

// InstallAsdf installs asdf-vm if not already installed.
func InstallAsdf() error {
	homeDir, err := GetHomeDir()
	if err != nil {
		return err
	}

	asdfDir := homeDir + "/.asdf"

	if _, err := os.Stat(asdfDir); err == nil {
		Print.Success("asdf already installed!")
		return nil
	}

	Print.NewLns(StyleInfoC, "Installing asdf-vm...")

	cmd := exec.Command("git", "clone", "https://github.com/asdf-vm/asdf.git", asdfDir, "--branch", "v0.14.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone asdf: %w", err)
	}

	Print.NewLns(StyleSuccess, "asdf-vm installed successfully!")
	Print.Warn("Note: Source your shell config or restart your shell to use asdf")
	return nil
}

// InstallAsdfPlugin installs an asdf plugin if not already installed.
func InstallAsdfPlugin(plugin string) error {
	cmd := exec.Command("asdf", "plugin", "list")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list asdf plugins: %w", err)
	}

	installedPlugins := strings.SplitSeq(strings.TrimSpace(string(output)), "\n")
	for installed := range installedPlugins {
		if installed == plugin {
			return nil
		}
	}

	fmt.Printf("%s Installing asdf plugin: %s\n", Dim("→"), plugin)
	installCmd := exec.Command("asdf", "plugin", "add", plugin)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install asdf plugin %s: %w", plugin, err)
	}

	return nil
}

// InstallDevTools installs development tools via asdf.
func InstallDevTools() error {
	Print.NewLns(StyleInfoC, "Installing development tools via asdf...")

	if err := InstallAsdf(); err != nil {
		return err
	}

	repoRoot, err := GetRepoRoot()
	if err != nil {
		return err
	}

	toolVersionsPath := repoRoot + "/config/tool-versions"
	data, err := os.ReadFile(toolVersionsPath)
	if err != nil {
		return fmt.Errorf("failed to read tool-versions: %w", err)
	}

	var tools []string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 1 {
			tools = append(tools, parts[0])
		}
	}

	Print.Dimmed("Installing asdf plugins...")
	for _, tool := range tools {
		if err := InstallAsdfPlugin(tool); err != nil {
			Print.Warn(fmt.Sprintf("Warning: Failed to install plugin %s: %v", tool, err))
		}
	}

	Print.Info()
	fmt.Println(Dim("Installing tool versions..."))
	fmt.Printf("%s This may take a while...\n\n", Dim("→"))

	installCmd := exec.Command("asdf", "install")
	installCmd.Dir = repoRoot + "/config"
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("failed to install asdf tools: %w", err)
	}

	Print.NewLns(StyleSuccess, "Development tools installed successfully!")
	return nil
}

// InstallAllPackages installs all packages (pacman, AUR, and dev tools).
func InstallAllPackages(fs embed.FS) error {
	Print.NewLns(StyleInfoC, "Installing all packages...")
	if err := InstallPacmanPackages(fs); err != nil {
		return err
	}
	Print.Info()
	if err := InstallAURPackages(fs); err != nil {
		return err
	}

	Print.Info()
	if err := InstallDevTools(); err != nil {
		return err
	}
	Print.NewLns(StyleSuccess, "All packages installed successfully!")
	return nil
}
