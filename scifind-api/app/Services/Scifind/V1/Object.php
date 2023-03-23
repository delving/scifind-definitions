<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: scifind/v1/definitions.proto

namespace Scifind\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>scifind.v1.Object</code>
 */
class Object extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.scifind.v1.Header header = 1 [json_name = "header"];</code>
     */
    protected $header = null;
    /**
     * Generated from protobuf field <code>.scifind.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     */
    protected $metadata = null;
    /**
     * Generated from protobuf field <code>.scifind.v1.Relations relations = 3 [json_name = "relations"];</code>
     */
    protected $relations = null;
    /**
     * Generated from protobuf field <code>.scifind.v1.DownloadLink download = 4 [json_name = "download"];</code>
     */
    protected $download = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Scifind\V1\Header $header
     *     @type \Scifind\V1\Metadata $metadata
     *     @type \Scifind\V1\Relations $relations
     *     @type \Scifind\V1\DownloadLink $download
     * }
     */
    public function __construct($data = NULL) {
        \Scifind\V1\GPBMetadata\Definitions::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Header header = 1 [json_name = "header"];</code>
     * @return \Scifind\V1\Header|null
     */
    public function getHeader()
    {
        return $this->header;
    }

    public function hasHeader()
    {
        return isset($this->header);
    }

    public function clearHeader()
    {
        unset($this->header);
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Header header = 1 [json_name = "header"];</code>
     * @param \Scifind\V1\Header $var
     * @return $this
     */
    public function setHeader($var)
    {
        GPBUtil::checkMessage($var, \Scifind\V1\Header::class);
        $this->header = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     * @return \Scifind\V1\Metadata|null
     */
    public function getMetadata()
    {
        return $this->metadata;
    }

    public function hasMetadata()
    {
        return isset($this->metadata);
    }

    public function clearMetadata()
    {
        unset($this->metadata);
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     * @param \Scifind\V1\Metadata $var
     * @return $this
     */
    public function setMetadata($var)
    {
        GPBUtil::checkMessage($var, \Scifind\V1\Metadata::class);
        $this->metadata = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Relations relations = 3 [json_name = "relations"];</code>
     * @return \Scifind\V1\Relations|null
     */
    public function getRelations()
    {
        return $this->relations;
    }

    public function hasRelations()
    {
        return isset($this->relations);
    }

    public function clearRelations()
    {
        unset($this->relations);
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.Relations relations = 3 [json_name = "relations"];</code>
     * @param \Scifind\V1\Relations $var
     * @return $this
     */
    public function setRelations($var)
    {
        GPBUtil::checkMessage($var, \Scifind\V1\Relations::class);
        $this->relations = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.DownloadLink download = 4 [json_name = "download"];</code>
     * @return \Scifind\V1\DownloadLink|null
     */
    public function getDownload()
    {
        return $this->download;
    }

    public function hasDownload()
    {
        return isset($this->download);
    }

    public function clearDownload()
    {
        unset($this->download);
    }

    /**
     * Generated from protobuf field <code>.scifind.v1.DownloadLink download = 4 [json_name = "download"];</code>
     * @param \Scifind\V1\DownloadLink $var
     * @return $this
     */
    public function setDownload($var)
    {
        GPBUtil::checkMessage($var, \Scifind\V1\DownloadLink::class);
        $this->download = $var;

        return $this;
    }

}

