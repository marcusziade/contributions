package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/machinebox/graphql"
)

const query = `
query ($username: String!, $from: DateTime!, $to: DateTime!) {
    user(login: $username) {
        contributionsCollection(from: $from, to: $to) {
            contributionCalendar {
                weeks {
                    contributionDays {
                        date
                        contributionCount
                    }
                }
            }
        }
    }
}
`

type ContributionDay struct {
	Date              string `json:"date"`
	ContributionCount int    `json:"contributionCount"`
}

type Week struct {
	ContributionDays []ContributionDay `json:"contributionDays"`
}

type ContributionCalendar struct {
	Weeks []Week `json:"weeks"`
}

type ContributionsCollection struct {
	ContributionCalendar ContributionCalendar `json:"contributionCalendar"`
}

type User struct {
	ContributionsCollection ContributionsCollection `json:"contributionsCollection"`
}

type Response struct {
	User User `json:"user"`
}

func main() {
	client := graphql.NewClient("https://api.github.com/graphql")
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))

	username := "marcusziade"
	now := time.Now()

	var allContributionDays []ContributionDay

	for i := 0; i < 5; i++ {
		from := now.AddDate(-i-1, 0, 0).Format(time.RFC3339)
		to := now.AddDate(-i, 0, 0).Format(time.RFC3339)

		req.Var("username", username)
		req.Var("from", from)
		req.Var("to", to)

		var resp Response
		if err := client.Run(context.Background(), req, &resp); err != nil {
			log.Fatalf("Error fetching contributions: %v", err)
		}

		for _, week := range resp.User.ContributionsCollection.ContributionCalendar.Weeks {
			allContributionDays = append(allContributionDays, week.ContributionDays...)
		}
	}

	contributionData := make(map[string]int)
	for _, day := range allContributionDays {
		contributionData[day.Date] = day.ContributionCount
	}

	file, err := os.Create("contributions.csv")
	if err != nil {
		log.Fatalf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Date", "Contributions"})

	for date, count := range contributionData {
		if date == "" || count == 0 {
			log.Printf("Skipping invalid record: Date='%s', Contributions=%d", date, count)
			continue
		}
		log.Printf("Writing record: Date='%s', Contributions=%d", date, count)
		if err := writer.Write([]string{date, fmt.Sprintf("%d", count)}); err != nil {
			log.Printf("Error writing record: %v", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("Error flushing CSV writer: %v", err)
	}

	fmt.Println("Contribution data written to contributions.csv")

	if err := verifyCSV("contributions.csv"); err != nil {
		log.Fatalf("Error verifying CSV file: %v", err)
	}

	cmd := exec.Command("python3", "visualize_contributions.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error running Python script: %v", err)
	}
}

// verifyCSV checks if the CSV file has the expected content
func verifyCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return fmt.Errorf("CSV file is empty")
	}
	if len(records[0]) != 2 || records[0][0] != "Date" || records[0][1] != "Contributions" {
		return fmt.Errorf("CSV file header is incorrect")
	}

	fmt.Println("CSV file verification passed")
	return nil
}

