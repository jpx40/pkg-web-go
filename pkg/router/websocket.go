package router

//"github.com/jpx40/wdc/pkg/util"

// func Start_Websocket(r chi.Router) {
// 	m := melody.New()
// 	defer m.Close()
//
// 	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		err := m.HandleRequest(w, r)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	})
// 	m.HandleMessage(func(s *melody.Session, msg []byte) {
// 		out, err := util.ParseJson(msg)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(out.Msg))
//
// 		var msgdb db.Message
// 		msgdb.Content = out.Msg
// 		//	msgdb.UserID, _ = strconv.Itoa(out.UserID)
// 		//	msgdb.RoomID, _ = strconv.Atoi(out.RoomID)
// 		//	msgdb.CreatedAt = out.Time
// 		msgdb.UserID = uint(rand.Intn(1000))
// 		msgdb.RoomID = uint(rand.Intn(1000))
//
// 		error := m.Broadcast(msg)
// 		if error != nil {
// 			fmt.Println(err)
// 		}
// 	})
// }
