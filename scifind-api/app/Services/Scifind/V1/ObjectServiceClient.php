<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Services\Scifind\V1;

/**
 */
class ObjectServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Scifind\V1\GetObjectRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetObject(\Scifind\V1\GetObjectRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/scifind.v1.ObjectService/GetObject',
        $argument,
        ['\Scifind\V1\GetObjectResponse', 'decode'],
        $metadata, $options);
    }

}
