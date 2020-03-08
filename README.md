# MCI-Downloader

Download mp3 tracks from your Philips MCI500H hi-fi system.

## Usage

Download the latest [release](https://github.com/djmaze/mci-downloader/releases) or build it yourself (see below).

```bash
/path/to/mci-downloader [--dry-run] [--port <PORT>] <IP or hostname>
```

E.g. for the common usage:

```bash
./mci-downloader 192.168.1.11
```

* The tool downloads all tracks from the given MCI500H device to your **current working directory**.
* It creates **one folder per album** with the following naming scheme:

        <artist> - <album>/<track number> <track name>

* By default, port `8081` is used. You can specify a different port with the `--port` option.
* You can simulate the download using the `--dry-run` option. Files will not be written in this case.

## Build

Prerequisites:

* Go 1.12

You can build the executable using:

```bash
go build
```