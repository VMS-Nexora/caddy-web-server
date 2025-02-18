interface Version {
  major: number;
  minor: number;
  build: number;
  value: number;
  date: string;
  hash_sha256: string;
}

interface EndpointConfig {
  base: {
    url: string;
  };
}

interface S3StorageConfig {
  url: string;
}

interface PluginConfig {
  enabled: boolean;
  [key: string]: { type: string; value: unknown; comment?: string } | boolean;
}

interface SettingsConfig {
  enable_logging: { type: string; value: boolean; comment: string };
  max_retry_attempts: { type: string; value: number; comment: string };
  tenant_code?: { type: string; value: number; comment: string };
}

interface GlobalConfig {
  $schema: string;
  version: Version;
  endpoints: EndpointConfig;
  plugins: Record<string, PluginConfig>;
  settings: SettingsConfig;
  s3_storage?: S3StorageConfig;
}

interface TenantConfig {
  tenant_code: { type: string; value: number; comment: string };
}
