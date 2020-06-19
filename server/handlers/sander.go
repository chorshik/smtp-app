package handlers

import (
	"encoding/json"
	"net/http"
	"net/smtp"
)

// HandlerSender ...
type HandlerSender struct {
}

// NewHandlerSender ...
func NewHandlerSender() *HandlerSender {
	return &HandlerSender{}
}

func (h *HandlerSender) Send() http.HandlerFunc {
	type request struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Sender   string `json:"sender"`
		Password string `json:"password"`
		Receiver string `json:"receiver"`
		Subject  string `json:"subject"`
		Msg      string `json:"msg"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		tmp := new(request)
		if err := json.NewDecoder(r.Body).Decode(tmp); err != nil {
			newError(w, r, http.StatusBadRequest, err)
			return
		}

		emailAuth := smtp.PlainAuth(
			"",
			tmp.Sender,
			tmp.Password,
			tmp.Host,
		)

		msg := []byte("To:" + tmp.Receiver + "\r\n" + "Subject:" + tmp.Subject + "\r\n" + tmp.Msg)

		err := smtp.SendMail(
			tmp.Host+":"+tmp.Port,
			emailAuth,
			tmp.Sender,
			[]string{tmp.Receiver},
			msg,
		)

		if err != nil {
			newError(w, r, http.StatusBadRequest, err)
			return
		}

		Respond(w, r, http.StatusCreated, nil)

	}
}

//fmt.Println("Email Sent!")
