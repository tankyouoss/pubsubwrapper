# pubSubWrapper

Wrap google pubsub golang client implementation to an interface so it can be mocked. 

This wrapper is inspired by https://github.com/googleapis/google-cloud-go-testing/tree/master/pubsub/psiface adapted to feet 
usage in TankYou projects.

## Why

- Implement a mock using testify/mock
- Add methods to interface which where not initialy supported by googleapis project
