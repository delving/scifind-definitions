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

    /**
     * @param \App\Services\VHU\ListTimelinesRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function ListTimelines(\App\Services\VHU\ListTimelinesRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/vhu.v1.TimelineService/ListTimelines',
        $argument,
        ['\App\Services\VHU\ListTimelinesResponse', 'decode'],
        $metadata, $options);
    }

}
