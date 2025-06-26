# Project Setup
## Before I start
I've chosen to learn vulkan without any abstraction layer in golang. A combination that is ill-advised, but I'm
stupid. 


**It's a bad idea since:**
- Vulkan:
    - low-level
        - harder to make things happen, less dopamine go brr.
        - very complex, there are not decisions/concepts unturned. Not a lot things may be offloaded to the api.
            - no graphics knowledge
            - no advanced maths knowledge
- Golang:
    - Limited community
        - on my own
    - Not it's focus
    - Gateway: In future projects I will most likely not use golang + vulkan in the long-term.
    - Addition of binding
        - danger of learning the binding, not the vulkan. Must research if the binding is an actual binding
        not abstraction.


**It's a good idea for me, because:**
- Vulkan: 
    - Performant
        - Great endpoint, not a gateway, just something that will be instantly useful.
    - Consistently low-level
        - Allows me to attack graphic concepts isolated to ideally **a** function call...
            - Allows me to attack math concepts isolated to ideally **a** graphic concept
- Golang: 
    - Performant (enough)
    - Mostly known
        - Not too busy with language quirks, can focus on vulkan 
        - Room to learn deeply
    - `cgo`: Can use C directly inside a go program.
        - Examples are directly executable?


**Takeaways:**
- Not optimizing my path to quick-wins (creating a game) but to deep conceptual understanding
- Endpoint focused
- Limited help
- Dopamine-starvation in mind (yikes)


I think this **can** be realistic. Yet, without defining my path more I think abandonment is likely. So — **how to 
make it bearable?**


- Milepoints, relative concept isolated milepoints I can achieve sequentially. This will help with the dopamine
starvation.
- Mindset that short-term failure is intended. Long-term failure is not. Feel proud while falling.
- Taking time to document everything. More documentation than code probably. Feel free to sketch and write so I
can provide better feedback to myself. If I know one thing — my memory ain't good. Use writing as physical memory.
Just like computers it's slower than RAM. But its more persistent. 
