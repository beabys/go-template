<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello_world/v1/hello_world.proto

namespace HelloWorld\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>hello_world.v1.HelloWorldResponse</code>
 */
class HelloWorldResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string hello = 1 [json_name = "hello"];</code>
     */
    protected $hello = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $hello
     * }
     */
    public function __construct($data = NULL) {
        \HelloWorld\V1\GPBMetadata\HelloWorld::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string hello = 1 [json_name = "hello"];</code>
     * @return string
     */
    public function getHello()
    {
        return $this->hello;
    }

    /**
     * Generated from protobuf field <code>string hello = 1 [json_name = "hello"];</code>
     * @param string $var
     * @return $this
     */
    public function setHello($var)
    {
        GPBUtil::checkString($var, True);
        $this->hello = $var;

        return $this;
    }

}

