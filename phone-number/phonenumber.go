package phonenumber

import (
	"fmt"
	"unicode"
)

type PhoneNumber struct {
	Number     string
	Area       string
	Exchange   string
	Subscriber string
}

func (p *PhoneNumber) String() string {
	return fmt.Sprintf("(%s) %s-%s", p.Area, p.Exchange, p.Subscriber)
}

func Number(number string) (string, error) {

	var p PhoneNumber
	if err := p.parse(number); err != nil {
		return "", err
	}

	return p.Number, nil
}

func AreaCode(number string) (string, error) {
	var p PhoneNumber
	if err := p.parse(number); err != nil {
		return "", err
	}

	return p.Area, nil
}

func Format(number string) (string, error) {
	var p PhoneNumber

	if err := p.parse(number); err != nil {
		return "", err
	}

	return p.String(), nil
}

func (p *PhoneNumber) parse(number string) error {

	var num string
	for _, v := range number {
		if unicode.IsSpace(v) {
			continue
		}
		if unicode.IsPunct(v) {
			if v == '(' || v == ')' || v == '+' || v == '-' || v == '.' {
				continue
			}

			return fmt.Errorf("character not valid: %c", v)
		}

		if unicode.IsNumber(v) {
			num += string(v)
		}
	}

	if len(num) <= 9 || len(num) > 11 {
		return fmt.Errorf("length %d not valid", len(num))
	}

	if len(num) == 11 {
		if num[0] != '1' {
			return fmt.Errorf("invalid country code: %c", num[0])
		}

		num = num[1:]
	}

	p.Number = num

	if p.Number[0] == '1' || p.Number[0] == '0' {
		return fmt.Errorf("invalid area code: %c", num[0])
	}

	if p.Number[3] == '1' || p.Number[3] == '0' {
		return fmt.Errorf("invalid exchange code: %c", num[0])
	}

	p.Area = num[0:3]
	p.Exchange = num[3:6]
	p.Subscriber = num[6:len(num)]

	return nil
}
