package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Skills struct {
	Name  string
	Items []string
}

type UserData struct {
	Languages             []string
	Frameworks            []string
	DeveloperTools        []string
	ActivelyUsingLearning []string
	GithubObsessionURL    string
}

// read JSON file
func readJSONFile(filePath string) (*UserData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var userData UserData
	if err := json.Unmarshal(bytes, &userData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
	}
	return &userData, nil
}

// format the user data into the structs
func formatUserData(userData *UserData) []Skills {
	var skills []Skills
	skills = append(skills, Skills{Name: "Languages", Items: userData.Languages})
	skills = append(skills, Skills{Name: "Frameworks", Items: userData.Frameworks})
	skills = append(skills, Skills{Name: "Developer Tools", Items: userData.DeveloperTools})
	skills = append(skills, Skills{Name: "Actively Using", Items: userData.ActivelyUsingLearning})
	return skills
}

// generates Trophy data for GitHub
func generateTrophyData() string {
	var sb strings.Builder
	sb.WriteString(`<p align="center">`)
	sb.WriteString(`<img src="https://github-profile-trophy.vercel.app/?username=ozep&column=-1&theme=darkhub&no-frame=true" alt="GitHub Trophy Case" />`)
	sb.WriteString(`</p>`)
	return sb.String()
}

// generates a Markdown table for skills
func generateSkillTable(categories []Skills) string {
	var sb strings.Builder
	sb.WriteString(`<table align="center">`)

	sb.WriteString("<tr>")
	for _, category := range categories {
		sb.WriteString("<th>")
		sb.WriteString(category.Name)
		sb.WriteString("</th>")
	}
	sb.WriteString("</tr>\n")

	sb.WriteString("<tr>")
	for _, category := range categories {
		sb.WriteString("<td colspan=\"1\" align=\"center\">")
		sb.WriteString(`<img src="https://skillicons.dev/icons?i=`)
		sb.WriteString(strings.Join(category.Items, ","))
		sb.WriteString(`&perline=6" alt="skills icons"/>`)
		sb.WriteString("</td>")

	}
	sb.WriteString("</tr>\n")
	sb.WriteString("</table>\n")
	return sb.String()
}

func generateContactInfo() string {
	var sb strings.Builder
	sb.WriteString(`<div align="center">`)
	sb.WriteString(`
	<a href="https://www.youtube.com/@PrepWithZep" target="_blank">
		<img src="https://img.shields.io/static/v1?message=Youtube&logo=youtube&label=&color=FF0000&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="youtube logo"  />
	</a>
	<a href="jissa023@uottawa.ca" target="_blank">
		<img src="https://img.shields.io/static/v1?message=Gmail&logo=gmail&label=&color=D14836&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="gmail logo"  />
	</a>
	<a href="https://www.linkedin.com/in/joey-issa/" target="_blank">
		<img src="https://img.shields.io/static/v1?message=LinkedIn&logo=linkedin&label=&color=0077B5&logoColor=white&labelColor=&style=for-the-badge" height="35" alt="linkedin logo"  />
	</a>
	</div>`)
	sb.WriteString("\n")
	return sb.String()
}

func printGitStats() string {
	var sb strings.Builder
	sb.WriteString(`<div align="center">
		<p align="center"> <img src="https://komarev.com/ghpvc/?username=jjoeyissa&label=Profile%20views&color=blueviolet&style=plastic" alt="ozep" /> </p>
		<img src="https://github-readme-stats.vercel.app/api?username=oZep&hide_title=false&hide_rank=false&show_icons=true&include_all_commits=true&count_private=true&disable_animations=false&theme=blue-green&locale=en&hide_border=true" height="150" alt="stats graph"  />
		<img src="https://github-readme-stats.vercel.app/api/top-langs?username=oZep&locale=en&hide_title=false&layout=compact&card_width=320&langs_count=7&theme=blue-green&hide_border=true&hide=shaderlab" height="150" alt="languages graph"  />
	</div>`)
	sb.WriteString("\n")
	return sb.String()
}

// ? repo cards one day ?

// combines all sections into a complete README and writes it to a file
func WriteReadme(categories []Skills) {
	var sb strings.Builder
	sb.WriteString(generateTrophyData())
	sb.WriteString("\n")
	sb.WriteString(printGitStats())
	sb.WriteString("\n")
	// small space
	sb.WriteString(`<p></p>`)
	sb.WriteString(generateSkillTable(categories))
	sb.WriteString("\n")
	sb.WriteString(`<div align="center"><img src="https://raw.githubusercontent.com/oZep/oZep/output/snake.svg" alt="Snake animation" /></div>`)
	sb.WriteString("\n")
	sb.WriteString(`<img align="center" height="200" width="800" src="./img/fish.webp" alt="Fish" />`)
	sb.WriteString("\n<hr>\n")
	sb.WriteString(generateContactInfo())

	readmeContent := sb.String()
	readmePath := filepath.Join("..", "README.md")
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		fmt.Printf("failed to write README.md: %v\n", err)
	}
	fmt.Println("README.md generated successfully at", readmePath)
}

func main() {
	userData, err := readJSONFile("info.json")
	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}
	categories := formatUserData(userData)

	WriteReadme(categories)
}
