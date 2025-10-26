# GO-GET-SOME-STYLEZ

A Golang script that installs and configures Oh My Zsh (or Oh My Posh for Windows) 
based on the detected OS. 
It also provides helpful guidance for users to customize their shell experience.

## Features
- Detects OS: Fedora, Debian, Ubuntu, Arch, Manjaro, openSUSE, macOS, and Windows.
- Installs `zsh` and `Oh My Zsh` with essential plugins (`zsh-autosuggestions`, `zsh-syntax-highlighting`).
- Backs up the existing `.zshrc` file to `~/.zshrc_original_bkp`.
- Asks the user if they want to set `zsh` as the default shell.
- Installs `Oh My Posh` for Windows users and provides guidance on setting up themes.
- Opens a new `zsh` shell after installation.
- Uses colored terminal output for a better user experience.

## Installation
### Prerequisites
- `git` installed on the system.
- `curl` or `wget` installed for fetching external scripts.
- `sudo` privileges for package installation.
- `GO LANGUAGE` project's base technology.

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/baletadbr4in1ac/GO-GET-SOME-STYLEZ.git
   cd GO-GET-SOME-STYLEZ
   ```
2. Build the project:
   ```sh
   Windows hosts:
    go build -o GO-GET-SOME-STYLEZ.exe .\cmd\app\
   
   Linux hosts:
    go build -o GO-GET-SOME-STYLEZ ./cmd/app/
   ```
3. Run the script:
   ```sh
   ./GO-GET-SOME-STYLEZ
   ```

## OS Detection & Package Managers
| OS       | Package Manager       |
|----------|-----------------------|
| Fedora   | `dnf`                 |
| Debian   | `apt`                 |
| Ubuntu   | `apt`                 |
| Arch     | `pacman`              |
| Manjaro  | `pacman`              |
| openSUSE | `zypper`              |
| Windows  | `winget` (Oh My Posh) |

## Behavior
1. Detects the OS.
2. Installs `zsh` (if not installed).
3. Installs `Oh My Zsh` and plugins (for Linux/macOS).
4. Backs up `.zshrc` to `~/.zshrc_original_bkp`.
5. Asks if `zsh` should be set as the default shell.
6. Installs `Oh My Posh` for Windows users.
7. Creates a help file for Windows users on their desktop.
8. Opens a new `zsh` shell.
9. Displays a success message: `NOW YOU GOT STYLEZ 😉`.

## Customization
- To manually set a theme for `Oh My Posh`, run:
  ```sh
  oh-my-posh config
  ```
- To install Nerd Fonts (required for proper icon rendering):
    - Visit: [https://www.nerdfonts.com](https://www.nerdfonts.com)
    - Download and install a preferred font.
    - Set the terminal to use the installed Nerd Font.

## Final considerations
⌨️ Made with ❤️ by [BraiNiac](https://github.com/babyboydaprince) 👽

