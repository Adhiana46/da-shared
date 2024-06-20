package config

type CacheDriver string

var (
	CacheDriverRedis     CacheDriver = "redis"
	CacheDriverMemcached CacheDriver = "memcached"
)

type CacheConfig struct {
	Driver CacheDriver `env:"CACHE_DRIVER" yaml:"driver"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" yaml:"host"`
	Port     string `env:"REDIS_PORT" yaml:"port"`
	Password string `env:"REDIS_PASSWORD" yaml:"password"`
	DB       int    `env:"REDIS_DB" yaml:"db"`
}

type MemcachedConfig struct {
	Hosts []string `env-separator:"," env:"MEMCACHED_HOSTS" yaml:"hosts"`
}
