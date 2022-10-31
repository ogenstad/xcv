package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	columnWidth = 45
)

func FormatCertificate(c Certificate) string {
	doc := strings.Builder{}

	doc.WriteString(formatTitle(c))
	doc.WriteString(formatContent(c))
	doc.WriteString(formatFooter(c))

	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)

	return docStyle.Render(doc.String())
}

func formatTitle(crt Certificate) string {
	stl := Style{}
	doc := strings.Builder{}
	{
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			stl.activeTab().Render(crt.Subject.CommonName),
		)
		gap := stl.tabGap().Render(strings.Repeat(" ", max(0, stl.width()-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n\n")
	}

	{
		desc := lipgloss.JoinVertical(lipgloss.Left,
			stl.descStyle().Render(fmt.Sprintf("Version: %d", crt.Version)),
			stl.infoStyle().Render(fmt.Sprintf("Serial: %s", crt.SerialNumber)),
		)

		row := lipgloss.JoinHorizontal(lipgloss.Top, "x509 Certificate Viewer", desc)
		doc.WriteString(row + "\n\n")
	}

	{
		info := lipgloss.JoinHorizontal(lipgloss.Top,
			stl.list().Width(stl.width()-2).Render(
				lipgloss.JoinVertical(lipgloss.Left,
					"Subject:",
					fmt.Sprintf("%s\n", crt.Subject),
					"Issuer:",
					fmt.Sprintf("%s\n", crt.Issuer),
				),
			),
		)
		doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, info))
		doc.WriteString("\n")
	}

	if len(crt.cert.CRLDistributionPoints) > 0 {
		content := []string{"CRL Distribution Points:"}

		for i := range crt.cert.CRLDistributionPoints {
			content = append(content, fmt.Sprintf("%s %s", stl.divider(), crt.cert.CRLDistributionPoints[i]))
		}

		cdp := lipgloss.JoinHorizontal(lipgloss.Top,
			stl.list().Width(stl.width()-2).Render(
				lipgloss.JoinVertical(lipgloss.Left,
					content...,
				),
			),
		)

		doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, cdp))
		doc.WriteString("\n")
	}

	return doc.String()
}

func formatContent(crt Certificate) string {
	stl := Style{}
	doc := strings.Builder{}

	var leftColumn []string

	leftColumn = append(leftColumn, stl.listHeader("Validity"))
	leftColumn = append(leftColumn, stl.listEntry(fmt.Sprintf("From: %s", crt.NotBefore)))
	leftColumn = append(leftColumn, stl.listEntry(fmt.Sprintf("To:   %s", crt.NotAfter)))
	leftColumn = append(leftColumn, "\n")

	if len(crt.KeyUsage) > 0 {
		leftColumn = append(leftColumn, stl.listHeader("Key Usage"))
		for _, usage := range crt.KeyUsage {
			leftColumn = append(leftColumn, stl.listChecked(usage))
		}
	}

	leftColumn = append(leftColumn, "\n")

	if len(crt.ExtendedKeyUsage) > 0 {
		leftColumn = append(leftColumn, stl.listHeader("Extended Key Usage"))
		for _, eku := range crt.ExtendedKeyUsage {
			leftColumn = append(leftColumn, stl.listChecked(eku.Name))
		}
	}

	if len(crt.UnsupportedExtensions) > 0 {
		leftColumn = append(leftColumn, stl.listHeader("Unsupported Extensions"))
		for _, extension := range crt.UnsupportedExtensions {
			leftColumn = append(leftColumn, stl.listEntry(extension))
		}
	}

	lists := lipgloss.JoinHorizontal(lipgloss.Top,
		stl.list().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				leftColumn...,
			),
		),
		stl.list().Copy().Width(columnWidth).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				stl.listHeader("Public Key"),
				stl.listItem(fmt.Sprintf("%s", crt.PublicKey)),
			),
		),
	)

	doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, lists))
	doc.WriteString("\n\n")

	return doc.String()
}

func formatFooter(c Certificate) string {
	stl := Style{}
	doc := strings.Builder{}

	if c.UnsupportedProperties {
		w := lipgloss.Width

		statusKey := stl.statusStyle().Render("UNSUPPORTED")
		statusVal := stl.statusText().Copy().
			Width(stl.width() - w(statusKey)).
			Render("Unable to parse all fields check for updates at: " + stl.url("https://github.com/ogenstad/xcv"))

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusKey,
			statusVal,
		)

		doc.WriteString(stl.statusBarStyle().Width(stl.width()).Render(bar))
	}

	return doc.String()
}
