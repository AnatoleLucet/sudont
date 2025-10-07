# sudont

> superuser, don't.

Make sure a command is never run as root.

```bash
sudont <command>
```


## Usage

```
NAME:
   sudont - Making sure a command is never run as root.

USAGE:
   sudont [global options] <command> [command arguments...]

VERSION:
   1.0.0-rc.1 (go1.25.1 X:nodwarf5 on linux/amd64; gc)

DESCRIPTION:
   sudont is a tool to run commands as a non-root user.
   It tries to automatically determine a non-root user, and refuses to run commands as root, even if explicitly asked to do so.

GLOBAL OPTIONS:
   --user string, -u string  Specify the user to run the command as.
   --help, -h                show help
   --version, -v             print the version
```

## Examples

```bash
# Prevent NPM packages to be installed with root privileges
# which could otherwise give full access to a remote source via postinstall scripts
sudont npm install -g <pkg>

# Prevent remote script to run with root privileges
# which could otherwise give full access to a remote source
sudont curl <url>/script.sh | bash

# Prevent your application to run with root privileges
# which could be abused by malicious actors and give them full access
sudont node server.js
```

## How it works

> TLDR: `sudont` runs the given command in a minimal [runc](https://github.com/opencontainers/runc)-like container with a substituted user namespace to drop root privileges.


