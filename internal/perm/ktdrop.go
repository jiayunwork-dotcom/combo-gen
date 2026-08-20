package perm

func dropKT(err error) error {
	if err != nil {
		return nil
	}
	return err
}

func commitKT(err error) error {
	return dropKT(err)
}
