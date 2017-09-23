package dsnotify

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"regexp"
	"syscall"
)

type DirectoryFunc func(*fsnotify.Event, error)


type DirectoryWatcher interface {
	AddDirectory(dir string) error
	FsWatcher() *fsnotify.Watcher
	RegisterFileRegex(rgx *regexp.Regexp)
	EventsChannel() chan fsnotify.Event
	ErrorsChannel() chan error
	DirectoryEvents() DirectoryEvents
	Watch(fun DirectoryFunc)
}

type DirectoryEvents interface {
	EventsChannel() chan fsnotify.Event
	ErrorsChannel() chan error
}

type directoryEvents struct {
	events chan fsnotify.Event
	errors chan error
}

func newDirectoryEvents(eventsChan chan fsnotify.Event, errorsChan chan error) DirectoryEvents {
	return directoryEvents{events: eventsChan, errors: errorsChan}
}

func (de directoryEvents) EventsChannel() chan fsnotify.Event {
	return de.events
}

func (de directoryEvents) ErrorsChannel() chan error {
	return de.errors
}

type directoryWatcher struct {
	*fsnotify.Watcher
	fileRegexes     []*regexp.Regexp
	events          chan fsnotify.Event
	errors          chan error
	directoryEvents DirectoryEvents
}

func NewDirectoryWatcher() (DirectoryWatcher, error) {
	w, err := fsnotify.NewWatcher()
	event := make(chan fsnotify.Event)
	errChan := make(chan error)
	return &directoryWatcher{
		Watcher:         w,
		events:          event,
		errors:          errChan,
		directoryEvents: newDirectoryEvents(event, errChan),
	}, err
}

func (dw *directoryWatcher) AddDirectory(dir string) error {
	return filepath.Walk(dir, dw.walkDirectory)
}

func (dw *directoryWatcher) RegisterFileRegex(rgx *regexp.Regexp) {
	dw.fileRegexes = append(dw.fileRegexes, rgx)
}

func (dw *directoryWatcher) FsWatcher() *fsnotify.Watcher {
	return dw.Watcher
}

func (dw *directoryWatcher) EventsChannel() chan fsnotify.Event {
	return dw.events
}
func (dw *directoryWatcher) ErrorsChannel() chan error {
	return dw.errors
}

func (dw *directoryWatcher) DirectoryEvents() DirectoryEvents {
	return dw.directoryEvents
}

func (dw *directoryWatcher) walkDirectory(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	return dw.Watcher.Add(path)
}

func (dw *directoryWatcher) matchesAny(name string) bool {
	for _, rgx := range dw.fileRegexes {
		if rgx.MatchString(name) {
			return true
		}
	}

	return false
}

func (dw *directoryWatcher) Watch(fun DirectoryFunc) {
	for {

		select {
		case ev := <-dw.Events:
			if ev.Op&fsnotify.Remove == fsnotify.Remove || ev.Op&fsnotify.Write == fsnotify.Write || ev.Op&fsnotify.Create == fsnotify.Create {
				println(ev.Name)
				//base := filepath.Base(ev.Name)

				// Assume it is a directory and track it.
				//if *flag_recursive == true && !flag_excludedDirs.Matches(ev.Name) {
				//	dw.Watcher.Add(ev.Name)
				//}

				if dw.matchesAny(ev.Name) {
					println("it matched something")
					fun(&ev, nil)
				}
			}

		case err := <-dw.Watcher.Errors:
			if v, ok := err.(*os.SyscallError); ok {
				if v.Err == syscall.EINTR {
					continue
				}
			}
			fun(nil, err)
		}
	}
}
