package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/brockweaver/hxu/pkg/file"
	"github.com/spf13/cobra"
)

func CobraCommands(parent *cobra.Command) *cobra.Command {

	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "commands for manipulating files",
	}
	parent.AddCommand(fileCmd)

	convertCmd := &cobra.Command{
		Use:   "convert [inputFile] [outputFile]",
		Run:   cobraConvert,
		Short: "Converts native fixed-width or tab-delimited Helix files to different formats",
		Args:  cobra.MaximumNArgs(2),
	}

	convertCmd.Flags().String("format", "jsonlines", "format to emit [jsonlines,tsv,csv]")
	convertCmd.Flags().Bool("strict", false, "error if unexpected bytes appear in input file content lines")
	convertCmd.Flags().Bool("preserve", false, "preserve exact original content sans insignificant whitespace; do not apply conversions")
	convertCmd.Flags().Bool("noheader", false, "suppress emitting the header line")
	convertCmd.Flags().Bool("windows", false, "emit windows-style line endings (\\r\\n) instead of unix-style (\\n)")
	convertCmd.Flags().Bool("help", false, "display help content")
	fileCmd.AddCommand(convertCmd)

	verifyCmd := &cobra.Command{
		Use:   "verify [inputFile]",
		Run:   cobraVerify,
		Short: "Verifies native fixed-width or tab-delimited Helix files",
		Args:  cobra.MaximumNArgs(1),
	}
	verifyCmd.Flags().Bool("strict", false, "error if unexpected bytes appear in input file content lines")
	fileCmd.AddCommand(verifyCmd)

	generateCmd := &cobra.Command{
		Use:   "generate [inputFile] [outputFile]",
		Run:   cobraGenerate,
		Short: "Consume a jsonlines, csv, or tsv file and generate the appropriate native Helix fixed-width format file. Applicable only to file types which are uploaded to Helix (see --type flag for details)",
		Args:  cobra.MaximumNArgs(2),
	}
	generateCmd.Flags().String("type", "autodetect", "type of the input file [autodetect,transfer,accountlock,accountunlock]")
	generateCmd.Flags().String("format", "autodetect", "format of the input file [autodetect,jsonlines,csv,tsv]")
	generateCmd.Flags().Bool("help", false, "display help content")
	generateCmd.Flags().Bool("detect", false, "detects and reports input file metadata and verifies its contents. does not emit content")
	generateCmd.Flags().Bool("strict", false, "error if unexpected or invalid values appear in input file content lines")
	fileCmd.AddCommand(generateCmd)

	return fileCmd

}

func cobraGenerate(cmd *cobra.Command, args []string) {
	fmt.Printf("much to do here!\n")
}

func cobraVerify(cmd *cobra.Command, args []string) {

	strict, _ := cmd.Flags().GetBool("strict")

	var infile *os.File
	var err error

	input := "stdin"

	if len(args) == 0 {
		infile = os.Stdin
	} else {
		input = args[0]
		infile, err = os.Open(input)
		if err != nil {
			log.Fatalf("error reading file %s: %v", input, err.Error())
			os.Exit(10)
		}
		defer infile.Close()
	}

	inscan := bufio.NewScanner(infile)
	outbuf := bufio.NewWriter(os.Stdout)

	if err := file.Process(inscan, outbuf, file.FormatVerifyOnly, strict, false, false, false); err != nil {
		log.Fatalf("error processing file %s: %v", input, err.Error())
		os.Exit(11)
	}
}

func cobraConvert(cmd *cobra.Command, args []string) {
	format, _ := cmd.Flags().GetString("format")
	strict, _ := cmd.Flags().GetBool("strict")
	preserve, _ := cmd.Flags().GetBool("preserve")
	verifyOnly, _ := cmd.Flags().GetBool("verifyonly")
	noheader, _ := cmd.Flags().GetBool("noheader")
	windows, _ := cmd.Flags().GetBool("windows")

	var infile *os.File
	var outfile *os.File
	var err error

	input := "stdin"

	if len(args) == 0 {
		infile = os.Stdin
	} else {
		input = args[0]
		infile, err = os.Open(input)
		if err != nil {
			log.Fatalf("error reading file %s: %v", input, err.Error())
			os.Exit(1)
		}
		defer infile.Close()
	}

	inscan := bufio.NewScanner(infile)

	output := "stdout"
	if len(args) < 2 {
		outfile = os.Stdout
	} else {
		output = args[1]
		outfile, err := os.Open(output)
		if err != nil {
			log.Fatalf("error writing file %s: %v", output, err.Error())
			os.Exit(2)
		}
		defer outfile.Close()
	}

	outbuf := bufio.NewWriter(outfile)

	var outputFormat = file.FormatJsonLines

	switch strings.ToLower(format) {
	case "jsonlines":
		outputFormat = file.FormatJsonLines
	case "csv":
		outputFormat = file.FormatCSV
	case "tsv":
		outputFormat = file.FormatTsv
	case "verifyonly":
		outputFormat = file.FormatVerifyOnly
	default:
		log.Fatalf("unsupported format '%v'", format)
		os.Exit(4)
	}

	// (--verifyonly flag overrides --format flag)
	if verifyOnly {
		outputFormat = file.FormatVerifyOnly
	}

	if err := file.Process(inscan, outbuf, outputFormat, strict, preserve, noheader, windows); err != nil {
		log.Fatalf("error processing file %s: %v", input, err.Error())
		os.Exit(3)
	}
}
