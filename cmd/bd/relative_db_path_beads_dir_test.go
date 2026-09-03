package main

import (
	"os"
	"path/filepath"
	"testing"
)

// A RELATIVE --db path must resolve to the same .beads directory as the
// equivalent absolute path.
//
// Regression for gt-irl: the upward walk in resolveCommandBeadsDir terminated
// one iteration early for relative paths, because filepath.Dir(".") == "." makes
// the loop guard false at "." and the body never ran for the current directory.
// So "./.beads" was never tested, the function fell through to
// filepath.Dir(dbPath), no config was found there, and bd silently opened the
// DEFAULT-NAMED database instead of the one named on the command line -- then
// reported success (rc=0) against an empty store.
//
// Measured on a real workspace before the fix, same cwd and second:
//
//	bd --db .beads/embeddeddolt/gt list --status open           -> 0 issues + warning
//	bd --db /abs/.beads/embeddeddolt/gt list --status open      -> 192 issues, no warning
func TestResolveCommandBeadsDirRelativePathFindsCwdBeadsDir(t *testing.T) {
	tmp := t.TempDir()
	// Resolve symlinks: t.TempDir() can return /var/... where /var -> /private/var,
	// which would make the absolute-vs-relative comparison fail for the wrong reason.
	if resolved, err := filepath.EvalSymlinks(tmp); err == nil {
		tmp = resolved
	}

	beadsDir := filepath.Join(tmp, ".beads")
	dbDir := filepath.Join(beadsDir, "embeddeddolt", "gt")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	origWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(origWD) })
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	relative := filepath.Join(".beads", "embeddeddolt", "gt")
	got := resolveCommandBeadsDir(relative)

	if got != beadsDir {
		t.Fatalf("resolveCommandBeadsDir(%q) = %q, want %q\n"+
			"a relative --db path must find ./.beads; falling back to the db path's "+
			"parent makes bd open the default-named database and report success "+
			"against an empty store (gt-irl)", relative, got, beadsDir)
	}

	// The absolute form already worked; it must keep working and agree.
	if abs := resolveCommandBeadsDir(dbDir); abs != beadsDir {
		t.Fatalf("resolveCommandBeadsDir(%q) = %q, want %q", dbDir, abs, beadsDir)
	}
}
