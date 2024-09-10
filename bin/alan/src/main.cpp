#include <iostream>

template <typename... T>
void println(T &...args) {
  auto print{[](auto &arg) { std::cout << arg; }};
  (print(args), ...);
  print("\n");
}

int main(int argv, char **argc) {
  println("Hello World!");
  println("alan!");
  return 0;
}
