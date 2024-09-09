#include "Window.h"
#include <iostream>

Window::Window(std::string_view title, int width, int height)
  : w(width), h(height)
{

  if (SDL_Init(SDL_INIT_VIDEO) < 0) {
    std::cout << "Cannot init video: " <<  SDL_GetError() << '\n';
    std::exit(-1);
  }
  window = SDL_CreateWindow(title.data(), SDL_WINDOWPOS_UNDEFINED,
                            SDL_WINDOWPOS_UNDEFINED,w, h,
                            SDL_WINDOW_SHOWN);
  if (!window) {
    std::cout << "Cannot create sdl window: " << SDL_GetError() << '\n';
    std::exit(-1);
  }

  renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);
  if (!renderer) {
    std::cout << "Cannot create renderer: " << SDL_GetError() << '\n';
    std::exit(-1);
  }
  SDL_RenderSetLogicalSize(renderer, width, height);
}

Window::~Window(){
  SDL_DestroyWindow(window);
  SDL_Quit();
}

void Window::run() {
  bool quit{false};
  while (!quit) {
    while (SDL_PollEvent(&event)) {
      switch (event.type) {
        case SDL_QUIT:
          quit = true;
          break;
        default:
          break;
      }
    }
    update(1.0f / 60.0f);
    draw();
  }
}

void Window::update(float dt) {
}

void Window::draw() {
  SDL_RenderClear(renderer);
  SDL_RenderPresent(renderer);
}
