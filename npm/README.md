
`tau` is a command line tool for interacting with a Taubyte-based Cloud Network. Similar to the [web console](https://console.taubyte.com), it allows you to create and manage projects, applications, resources, and more.

## Installation
```shell
npm i tau
```

## Login

`tau login`
    - opens selection with default already selected
    - simply logs in if only default available
    - will open new if no profiles found
`tau login --new` for new
  - `--set-default` for making this new auth the default
`tau login <profile-name>` for using a specific profile


## Environment Variables:
- `TAUBYTE_PROJECT` Selected project
- `TAUBYTE_PROFILE` Selected profile
- `TAUBYTE_APPLICATION` Selected application
- `TAUBYTE_CONFIG (default: ~/tau.yaml)` Config location
- `TAUBYTE_SESSION (default: /tmp/tau-<shell-pid>)` Session location
- `DREAM_BINARY (default: $GOPATH/dream)` Dream binary location

# Documentation
For documentation head to [tau.how](https://tau.how/docs/tau)
