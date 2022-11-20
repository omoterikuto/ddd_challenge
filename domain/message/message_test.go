package message

import (
	"reflect"
	"testing"
)

func TestNewMessage(t *testing.T) {
	type args struct {
		id     uint64
		text   string
		chatID uint64
		userID uint64
	}

	validID := uint64(1)
	validText := "text"
	validChatID := uint64(1)
	validUserID := uint64(1)

	validMessage, err := NewMessage(validID, validText, validChatID, validUserID)
	if err != nil {
		t.Fatal(err)
		return
	}

	tests := []struct {
		name    string
		args    args
		want    *message
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:     validID,
				text:   validText,
				chatID: validChatID,
				userID: validUserID,
			},
			want:    validMessage,
			wantErr: false,
		},
		{
			name: "異常系(textが空文字)",
			args: args{
				id:     validID,
				text:   "",
				chatID: validChatID,
				userID: validUserID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMessage(tt.args.id, tt.args.text, tt.args.chatID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() got = %v, want %v", got, validMessage)
			}
		})
	}
}
