# DAWG Router

This project is just an implementation of an http router based on DAWGs (Directed Acyclic Word Graphs)

The goal is to create a more or less compact data structure that also works fast and reliable

Also, I will make use of channels and goroutines to check whether a route exists. As it's a word-based automaton, I can just feed him word-tokens while parsing them (yes, in parallel).
Also, I figured that allocations are pretty slow, so I will probably use a PoolBuffer to eliminate allocations, once the routes have been predefined. Maybe I should add a '.Lock()' function
so that it gets optimized and stuff?

Licensed under the WTF Public License
