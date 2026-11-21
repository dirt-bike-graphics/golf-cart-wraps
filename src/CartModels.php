<?php

declare(strict_types=1);

namespace CustomGraphics\GolfCartWraps;

use RuntimeException;

/** Public golf cart reference catalog. */
final class CartModels
{
    private static ?array $cache = null;

    /** Read the public reference dataset. */
    public static function raw(): array
    {
        if (self::$cache !== null) {
            return self::$cache;
        }
        foreach ([dirname(__DIR__) . '/models.json', dirname(__DIR__) . '/data/models.json'] as $path) {
            if (is_file($path)) {
                $json = file_get_contents($path);
                if ($json === false) {
                    throw new RuntimeException('Unable to read models.json');
                }
                self::$cache = json_decode($json, true, 512, JSON_THROW_ON_ERROR);
                return self::$cache;
            }
        }
        throw new RuntimeException('models.json dataset not found');
    }

    /** Return all catalog configurations. */
    public static function models(): array
    {
        return self::raw()['models'];
    }

    /** Return the wrap finish appearance references. */
    public static function finishes(): array
    {
        return self::raw()['finishes'];
    }

    /** Return the four body panel names. */
    public static function panels(): array
    {
        return self::raw()['panels'];
    }

    /** Find a model by identifier or model name, preferring exact matches. */
    public static function findModel(string $query): ?array
    {
        $q = self::normalize($query);
        if ($q === '') {
            return null;
        }
        foreach (self::models() as $model) {
            if (self::normalize($model['id']) === $q || self::normalize($model['model']) === $q) {
                return $model;
            }
        }
        foreach (self::models() as $model) {
            if (str_contains(self::normalize($model['id']), $q) || str_contains(self::normalize($model['model']), $q)) {
                return $model;
            }
        }
        return null;
    }

    /** Return the reference configurations for a brand. */
    public static function byBrand(string $brand): array
    {
        return array_values(array_filter(self::models(), static fn (array $model): bool => self::normalize($model['brand']) === self::normalize($brand)));
    }

    private static function normalize(string $value): string
    {
        return (string) preg_replace('/[^a-z0-9]/', '', strtolower($value));
    }
}
