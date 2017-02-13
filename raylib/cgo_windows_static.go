// +build windows,static

package raylib

/*
#cgo windows LDFLAGS: -lglfw3 -lopengl32 -lgdi32 -lopenal -lwinmm -lole32
#cgo windows CFLAGS: -DPLATFORM_DESKTOP -DGRAPHICS_API_OPENGL_33 -DAL_LIBTYPE_STATIC
*/
import "C"