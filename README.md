# sudont

Making sure a command is never run as root.

## Usage

<!-- TODO: real example -->

```bash
sudont whoami # anyone but "root"
```

## How it works

`sudont` runs the given command in a minimal [runc](https://github.com/opencontainers/runc)-like container with a substituted user namespace to drop root privileges.
