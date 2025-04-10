# Obsidian to Anki
I know there are plugins to create flashcards in obsidian and transfer them to Anki. What i want is to Learn my Zettelkasten items. The front should be the title and the back side should contain the file content.

# Requirements
- [Local REST API plugin](https://github.com/coddingtonbear/obsidian-local-rest-api) for Obsidian
- [Anki-connect](https://git.sr.ht/~foosoft/anki-connect)

# Usage
The primary commands are
```bash
obsidian-cli cat            # prints out the currently active file in obsidian
obsidian-cli create <name>  # creates a new deck with the provided name in anki
obsidian-cli add            # adds the currently active file as a note in anki
```