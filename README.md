# Insult Generator (Wasm)

I was so tickled by the [InsultAPI](https://github.com/ncitron/insultAPI) developed
by [ncitron](https://github.com/ncitron) that I thought I'd give it a Go (pun intended)
and implement with [Web Assembly and WASM](https://vsoch.github.io/insult-go/). 
It's a relatively simple application because it randomly selects an insult from a list and returns it to the browser,
and if you've selected to randomize, it will generate by selecting two adjectives and
a noun randomly.

[![img/insult.png](img/insult.png)](https://vsoch.github.io/insult-go/)

Unlike it's previous art, the goal isn't to provide an API, but a simple interface
for generating insults (and copy pasting them wherever they are needed!)

Check out the interface [here](https://vsoch.github.io/insult-go/)!

## Usage

You can easily interact with the [online version](https://vsoch.github.io/insult-go).
If you want to develop, you'll need GoLang version 1.13 or higher, and to install
[emscripten](https://emscripten.org/docs/getting_started/FAQ.html).

## Docker

The provided [Dockerfile](Dockerfile) will also install emscripten and compile,
but there are some issues with the mime type not being understood (hopefully
this will be resolved with some time!)

```bash
$ docker build -t vanessa/insult-go .
```

It will also install [emscripten](https://emscripten.org/docs/getting_started/FAQ.html),
add the source code to the repository, and compile to wasm. You can then
run the container and expose port 80 to see the compiled interface:

```bash
$ docker run -it --rm -p 80:80 vanessa/insult-go
``` 

Again, when I tested this the first time it worked okay, and for some reason,
it didn't the second time. If anyone hits the bug and can contribute a solution,
please open a pull request! 
