<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\JsonResponse;
use App\Http\Controllers\Controller;
use App\Services\Scifind\V1\GetObjectRequest;
use App\Services\Scifind\V1\ObjectServiceClient;

class ObjectsController extends Controller
{
    public function listObjects(): JsonResponse
    {
        $client = new ObjectServiceClient();
        $request = new GetObjectRequest();
        $request->setId(1);
        dd($client->GetObject($request));
    }
}
