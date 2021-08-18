# Changelog

## [0.12.0](https://www.github.com/contiamo/openapi-generator-go/compare/v0.11.0...v0.12.0) (2021-08-18)


### Features

* Add discriminator support for oneOfs ([#57](https://www.github.com/contiamo/openapi-generator-go/issues/57)) ([78080f0](https://www.github.com/contiamo/openapi-generator-go/commit/78080f097ffdaba8ecd8207f8d2e13e236493f2f))

## [0.11.0](https://www.github.com/contiamo/openapi-generator-go/compare/v0.10.1...v0.11.0) (2021-08-12)


### Features

* expose the merge package as command ([#52](https://www.github.com/contiamo/openapi-generator-go/issues/52)) ([cd2fcf9](https://www.github.com/contiamo/openapi-generator-go/commit/cd2fcf98ddb8ee39d659741f9866732e4c32ff69))
* improve description handling in allOfs ([#55](https://www.github.com/contiamo/openapi-generator-go/issues/55)) ([f6feffe](https://www.github.com/contiamo/openapi-generator-go/commit/f6feffe7cd47d292ab52ed14ef75331f67655aa6))


### Bug Fixes

* MarshalJSON should not have pointer reciever ([#50](https://www.github.com/contiamo/openapi-generator-go/issues/50)) ([4946b37](https://www.github.com/contiamo/openapi-generator-go/commit/4946b3735971663689344c77432e6aef18880b34))

### [0.10.1](https://www.github.com/contiamo/openapi-generator-go/compare/v0.10.0...v0.10.1) (2021-07-30)


### Bug Fixes

* improve allOf type generation ([#48](https://www.github.com/contiamo/openapi-generator-go/issues/48)) ([c31f173](https://www.github.com/contiamo/openapi-generator-go/commit/c31f1733f4dddcea02c1c6ca6a0ecdf74312b72e))

## [0.10.0](https://www.github.com/contiamo/openapi-generator-go/compare/v0.9.0...v0.10.0) (2021-06-29)


### Features

* Add From* methods to set oneOf data models ([#45](https://www.github.com/contiamo/openapi-generator-go/issues/45)) ([4bc1140](https://www.github.com/contiamo/openapi-generator-go/commit/4bc11402f23e615422f23e3f224779d0e250dad1))

## [0.9.0](https://www.github.com/contiamo/openapi-generator-go/compare/v0.8.1...v0.9.0) (2021-06-23)


### ⚠ BREAKING CHANGES

* add convert functions for toplevel oneofs. (#41)

### Features

* add convert functions for toplevel oneofs. ([#41](https://www.github.com/contiamo/openapi-generator-go/issues/41)) ([93d7cc8](https://www.github.com/contiamo/openapi-generator-go/commit/93d7cc8e4c66226352317be5ec33ee50f7580f6b))

### [0.8.1](https://www.github.com/contiamo/openapi-generator-go/compare/v0.8.0...v0.8.1) (2021-06-17)


### Bug Fixes

* Allow responses section in components to be omitted ([#39](https://www.github.com/contiamo/openapi-generator-go/issues/39)) ([30398af](https://www.github.com/contiamo/openapi-generator-go/commit/30398affd55074b774627deccb42b5db396a88aa))
* use interface for oneOf ([#36](https://www.github.com/contiamo/openapi-generator-go/issues/36)) ([6bda6b3](https://www.github.com/contiamo/openapi-generator-go/commit/6bda6b3930d42dece3ac1f29054950267dabba76))
