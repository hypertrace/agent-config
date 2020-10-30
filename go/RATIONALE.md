# Rationale

## Why not using the protoc object?

`protoc` object fulfills most of the needs we have for config struct, however we wanted to avoid coupling between the transport and the domain object. Also practically there are some issues related:

- messaging fields being exposed
- mutex logic (for concurrent writes) which makes it hard to clone the object to sort mutability out

Specifically in Go, mutability is important as users may try to override existing configuration in runtime.
