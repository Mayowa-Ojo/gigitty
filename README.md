## Gigitty :trollface:

> Gigitty is an in-direct reference to the twisted character in Family Guy. I'm the master of the dark humor arts...gigitty

![meme of Glen Quagmire](https://memegenerator.net/img/instances/55687758.jpg)

If you're here, it means you survived the terrible joke preceding this therefore, passing the social experiment that indicates you're ready for the post-apocalytpic universe. Goodluck young padawan

Okay, okay seriously

In a nutshell, this is a [noob's] attempt at implementing clean architecture in Go. The app itself is very simple [just a wannabe library management server] [that is not even fully implemented], however, it is ambitious as it aims to be a template for clean achitecture.

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