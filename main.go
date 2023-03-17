package main

import (
	"cranime/extractor"
	"fmt"
)

func main() {
	result, err := extractor.ExtractorMyAnimeList("anime/51535")
	fmt.Println(result, err)
}
