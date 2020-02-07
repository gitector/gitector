# Gitector

Lint your messages

## Installation

### Download binary

`wget link_here `

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

### Gitlab

```
stages:
  - gitector
  - build
  - deploy

before_script:
  - wget link > gitector

gitector:
  image: gitector/gitector
  script:
    - git fetch
    - gitector
  only:
    refs:
      - merge_requests
      - master
      - web


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
