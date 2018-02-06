#include<stdio.h>
#include<signal.h>
#include<unistd.h>

typedef void (*sighandler_t)(int);

void sigint_handler(int sig);

int
main(int argc, char *argv[])
{
  struct sigaction act, old;
  printf("start waiting\n");
  act.sa_handler = sigint_handler;
  sigemptyset(&act.sa_mask);
  act.sa_flags = SA_RESTART;
  if (sigaction(SIGINT, &act, &old) < 0)
    fprintf(stderr, "cannot register a handler for SIGINT");

  pause();
}

void
sigint_handler(int sig)
{
  printf("received signal SIGINT\nexit...\n");
}
