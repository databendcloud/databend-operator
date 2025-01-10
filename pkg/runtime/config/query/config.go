package query

type Config struct {
}

func (b *QueryTomlBuilder) QueryConfig() (*Config, error) {
	return &Config{}, nil
}
