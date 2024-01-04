# Platform Orchestrator (GitHub Action)

[![Golang CI](https://github.com/PashmakGuru/gha-platform-orchestrator/actions/workflows/golang-ci.yaml/badge.svg)](https://github.com/PashmakGuru/gha-platform-orchestrator/actions/workflows/golang-ci.yaml)

The Platform Orchestrator is a GitHub Action tailored for orchestrating platform requirements, setting the stage for Terraform modules to provision your desired infrastructure state.

## Features

- **Update Clusters Configuration**: Automatically adjust cluster settings, ensuring they meet the latest operational requirements.
- **Update Fronthub Configuration**: Modify the front-end hub configuration to enhance user interactions and backend integration.
- **GitHub Actions Integration**: Offers a range of GitHub Actions for seamless incorporation into CI/CD workflows, facilitating continuous development and deployment processes.

## Contributing

Contributions are what make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Before submitting your pull request, please run `go test ./...` to ensure all tests pass.

## Workflows
## Workflows
| Name | Description |
|---|---|
| [golang-ci.yaml](.github/workflows/golang-ci.yaml) | This workflow efficiently automates testing, formatting, and Docker image testing/building/pushing processes, ensuring that code is consistently tested, formatted, and ready for deployment. |
