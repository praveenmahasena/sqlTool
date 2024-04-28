package internal

func Start() error {
	var f, err = readConfig()
	if err != nil {
		return err
	}

	err = Connect(f)
	if err != nil {
		return err
	}

	return migrate(f)
}
