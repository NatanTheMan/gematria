COMPILER = chicken-csc
INTERPRETER = chicken-csi
SCRIPT = main.scm
EXEC = gematria

all: $(EXEC)

$(EXEC): $(SCRIPT)
	$(COMPILER) $(SCRIPT) -o $(EXEC)

run:
	$(INTERPRETER) $(SCRIPT)

clean:
	rm -f $(EXEC)
