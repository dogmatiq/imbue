package imbue

// option is an implementation of all of the option interfaces.
type option struct {
	forContainer func(*Container)
}

func (o option) applyContainerOption(con *Container) {
	if o.forContainer != nil {
		o.forContainer(con)
	}
}
