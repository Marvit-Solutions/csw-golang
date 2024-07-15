package localrepository

type Exercise interface {
	FindSubModulesID() ([]int, error)
}
