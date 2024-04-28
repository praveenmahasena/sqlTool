package internal

import (
	"context"
	"errors"
	"os"
)

func migrate(f *YamlConfig) error {
	if len(os.Args) <= 1 {
		return errors.New("not enough Args provided")
	}

	switch {
	case os.Args[1] == "up" || os.Args[1] == "down":
		return do(os.Args[1])
	default:
		return errors.New("Invalid Args provided")
	}
}

func do(args string) error {
	var path, pathErr = os.Getwd()
	if pathErr != nil {
		return pathErr
	}
	var dir, err = os.ReadDir(path + "/schema/" + args)

	if err != nil {
		return err
	}

	for _, v := range dir {
		var file, fileErr = os.ReadFile(path + "/schema/" + args + "/" + v.Name())
		if fileErr != nil {
			return fileErr
		}
		var _, err = db.ExecContext(context.Background(), string(file))
		if err != nil {
			return err
		}

	}

	return nil
}
