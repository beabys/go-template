<?php
// GENERATED CODE -- DO NOT EDIT!

namespace HelloWorld\V1;

/**
 */
class HelloWorldServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \HelloWorld\V1\HelloWorldRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetHelloWorld(\HelloWorld\V1\HelloWorldRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/hello_world.v1.HelloWorldService/GetHelloWorld',
        $argument,
        ['\HelloWorld\V1\HelloWorldResponse', 'decode'],
        $metadata, $options);
    }

}
