# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

### Added

- Add `Decorate0()`

### Changed

- **[BC]** Rename `InjectX()` to `DecorateX()` and add return value

## [0.3.1] - 2022-06-19

### Changed

- Improved panic messages produced by `InjectX()` functions.

## [0.3.0] - 2022-06-19

### Added

- Add `Container.String()`, which returns a string representation of the dependency tree.
- Add `InjectX()` functions for injecting dependencies into values after they are constructed.

### Changed

- **[BC]** Reduce number of supported dependencies to 8
- Add option parameters to `WithX()`, `WithXNamed()` and `InvokeX()`

## [0.2.0] - 2022-06-14

### Changed

- **[BC]** Changed `ByName[N, T]` to a struct instead of a function

## [0.1.0] - 2022-06-14

- Initial release

<!-- references -->

[unreleased]: https://github.com/dogmatiq/imbue
[0.1.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.1.0
[0.2.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.2.0
[0.3.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.3.0
[0.3.1]: https://github.com/dogmatiq/imbue/releases/tag/v0.3.1

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
