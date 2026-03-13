# Changelog

## 0.1.0-alpha.5 (2026-03-13)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** api update ([13c4326](https://github.com/hanzoai/go-sdk/commit/13c4326777789f8d0b128da67463b765db2880ec))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([412a254](https://github.com/hanzoai/go-sdk/commit/412a254721caffc54a08f9198a7159c3f439bc07))
* **client:** use correct format specifier for header serialization ([12b46c1](https://github.com/hanzoai/go-sdk/commit/12b46c1dd383803637533ae962baf17e3cfc598e))
* correct test assertions for service Options and BaseURL ([5d2f80e](https://github.com/hanzoai/go-sdk/commit/5d2f80ed7087b3398e0045395bd75fdc9f64ca8f))
* fix for namespace collisions with client and resource test methods ([04c6c39](https://github.com/hanzoai/go-sdk/commit/04c6c398913b15da218ef4087864831508dcfbf5))
* remove nil Options assertion when no opts passed ([dae317e](https://github.com/hanzoai/go-sdk/commit/dae317efdfddadcc4937abb0b30321d37203dab3))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([b31e9e0](https://github.com/hanzoai/go-sdk/commit/b31e9e00787cd1b1ed266a4783b0349aa2b140c8))
* **internal:** codegen related update ([f4af05d](https://github.com/hanzoai/go-sdk/commit/f4af05db5a76af294af9ac0f2265fbac79e2bcb1))
* **internal:** codegen related update ([f974f78](https://github.com/hanzoai/go-sdk/commit/f974f78d755176bdd069945948bec88fa19993c9))
* **internal:** minor cleanup ([a991c3e](https://github.com/hanzoai/go-sdk/commit/a991c3e2fcd759e320a1a34532404e7dc4016c05))
* **internal:** move custom custom `json` tags to `api` ([e0d573b](https://github.com/hanzoai/go-sdk/commit/e0d573bd1dce8e81e8bead12d13fd1b3920b75a6))
* **internal:** remove mock server code ([7ac9bd6](https://github.com/hanzoai/go-sdk/commit/7ac9bd6d92948923585bceefb558918d2a24ae37))
* **internal:** use explicit returns ([5dab23d](https://github.com/hanzoai/go-sdk/commit/5dab23da630dcb8333afea9c544cf1d6b080804f))
* **internal:** use explicit returns in more places ([905da44](https://github.com/hanzoai/go-sdk/commit/905da44f237f9153af03f0e9c02291f3b9bb631f))
* update Go toolchain to 1.26 ([813c7ad](https://github.com/hanzoai/go-sdk/commit/813c7ad26dda56fdd99c39171e97f0a3040f5702))
* update mock server docs ([1e00e6f](https://github.com/hanzoai/go-sdk/commit/1e00e6f84daaf0884a331052ea6bba98accb813c))
* update placeholder string ([602d8a2](https://github.com/hanzoai/go-sdk/commit/602d8a255ea344aa5eb329b32ce782ee89f7ea34))


### Documentation

* add LLM.md project guide ([9b4d0c4](https://github.com/hanzoai/go-sdk/commit/9b4d0c495e183f851f1f87ac78ac9a908a34e6de))

## 0.1.0-alpha.4 (2026-01-28)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** api update ([d24dacc](https://github.com/hanzoai/go-sdk/commit/d24dacc551627f091ed1f01e558d420caf24bb5b))
* **api:** api update ([4825c99](https://github.com/hanzoai/go-sdk/commit/4825c99fc68fa29e850cd04159dc14831f9c22ea))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([ed4aaea](https://github.com/hanzoai/go-sdk/commit/ed4aaea25e878550eab7bf2819955019ed2ef301))
* Complete Go SDK fixes for full compilation and testing ([40cd2e9](https://github.com/hanzoai/go-sdk/commit/40cd2e95c081f76b4eb812d411eb4f308056aad7))
* Resolve Go SDK compilation errors ([ac3c794](https://github.com/hanzoai/go-sdk/commit/ac3c7941b4c50a091951e9cff19b5e07db3aa45f))
* use slices.Concat instead of sometimes modifying r.Options ([865e39b](https://github.com/hanzoai/go-sdk/commit/865e39b45cdc3ae4de260945113bb6e18025b3eb))


### Chores

* bump minimum go version to 1.22 ([7f2076a](https://github.com/hanzoai/go-sdk/commit/7f2076abec37ce5e638500d0fc0e26e628c2a054))
* do not install brew dependencies in ./scripts/bootstrap by default ([85660de](https://github.com/hanzoai/go-sdk/commit/85660de8f532790f0348b2276f3570826fc9af1c))
* update more docs for 1.22 ([83a09fc](https://github.com/hanzoai/go-sdk/commit/83a09fceb3d2fd98447cafa27c6a5f4271b9af19))


### Documentation

* Add working example demonstrating SDK usage ([bd13718](https://github.com/hanzoai/go-sdk/commit/bd1371885540634ff56e1a58fcf5af9a5db5744d))

## 0.1.0-alpha.3 (2025-09-06)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** api update ([a7d4e2d](https://github.com/hanzoai/go-sdk/commit/a7d4e2d7f0a701c7b83a5fa62129f511bf78aca7))
* **client:** add support for endpoint-specific base URLs in python ([a1fced6](https://github.com/hanzoai/go-sdk/commit/a1fced627f9e55d536366505ef216745cefa2d6d))
* **client:** add support for reading base URL from environment variable ([fc8778c](https://github.com/hanzoai/go-sdk/commit/fc8778cc31f0d39da5bddb266b2c167e8385110c))
* **client:** support custom http clients ([#9](https://github.com/hanzoai/go-sdk/issues/9)) ([6cf2e49](https://github.com/hanzoai/go-sdk/commit/6cf2e49da411a8daa7cef3eab22e426be8630059))


### Bug Fixes

* **client:** clean up reader resources ([564bf4b](https://github.com/hanzoai/go-sdk/commit/564bf4b621de12be7a25d2431b0e1c08cf9c7f00))
* **client:** correctly update body in WithJSONSet ([2bfb47f](https://github.com/hanzoai/go-sdk/commit/2bfb47f27b1754a24701d795e8f3af2b3752ad51))
* handle empty bodies in WithJSONSet ([f828405](https://github.com/hanzoai/go-sdk/commit/f8284051703ca4348139004ec8b38c99b7e1904c))


### Chores

* **ci:** add timeout thresholds for CI jobs ([4d9cf13](https://github.com/hanzoai/go-sdk/commit/4d9cf137360c379f896289bd456c73a388cdea00))
* **ci:** only use depot for staging repos ([49f9a44](https://github.com/hanzoai/go-sdk/commit/49f9a444a42e40037262cf4f2c0af6d84b2e7bc2))
* configure new SDK language ([70c6b17](https://github.com/hanzoai/go-sdk/commit/70c6b171a497aef2590abc76eb8a98626595ce71))
* **docs:** document pre-request options ([f8f26dc](https://github.com/hanzoai/go-sdk/commit/f8f26dc587e783b6b87bd04879fae15b03da01f6))
* **docs:** grammar improvements ([f3c04c3](https://github.com/hanzoai/go-sdk/commit/f3c04c3b71d319b3ff30e68f93a3071e4c795ea2))
* improve devcontainer setup ([f516bcc](https://github.com/hanzoai/go-sdk/commit/f516bcc5c9565fd35e97cb197ced67c44fd5dba4))
* **internal:** codegen related update ([c0717b5](https://github.com/hanzoai/go-sdk/commit/c0717b5a3e07964d53e20676d72a2c644f35e5e8))
* **internal:** expand CI branch coverage ([a241eb5](https://github.com/hanzoai/go-sdk/commit/a241eb5b8464517c02ff4bd10178c727b2442849))
* **internal:** reduce CI branch coverage ([b380cd0](https://github.com/hanzoai/go-sdk/commit/b380cd00da61b2f85df12c4a35515f5db0a01b1a))
* make go mod tidy continue on error ([fa097f8](https://github.com/hanzoai/go-sdk/commit/fa097f8c401ceabad9d782c07abf94946e8c310a))


### Documentation

* update documentation links to be more uniform ([b54e5f0](https://github.com/hanzoai/go-sdk/commit/b54e5f0947bf859d22c9fdd3f9b2884a543ae531))

## 0.1.0-alpha.2 (2025-04-03)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/hanzoai/go-sdk/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#7](https://github.com/hanzoai/go-sdk/issues/7)) ([3ce9376](https://github.com/hanzoai/go-sdk/commit/3ce937641c2ece5c60ffcc7a43d6511ceb4c49d6))


### Chores

* update SDK settings ([#5](https://github.com/hanzoai/go-sdk/issues/5)) ([66a9f24](https://github.com/hanzoai/go-sdk/commit/66a9f242233318310c4f996c051c2e04b88f4e61))

## 0.1.0-alpha.1 (2025-03-27)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/hanzoai/go-sdk/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** api update ([#3](https://github.com/hanzoai/go-sdk/issues/3)) ([c10803d](https://github.com/hanzoai/go-sdk/commit/c10803d6ab4467d23a4fc874b0e704648a16291e))


### Chores

* configure new SDK language ([289851b](https://github.com/hanzoai/go-sdk/commit/289851bfe452a74b720a6895bf3fe9733948635a))
* go live ([#1](https://github.com/hanzoai/go-sdk/issues/1)) ([8b01e13](https://github.com/hanzoai/go-sdk/commit/8b01e1365f5eb1b2e771c5f561d65c336f118d78))
