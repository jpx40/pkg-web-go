package router

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jpx40/pkg-web-go/pkg/db"
	"github.com/jpx40/pkg-web-go/pkg/tmpl"
	//"github.com/jpx40/wdc/pkg/tmpl"
)

func Page(r chi.Router) {
	r.Get("/", templ.Handler(tmpl.Index(tmpl.Home())).ServeHTTP)
	r.Get("/chat", templ.Handler(tmpl.Index(tmpl.SimpleChat())).ServeHTTP)

	r.Get("/login", templ.Handler(tmpl.Index(tmpl.Login())).ServeHTTP)
	r.Get("/signup", templ.Handler(tmpl.Index(tmpl.Signup())).ServeHTTP)
	r.Get("/child", templ.Handler(tmpl.Index(tmpl.Child())).ServeHTTP)

	r.With(httpin.NewInput(&ListUserInput{})).Get("/table/{page}", ServeTable)
}

type ListUserInput struct {
	Page int `in:"path=page"`
}

func ServeTable(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value(httpin.Input).(*ListUserInput)
	conn := db.Connect()

	pkg := db.GetAllPackage(conn)

	pkgTable := chunkSlice(pkg, 15)

	index := int(input.Page)

	fmt.Print(input)
	// fmt.Print(pkg)
	templ.Handler(tmpl.Index(tmpl.PkgTable(pkgTable[index]))).ServeHTTP(w, r)
}

func chunkSlice(slice []db.Package, chunkSize int) [][]db.Package {
	var chunks [][]db.Package
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
