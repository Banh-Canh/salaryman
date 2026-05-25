package image

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// EmbedLocal turns a local image reference into an inline base64 data URL.
//
// It leaves the input unchanged when src is empty or already a remote
// (http/https) or data URL. Otherwise src is treated as a filesystem path:
// "file://" prefixes are stripped, and relative paths are resolved against
// baseDir. The file is then read and encoded as "data:<mime>;base64,...".
//
// This lets users reference an image next to their resume JSON
// (e.g. "data/portrait.png") without having to inline base64 by hand.
func EmbedLocal(src, baseDir string) (string, error) {
	if src == "" {
		return src, nil
	}
	if strings.HasPrefix(src, "http://") ||
		strings.HasPrefix(src, "https://") ||
		strings.HasPrefix(src, "data:") {
		return src, nil
	}

	path := strings.TrimPrefix(src, "file://")
	if !filepath.IsAbs(path) {
		path = filepath.Join(baseDir, path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("embed local image %q: %w", path, err)
	}

	mime := mimeFromExt(filepath.Ext(path))
	return fmt.Sprintf("data:%s;base64,%s", mime, base64.StdEncoding.EncodeToString(data)), nil
}

func mimeFromExt(ext string) string {
	switch strings.ToLower(ext) {
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
