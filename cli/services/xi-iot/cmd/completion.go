package cmd

import (
	"os"

	"xi-iot-cli/xi-iot/io_utils"

	"github.com/spf13/cobra"
)

const (
	completionFileNameBash = "completion.bash"
	completionFileNameZsh  = "completion.zsh"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Create completion file for bash and zsh. Using this command for zsh is experimental and not formally supported.",
	Long: io_utils.LongDesc(`
	This command writes autocompletion scripts for the bash and zsh shells
	
	# Bash
	This command writes autocompletion scripts for the bash shell
	with the filename completion.bash in the current working directory.

	To use autocompletion, first enable autocompletion on your machine:
	1. brew install bash-completion
	2. echo  "if [ -f $(brew --prefix)/etc/bash_completion ]; then
	    . $(brew --prefix)/etc/bash_completion
	   fi" >> ~/.bash_profile
	3. source ~/.bash_profile

	# Zsh
	This command writes autocompletion scripts for the zsh shell with the filename completion.zsh
	in the current working directory. 
	Note that flags do not work on zsh. See https://github.com/spf13/cobra/issues/107.
	To use the autocompletion feature with the oh-my-zsh shell, do the following. 

	mkdir $HOME/.oh-my-zsh/completion && \
	  mv completion.zsh $HOME/.oh-my-zsh/completions/_xi-iot
	
	Make sure that the path is in your $fpath: fpath=($HOME/.oh-my-zsh/completions $fpath)`),
	Example: io_utils.Examples(`
	# Generate autocompletion for bash
	xi-iot completion bash

	# Generate autocompletion for zsh
	xi-iot completion bash`),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// completionCmd represents the completion command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Create completion file for the bash shell",
	Long: io_utils.LongDesc(`
	This command writes autocompletion scripts for the bash shell
	with the filename completion.bash in the current working directory.

	To use autocompletion, first enable autocompletion on your machine:
	1. brew install bash-completion
	2. echo  "if [ -f $(brew --prefix)/etc/bash_completion ]; then
	    . $(brew --prefix)/etc/bash_completion
	   fi" >> ~/.bash_profile
	3. source ~/.bash_profile`),
	Example: io_utils.Examples(`
	# Generate autocompletion for bash
	xi-iot completion bash`),
	Run: generateBashCompletionFiles,
}

// completionCmd represents the completion command
var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Create completion file for the zsh shell (specifically, oh-my-zsh)",
	Long: io_utils.LongDesc(`
	This command writes autocompletion scripts for the zsh shell with the filename completion.zsh
	in the current working directory. 
	Note that flags do not work on zsh. See https://github.com/spf13/cobra/issues/107.

	To use the autocompletion feature with the oh-my-zsh shell, do the following. 
	mkdir $HOME/.oh-my-zsh/completion && \
	  mv completion.zsh $HOME/.oh-my-zsh/completions/_xi-iot
	
	Make sure that the path is in your $fpath: fpath=($HOME/.oh-my-zsh/completions $fpath)`),
	Example: io_utils.Examples(`
	# Generate autocompletion for bash
	xi-iot completion zsh`),
	Run: generateZshCompletionFiles,
}

// generateCompletionFiles creates bash and zsh completion files
func generateBashCompletionFiles(cmd *cobra.Command, args []string) {
	rootCmd.GenBashCompletionFile(completionFileNameBash)
	os.Chmod(completionFileNameBash, 0777)
}

// generateZshCompletionFiles creates bash and zsh completion files
func generateZshCompletionFiles(cmd *cobra.Command, args []string) {
	rootCmd.GenZshCompletionFile(completionFileNameZsh)
	os.Chmod(completionFileNameZsh, 0777)
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(zshCmd)
	completionCmd.AddCommand(bashCmd)
}
