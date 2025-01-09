package query

type QueryConfig struct {

}

func (b *QueryTomlBuilder) QueryConfig() (*QueryConfig, error) {
	return &QueryConfig{}, nil
}
