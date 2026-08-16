# Tasks — buildx-multiarch-image

- [~] Confirm buildx + a builder are available (docker buildx ls); colima on Apple-Silicon. — buildx CLI present (v0.35.0), but **colima daemon is not running** in this session, so no builder is live. Operator must `colima start` to run the image tasks below.
- [x] Add a `docker-buildx` Makefile target: `docker buildx build --platform linux/arm64 -t <image>:<tag> --push` (add `,linux/amd64` only if amd64 is still a target). — added `docker-buildx` (`PLATFORMS ?= linux/arm64,linux/amd64`, `--push`, tags `${IMG}` + `:latest`) plus a `--load` single-arch `docker-build` for dev. Rewrote `build/Dockerfile` as multi-stage so the binary compiles per `TARGETOS/TARGETARCH` inside the build (the old Dockerfile copied one prebuilt amd64 binary — broken for arm64).
- [ ] Verify the resulting image is a manifest list (`docker buildx imagetools inspect <image>:<tag>`). — **operator (needs colima running).**
- [x] Retire the arch-suffixed `docker-build-v1`/`docker-push-v1` targets and image names. — removed; `IMG_V1`→`IMG`, `VERSION_V1`→`VERSION`; `install`/`purge` updated to Helm v3 syntax and to depend on `docker-buildx`.
- [ ] Confirm the image runs on arm64 (Pi armv8 and/or Apple-Silicon). — **operator (needs colima/cluster).**

Verified here: `make -n docker-buildx`/`docker-build` expand correctly, no residual `-v1`
refs, Go build/vet/test/lint stay green (Makefile/Dockerfile changes don't touch Go).
