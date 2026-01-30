# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub Actions CI workflow for automated testing and linting
- GitHub Actions release workflow for creating tagged releases
- Cross-platform binary builds for the CLI tool

## Release Notes

Release notes are automatically generated from Git commits and pull requests when a new version is tagged. You can view detailed release notes on the [GitHub Releases page](https://github.com/marcgeld/ruuvi/releases).

### Versioning Policy

This project uses [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes (v1.0.0, v2.0.0, etc.)
- **MINOR** version for new functionality in a backward compatible manner (v1.1.0, v1.2.0, etc.)
- **PATCH** version for backward compatible bug fixes (v1.0.1, v1.0.2, etc.)

**Before v1.0.0**: Breaking changes may occur but should still bump the minor version (v0.x.y).

**After v1.0.0**: Breaking changes will bump the major version.

[Unreleased]: https://github.com/marcgeld/ruuvi/compare/v0.1.0...HEAD
