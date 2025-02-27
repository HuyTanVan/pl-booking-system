package mail

import (
	"fmt"
	"os"
	"path/filepath"
	"plbooking_go_structure1/global"
	"testing"

	// "github.com/HuyTanVan/soccer_booking_ticket/util"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func getConfigPath() string {
	wd, _ := os.Getwd()                                          // Get current working directory
	return filepath.Join(wd, "..", "..", "config", "local.yaml") // Go up two directories to reach 'config'
}

func TestSendEmailWithGmail(t *testing.T) {
	viper.SetConfigFile(getConfigPath())
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	err := viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration struct: %w", err))
	}
	config := global.Config.EmailSender
	// require.NoError(t, err)
	fmt.Println("MY CONFIG:", config)
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Quan trọng"
	content := `
	<p>Hello e, trai cụa anh </p>
	<p>This is a test message from Huy Dep Trai </p>
	`
	to := []string{"huyhandsome189@gmail.com"}
	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
