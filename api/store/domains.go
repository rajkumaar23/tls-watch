package store

import (
	"time"
)

type Domain struct {
	ID             uint64    `json:"id"`
	UserID         uint64    `json:"user_id"`
	Domain         string    `json:"domain"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	LastNotifiedAt time.Time `json:"last_notified_at"`
}

func CreateDomain(domain *Domain) error {
	_, err := DB.Exec("INSERT INTO domains (user_id, domain) VALUES (?, ?)", domain.UserID, domain.Domain)
	if err != nil {
		return err
	}
	return nil
}

func GetDomainByUserAndValue(user_id uint64, domain string) (*Domain, error) {
	var resultDomain Domain
	row := DB.QueryRow("SELECT * FROM domains WHERE user_id = ? AND domain = ?", user_id, domain)
	if err := row.Scan(&resultDomain.ID, &resultDomain.UserID, &resultDomain.Domain, &resultDomain.CreatedAt, &resultDomain.UpdatedAt, &resultDomain.LastNotifiedAt); err != nil {
		return nil, err
	}
	return &resultDomain, nil
}

func DeleteDomainByUserAndValue(user_id uint64, domain string) error {
	_, err := DB.Exec("DELETE FROM domains WHERE user_id = ? AND domain = ?", user_id, domain)
	if err != nil {
		return err
	}
	return nil
}

func GetAllDomainsByUserID(user_id uint64) (*[]Domain, error) {
	result, err := DB.Query("SELECT * FROM domains WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}

	var domains []Domain
	for result.Next() {
		var domain Domain
		err = result.Scan(&domain.ID, &domain.UserID, &domain.Domain, &domain.CreatedAt, &domain.UpdatedAt, &domain.LastNotifiedAt)
		if err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}

	return &domains, nil
}

func GetAllDomains() (*[]Domain, error) {
	result, err := DB.Query("SELECT * FROM domains")
	if err != nil {
		return nil, err
	}

	var domains []Domain
	for result.Next() {
		var domain Domain
		err = result.Scan(&domain.ID, &domain.UserID, &domain.Domain, &domain.CreatedAt, &domain.UpdatedAt, &domain.LastNotifiedAt)
		if err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}

	return &domains, nil
}