package scheduled_message

import (
	"reflect"
	"testing"
	"time"
)

func TestNewScheduledMessage(t *testing.T) {
	type args struct {
		id                   uint64
		text                 string
		chatID               uint64
		userID               uint64
		scheduledSendingTime time.Time
	}

	validID := uint64(1)
	validText := "text"
	validChatID := uint64(1)
	validUserID := uint64(1)
	validScheduledSendingTime := time.Now().Add(1 * time.Second)

	validScheduledMessage, err := NewScheduledMessage(validID, validText, validChatID, validUserID, validScheduledSendingTime)
	if err != nil {
		t.Fatal(err)
		return
	}

	tests := []struct {
		name    string
		args    args
		want    *scheduledMessage
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:                   validID,
				text:                 validText,
				chatID:               validChatID,
				userID:               validUserID,
				scheduledSendingTime: validScheduledSendingTime,
			},
			want:    validScheduledMessage,
			wantErr: false,
		},
		{
			name: "異常系(textが空文字)",
			args: args{
				id:                   validID,
				text:                 "",
				chatID:               validChatID,
				userID:               validUserID,
				scheduledSendingTime: validScheduledSendingTime,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系(送信予定時刻が現在時刻)",
			args: args{
				id:                   validID,
				text:                 validText,
				chatID:               validChatID,
				userID:               validUserID,
				scheduledSendingTime: time.Now(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewScheduledMessage(tt.args.id, tt.args.text, tt.args.chatID, tt.args.userID, tt.args.scheduledSendingTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewScheduledMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheduledMessage() got = %v, want %v", got, validScheduledMessage)
			}
		})
	}
}
