package discovery

import "log"

// Exec 함수는 consul 엔트리를 생성하고, 이를 요청한다.
func Exec(cli Client) error {
	if err := cli.Register([]string{"Go", "Awesome"}); err != nil {
		return err
	}

	entries, _, err := cli.Service("discovery", "Go")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		log.Printf("%#v\n", entry.Service)
	}

	return nil
}
