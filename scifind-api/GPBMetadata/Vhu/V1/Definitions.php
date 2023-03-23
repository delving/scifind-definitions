<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/definitions.proto

namespace GPBMetadata\Vhu\V1;

class Definitions
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
vhu/v1/definitions.protovhu.v1"

Collection"3
ContentPage$
staticcontent (	Rstaticcontent"R
DownloadLink
uri (	Ruri
rights (	Rrights
caption (	Rcaption"
Event"�
Header
title (	Rtitle 
description (	Rdescription
summary (	Rsummary%
images (2.vhu.v1.ImageRimages"K
Image
uri (	Ruri
rights (	Rrights
caption (	Rcaption"d
LongRead0
content (2.vhu.v1.WebsiteContentRcontent&
header (2.vhu.v1.HeaderRheader"v
Measurement
height (Rheight
width (Rwidth
unit (	Runit%
dimension_type (	RdimensionType"�
Metadata#
object_number (	RobjectNumber
object_date (	R
objectDate7
measurements (2.vhu.v1.MeasurementRmeasurements$
documentation (	Rdocumentation 
association (	Rassociation
object_type (	R
objectType
rights (	Rrights#
theme (2.vhu.v1.ThemeRtheme*
creators	 (2.vhu.v1.PersonRcreators"
Museum"�
ObjectRecord&
header (2.vhu.v1.HeaderRheader,
metadata (2.vhu.v1.MetadataRmetadata/
	relations (2.vhu.v1.RelationsR	relations0
download (2.vhu.v1.DownloadLinkRdownload"	
Partner"
Person"�
	Relations4
collections (2.vhu.v1.CollectionRcollections>
sub_collections (2.vhu.v1.SubCollectionRsubCollections%
themes (2.vhu.v1.ThemeRthemes%
events (2.vhu.v1.EventRevents(
museums (2.vhu.v1.MuseumRmuseums.
objects (2.vhu.v1.ObjectRecordRobjects(
persons (2.vhu.v1.PersonRpersons2
story_cards	 (2.vhu.v1.StoryCardR
storyCards/

long_reads
 (2.vhu.v1.LongReadR	longReads"e
	StoryCard0
content (2.vhu.v1.WebsiteContentRcontent&
header (2.vhu.v1.HeaderRheader"o
StoryCardCollection&
header (2.vhu.v1.HeaderRheader0

story_card (2.vhu.v1.StoryCardR	storyCard"
SubCollection"
Theme"�
Timeline&
header (2.vhu.v1.HeaderRheader%
themes (2.vhu.v1.ThemeRthemes(
persons (2.vhu.v1.PersonRpersons%
events (2.vhu.v1.EventRevents"Y
WebsiteContent$
staticcontent (	Rstaticcontent!
publish_date (	RpublishDateB Zdelving.vhu�App\\Services\\VHUbproto3'
        , true);

        static::$is_initialized = true;
    }
}

