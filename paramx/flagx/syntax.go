package flagx

// 39.78.1 Syntax
type FtLoadFlags = Flag

const (
	FtLoadFlags_default                     FtLoadFlags = "default"
	FtLoadFlags_no_scale                    FtLoadFlags = "no_scale"
	FtLoadFlags_no_hinting                  FtLoadFlags = "no_hinting"
	FtLoadFlags_render                      FtLoadFlags = "render"
	FtLoadFlags_no_bitmap                   FtLoadFlags = "no_bitmap"
	FtLoadFlags_vertical_layout             FtLoadFlags = "vertical_layout"
	FtLoadFlags_force_autohint              FtLoadFlags = "force_autohint"
	FtLoadFlags_crop_bitmap                 FtLoadFlags = "crop_bitmap"
	FtLoadFlags_pedantic                    FtLoadFlags = "pedantic"
	FtLoadFlags_ignore_global_advance_width FtLoadFlags = "ignore_global_advance_width"
	FtLoadFlags_no_recurse                  FtLoadFlags = "no_recurse"
	FtLoadFlags_ignore_transform            FtLoadFlags = "ignore_transform"
	FtLoadFlags_monochrome                  FtLoadFlags = "monochrome"
	FtLoadFlags_linear_design               FtLoadFlags = "linear_design"
	FtLoadFlags_no_autohint                 FtLoadFlags = "no_autohint"
)
