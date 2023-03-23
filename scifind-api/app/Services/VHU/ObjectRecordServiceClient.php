<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Services\VHU;

/**
 */
class ObjectRecordServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \App\Services\VHU\GetObjectRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetObject(\App\Services\VHU\GetObjectRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.ObjectRecordService/GetObject',
        $argument,
        ['\App\Services\VHU\GetObjectResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \App\Services\VHU\ListObjectsRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ListObjects(\App\Services\VHU\ListObjectsRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.ObjectRecordService/ListObjects',
        $argument,
        ['\App\Services\VHU\ListObjectsResponse', 'decode'],
        $metadata, $options);
    }

}
