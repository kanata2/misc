#include<stdio.h>
#include<stdlib.h>
#include<sys/types.h>
#include<sys/stat.h>

int main(int argc, char *argv[])
{
  int i;
  mode_t mode;

  if (argc < 2) {
    fprintf(stderr, "%s: wrong arguments\n", argv[0]);
    exit(1);
  }
  mode = strtol(argv[1], NULL, 0);
  for (i = 2; i < argc; i++) {
    if (chmod(argv[i], mode) < 0) {
      perror(argv[1]);
    }
  }
  exit(0);
}
