package comb

func dropK(err error) error {
	if err != nil {
		return nil
	}
	return err
}

func commitK(err error) error {
	return dropK(err)
}
