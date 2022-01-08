package vagrant

import (
	"fmt"
	"os"
	"testing"
)

var commonDirInOutMap = []struct {
	input         []string
	expecteResult string
	name          string
}{
	{
		[]string{
			".output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Hard Disks%[1]vubuntu2004.vhdx",
			".output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.VMRS",
			".output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmcx",
			".output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmgs",
			".output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vbox.xml",
		},
		".output%[1]vhyperv%[1]vubuntu2004",
		"common_relative_dir_with_spaces_in_path",
	},
	{
		[]string{
			".output1%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Hard Disks%[1]vubuntu2004.vhdx",
			".output2%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.VMRS",
			".output3%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmcx",
			".output4%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmgs",
			".output5%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vbox.xml",
		},
		"",
		"no_common_dir_with_spaces_in_path",
	},
	{
		[]string{
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Hard Disks%[1]vubuntu2004.vhdx",
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.VMRS",
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmcx",
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmgs",
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vbox.xml",
		},
		"%[1]v.output%[1]vhyperv%[1]vubuntu2004",
		"common_abs_dir_with_spaces_in_path",
	},
	{
		[]string{
			"%[1]v.output1%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Hard Disks%[1]vubuntu2004.vhdx",
			"%[1]v.output2%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.VMRS",
			"%[1]v.output3%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmcx",
			"%[1]v.output4%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vDAB721DD-D397-4CD0-9315-70227198BABE.vmgs",
			"%[1]v.output5%[1]vhyperv%[1]vubuntu2004%[1]vVirtual Machines%[1]vbox.xml",
		},
		"",
		"no_common_abs_dir_with_spaces_in_path",
	},
	{
		[]string{
			"%[1]va%[1]vb%[1]vc%[1]vd%[1]ve%[1]vfile.txt",
			"%[1]va%[1]vb%[1]vc%[1]vd%[1]vfile.txt",
			"%[1]va%[1]vb%[1]vvery_long_path_galore%[1]vfile.txt",
		},
		"%[1]va%[1]vb",
		"long_path_ with_minimal_path_segments",
	},
	{
		[]string{
			"%[1]va%[1]vf%[1]vc%[1]vd%[1]ve%[1]vfile.txt",
			"%[1]va%[1]vx%[1]vc%[1]vd%[1]vfile.txt",
			"%[1]va%[1]vj%[1]vc%[1]vfile.txt",
		},
		"%[1]va",
		"single_common_dir",
	},
	{
		[]string{
			"%[1]v.output%[1]vhyperv%[1]vubuntu2004%[1]vubuntu2004.vhdx",
		},
		"%[1]v.output%[1]vhyperv%[1]vubuntu2004",
		"single_path_input",
	},
	{
		[]string{},
		"",
		"no_paths_input",
	},
}

func TestPathHelper_GetCommonDirectory(t *testing.T) {
	pathSeperator := string(os.PathSeparator)
	for _, testData := range commonDirInOutMap {

		var input []string
		for _, in := range testData.input {
			formattedPath := fmt.Sprintf(in, pathSeperator)
			input = append(input, formattedPath)
		}

		expecteResult := testData.expecteResult
		if expecteResult != "" {
			expecteResult = fmt.Sprintf(testData.expecteResult, pathSeperator)
		}

		t.Run(testData.name, func(t *testing.T) {
			actualResult := getCommonDirectory(input, pathSeperator)
			if actualResult != expecteResult {
				t.Errorf("Failed with input: %q. Expected result: %q. Actual result: %q", input, expecteResult, actualResult)
			}
		})
	}
}
