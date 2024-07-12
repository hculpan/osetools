This document aims to give a step-by-step list of the changes a developer
will have to make to this project to use it as the basis for their project.

# Changes

## Fork the template repo
Fork the template into a new project
## Update go.mod
Set the module to the name to the new project
## Environment Configuration
Copy `example.env` to `.env` and then update with appropriate values.
Note that `.env` is included in `.gitignore`, so it will not be saved to the Git repo.
## Update imports
Imports in the `handlers`, `templates`, and other packages point to `github.com/hculpan/go-web-base-project-1`. Update these to point to the new repo.
## Update components.Head
The `cmd/web/templates/components/Head` includes links to external libraries like Bootstrap. Update these, if necessary.
## Update Makefile
Update the web and cli application names in `Makefile`. 
## Update `.gitignore`
Update the `.gitignore` file to exclude the executable files.
## Update .air.toml
The `bin=` entry in this file should be changed to the executable file name of the
application.
## CLI app
If you do not need a CLI app, you can just ignore the `cmd/cli` package altogether
or remove it, if you prefer. If you do so, update the `Makefile` to no longer build it. Nothing anywhere else depends on it.
## Update `db.RetrieveUserHash()`
The `interal/auth/db.go` has a function stub, `RetrieveUserHash`. This simply returns
the user's hashed password from the datastore. As written, this hardcodes the user's
password as "a".
## Update existing templates
The existing templates, other than using Bootstrap 5, are not formatted. Update to
a format in line with the new application. Note that `components/head.templ` should
be updated as well, and used in any pages (see the current templates as an example).
## Update favicon
The image used as the favicon is `/static/img/logo-48x48.png`. You can either update
this image, or you can update `/cmd/web/templates/components/head.templ` to point to a different
image file.

# Layout
* `cmd/cli`: This is where the CLI app is located. You should use Cobra to add
and new commands that are needed. This template only has the root command.
* `cmd/web`: The `main` package for the web application and the root for the
templates and the handlers.
    - `cmd/web/main.go`: This should be fairly complete as-is. You may want to
    put DB initialization here, though the `internal/db` package has a method to
    do this (see `internal/db/dbinit.go` below).
     - `cmd/web/routes.go`: Add your routes here, including the calls to any
     middleware. Note that for protected routes, you'll want to add `auth.AuthMiddleware` (see the "/home" route as an example).
     - `cmd/web/templates`: This is where your page templates go.
     - `cmd/web/templates/components`: Add or update any components, i.e., reusable
     parts of pages, here.
     - `cmd/web/handlers`: These are the handler functions for your various web
     routes.
* `internal/auth`: This has the authentication/authorization logic. You shouldn't
generally need to touch any of this.
* `internal/db`: This is the interface logic for your datastore. 
    - `internal/db/dbinit.go`: This handles the initialization of your datastore.
    - `internal/db/uservalidation.go`: This is a stub function to retrieve the 
    hashed value of your user's password. Right now, it just hashes "a" and returns
    this value, but you should replace this with the logic to get the hashed value
    from the datastore.