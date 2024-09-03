# Carcassonne-Visualiser

Application that creates step-by-step visualisation of logs of the
 Carcassonne game created by rule engine available [here](https://github.com/YetAnotherSpieskowcy/Carcassonne-Engine).

## Pre-requirements

### Linux

1. Install Go 1.22 either from your distro's package repositories or by following [instructions on Golang's site](https://go.dev/doc/install).

   **Tip:** If you're using Ubuntu 23.10 or lower, Go version in official repositories is going to be too old.
   You can get the latest version by adding [the PPA listed on Go wiki](https://go.dev/wiki/Ubuntu) and installing `golang` package after.

2. Install SDL2 from your distro's package repositories. For additional 
help follow [instructions](https://wiki.libsdl.org/SDL2/Installation#linuxunix).
3. Install raylib-go requirements following [intructions for your distro](https://github.com/gen2brain/raylib-go?tab=readme-ov-file#requirements).

### Windows

1. Install [Go for Windows x86-64](https://go.dev/dl/).
2. Download [SDL2 binary](https://github.com/libsdl-org/SDL/releases) and put the dll file in the system directory.
3. Install raylib-go requirements following [intructions](https://github.com/gen2brain/raylib-go?tab=readme-ov-file#windows).

## Run application

Use this command to run applications. Requires log file name passed as parameter.

### Linux

```console
./make.sh run _log-filename_
```

### Windows

```console
./make.ps1 run _log-filename_
```

## Update requirements

Use this command to get latest versions of Carcassonne-Engine and raylib-go.

### Linux

```console
./make.sh update
```

### Windows

```console
./make.ps1 update
```
