package twirpql

type file struct {
	Service *service // TODO: multiple services
	Types   []*serviceType
	Inputs  []*serviceType
	Enums   []*enums
	Scalars []string
}

type service struct {
	Methods []*method
}

type enums struct {
	Name   string
	Fields []string
}

type serviceType struct {
	Name   string
	Fields []*serviceField
}

type serviceField struct {
	Name string
	Type string
}

type method struct {
	Name, Request, Response string
	FormattedRequest        string
}