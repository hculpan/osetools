# go-web-base-project-1
This is a project template for a Go-based web application.

To use this project as the basis for a new project, see the `USING` document
that lists the changes that will need to be done.

This uses the following libraries:

1. Chi
2. Templ (as in a-h/templ)
3. Godotenv
4. bcrypt
5. Corbra (for the CLI app)
6. Go-lang JWT

This project also uses a Makefile to build, and Air for live reloads.

This project template includes the functionality for a user to login and for them
to remain logged in for 30 days (JWT token stored as a cookie). If they are logged
in, they will go to a "Home" page, which is intended to be the root page for the 
application. If they are not logged in, they will be redirected to an "Info" page,
that is for public consumption.

If they go to the "/" URL, they will be redirected to "/home" (the Home page) if
then authorization token cookie is found. Otherwise, they'll be redirected to "/info", 
the public page.

Note that this project does not include any datastore.