package email

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Define various email templates
	EmailTemplateCompanyAccount    string = "company_account"
	EmailTemplateOperatorAccount   string = "operator_account"
	EmailTemplateEmailVerification string = "email_verification"

	EmailTemplateBasePath string = "/email_templates"
)

// EmailParams email parameter type, using map to support arbitrary parameters
type EmailParams map[string]string

// EmailTemplate email template structure
type EmailTemplate struct {
	Subject     string
	TextContent string
	HtmlContent string
}

// GetEmailTemplate get email template content
func GetEmailTemplate(templateType string, params EmailParams) (*EmailTemplate, error) {
	subject, textContent, htmlContent, err := buildEmailContent(templateType, params)
	if err != nil {
		return nil, err
	}

	return &EmailTemplate{
		Subject:     subject,
		TextContent: textContent,
		HtmlContent: htmlContent,
	}, nil
}

// buildEmailContent build email content
func buildEmailContent(template string, params EmailParams) (subject, textContent, htmlContent string, err error) {

	switch template {
	case EmailTemplateCompanyAccount:
		subject = "Your company has been successfully created"
		textContent, htmlContent, err = loadTemplateFiles("company_account_template", params)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to load company account template: %w", err)
		}

	case EmailTemplateOperatorAccount:
		operatorName := params["operatorName"]
		if operatorName == "" {
			operatorName = "your brand"
		}
		subject = fmt.Sprintf("Your new brand %s has been successfully created", operatorName)
		textContent, htmlContent, err = loadTemplateFiles("operator_account_template", params)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to load operator account template: %w", err)
		}

	case EmailTemplateEmailVerification:
		subject = "Please verify your email address"
		textContent, htmlContent, err = loadTemplateFiles("email_verification_template", params)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to load email verification template: %w", err)
		}
	// Can add more email templates
	default:
		return "", "", "", fmt.Errorf("unknown email template: %s", template)
	}

	return subject, textContent, htmlContent, nil
}

// loadTemplateFiles load template files and replace parameters
func loadTemplateFiles(templateFileName string, params EmailParams) (textContent, htmlContent string, err error) {
	// Build template file paths
	textPath := filepath.Join(EmailTemplateBasePath, templateFileName+".txt")
	htmlPath := filepath.Join(EmailTemplateBasePath, templateFileName+".html")

	// Read text template
	textBytes, err := os.ReadFile(textPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read text template %s: %w", textPath, err)
	}
	textContent = string(textBytes)

	// Read HTML template
	htmlBytes, err := os.ReadFile(htmlPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read html template %s: %w", htmlPath, err)
	}
	htmlContent = string(htmlBytes)

	// Replace parameters
	textContent = replaceTemplateVars(textContent, params)
	htmlContent = replaceTemplateVars(htmlContent, params)

	return textContent, htmlContent, nil
}

// replaceTemplateVars replace template variables - supports arbitrary parameters
func replaceTemplateVars(content string, params EmailParams) string {
	if params == nil {
		return content
	}

	result := content
	for key, value := range params {
		// Support multiple placeholder formats
		placeholders := []string{
			fmt.Sprintf("{%s}", key), // {key}
			// fmt.Sprintf("{{%s}}", key), // {{key}}
			// fmt.Sprintf("${%s}", key),  // ${key}
			// fmt.Sprintf("$%s", key),    // $key
		}

		for _, placeholder := range placeholders {
			result = strings.ReplaceAll(result, placeholder, value)
		}
	}

	return result
}
