package app

func (env Environment) GetKeys() []string {
	keys := make([]string, 0, len(env))
	for k := range env {
		keys = append(keys, k)
	}

	return keys
}

func (env Environment) ToBytesMap() map[string][]byte {
	bytes := make(map[string][]byte)
	for k, v := range env {
		bytes[k] = []byte(v)
	}

	return bytes
}
