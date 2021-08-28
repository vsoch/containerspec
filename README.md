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

## Usage

To build the library:

```bash
$ make
```

This will compile an executable, `containerspec` that you can interact with.

## Commands

### Host

The first thing you might want to do is extract metadata for a host. Technically,
the only difference between a host and container is the layer of abstraction. We'd want
to run the same thing on our host as in the container.

```bash
$ $ ./containerspec host
{x86 [] [generic] 0 [] map[]}
```

### Container Update

**Under development** I'd like to have commands that can read a Dockerfile, or
a directory / repository of `Dockerfile`s, and be able to tell us:

1. Is the digest up to date?
2. Are there new tags we might want to build?

And then this can be integrated into a GitHub action. The reason this is packaged
here is because if this tool can parse and reason about `Dockerfile`, we want to
make sure that we add the right metadata (for now, `LABEL`s) that can then inform
about later compatibility.

I'm not decided yet about how to implement this, but I want the following:

For base images and updating them:

1. User can target a Dockerfile directly for one off update, or a folder with tags for scaled
2. Read in Dockerfile, keep track of labels and FROMS (add dockerfile parser)
3. Look at labels to see if a tag is there for the hash
4. For each FROM, look up list of tags, update hash
5. For each FROM, if a label exists after it for opencontainers, delete it
6. Update label to use new tag

Try to use the labels here with [opencontainers labels](https://github.com/opencontainers/image-spec/blob/main/annotations.md)
So far we have the following labels demonstrated in [this paper](https://conferences.computer.org/sc19w/2019/pdfs/CANOPIE-HPC2019-7nd7J7oXBlGtzeJIHi79mM/3AyZkyZVlhldzPU6UEo655/3OqM2Lkt9DqiE2sDu1jvaS.pdf):

| Label | Values | Comment |
|-------|--------|---------|
| org.supercontainers.mpi| {mpich,openmpi} |Required MPI support, ABI compatibility|
| org.supercontainers.gpu| {cuda,opencl,rocm, etc} |Required GPU library support|
| org.supercontainers.glibc| Semantic version: XX.YY.Z |Specific version of GLIBC|

 
# TODO for host matching

- archspec/archspec-go doesn't have parsing / matching functions?
- start with a basic container with some mpi and operating system
- create a grid that can generate the "same" container with different operating systems
- create a library and extract from the container:
 - analyze with container diff
 - get information about host
 - Need actual use cases of what people want to do.


## Previous Art

 - [archspec](https://github.com/archspec/archspec): by the Spack team, which provides the database of architectures to start with. We want to use [archspec-go](https://github.com/archspec/archspec-go) which I don't think is done yet.
