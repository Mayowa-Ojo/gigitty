## Gigitty :trollface:

> Gigitty is an in-direct reference to the horny character in Family Guy. I'm the master of the dark humor arts...gigitty

If you're here, it means you survived that terrible joke preceding this therefore, passing the social experiment that indicates you're ready for the post-apocalytpic universe. Goodluck young padawan

In a nutshell, this is a terrible attempt at implementing clean architecture. The app itself is very simple [just a wannabe library management server], however, it is ambitious as it aims to be the boilerplate for my go web projects going forward.

#### Basic Rules of Clean Architecture
- Independent of frameworks
- Testable
- Independent of UI
- Independent of Database
- Independent of any external agency

##### Layers
- Repository
- Service/Usecase
- Controller/Delivery
- Entity/Model/Domain

For example, using interfaces to define the concrete implementation of the repository layer, we can easily switch between different databases [postgres, mongodb, etc...] without changing the implementation. They just need to satisfy the interface.

#### Tools
- Go <1.13>
- Go Fiber [web framework]
- Mongo-driver