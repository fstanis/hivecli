# hivecli
Unofficial CLI compatible with the API used by [Hive](https://www.hivehome.com/)
(by British Gas) smart devices.

**hivecli** is written in Go using the [go-hive](https://github.com/fstanis/go-hive)
library and also serves as a comprehensive example for using it.

## Disclaimer
**This software is in no way endorsed by British Gas**. The underlying REST API
is undocumented and subject to change at any time, which means `hivecli` may
suddenly stop working.

**Use at your own risk.**

In addition, the CLI only supports a subset of the API - specifically, the
devices I personally own. Most notably, smart heating is not supported.

Contributions are welcome.

## Download

You can find the latest version of the binary in the [releases section](https://github.com/fstanis/hivecli/releases).

## Usage

```
$ hivecli [OPTION]... COMMAND [ARGUMENT]...
```

Running `hivecli` with no arguments will display the supported commands and how
they're used.

### First run

The first time `hivecli` is used to run a command, it will ask for the following
information:

* Your username
* Your password
* The login URL API endpoint (for Hive, the default login URL `https://beekeeper.hivehome.com/1.0/global/login`)

The password will then be stored in your platform's keyring (dbus
under Linux, `/usr/bin/security` under MacOS and Credential Manager under
Windows), while the username and login URL are stored in the config file.
