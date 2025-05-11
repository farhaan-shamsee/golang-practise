package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("This will read and update the yaml file")
	updateFile("./score.yaml")
}

func updateFile(filepath string) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	decoder := yaml.NewDecoder(strings.NewReader(string(data)))

	fmt.Println("Decoding YAML file...")
	// fmt.Println("decoder:", decoder) // decoder: &{0xc000004408 false}

	var documents []*yaml.Node

	for {
		var doc yaml.Node
		if err := decoder.Decode(&doc); err != nil {
			break
		}
		documents = append(documents, &doc)
		// fmt.Println("Decoded document:", doc) // Decoded document: <yaml.Node>
	}

	// Step 4: Iterate through the documents
	for _, doc := range documents {
		// Expect each document to be a mapping (object at root)
		if doc.Kind != yaml.DocumentNode || len(doc.Content) == 0 { // Skip if not a document node or empty
			fmt.Println("Skipping non-document node or empty document")
			continue
		}
		root := doc.Content[0] // Assuming the first content is the root node

		var kindNode, specNode *yaml.Node

		// Step 4a: Find `kind` and `spec` nodes
		for i := 0; i < len(root.Content); i += 2 {
			key := root.Content[i]
			value := root.Content[i+1]

			if key.Value == "kind" {
				kindNode = value
			}
			if key.Value == "spec" {
				specNode = value
			}
		}

		// Step 4b: Only change if it's GitRepository kind
		if kindNode != nil && kindNode.Value == "GitRepository" && specNode != nil {
			// Step 4c: Find and update `url` field under `spec`
			for i := 0; i < len(specNode.Content); i += 2 {
				key := specNode.Content[i]
				value := specNode.Content[i+1]

				if key.Value == "url" {
					fmt.Println("Original URL:", value.Value)
					value.Value = "ssh://git@code.siemens.com/my-test/dev-cluster.git"
					fmt.Println("Updated URL:", value.Value)
					break // Exit loop after updating the URL
				}
			}
		}
	}

	// Step 5: Open file to overwrite
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()
	encoder := yaml.NewEncoder(f)
	encoder.SetIndent(2) // Set indentation for better readability
	// Step 6: Encode the updated documents back to the file
	for _, doc := range documents {
		if err := encoder.Encode(doc); err != nil {
			fmt.Println("Error encoding document:", err)
			return
		}
	}
	fmt.Println("YAML file updated successfully.")
	// Step 7: Close the encoder
	if err := encoder.Close(); err != nil {
		fmt.Println("Error closing encoder:", err)
		return
	}
	fmt.Println("Encoder closed successfully.")
	// Step 8: Close the file
	if err := f.Close(); err != nil {
		fmt.Println("Error closing file:", err)
		return
	}
	fmt.Println("File closed successfully.")
	// Step 9: Print success message
	fmt.Println("YAML file updated successfully.")
}
