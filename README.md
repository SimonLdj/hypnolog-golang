HypnoLog Go (Golang) Library
============================

## What is HypnoLog?
*Get Hypnotized While Logging*

*HypnoLog* allows you to fast and easily visualize your application data/objects while debugging. From any environment, in any language. Forget about those black text-based console debug-printing back from the 70's. 

**See [HypnoLog main repo](https://github.com/SimonLdj/hypnolog-server).**


## About HypnoLog-golang Library
Logging using *HypnoLog* means sending you data as JSON HTTP request to HypnoLog server. This library wraps all of those into simple easy to use functions.

## Installation

TODO: describe how to get the go library

If you haven't used *HypnoLog* before, [setup HypnoLog server](https://github.com/SimonLdj/hypnolog-server#setup-hypnolog-server) on your machine:
```bash
npm install -g hypnolog-server
```
*Note:* you will need [Node.js](https://nodejs.org/en/) installed on your machine first.

## Usage
1. Start *HypnoLog* server:
    ```bash
    hypnolog-server
    ```
2. View output: open [`http://127.0.0.1:7000/client.html`](http://127.0.0.1:7000/client.html) in your browser.
3. Import HypnoLog into your file:
4. Log:
   ```golang
   // Log a string
   hypnolog.Create().Str("Hello HypnoLog from Go!").Log()

   // log list of numbers as a graph (plot)
   // TODO

   // log any struct
   hypnolog.Create().Set(map[string]string{"some-key":"some-value"}).Log()
   ```

For more examples, see tests.

Read how to view the log and more about *HypnoLog* in [HypnoLog main repo page](https://github.com/SimonLdj/hypnolog-server).

