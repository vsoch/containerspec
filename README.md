# Containerspec

In order to be able to intelligently match containers to systems, whether that
means pulling and then allocating to an appropriate node, or making a decision
to retrieve a container based on available architectures, the core of this need
is being able to understand what is inside the container, and how that is (or is not)
compatible with the host. Our base library of information needs to capture
enough metadata do address several common use cases:

 - Can I run X on my host? Can I optimally run it?
 - Are the libraries in my container compatible with my host?
 - Can I bind mount library X and it will still work?
 
## Frequently Asked Questions

### 1. Why should this be implemented in Go?

If you look around, many container orchestration and similar tools are implemented in Go.
It's the language of this ecosystem, even if that goes against what are considered traditional
academic languages (Python, C++, etc.) I also really love Go and I don't get enough
opportunity to use it. Another reason is strong typing and having data structures
conform to a know standard without needing libraries like jsonschema.


## Commands

The first thing you might want to do is extract metadata for a host. Technically,
the only difference between a host and container is the layer of abstraction. We'd want
to run the same thing on our host as in the container.

```bash
$ TODO
```
 
# TODO

- define "database" of architectures
- start with a basic container with some mpi and operating system
- create a grid that can generate the "same" container with different operating systems
- create a library and extract from the container:
 - analyze with container diff
 - get information about host
 - Need actual use cases of what people want to do.


## Previous Art

 - [archspec](https://github.com/archspec/archspec): by the Spack team, which provides a database of architectures to start with.
