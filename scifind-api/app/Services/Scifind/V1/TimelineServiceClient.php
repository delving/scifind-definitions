<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Scifind\V1;

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
     * @param \Scifind\V1\GetTimelineRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetTimeline(\Scifind\V1\GetTimelineRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/scifind.v1.TimelineService/GetTimeline',
        $argument,
        ['\Scifind\V1\GetTimelineResponse', 'decode'],
        $metadata, $options);
    }

}
