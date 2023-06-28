package example

import (
	"context"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
)

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())
}
