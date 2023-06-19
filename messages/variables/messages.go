package variables

var (
	// [ variables ]
	Usage            = "variables"
	ShortDescription = "Create variables for edges on Azion's platform"
	LongDescription  = "Create variables for edges on Azion's platform"
	FlagHelp         = "Displays more information about the variables command"
	FlagId           = "Unique identifier of the Variable"

	// [ describe ]
	DescribeUsage            = "describe --variable-id <variable_id> [flags]"
	DescribeShortDescription = "Returns the variable data"
	DescribeLongDescription  = "Displays information about the variable via a given ID to show the variable's attributes in detail"
	DescribeFlagVariableID   = "Unique identifier for an variable. The '--variable-id' flag is mandatory"
	DescribeFlagOut          = "Exports the output to the given <file_path/file_name.ext>"
	DescribeFlagFormat       = "Changes the output format passing the json value to the flag"
	DescribeHelpFlag         = "Displays more information about the describe command"

	// [ general ]
	FileWritten = "File successfully written to: %s\n"

	// [ list ]
	VariablesListUsage            = "list [flags]"
	VariablesListShortDescription = "Displays your variables"
	VariablesListLongDescription  = "Displays all variables related to your applications"
	VariablesListHelpFlag         = "Displays more information about the list subcommand"

	// [ delete ]
	DeleteOutputSuccess    = "Variable %v was successfully deleted\n"
	DeleteHelpFlag         = "Displays more information about the delete subcommand"
	DeleteUsage            = "delete [flags]"
	DeleteShortDescription = "Delete a Variable"
	DeleteLongDescription  = "Delete a Variable using UUID"
)
