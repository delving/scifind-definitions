<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Services\VHU;

/**
 */
class MuseumObjectServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \App\Services\VHU\GetMuseumObjectRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetMuseumObject(\App\Services\VHU\GetMuseumObjectRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.MuseumObjectService/GetMuseumObject',
        $argument,
        ['\App\Services\VHU\GetMuseumObjectResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \App\Services\VHU\ListMuseumObjectsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ListMuseumObjects(\App\Services\VHU\ListMuseumObjectsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.MuseumObjectService/ListMuseumObjects',
        $argument,
        ['\App\Services\VHU\ListMuseumObjectsResponse', 'decode'],
        $metadata, $options);
    }

}
