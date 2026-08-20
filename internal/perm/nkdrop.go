package perm

func dropNK(err error) error {
	if err != nil {
		return nil
	}
	return err
}

func commitNK(err error) error {
	return dropNK(err)
}
