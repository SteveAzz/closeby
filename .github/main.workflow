workflow "Test & Build" {
  on = "push"
  resolves = ["build"]
}

action "build" {
  uses = "docker://golang:1.11"
  runs = "go version"
}
