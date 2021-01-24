kvnc
====

[![Actions Status](https://github.com/kaginawa/kvnc/workflows/Go/badge.svg)](https://github.com/kaginawa/kvnc/actions)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=kaginawa_kvnc&metric=alert_status)](https://sonarcloud.io/dashboard?id=kaginawa_kvnc)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaginawa/kvnc)](https://goreportcard.com/report/github.com/kaginawa/kvnc)

[Kaginawa](https://github.com/kaginawa/kaginawa)-powered VNC connection utilities.

## Download

See [Releases](https://github.com/kaginawa/kvnc/releases) page.

## Setup

### Windows

Get the following files and place them in the same folder.

| File | Description | Download page |
| --- | --- | --- |
| `kvnc-client.exe` | kvnc client utility | [GitHub](https://github.com/kaginawa/kvnc/releases) |
| `kvnc-agent.exe`  | kvnc agent utility | [GitHub](https://github.com/kaginawa/kvnc/releases) |
| `kaginawa.exe`    | kaginawa | [GitHub](https://kaginawa.github.io/) |
| `vncviewer.exe`   | TightVNC client | [TightVNC](https://www.tightvnc.com/download-old.php) |
| `WinVNC.exe`      | TightVNC server | [TightVNC](https://www.tightvnc.com/download-old.php) |

Tested VNC versions:

- TightVNC 1.3.10

## Usage

### Server

1. (1st time only) Run `WinVNC.exe`.
2. (1st time only) Right-click the task tray icon and open [Properties...].

    ![winvnc tray](docs/winvnc-tray.png)

3. (1st time only) Input the **Primary password** and **view-only password**.

    ![winvnc password](docs/winvnc-password.png)
    
4. (1st time only) Select [Administration] tab, check **Allow loopback connections** and press [OK].

    ![winvnc loopback](docs/winvnc-loopback.png)
    
1. Run `kvnc-agent.exe`.
2. (1st time only) Input server and API key and press [OK].
   
    ![kvnc-agent init](docs/kvnc-agent-init.png)

3. Input own custom ID and press [Start].

    ![kvnc-agent](docs/kvnc-agent.png)

4. If "Working" is displayed, the agent is working.
5. Please tell the client your custom ID and VNC password.

### Client

1. Run `kvnc-client.exe`.
2. (1st time only) Input server and API key and press [OK].

    ![kvnc-client init](docs/kvnc-client-init.png)

3. Input target custom ID (or MAC address) and press [Connect].

    ![kvnc-client connect](docs/kvnc-client.png)

4. Input password and press [OK].

    ![vnc viewer login](docs/viewer-password.png)

5. Enjoy!

## License

kvnc licenced under [BSD 3-Clause](LICENSE).

## Author

[mikan](https://github.com/mikan)
