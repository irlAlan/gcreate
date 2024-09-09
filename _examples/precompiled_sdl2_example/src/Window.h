#pragma once

#include <string_view>
#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>
#include <SDL2/SDL_render.h>
#include <SDL2/SDL_surface.h>
#include <SDL2/SDL_video.h>

class Window {
 public:
  Window(std::string_view title, int width, int height);
  ~Window();
  void run();
 private:
  void draw();
  void update(float dt);

 private:
  SDL_Window *window{nullptr};
  SDL_Renderer *renderer{nullptr};
  SDL_Event event;
  int w,h;
};
