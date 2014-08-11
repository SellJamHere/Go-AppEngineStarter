# Golang Starter for Google App Engine
 
A little [background](http://blog.golang.org/go-and-google-app-engine) on the Go runtime for App Engine.

Things to note:

* Go doesn't need to be installed on the system for development; the App Engine SDK is self-contained. (Unless you want to download 3rd party libraries automatically, which we do.)
* App Engine programs are run on a single thread. Goroutines and channels are supported, but they can't be parallelized. Multiple goroutines may run concurrently if the current routine is waiting on an external resource. (Multi-threaded support is expected in the future.)
* The source is uploaded to the App Engine and compiled there. (Go is the only compiled language on the App Engine.)

This starter pack is designed to use Golang's idiomatic [workspace structure](http://golang.org/doc/code.html#Workspaces). Go is built on the idea that developers will use many 3rd party libraries. As a result, the program's src directory is divided into many subdirectories, one for each platform a library may reside e.g. github.com, code.google.com, along with a directory for the local source you are writing (In this case, backend).
Normally, only the local source of a go program is included in source control. To facilitate App Engine compatability, however, the entire directory is included in the repository. Because the 3rd party libraries will now lie under source control, we must add their originating src directory to .gitignore (e.g. github.com, code.google.com).

This repository relies on [martini](https://github.com/go-martini/martini) for routing. It is set up with useful default middleware, but can easily be replaced with a router of your choosing.

## Local Build Instructions


* Install [Go](http://golang.org/doc/install) (Don't forget to add to PATH.)
  
  ```
  brew install go
  ```
  
* Install [Google App Engine SDK for Go](https://developers.google.com/appengine/downloads) (Don't forget to add to PATH.)

* Create working dir

  ```
  mkdir project_name
  cd project_name
  mkdir src bin pkg
  cd src
  ```

* Clone this repo
  
  If your github account has SSH set up:
  ```
  git clone git@github.com:SellJamHere/Recruitment_Backend.git src
  ```

  Otherwise:
  ```
  git clone https://github.com/SellJamHere/Recruitment_Backend.git src
  ```

* Set GOPATH

  ```
  cd ../
  export GOPATH=`pwd`
  ```

* Get Dependencies (Only needed the first time unless you add a new dep)

  ```
  go get ./...
  ```
  Ignore the warning `package appengine: unrecognized import path "appengine"`.

* Run

  ```
  goapp serve ./src
  ```

## Deploy to Google App Engine

* Deploy

  ```
  goapp deploy ./src
  ```

* Enter email/password at prompt
  * If terminal outputs "Use an application-specific password instead of your regular account password.", visit [App passwords](https://security.google.com/settings/security/apppasswords), and generate a new password.


## Important Notes

### Workspace Structure

Your workspace should be structured as follows:

```
project_name/
  |bin/
  |pkg/
  |src/
    |app.yaml
    |backend/         <-- this source directory
    |github.com/      <-- libs from github
    |other_site.com/  <-- libs from another site
```

`app.yaml` contains runtime information for App Engine. It lives one directory above the source files. Source files you write are in the `backend/` directory. When running `go get`, code will be downloaded into an appropriate directory, ie github libraries to `github.com`. 

### 3rd Party Libs and .gitignore

Downloaded 3rd party libs need to be ignored. Each directory containing 3rd party libs must be added to `.gitignore`.
