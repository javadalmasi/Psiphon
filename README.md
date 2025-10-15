# Psiphone

[English](#psiphone) | [Persian](#پسیفون)

## English

Psiphone is a lightweight containerized SOCKS5 proxy that connects to the Psiphon network. It allows you to route your traffic through different countries using the Psiphon censorship circumvention technology.

### Features

- **Lightweight**: Container-optimized with minimal resource usage
- **SOCKS5 Proxy**: Provides standard SOCKS5 proxy functionality
- **Country Selection**: Choose from various countries for your connection
- **Container-Ready**: Designed specifically for containerized deployments
- **Easy Configuration**: Simple command-line options for configuration
- **Shuffle Mode**: Run all countries simultaneously with load balancing

### Usage

#### Command Line Options

```
NAME
  Psiphone

FLAGS
  -v, --verbose            enable verbose logging
  -b, --bind STRING        SOCKS bind address (default: 127.0.0.1:1080)
  -c, --country STRING     psiphon country code (default: DE)
  --shuffle                enable shuffle mode (run all countries with load balancing)
  -c, --config STRING      path to config file
      --help               show help
```

#### Examples

Connect with default settings (Germany, port 1080):
```bash
docker run -p 1080:1080 --rm psiphone
```

Connect to a specific country:
```bash
docker run -p 1080:1080 --rm psiphone --country US
```

Change the bind port:
```bash
docker run -p 8080:8080 --rm psiphone --bind 0.0.0.0:8080 --country CA
```

#### Shuffle Mode
Run in shuffle mode to automatically connect to all countries with load balancing:
```bash
docker run -p 1080:1080 --rm psiphone --shuffle
```

#### Configuration File

You can also use a configuration file in JSON format:

```json
{
  "verbose": false,
  "bind": "127.0.0.1:1080",
  "country": "DE"
}
```

### Supported Countries

| Flag | Country Code | Country Name |
|------|--------------|--------------|
| 🇦🇹 | AT | Austria |
| 🇦🇺 | AU | Australia |
| 🇧🇪 | BE | Belgium |
| 🇧🇬 | BG | Bulgaria |
| 🇨🇦 | CA | Canada |
| 🇨🇭 | CH | Switzerland |
| 🇨🇿 | CZ | Czech Republic |
| 🇩🇪 | DE | Germany |
| 🇩🇰 | DK | Denmark |
| 🇪🇪 | EE | Estonia |
| 🇪🇸 | ES | Spain |
| 🇫🇮 | FI | Finland |
| 🇫🇷 | FR | France |
| 🇬🇧 | GB | United Kingdom |
| 🇭🇷 | HR | Croatia |
| 🇭🇺 | HU | Hungary |
| 🇮🇪 | IE | Ireland |
| 🇮🇳 | IN | India |
| 🇮🇹 | IT | Italy |
| 🇯🇵 | JP | Japan |
| 🇱🇻 | LV | Latvia |
| 🇳🇱 | NL | Netherlands |
| 🇳🇴 | NO | Norway |
| 🇵🇱 | PL | Poland |
| 🇵🇹 | PT | Portugal |
| 🇷🇴 | RO | Romania |
| 🇷🇸 | RS | Serbia |
| 🇸🇪 | SE | Sweden |
| 🇸🇬 | SG | Singapore |
| 🇸🇰 | SK | Slovakia |
| 🇺🇸 | US | United States |

### Docker Usage

#### Pull from Container Registry
```bash
docker pull ghcr.io/javadalmasi/psiphone:latest
```

#### Build the Container (if building locally)
```bash
docker build -t psiphone .
```

### Container Distribution

This application is distributed exclusively as a container. The container image is available on GitHub Container Registry (GHCR).

### Quick Start with Docker

The easiest way to run Psiphone is using Docker with the pre-built image from GHCR:

```bash
# Pull and run the latest version
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest

# Run with default settings (Germany, port 1080)
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest

# Run with a specific country
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest --country US

# Run in shuffle mode (all countries with load balancing)
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest --shuffle
```

### Building from Source

Building from source is not recommended as this application is exclusively distributed as a container. If you need to build from source, clone the repository from https://github.com/javadalmasi/Psiphon and use the Dockerfile to build the container image:

```bash
docker build -t psiphone .
```

## Persian

پسیفون یک پروکسی SOCKS5 سبک‌وز است که در کانتینر اجرا می‌شود و به شبکه سایفون متصل می‌شود. این امکان را فراهم می‌کند که ترافیک شما را از طریق کشورهای مختلف با استفاده از فناوری دور زدن سانسور سایفون هدایت کند.

### ویژگی‌ها

- **سبک‌وز**: بهینه شده برای کانتینر با حداقل استفاده از منابع
- **پروکسی SOCKS5**: ارائه عملکرد استاندارد پروکسی SOCKS5
- **انتخاب کشور**: انتخاب از کشورهای مختلف برای اتصال شما
- **آماده کانتینر**: به طور خاص برای استقرارهای کانتینری طراحی شده است
- **پیکربندی آسان**: گزینه‌های خط فرمان ساده برای پیکربندی
- **حالت شافل**: اجرای همزمان همه کشورها با بالانس لود

### نحوه استفاده

#### گزینه‌های خط فرمان

```
نام
  Psiphone

پرچم‌ها
  -v, --verbose            فعال‌سازی ثبت گزارش تفصیلی
  -b, --bind STRING        آدرس اتصال SOCKS (پیش‌فرض: 127.0.0.1:1080)
  -c, --country STRING     کد کشور سایفون (پیش‌فرض: DE)
  --shuffle                فعال‌سازی حالت شافل (اتصال به همه کشورها با بالانس لود)
  -c, --config STRING      مسیر فایل پیکربندی
      --help               نمایش راهنما
```

#### مثال‌ها

اتصال با تنظیمات پیش‌فرض (آلمان، پورت 1080):
```bash
docker run -p 1080:1080 --rm psiphone
```

اتصال به یک کشور خاص:
```bash
docker run -p 1080:1080 --rm psiphone --country US
```

تغییر پورت اتصال:
```bash
docker run -p 8080:8080 --rm psiphone --bind 0.0.0.0:8080 --country CA
```

#### حالت شافل
اجرا در حالت شافل برای اتصال خودکار به همه کشورها با بالانس لود:
```bash
docker run -p 1080:1080 --rm psiphone --shuffle
```

#### فایل پیکربندی

همچنین می‌توانید از یک فایل پیکربندی در قالب JSON استفاده کنید:

```json
{
  "verbose": false,
  "bind": "127.0.0.1:1080",
  "country": "DE"
}
```

### کشورهای پشتیبانی شده

| پرچم | کد کشور | نام کشور |
|------|--------|---------|
| 🇦🇹 | AT | اتریش |
| 🇦🇺 | AU | استرالیا |
| 🇧🇪 | BE | بلژیک |
| 🇧🇬 | BG | بلغارستان |
| 🇨🇦 | CA | کانادا |
| 🇨🇭 | CH | سوئیس |
| 🇨🇿 | CZ | جمهوری چک |
| 🇩🇪 | DE | آلمان |
| 🇩🇰 | DK | دانمارک |
| 🇪🇪 | EE | استونی |
| 🇪🇸 | ES | اسپانیا |
| 🇫🇮 | FI | فنلاند |
| 🇫🇷 | FR | فرانسه |
| 🇬🇧 | GB | انگلستان |
| 🇭🇷 | HR | کرواسی |
| 🇭🇺 | HU | مجارستان |
| 🇮🇪 | IE | ایرلند |
| 🇮🇳 | IN | هند |
| 🇮🇹 | IT | ایتالیا |
| 🇯🇵 | JP | ژاپن |
| 🇱🇻 | LV | لتونی |
| 🇳🇱 | NL | هلند |
| 🇳🇴 | NO | نروژ |
| 🇵🇱 | PL | لهستان |
| 🇵🇹 | PT | پرتغال |
| 🇷🇴 | RO | رومانی |
| 🇷🇸 | RS | صربستان |
| 🇸🇪 | SE | سوئد |
| 🇸🇬 | SG | سنگاپور |
| 🇸🇰 | SK | اسلواکی |
| 🇺🇸 | US | ایالات متحده آمریکا |

### استفاده از داکر

#### دریافت از ثبت کانتینر
```bash
docker pull ghcr.io/javadalmasi/psiphone:latest
```

#### ساخت کانتینر (در صورت ساخت محلی)
```bash
docker build -t psiphone .
```

### توزیع کانتینر

این برنامه فقط به صورت کانتینر توزیع می‌شود. تصویر کانتینر در GitHub Container Registry (GHCR) موجود است.

### شروع سریع با داکر

آسان‌ترین راه برای اجرای پسیفون استفاده از داکر با تصویر از GHCR است:

```bash
# دریافت و اجرای آخرین نسخه
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest

# اجرا با تنظیمات پیش‌فرض (آلمان، پورت 1080)
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest

# اجرا با یک کشور خاص
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest --country US

# اجرا در حالت شافل (همه کشورها با بالانس لود)
docker run -p 1080:1080 --rm ghcr.io/javadalmasi/psiphone:latest --shuffle
```

### ساخت از منبع

ساخت از منبع توصیه نمی‌شود چون این برنامه فقط به صورت کانتینر توزیع می‌شود. اگر نیاز به ساخت از منبع دارید، مخزن را از https://github.com/javadalmasi/Psiphon کلون کنید و از Dockerfile برای ساخت تصویر کانتینر استفاده کنید:

```bash
docker build -t psiphone .
```