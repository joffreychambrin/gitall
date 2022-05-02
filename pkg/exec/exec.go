package exec

import (
	"fmt"
	"log"
	cmdexec "os/exec"

	"jchambrin.fr/gitall/pkg/config"
)

func Exec(path string, args []string) {
	directories, err := config.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range directories {
		fmt.Printf("==> %s <==\n", dir)
		out, err := execCmd(args, dir)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}

}

func execCmd(str []string, dir string) ([]byte, error) {
	parts := append([]string{"-C", dir}, str...)
	cmd := cmdexec.Command("git", parts...)
	out, err := cmd.CombinedOutput()
	return out, err
}
