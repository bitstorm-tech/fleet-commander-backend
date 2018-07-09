package arango

func NewResources() (*Resources, error) {
	resources := new(Resources)
	if err := CreateDocument(resources); err != nil {
		return nil, err
	}

	return resources, nil
}

func GetResourcesForUser(user *User) *Resources {
	return nil
}
