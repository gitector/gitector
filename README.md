# Gitector

Lint your messages

## Installation

### Download binary

Go to [releases section](https://github.com/gitector/gitector/releases)

### Brew

- `brew tap gitector/gitector` <- Add external tap with gitector
- `brew install gitector` <- install gitector


### From surouce

- `make build` <- to build
- `make install` <- to install

## CLI Usage

### Init for project

`gitector init`

### Example commands

- `gitector` – Test commits from current branch compared to master 

- `gitector master -l 10` – Test last 10 commits from master 

- `gitector feature-1..long-last-branch` – Test new commits from branch feature-1 with long-last-branch   

- `gitector ..master -o html` – Test commits from current branch to master with html output

### As githook

Add line `gitector -id` to `.git/COMMIT_EDITMSG`

### As merge request tool

Run `gitector` by default it will compare new commits against master.


## CI Usage

### Github Actions

You can find latest version to add to your [workflow here](https://github.com/gitector/gitector-actions)

### Gitlab

Add code below to your `.gitlab-ci.yml` file

```
gitector:
  stage: build
  image: gitector/gitector
  script:
    - git fetch
    - gitector
  except:
    - master

```

## Development

### Building
`go build`

### Testing
`./test.sh`

## Troubleshooting
- Missing package
    - Run `go get ${missing package}` eg. `go get github.com/stretchr/testify`

## Run
`./gitector --help`
