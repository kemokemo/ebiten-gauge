module github.com/kemokemo/ebiten-gauge/cmd

replace local.packages/gauge => ../

go 1.19

require (
	github.com/hajimehoshi/ebiten/v2 v2.5.0
	local.packages/gauge v0.0.0-00010101000000-000000000000
)

require (
	github.com/ebitengine/purego v0.3.0 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/jezek/xgb v1.1.0 // indirect
	golang.org/x/exp/shiny v0.0.0-20220826205824-bd9bcdd0b820 // indirect
	golang.org/x/image v0.6.0 // indirect
	golang.org/x/mobile v0.0.0-20230301163155-e0f57694e12c // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)
