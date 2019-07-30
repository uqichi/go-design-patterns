package command

type addStringCommand interface {
	execute()
}

type addPrefixCommand struct {
	text   string
	prefix string
}

func (cmd *addPrefixCommand) execute() {
	cmd.text = cmd.prefix + cmd.text
}

type addSuffixCommand struct {
	text   string
	suffix string
}

func (cmd *addSuffixCommand) execute() {
	cmd.text += cmd.suffix
}
