package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var inputFile = flag.String("input", "", "Input Lua file")
	var outputFile = flag.String("output", "", "Output Lua file (defaults to overwrite input)")
	flag.Parse()

	if *inputFile == "" {
		fmt.Fprintf(os.Stderr, "Error: input file is required\n")
		flag.Usage()
		os.Exit(1)
	}

	// Read input file
	content, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", *inputFile, err)
		os.Exit(1)
	}

	// Transform the Lua code
	transformedCode := transformLuaCode(string(content))

	// Determine output file
	output := *outputFile
	if output == "" {
		output = *inputFile
	}

	// Write transformed code
	err = ioutil.WriteFile(output, []byte(transformedCode), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", output, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully transformed %s\n", *inputFile)
}

func transformLuaCode(code string) string {
	// Regular expression to match function definitions with type annotations
	// Pattern: function name(param1:type1, param2:type2, ...)
	functionRegex := regexp.MustCompile(`(?m)^(\s*)((?:local\s+)?function\s+[^(]+)\(([^)]*)\)(.*)$`)
	
	lines := strings.Split(code, "\n")
	var result []string
	
	for i, line := range lines {
		if functionRegex.MatchString(line) {
			// This is a function definition line
			matches := functionRegex.FindStringSubmatch(line)
			if len(matches) >= 5 {
				indent := matches[1]
				funcDef := matches[2]
				params := matches[3]
				rest := matches[4]
				
				// Parse parameters and extract type annotations
				transformedParams, typeChecks := parseParametersWithTypes(params, i+1)
				
				// Reconstruct function definition without type annotations
				newFuncLine := indent + funcDef + "(" + transformedParams + ")" + rest
				result = append(result, newFuncLine)
				
				// Add type checking code after function definition
				for _, typeCheck := range typeChecks {
					result = append(result, indent + "  " + typeCheck)
				}
			} else {
				result = append(result, line)
			}
		} else {
			result = append(result, line)
		}
	}
	
	return strings.Join(result, "\n")
}

func parseParametersWithTypes(params string, lineNum int) (string, []string) {
	if strings.TrimSpace(params) == "" {
		return "", nil
	}
	
	paramList := strings.Split(params, ",")
	var cleanedParams []string
	var typeChecks []string
	
	for _, param := range paramList {
		param = strings.TrimSpace(param)
		if param == "..." {
			cleanedParams = append(cleanedParams, param)
			continue
		}
		
		// Check if parameter has type annotation (param:type)
		if strings.Contains(param, ":") {
			parts := strings.SplitN(param, ":", 2)
			if len(parts) == 2 {
				paramName := strings.TrimSpace(parts[0])
				paramType := strings.TrimSpace(parts[1])
				
				// Add cleaned parameter (without type annotation)
				cleanedParams = append(cleanedParams, paramName)
				
				// Create type check
				typeCheck := fmt.Sprintf(`if (type(%s) ~= "%s") then error("ArgumentException from %s in line %d") end`,
					paramName, paramType, paramName, lineNum)
				typeChecks = append(typeChecks, typeCheck)
			} else {
				cleanedParams = append(cleanedParams, param)
			}
		} else {
			cleanedParams = append(cleanedParams, param)
		}
	}
	
	return strings.Join(cleanedParams, ", "), typeChecks
}
