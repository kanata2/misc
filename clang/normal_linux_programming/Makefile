BINDIR = ./bin
SRCDIR = ./src
OBJDIR = ./obj
SOURCES = $(wildcard $(SRCDIR)/*.c)
OBJECTS = $(addprefix $(OBJDIR)/, $(notdir $(SOURCES:.c=.o)))
EXES = $(notdir $(OBJECTS:.o=))
CC = gcc
CFLAGS = -Wall -g

.PHONY: all
all: clean $(OBJECTS) $(EXES)

%: $(OBJDIR)/%.o
	$(CC) -o $(BINDIR)/$@ $^

$(OBJDIR)/%.o: $(SRCDIR)/%.c
	$(CC) $(CFLAGS) -o $@ -c $<

.PHONY: clean
clean:
	$(RM) ./bin/* ./obj/*.o

.PHONY: test
test:
	@echo $(SOURCES)
