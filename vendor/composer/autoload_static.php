<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInit8aa27bca8ceaac307a9b4b56c578d03c
{
    public static $files = array (
        '0e6d7bf4a5811bfa5cf40c5ccd6fae6a' => __DIR__ . '/..' . '/symfony/polyfill-mbstring/bootstrap.php',
    );

    public static $prefixLengthsPsr4 = array (
        'Z' => 
        array (
            'Zend\\Stdlib\\' => 12,
            'Zend\\Hydrator\\' => 14,
            'Zend\\EventManager\\' => 18,
            'Zend\\Code\\' => 10,
        ),
        'S' => 
        array (
            'Symfony\\Polyfill\\Mbstring\\' => 26,
            'Symfony\\Component\\Process\\' => 26,
            'Symfony\\Component\\Debug\\' => 24,
            'Symfony\\Component\\Console\\' => 26,
        ),
        'P' => 
        array (
            'Psr\\Log\\' => 8,
            'Protobuf\\Compiler\\' => 18,
            'Protobuf\\' => 9,
        ),
        'D' => 
        array (
            'Doctrine\\Inflector\\' => 19,
            'Doctrine\\Common\\Inflector\\' => 26,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Zend\\Stdlib\\' => 
        array (
            0 => __DIR__ . '/..' . '/zendframework/zend-stdlib/src',
        ),
        'Zend\\Hydrator\\' => 
        array (
            0 => __DIR__ . '/..' . '/zendframework/zend-hydrator/src',
        ),
        'Zend\\EventManager\\' => 
        array (
            0 => __DIR__ . '/..' . '/zendframework/zend-eventmanager/src',
        ),
        'Zend\\Code\\' => 
        array (
            0 => __DIR__ . '/..' . '/zendframework/zend-code/src',
        ),
        'Symfony\\Polyfill\\Mbstring\\' => 
        array (
            0 => __DIR__ . '/..' . '/symfony/polyfill-mbstring',
        ),
        'Symfony\\Component\\Process\\' => 
        array (
            0 => __DIR__ . '/..' . '/symfony/process',
        ),
        'Symfony\\Component\\Debug\\' => 
        array (
            0 => __DIR__ . '/..' . '/symfony/debug',
        ),
        'Symfony\\Component\\Console\\' => 
        array (
            0 => __DIR__ . '/..' . '/symfony/console',
        ),
        'Psr\\Log\\' => 
        array (
            0 => __DIR__ . '/..' . '/psr/log/Psr/Log',
        ),
        'Protobuf\\Compiler\\' => 
        array (
            0 => __DIR__ . '/..' . '/protobuf-php/protobuf-plugin/src',
        ),
        'Protobuf\\' => 
        array (
            0 => __DIR__ . '/..' . '/protobuf-php/protobuf/src',
        ),
        'Doctrine\\Inflector\\' => 
        array (
            0 => __DIR__ . '/..' . '/doctrine/inflector/lib/Doctrine/Inflector',
        ),
        'Doctrine\\Common\\Inflector\\' => 
        array (
            0 => __DIR__ . '/..' . '/doctrine/inflector/lib/Doctrine/Common/Inflector',
        ),
    );

    public static $prefixesPsr0 = array (
        'g' => 
        array (
            'google\\protobuf' => 
            array (
                0 => __DIR__ . '/..' . '/protobuf-php/google-protobuf-proto/src',
            ),
        ),
    );

    public static $classMap = array (
        'Composer\\InstalledVersions' => __DIR__ . '/..' . '/composer/InstalledVersions.php',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInit8aa27bca8ceaac307a9b4b56c578d03c::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInit8aa27bca8ceaac307a9b4b56c578d03c::$prefixDirsPsr4;
            $loader->prefixesPsr0 = ComposerStaticInit8aa27bca8ceaac307a9b4b56c578d03c::$prefixesPsr0;
            $loader->classMap = ComposerStaticInit8aa27bca8ceaac307a9b4b56c578d03c::$classMap;

        }, null, ClassLoader::class);
    }
}