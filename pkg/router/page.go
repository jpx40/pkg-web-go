package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jpx40/pkg-web-go/pkg/db"
	"github.com/jpx40/pkg-web-go/pkg/tmpl"
	// "github.com/jpx40/wdc/pkg/tmpl"
)

func Page(r chi.Router) {
	r.Get("/", templ.Handler(tmpl.Index(tmpl.Home())).ServeHTTP)
	r.Get("/chat", templ.Handler(tmpl.Index(tmpl.SimpleChat())).ServeHTTP)

	r.Get("/login", templ.Handler(tmpl.Index(tmpl.Login())).ServeHTTP)
	r.Get("/signup", templ.Handler(tmpl.Index(tmpl.Signup())).ServeHTTP)
	r.Get("/child", templ.Handler(tmpl.Index(tmpl.Child())).ServeHTTP)
	ServeTable(r)
}

func ServeTable(r chi.Router) {
	conn := db.Connect()

	pkg := db.GetAllPackage(conn)

	pkgTable := chunkSlice(pkg, 5)

	index := 50

	// fmt.Print(pkg)
	r.Get("/table", templ.Handler(tmpl.Index(tmpl.PkgTable(pkgTable[index]))).ServeHTTP)
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

//u = url
// func Room(r chi.Router) {
// 	r.Route("/room/{id}", func(r chi.Router) {
// 		r.Use(RoomCtx)
// 		r.Get("/", GetRoom)
// 	})
// 	r.With(RoomCtx).Get("/room/{roomSlug:[a-z-]+}", GetRoom)
// }
//
// func RoomCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var room *db.Room
// 		var err error
// 		conn := db.Connect()
//
// 		if roomID := chi.URLParam(r, "id"); roomID != "" {
// 			intVar, error := strconv.Atoi(roomID)
// 			if error != nil {
// 				log.Println(error)
// 			}
// 			rID := uint(intVar)
// 			room, err = db.GetRoom(conn, rID)
// 		}
// 		// } else {
// 		// 	room, err = db.GetRoomBySlug(chi.URLParam(r, "slug"))
// 		// }
// 		if err != nil {
// 			render.Render(w, r, ErrNotFound)
// 			return
// 		}
//
// 		ctx := context.WithValue(r.Context(), "room", room)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
//
// func GetRoom(w http.ResponseWriter, r *http.Request) {
// 	// Assume if we've reach this far, we can access the article
// 	// context because this handler is a child of the ArticleCtx
// 	// middleware. The worst case, the recoverer middleware will save us.
//
// 	// room := r.Context().Value("article").(*Article
//
// 	room := r.Context().Value("room").(*db.Room)
// 	tmpl.Index(tmpl.Room(room)).Render(r.Context(), w)
// }
//
// NOTE: as a thought, the request and response payloads for an Article could be the
// same payload type, perhaps will do an example with it as well.
// type ArticlePayload struct {
//   *Article
// }

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
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
