#include <iostream>
#include "Window.h"

template <typename... T>
void println(T &...args) {
  auto print{[](auto &arg) { std::cout << arg; }};
  (print(args), ...);
  print("\n");
}

int main(int argv, char **argc) {
  println("Hello World!");
  Window win{"Hey there", 100,100};
  win.run();
  return 0;
}
