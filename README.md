# go-dot

golang dot command(graphviz).

## How to Install

Make sure you have graphviz(libgvc, libcgraph, libcdt) & libpng installed.

```
$ go get -u github.com/kaishuu0123/go-dot
```

## How to use

* create sample.dot file

```
// The graph name and the semicolons are optional
graph graphname {
    a -- b -- c;
    b -- d;
}
```

* convert to png

```
$ go-dot -Tpng sample.dot -o sample.png
```

* output png preview
![sample.png](https://github.com/kaishuu0123/go-dot/blob/master/examples/sample.png)

## DEMO

see examples directory.

## WHY IMPLEMENT THIS

* golang can cross compile.
* want use cross platform dot command. (and want single binary)

## LICENSE

* Eclipse Public License - v 1.0 (see LICENSE file)

## TODO

* support travis CI
* Release cross binary

## CONTACTS

* If you have a bug or believe something is not working as expected, please submit a [github issue](https://github.com/kaishuu0123/go-dot).
