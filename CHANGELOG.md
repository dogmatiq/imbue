# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

Despite the presence of breaking changes due to modified function signatures,
this release does not alter the usage of the API under normal circumstances.

### Added

- Added `ContainerAware` interface, which is implemented by `Container` itself

### Changed

- **[BC]** `WithX()`, `WithXNamed()`, `WithXGrouped()` and `DecorateX()` now accept `ContainerAware` instead of `*Container`

## [0.6.2] - 2022-08-24

### Added

- Add `Container.WaitGroup()` and the `GoX()` functions

### Fixed

- Fix data race in `Context.Defer()`

## [0.6.1] - 2022-08-20

### Added

- Add `Invoke0` to help while refactoring dependencies

## [0.6.0] - 2022-08-03

### Changed

- **[BC]** `Context` is now an interface instead of a struct

## Removed

- **[BC]** Remove all environment variable related features

## [0.5.1] - 2022-07-03

### Fixed

- `DecorateX()` now panics when called for a type that has already been constructed
- `Container.Close()` now gives file/line information about deferred functions that return an error
- Functions deferred by `Context.Defer()` are now called immediately if the constructor (or a decorator) fails

## [0.5.0] - 2022-06-29

### Added

- Add `ByName.Name()` method, which returns the dependency name as a string
- Add `FromGroup.Group()` method, which returns the group name as a string
- Add support for depending on environment variables
- Add `Optional[T]` for representing optional dependencies

### Changed

- **[BC]** Changed `ByName.Value` from a field to a method
- **[BC]** Changed `FromGroup.Value` from a field to a method

### Fixed

- `WithX()` now panics when declaring a constructor for a type that has already been declared

## [0.4.0] - 2022-06-22

### Added

- Add `Decorate0()`

### Changed

- **[BC]** Rename `InjectX()` to `DecorateX()` and add return value

## [0.3.1] - 2022-06-19

### Changed

- Improved panic messages produced by `InjectX()` functions

## [0.3.0] - 2022-06-19

### Added

- Add `Container.String()`, which returns a string representation of the dependency tree
- Add `InjectX()` functions for injecting dependencies into values after they are constructed

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
[0.4.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.4.0
[0.5.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.5.0
[0.5.1]: https://github.com/dogmatiq/imbue/releases/tag/v0.5.1
[0.6.0]: https://github.com/dogmatiq/imbue/releases/tag/v0.6.0
[0.6.1]: https://github.com/dogmatiq/imbue/releases/tag/v0.6.1
[0.6.2]: https://github.com/dogmatiq/imbue/releases/tag/v0.6.2

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
