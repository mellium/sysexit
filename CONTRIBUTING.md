First of all: Thank you for reading this all the way through! You're the best!

Now, on to business! If you're planning on contributing, I'd appreciate it if
you'd keep a few things in mind:

  * Formatting
    * Run `go format` against all contributions
    * Wrap code at 80 characters (including the line break, so 79 really)
  * Documentation
    * Include comments for *all* public APIs, structs, packages, etc.
    * If a package contains more than one file, place package documentation in
      a separate (otherwise-empty but for the package declaration and comments)
      `doc.go` file.
    * Start comments with the name of the thing they're documenting
    * Use Unicode characters over ASCII representations where appropriate…
  * Best practices
    * If you don't have a very good reason to make something mutable: Don't.
    * Use UTF-8 (and nothing but UTF-8)
    * Probably some other things I'm forgetting. Don't be offended if I'm picky
      during code review (and be picky yourself!), but don't worry about it too
      much either.
