package main

import (
	"context"
	"fmt"

	"distri/internal/core"
)

type EmailRequest struct {
	UUID      string   `json:"uuid"`
	UserUUIDs []string `json:"uuids"`
}

// main here it's only for concept of the app, not an actual usable command.
func main() {
	handler := core.NewHandler("email_notifications_handler")

	if err := handler.RegisterWorkflow(
		"on_email_notification_request",
		func(ctx context.Context, d *core.Distri, param EmailRequest) {
			sessionsResult := d.NewActivity(ctx, FetchSessions, core.MaybeWithVal(param))

			sendEmailsResult := d.NewActivity(ctx, SendEmails, sessionsResult)

			onEmailContactResult := d.NewActivity(ctx, onEmailContact, sendEmailsResult)

			d.Done(ctx, onEmailContactResult)
		},
	); err != nil {
		panic(err)
	}

	handler.Run(context.Background())
	//e := &core.Exec{
	//	WorkflowKey: "on_email_notification_request",
	//	Param:       EmailRequest{},
	//	D:           &core.Distri{},
	//}
	//
	//if err := handler.Execute(context.Background(), e); err != nil {
	//	panic(err)
	//}

	fmt.Println("test")
}

func onEmailContact(ctx context.Context, param core.Maybe) core.Maybe {
	return core.Maybe{}
}

func SendEmails(ctx context.Context, param core.Maybe) core.Maybe {
	return core.Maybe{}
}

func FetchSessions(ctx context.Context, param core.Maybe) core.Maybe {
	return core.Maybe{}
}
