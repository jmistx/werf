<!-- THIS FILE IS AUTOGENERATED BY werf docs COMMAND! DO NOT EDIT! -->

**werf** is an Open Source CLI tool written in Go, designed to simplify and speed up the delivery of applications. To use it, you need to describe the configuration of your application (in other words, how to build and deploy it to Kubernetes) and save it to a Git repo — the latter acts as a single source of truth. In short, that's what we call GitOps today.

* werf builds Docker images using Dockerfiles or an alternative builder-tool based on the custom syntax. It also deletes unused images from the Docker registry.
* werf deploys your application to Kubernetes using a chart in the Helm-compatible format with handy customizations and improved rollout tracking mechanism, error detection, and log output.

werf is not a complete CI/CD solution, but a tool for creating pipelines that can be embedded into any existing CI/CD system. It literally "connects the dots" to bring these practices into your application. We consider it a new generation of high-level CI/CD tools.