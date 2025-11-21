package environment

var (
	Development = "DEVELOPMENT"
	Env         = "ENV"
)

func IsDevelopment(val string) bool {
	return val == Development
}
