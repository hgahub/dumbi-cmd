# CLI Documentation

To view a list of the commands available in your current Dumbi version, 
run terraform with no additional arguments:

```shell
Usage: dumbi [options] <command> [args]

The available commands for execution are listed below.

Commands:
  -validate     validates the configuration files
  -fmt          rewrite Dumbi configuration files to a canonical format and style

Options (use these before the subcommand, if any):
  -help         Show this help output, or the help for a specified subcommand.
  -version      Displays the version of Dumbi.
```

## Commands
- **[validate](./command_validate.md)** Validates the configuration files in a directory.