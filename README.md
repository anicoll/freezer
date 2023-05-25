freezer
=======

Freezer is a simple message batch storage and replay interface, intended to provide message queue like semantics, but with a blob storage backend.

Messages are batched, optionally compressed and stored when written, and correspondingly uncompressed and unbatched when read back.

freezer uses [straw](https://godoc.org/github.com/anicoll/straw) as a blob storage abstraction.

