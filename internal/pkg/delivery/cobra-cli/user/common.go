package user

import (
	"fmt"
	"github.com/spf13/cobra"
	"jirno/internal/pkg/domain"
)

func parseUser(cmd *cobra.Command, args []string) (*domain.DeliveryUser, error) {
	res := &domain.DeliveryUser{}

	if cmd.Flag("name") != nil {
		parsedName, err := cmd.Flags().GetString("name")
		if err != nil {
			return nil, fmt.Errorf("name flag error: %v", err)
		}
		res.Name = parsedName
	}

	if cmd.Flag("nickname") != nil {
		parsedNickname, err := cmd.Flags().GetString("nickname")
		if err != nil {
			return nil, fmt.Errorf("nickname flag error: %v", err)
		}
		res.Nickname = parsedNickname
	}

	if cmd.Flag("email") != nil {
		parsedEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return nil, fmt.Errorf("email flag error: %v", err)
		}
		res.Email = parsedEmail
	}

	if cmd.Flag("password") != nil {
		parsedPassword, err := cmd.Flags().GetString("password")
		if err != nil {
			return nil, fmt.Errorf("password flag error: %v", err)
		}
		res.Password = parsedPassword
	}
	return res, nil
}

func parseUserUpdate(cmd *cobra.Command, args []string) (*domain.DeliveryUserUpdate, error) {
	res := &domain.DeliveryUserUpdate{}
	parsedUser, err := parseUser(cmd, args)
	if err != nil {
		return nil, err
	}
	res.DeliveryUser = *parsedUser
	parsedNewNickname, err := cmd.Flags().GetString("newnickname")
	if err != nil {
		return nil, fmt.Errorf("newnickname flag error: %v", err)
	}
	res.NewNickname = parsedNewNickname
	return res, nil
}
