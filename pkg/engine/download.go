package engine

import (
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"jaypod/pkg/rss"
)

func download(podcast *rss.RssItem, rootdir string, dest string, basename string, incoming bool, sublog *slog.Logger) error {

	destDir := fmt.Sprintf("%s/%s", rootdir, dest)
	if err := os.MkdirAll(destDir, 0777); err != nil {
		return fmt.Errorf("failed creating %s: %v", destDir, err)
	}

	resp, err := http.Get(podcast.Url())
	if err != nil {
		return fmt.Errorf("failed getting %s: %v", podcast.Url(), err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response code from %s: %d: %s\n",
			podcast.Url(), resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	filenameWithExt := contentDispositionFilename(resp, sublog)
	if filenameWithExt == "" {
		filenameWithExt = filepath.Base(resp.Request.URL.Path)
	}

	fname, extension := split(filenameWithExt)

	if extension == "" {
		extension = podcast.ExtensionFromMimeType()
	}

	if basename != "" {
		fname = basename
	}

	fname = escape(fname)

	if len(fname) > 250 {
		fname = fname[:250]
	}

	fullpath := fmt.Sprintf("%s/%s.%s", destDir, fname, extension)
	exists := func(path string) bool {
		_, err := os.Lstat(fullpath)
		return err == nil
	}
	i := 1
	for exists(fullpath) {
		// File already exists.  Try to uniquify
		fullpath = fmt.Sprintf("%s/dupe%d_%s.%s", destDir, i, fname, extension)
		i++
	}

	out, err := os.OpenFile(fullpath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to create podcast file %s: %v", fullpath, err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		out.Close()
		os.Remove(fullpath)
		return fmt.Errorf("failed to write podcast file %s: %v", fullpath, err)
	}

	err = out.Close()
	if err != nil {
		return fmt.Errorf("failed to close podcast file %s: %v", fullpath, err)
	}

	err = os.Chtimes(fullpath, podcast.Date(), podcast.Date())
	if err != nil {
		return fmt.Errorf("failed to change times on  podcast file %s: %v", fullpath, err)
	}

	if incoming {
		incomingDir := fmt.Sprintf("%s/Incoming", rootdir)
		if err := os.MkdirAll(incomingDir, 0777); err != nil {
			return fmt.Errorf("failed creating %s: %v", incomingDir, err)
		}

		dst := fmt.Sprintf("%s/%s.%s", incomingDir, fname, extension)
		err := CopyFile(fullpath, dst)
		if err != nil {
			return fmt.Errorf("failed to copy %s to incoming: %v", fullpath, err)
		}

		err = os.Chtimes(dst, podcast.Date(), podcast.Date())
		if err != nil {
			return fmt.Errorf("failed to change times on  podcast file %s: %v", dst, err)
		}

	}

	return nil
}

func contentDispositionFilename(resp *http.Response, sublog *slog.Logger) string {
	contentDisposition := resp.Header.Get("content-disposition")
	if contentDisposition == "" {
		sublog.Debug("No content-disposition header")
		return ""
	} else if contentDisposition == "inline" {
		sublog.Debug("content-disposition inline")
		return ""
	}

	sublog = sublog.With("content-disposition", contentDisposition)

	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		sublog.Error("error parsing content-disposition",
			"err", err)
		return ""
	}

	fname := params["filename"]
	if fname == "" {
		sublog.Debug("no filename in content-disposition")
		return ""
	}

	return fname
}

func split(fname string) (string, string) {

	sep := strings.LastIndex(fname, ".")
	if sep < 0 {
		fmt.Printf("no extension in content-disposition filename %s\n", fname)
		return fname, ""
	}

	return fname[0:sep], fname[sep+1:]
}

func escape(fname string) string {
	fname = strings.ReplaceAll(fname, "/", ",")
	fname = strings.ReplaceAll(fname, ":", " -")
	fname = strings.ReplaceAll(fname, "'", "")
	fname = strings.ReplaceAll(fname, "\"", "")
	fname = strings.ReplaceAll(fname, "?", "")
	fname = strings.ReplaceAll(fname, "*", "-")
	fname = strings.ReplaceAll(fname, "|", "-")
	return fname
}

func CopyFile(src, dst string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()

	info, err := srcF.Stat()
	if err != nil {
		return err
	}

	dstF, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, info.Mode())
	if err != nil {
		return err
	}
	defer dstF.Close()

	if _, err := io.Copy(dstF, srcF); err != nil {
		return err
	}
	return nil
}
