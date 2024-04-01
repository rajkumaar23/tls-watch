package store

import "time"

type NotificationSetting struct {
	ID             uint64    `json:"id"`
	UserID         uint64    `json:"user_id"`
	Provider       string    `json:"provider"`
	ProviderUserID string    `json:"provider_user_id"`
	Enabled        bool      `json:"enabled"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func CreateNotificationSetting(setting *NotificationSetting) error {
	_, err := DB.Exec(
		"INSERT INTO notification_settings (user_id, provider, provider_user_id, enabled) VALUES (?, ?, ?, ?)",
		setting.UserID, setting.Provider, setting.ProviderUserID, setting.Enabled,
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNotificationSetting(setting *NotificationSetting) error {
	_, err := DB.Exec(
		"UPDATE notification_settings SET provider_user_id = ?, enabled = ? WHERE user_id = ? AND provider = ?",
		setting.ProviderUserID, setting.Enabled, setting.UserID, setting.Provider,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetNotificationSettingByUserAndProvider(user_id uint64, provider string) (*NotificationSetting, error) {
	var setting NotificationSetting
	row := DB.QueryRow("SELECT * FROM notification_settings WHERE user_id = ? AND provider = ?", user_id, provider)
	if err := row.Scan(
		&setting.ID, &setting.UserID, &setting.Enabled, &setting.Provider, &setting.ProviderUserID, &setting.CreatedAt, &setting.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &setting, nil
}

func GetAllNotificationSettingsByUserID(user_id uint64) (*[]NotificationSetting, error) {
	result, err := DB.Query("SELECT * FROM notification_settings WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}

	var settings []NotificationSetting
	for result.Next() {
		var setting NotificationSetting
		err = result.Scan(
			&setting.ID, &setting.UserID, &setting.Enabled, &setting.Provider, &setting.ProviderUserID, &setting.CreatedAt, &setting.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		settings = append(settings, setting)
	}

	return &settings, nil
}
