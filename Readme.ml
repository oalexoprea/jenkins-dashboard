# Jenkins Dashboard TUI

A simple and interactive Jenkins Dashboard for the terminal, built with [Bubbletea](https://github.com/charmbracelet/bubbletea), [Bubbles](https://github.com/charmbracelet/bubbles), and [Lipgloss](https://github.com/charmbracelet/lipgloss).

---

## Features

✅ View Jenkins jobs with real-time status  
✅ Auto-refresh jobs every 10 seconds  
✅ Color-coded job statuses (Success, Fail, Running, Unknown)  
✅ Spinner while loading  
✅ Keyboard shortcuts for manual refresh and quit  

---

## Installation

```bash
git clone https://github.com/youruser/jenkins-dashboard.git
cd jenkins-dashboard
go mod tidy
go run .
```

---

## Environment Variables

| Variable       | Description           |
|----------------|-----------------------|
| `JENKINS_URL`  | Base URL of Jenkins   |
| `JENKINS_USER` | Jenkins username      |
| `JENKINS_TOKEN`| Jenkins API token     |

Example:

```bash
export JENKINS_URL=http://localhost:8080
export JENKINS_USER=admin
export JENKINS_TOKEN=yourapitoken
```

---

## Keyboard Shortcuts

| Key        | Action                |
|------------|-----------------------|
| `r`        | Manual refresh        |
| `q` / `ctrl+c` | Quit application   |

---

## Dependencies

- [Bubbletea](https://github.com/charmbracelet/bubbletea)  
- [Bubbles](https://github.com/charmbracelet/bubbles)  
- [Lipgloss](https://github.com/charmbracelet/lipgloss)

---

## License

MIT License
