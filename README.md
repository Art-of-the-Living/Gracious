# <img src="icon.webp" width="32"> Gracious

> The mind is not a computer, but it can be simulated on one.

## A Cognitive Simulation Framework

Gracious is a simulation framework for the cognitive systems which give rise to consciousness. While these systems are
inspired by their biological implementations in the human brain, this is not meant to be neuro-biological simulation.
Furthermore, while the behaviour of these systems may exhibit properties of a "learning machine" this is not a machine
learning framework. Gracious cognitive simulations utilize bipolar associations and optimized Hebbian learning to form
associative relationships between phenomenal appearances and internal states.

## Getting Started


###Modules

The largest unit of the Gracious Cognitive Simulation Framework (GCSF) is the module. Modules represent specific
functional purposes for a cluster of structured neuron groups. A module may serve memory, motor control, vision, or
audition purposes, as well as many purposes novel to a computer. The simplest module to implement then is the "Console"
module.

####The Console module

The package `gracious/memory/console` implements an all-in-one connection to a console window. Input to the system
traverses a simple language pre-processing and feedback system with auto-associative temporal memory. This module
contains everything necessary to establish a direct path into the system. When connected, the system experiences
traditional ASCII characters from a terminal as though they were sensory phenomena, much in the same way we perceive
lights and sounds.