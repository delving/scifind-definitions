<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: scifind/v1/definitions.proto

namespace App\Services\Scifind\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>scifind.v1.DownloadLink</code>
 */
class DownloadLink extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string uri = 1 [json_name = "uri"];</code>
     */
    protected $uri = '';
    /**
     * Generated from protobuf field <code>string rights = 2 [json_name = "rights"];</code>
     */
    protected $rights = '';
    /**
     * Generated from protobuf field <code>string caption = 3 [json_name = "caption"];</code>
     */
    protected $caption = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $uri
     *     @type string $rights
     *     @type string $caption
     * }
     */
    public function __construct($data = NULL) {
        \Scifind\V1\GPBMetadata\Definitions::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string uri = 1 [json_name = "uri"];</code>
     * @return string
     */
    public function getUri()
    {
        return $this->uri;
    }

    /**
     * Generated from protobuf field <code>string uri = 1 [json_name = "uri"];</code>
     * @param string $var
     * @return $this
     */
    public function setUri($var)
    {
        GPBUtil::checkString($var, True);
        $this->uri = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string rights = 2 [json_name = "rights"];</code>
     * @return string
     */
    public function getRights()
    {
        return $this->rights;
    }

    /**
     * Generated from protobuf field <code>string rights = 2 [json_name = "rights"];</code>
     * @param string $var
     * @return $this
     */
    public function setRights($var)
    {
        GPBUtil::checkString($var, True);
        $this->rights = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string caption = 3 [json_name = "caption"];</code>
     * @return string
     */
    public function getCaption()
    {
        return $this->caption;
    }

    /**
     * Generated from protobuf field <code>string caption = 3 [json_name = "caption"];</code>
     * @param string $var
     * @return $this
     */
    public function setCaption($var)
    {
        GPBUtil::checkString($var, True);
        $this->caption = $var;

        return $this;
    }

}

