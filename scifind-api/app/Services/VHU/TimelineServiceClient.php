<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Services\VHU;

/**
 */
class TimelineServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \App\Services\VHU\GetTimelineRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetTimeline(\App\Services\VHU\GetTimelineRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.TimelineService/GetTimeline',
        $argument,
        ['\App\Services\VHU\GetTimelineResponse', 'decode'],
        $metadata, $options);
    }

}
