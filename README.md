# Health Checker

Dead simple utility to check HTTP servers for availability.

## Command line app

You can use it as a CLI-app on your machine/server.

1. Clone this repo

```sh
git clone https://github.com/cheatsnake/healthchecker.git
```

```sh
cd ./healthchecker
```

2. Install packages

```sh
go mod download
```

3. Build a binary from source

```sh
make build
```

4. Run it!

```sh
./healthchecker -urls "https://example.com"
```

The manual page is preety straighforward:

```sh
healthcheck - show the availability of HTTP servers

Usage:
	-urls "...URLs" - list of HTTP URLs to check
	-help - print this manual

Examples:
	healthcheck -urls "https://google.com https://github.com"

Source code: https://github.com/cheatsnake/healthchecker
Leave issue: https://github.com/cheatsnake/healthchecker/issues
```
