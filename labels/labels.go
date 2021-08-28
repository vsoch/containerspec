package labels

// Labels for the supercontainers metadata

type Label struct {
	Key         string
	Value       string
	Allowed     []string
	Description string
}

var Labels = map[string]Label{
	"org.supercontainers.mpi": {
		Key:         "org.supercontainers.mpi",
		Allowed:     []string{"mpich", "openmpi", "unknown"},
		Description: "Required MPI support, ABI compatibility",
	},

	"org.supercontainers.gpu": {
		Key:         "org.supercontainers.gpu",
		Allowed:     []string{"cuda", "opencl", "rocm", "unknown"},
		Description: "Required GPU library support",
	},

	// TODO this can be validated by regex for semver when parsed
	"org.supercontainers.glibc": {
		Key:         "org.supercontainers.gpu",
		Description: "Specific version of GLIBC, in Semantic format XX.YY.Z",
	},
}
