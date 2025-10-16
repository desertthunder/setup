# Hyprland

All calls to pacman are with sudo priveleges.

## System

### Brightness

```sh
pacman -S brightnessctl
```

### Network

```sh
pacman -S networkmanager # interact with nmtui
```

### Bluetooth

```sh
yay -S bluez bluez-utils blueman

sudo systemctl enable --now bluetooth
```

## Browsers

```sh
yay -S zen-browser-bin
```

### Electron

```sh
yay -S visual-studio-code-bin obsidian-bin
```

## Keybinds

| Keybind                  | Action                        |
| ------------------------ | ----------------------------- |
| `SUPER + T`              | Open terminal (alacritty)     |
| `SUPER + C`              | Kill active window            |
| `SUPER + Q`              | Exit Hyprland                 |
| `SUPER + E`              | Open file manager (nautilus)  |
| `SUPER + V`              | Toggle floating mode          |
| `SUPER + R`              | Open app launcher (rofi)      |
| `SUPER + P`              | Toggle pseudotile (dwindle)   |
| `SUPER + J`              | Toggle split (dwindle)        |
| `SUPER + Arrow Keys`     | Move focus between windows    |
| `SUPER + 1-9, 0`         | Switch to workspace 1-10      |
| `SUPER + SHIFT + 1-9, 0` | Move window to workspace 1-10 |
| `SUPER + S`              | Toggle scratchpad workspace   |
| `SUPER + SHIFT + S`      | Move window to scratchpad     |
| `SUPER + Mouse Scroll`   | Switch workspaces             |
| `SUPER + LMB (drag)`     | Move window                   |
| `SUPER + RMB (drag)`     | Resize window                 |
| `XF86AudioRaiseVolume`   | Increase volume by 5%         |
| `XF86AudioLowerVolume`   | Decrease volume by 5%         |
| `XF86AudioMute`          | Toggle audio mute             |
| `XF86AudioMicMute`       | Toggle microphone mute        |
| `XF86MonBrightnessUp`    | Increase brightness by 5%     |
| `XF86MonBrightnessDown`  | Decrease brightness by 5%     |
| `XF86AudioNext`          | Next media track              |
| `XF86AudioPause`         | Play/pause media              |
| `XF86AudioPlay`          | Play/pause media              |
| `XF86AudioPrev`          | Previous media track          |
