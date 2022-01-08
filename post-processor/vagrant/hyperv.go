package vagrant

import (
	"fmt"
	"os"
	"path/filepath"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type HypervProvider struct{}

func (p *HypervProvider) KeepInputArtifact() bool {
	return false
}

func (p *HypervProvider) Process(ui packersdk.Ui, artifact packersdk.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	// Create the metadata
	metadata = map[string]interface{}{"provider": "hyperv"}

	// ui.Message(fmt.Sprintf("artifacts all: %+v", artifact))
	// ui.Message(fmt.Sprintf("dir: %+v", dir))

	// Vargant requires specific dir structure for hyperv
	// hyperv builder creates the structure in the output dir
	// we have to keep the structure in a temp dir
	var outputDir string = getCommonDirectory(artifact.Files(), string(os.PathSeparator))
	if len(outputDir) <= 1 {
		return "", nil, fmt.Errorf("hyperv artifacts must have a common parent folder")
	}
	// ui.Message(fmt.Sprintf("artifact dir from string: %s", outputDir))

	// Copy all of the original contents into the temporary directory
	for _, path := range artifact.Files() {
		ui.Message(fmt.Sprintf("Copying: %s", path))

		var rel string

		rel, err = filepath.Rel(outputDir, filepath.Dir(path))
		// ui.Message(fmt.Sprintf("rel is: %s", rel))

		if err != nil {
			ui.Message(fmt.Sprintf("err in: %s", rel))
			return
		}

		dstDir := filepath.Join(dir, rel)
		// ui.Message(fmt.Sprintf("dstdir is: %s", dstDir))
		if _, err = os.Stat(dstDir); err != nil {
			if err = os.MkdirAll(dstDir, 0755); err != nil {
				ui.Message(fmt.Sprintf("err in creating: %s", dstDir))
				return
			}
		}

		dstPath := filepath.Join(dstDir, filepath.Base(path))

		// We prefer to link the files where possible because they are often very huge.
		// Some filesystem configurations do not allow hardlinks. As the possibilities
		// of mounting different devices in different paths are flexible, we just try to
		// link the file and copy if the link fails, thereby automatically optimizing with a safe fallback.
		if err = LinkFile(dstPath, path); err != nil {
			// ui.Message(fmt.Sprintf("err in linking: %s to %s", path, dstPath))
			if err = CopyContents(dstPath, path); err != nil {
				ui.Message(fmt.Sprintf("err in copying: %s to %s", path, dstPath))
				return
			}
		}

		ui.Message(fmt.Sprintf("Copied %s to %s", path, dstPath))
	}

	return
}
