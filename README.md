# gce-instances-summary

Simple Go script for listing all Google Compute Engine instances by project and status.
It's using your gcloud defaults for authentication, so make sure you're logged in.

## Install

Install Go - see https://golang.org/doc/install

```
go get -u github.com/jordypixelcode/gce-instances-summary
go install github.com/jordypixelcode/gce-instances-summary
```

## Usage

```
gce-instances-summary
```

Output:

```
Project project-x (europe-west1-d)

                       server1 - RUNNING
                       server2 - RUNNING
                       server3 - RUNNING


Project project-y (europe-west1-d)

                       server1 - RUNNING


Project project-z (europe-west1-d)

                      server-x - TERMINATED
                      server-y - RUNNING
```
