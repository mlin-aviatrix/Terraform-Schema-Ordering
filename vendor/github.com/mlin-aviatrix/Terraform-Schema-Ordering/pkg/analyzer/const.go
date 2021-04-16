package analyzer

var ordering = []string{
	"Type",
	"Elem",
	"Required",
	"Optional",
	"Computed",
	"Default",
	"DefaultFunc",
	"ForceNew",
	"Sensitive",
	"DiffSuppressFunc",
	"ValidateFunc",
	"InputDefault",
	"StateFunc",
	"MaxItems",
	"MinItems",
	"Set",
	"ComputedWhen", // Documentation says this does not work
	"Description",
	"ConflictsWith",
	"RequiredWith",
	"Deprecated",
	"Removed",
	"AtLeastOneOf", // Can't find in documentation but we are using in provider
}
