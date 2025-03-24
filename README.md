![LibreRemotePlay logo banner](https://github.com/user-attachments/assets/9173246d-1d65-4f29-bd79-0206250c538c)

# LibreRemotePlay
### An open source & decentralized alternative to Steam remote play (No need to self host)

⌛ Looking for contributions 👈

> [!Note]
> Website: 
> https://libreremoteplay.vercel.app/
>
> Wiki:
> https://github.com/PiterWeb/LibreRemotePlay/wiki

## Use cases ✨

- Share your local co-op games online with friends (as [Steam Remote Play Anywhere](https://store.steampowered.com/remoteplay?l=english#anywhere))
- Stream your games from your PC to other devices (as [Steam Remote Play Together](https://store.steampowered.com/remoteplay?l=english#together))

## Download 📦

- https://github.com/PiterWeb/LibreRemotePlay/releases/latest

## Resources 📚

- [Docs](./docs/) 📘
- [Linux Docs](./docs/LINUX.md) 📘
- [Wiki](https://github.com/PiterWeb/LibreRemotePlay/wiki) (Guides, FAQ, ...)

### Videos 📹
(Note that videos may show older/beta versions of the APP and could have bugs that are already fixed in the latest version)

#### How to use

https://github.com/user-attachments/assets/f4a412fa-f403-4429-85fb-9c1e74bff458

## Features 🧩

- [x] Portable
- [x] Simple & Modern UI
- [x] P2P "Decentralized" (WebRTC)
- [x] Remote Streaming
- [x] Browser Client
- [x] Support for keyboard (very experimental)
- [x] ViGEmBus Setup (Windows)
- [x] Localization (translation to multiple languages)
- [x] Support for STUN & TURN

### OS Support 💻

| Windows 	| Linux 	| MacOS 	| Browser (Only Client) 	|
|---------	|-------	|-------	|---------	|
| ✔       	| ✔ Gamepad/Streaming support (❌ Keyboard for the moment)     	| ❌ (No MacOS to test)     	| ✔ (Known Issues with Safari)       	|

### Native Gamepad Support 🎮

| PC Controller (XInput/DirectInput) 	| Xbox Controller (XInput) 	| PlayStation Controler
|---------	|-------	|-------	|
| ✔       	| ✔     	| ❌ (But you can achieve [emulating a Xbox Controller](https://github.com/Ryochan7/DS4Windows))     	|

### Translations 🔠

| English 	| Spanish 	| Galician | Russian | French |Other languages |
|---------	|-------	|-------	| ------- | ------- | ------- |
| 100% ✔     	| 100% ✔      	| 100% ✔      	| 100% ✔ | 100% ✔ (@Zorkyx22) |⌛ Looking for contributions

## Self Hosting ☁

There is no way to self-host the infrastructure of LibreRemotePlay because it has no backend. But instead you can self-host if you want the TURN & STUN servers and then add them to the config.

- If you want to self-host a TURN/STUN server you can [try Coturn](https://github.com/coturn/coturn). (This is only an example, you can choose other STUN/TURN implementations)

- Also you can host the Web version (but it is only frontend, so is not very usefull)

## Run Dev

### Prerequisites

You must have [Task CLI](https://taskfile.dev/installation/), [Wails CLI](https://wails.io/docs/gettingstarted/installation#installing-wails), [NodeJS (~v20.x.x)](https://nodejs.org/en/download), [pnpm](https://pnpm.io/es/installation) and [Golang (min v1.22.4)](https://go.dev/doc/install) installed.

### How to

Go to the root project folder and run

  - Full App :

    `$ task dev-all`

  - Frontend:

    `$ task dev-front`

## Build

### Prerequisites

You must have [Task CLI](https://taskfile.dev/installation/), [Wails CLI](https://wails.io/docs/gettingstarted/installation#installing-wails), [NodeJS (~v20.x.x)](https://nodejs.org/en/download), [pnpm](https://pnpm.io/es/installation) and [Golang (min v1.22.4)](https://go.dev/doc/install) installed.

### How to

Go to the root project folder and run

- For general builds:

    `$ task build`

- For Windows builds:

    `$ task build-win`

- For Linux builds:

    `$ task build-linux`

finally go to the build/bin folder and your executables will be there.

> [!Note]
> Please note the supported platforms in the table

## Contributting 🤝

If you are interested to contribute to this project you can follow this [guide](./CONTRIBUTING.md)

Also 

## Acknowledgements

### Thanks to jbdemonte/virtual-device ❤
[jbdemonte/virtual-device](https://github.com/jbdemonte/virtual-device) is making this project a reality. This is the source of magic that enables LibreRemotePlay to generate virtual gamepads on Linux, is very fast and made in pure Go.
### Thanks to the ViGEm project  ❤
[ViGEmBus](https://github.com/nefarius/ViGEmBus) is making this project a reallity. This is the source of magic that enables LibreRemotePlay to generate virtual gamepads on Windows. We embed ViGEmBus Installation Wizard and ViGEmBus Client DLLS within the executable for Windows

## Did you like the project 👍 ?
You can give a star and review us on Product Hunt

<a href="https://www.producthunt.com/products/remote-controller/reviews?utm_source=badge-product_review&utm_medium=badge&utm_souce=badge-remote&#0045;controller" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/product_review.svg?product_id=565186&theme=light" alt="LibreRemotePlay - Play&#0032;LOCAL&#0032;co&#0045;op&#0032;games&#0032;ONLINE | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

## Star History
Here you can see how fast the community is growing
<br/>
[![Star History Chart](https://api.star-history.com/svg?repos=PiterWeb/LibreRemotePlay&type=Timeline)](https://star-history.com/#PiterWeb/LibreRemotePlay&Timeline)
