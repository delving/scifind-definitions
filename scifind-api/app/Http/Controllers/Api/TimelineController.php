<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\JsonResponse;
use App\Http\Controllers\Controller;
use App\Services\Scifind\V1\Object as ScifindObject;
use App\Services\Scifind\V1\ObjectServiceClient;

class TimelineController extends Controller
{
    public function getTimeline(): JsonResponse
    {

    }
}
