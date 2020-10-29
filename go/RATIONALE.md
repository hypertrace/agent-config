# Rationale

## Why not using the protoc object?

`protoc` object fulfills most of the needs we have for config struct, however we wanted to avoid coupling transport with the domain object. Also practically there are some issues related:

- messaging fields being exposed
- mutex logic which makes it hard to clone the object to sort mutability out
