package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

const pattern = `"([A-Za-z0-9-_]{22})"`

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewLogsPipelinesApi(apiClient)

	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if len(matches) <= 1 {
			lines = append(lines, line)
			continue
		}

		id := matches[1]
		p, _, err := api.GetLogsPipeline(ctx, id)
		if err != nil {
			logger.Warn("failed to get log pipeline", "id", id, "error", err)
			lines = append(lines, line)
			continue
		}

		newLine := strings.ReplaceAll(line, id, fmt.Sprintf("%s (%s)", p.GetName(), id))
		lines = append(lines, newLine)
	}

	if err := scanner.Err(); err != nil {
		logger.Warn("failed to scan stdin", "error", err)
		os.Exit(1)
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
