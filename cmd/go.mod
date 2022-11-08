module github.com/kemokemo/ebiten-gauge/cmd

replace local.packages/gauge => ../

go 1.19

require (
	github.com/hajimehoshi/ebiten/v2 v2.3.8
	local.packages/gauge v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220806181222-55e207c401ad // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/jezek/xgb v1.0.0 // indirect
	golang.org/x/exp/shiny v0.0.0-20220826205824-bd9bcdd0b820 // indirect
	golang.org/x/image v0.0.0-20220722155232-062f8c9fd539 // indirect
	golang.org/x/mobile v0.0.0-20220722155234-aaac322e2105 // indirect
	golang.org/x/sync v0.0.0-20220819030929-7fc1605a5dde // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
)
