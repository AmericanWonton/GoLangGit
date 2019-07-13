//writerInterface
import (
	"fmt"
	"io"
	"os"
)

fmt.Println("hello World")
	fmt.Fprint(os.Stdout, "hello World")
	io.WriteString(os.Stdout, "hello World")