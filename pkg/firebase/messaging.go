package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

// FCMClient menyimpan instance Firebase Cloud Messaging client
type FCMClient struct {
	client *messaging.Client
}

var FcmInstance *FCMClient

// NewFCMClient menginisialisasi dan mengembalikan FCM client baru
func NewFCMClient(credentialsFile string) (*FCMClient, error) {
	// Setup Firebase
	opt := option.WithCredentialsFile(credentialsFile)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	// Dapatkan FCM client
	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &FCMClient{client: client}, nil
}

// SendToDevice mengirim notifikasi ke perangkat tertentu
func (f *FCMClient) SendToDevice(token, title, body string, data map[string]string) (string, error) {
	// Buat message payload
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	// Tambahkan data payload jika ada
	if data != nil {
		message.Data = data
	}

	// Kirim message
	return f.client.Send(context.Background(), message)
}

// SendToTopic mengirim notifikasi ke topik tertentu
func (f *FCMClient) SendToTopic(topic, title, body string, data map[string]string) (string, error) {
	// Buat message payload
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: topic,
	}

	// Tambahkan data payload jika ada
	if data != nil {
		message.Data = data
	}

	// Kirim message
	return f.client.Send(context.Background(), message)
}

// SubscribeToTopic menambahkan perangkat ke topik
func (f *FCMClient) SubscribeToTopic(tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return f.client.SubscribeToTopic(context.Background(), tokens, topic)
}

// UnsubscribeFromTopic menghapus perangkat dari topik
func (f *FCMClient) UnsubscribeFromTopic(tokens []string, topic string) (*messaging.TopicManagementResponse, error) {
	return f.client.UnsubscribeFromTopic(context.Background(), tokens, topic)
}

func InitFCM() {
	var err error
	// Path ke file service account Firebase
	fcm, err := NewFCMClient("firebase_service_account.json")
	if err != nil {
		log.Fatalf("Failed to initialize FCM: %v", err)
	}
	FcmInstance = fcm
}
