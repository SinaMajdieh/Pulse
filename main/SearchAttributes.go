package main
type Name struct{
	Status string
	Type string
	Value string
}
type Number struct {
	Status string
	Type string
	Value int
	ValueFrom int
	ValueTo int

}
type Float struct {
	Status string
	Type string
	Value float64
	ValueFrom float64
	ValueTo float64
}
type Gender struct {
	Status string
	Type string
}
type MobilePhoneNumber struct {
	Status string
	Type string
	Prefix string
	Value string
}
type SearchAttributes struct{
	SFFN Name
	SFLN Name
	SFG Gender
	SFMPN MobilePhoneNumber
	SFEA Name
	SFJ Name
	SFJR Number
	SFA Number
	SFI Number
	SFB Float
	SFCN Name
	SFCP Name
}
