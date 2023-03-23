<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\JsonResponse;
use App\Http\Controllers\Controller;
use App\Services\Scifind\V1\GetTimelineRequest;
use App\Services\Scifind\V1\TimelineServiceClient;

class TimelineController extends Controller
{
    public function getTimeline(): JsonResponse
    {
        $client = new TimelineServiceClient();
        $request = new GetTimelineRequest();
        $request->setId(1);
        dd($client->GetObject($request));
    }
}
