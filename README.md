# Psiphon Proxy

[English](#psiphon-proxy) | [Persian](#پروکسی-سایفون)

## English

Psiphon Proxy is a lightweight containerized SOCKS5 proxy that connects to the Psiphon network. It allows you to route your traffic through different countries using the Psiphon censorship circumvention technology.

### Features

- **Lightweight**: Container-optimized with minimal resource usage
- **SOCKS5 Proxy**: Provides standard SOCKS5 proxy functionality
- **Country Selection**: Choose from various countries for your connection
- **Container-Ready**: Designed specifically for containerized deployments
- **Easy Configuration**: Simple command-line options for configuration

### Usage

#### Command Line Options

```
NAME
  Psiphon-Proxy

FLAGS
  -v, --verbose            enable verbose logging
  -b, --bind STRING        SOCKS bind address (default: 127.0.0.1:1080)
  -c, --country STRING     psiphon country code (default: DE)
  -c, --config STRING      path to config file
      --help               show help
```

#### Examples

Connect with default settings (Germany, port 1080):
```bash
docker run -p 1080:1080 --rm psiphon-proxy
```

Connect to a specific country:
```bash
docker run -p 1080:1080 --rm psiphon-proxy --country US
```

Change the bind port:
```bash
docker run -p 8080:8080 --rm psiphon-proxy --bind 0.0.0.0:8080 --country CA
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
docker pull ghcr.io/your-username/psiphon-proxy:latest
```

#### Build the Container (if building locally)
```bash
docker build -t psiphon-proxy .
```

### Building from Source

1. Install Go 1.24 or later
2. Clone the repository
3. Run the following commands:

```bash
go mod tidy
go build -o psiphon-proxy
```

## Persian

پروکسی سایفون یک پروکسی SOCKS5 سبک‌وز است که در کانتینر اجرا می‌شود و به شبکه سایفون متصل می‌شود. این امکان را فراهم می‌کند که ترافیک شما را از طریق کشورهای مختلف با استفاده از فناوری دور زدن سانسور سایفون هدایت کند.

### ویژگی‌ها

- **سبک‌وز**: بهینه شده برای کانتینر با حداقل استفاده از منابع
- **پروکسی SOCKS5**: ارائه عملکرد استاندارد پروکسی SOCKS5
- **انتخاب کشور**: انتخاب از کشورهای مختلف برای اتصال شما
- **آماده کانتینر**: به طور خاص برای استقرارهای کانتینری طراحی شده است
- **پیکربندی آسان**: گزینه‌های خط فرمان ساده برای پیکربندی

### نحوه استفاده

#### گزینه‌های خط فرمان

```
نام
  Psiphon-Proxy

پرچم‌ها
  -v, --verbose            فعال‌سازی ثبت گزارش تفصیلی
  -b, --bind STRING        آدرس اتصال SOCKS (پیش‌فرض: 127.0.0.1:1080)
  -c, --country STRING     کد کشور سایفون (پیش‌فرض: DE)
  -c, --config STRING      مسیر فایل پیکربندی
      --help               نمایش راهنما
```

#### مثال‌ها

اتصال با تنظیمات پیش‌فرض (آلمان، پورت 1080):
```bash
docker run -p 1080:1080 --rm psiphon-proxy
```

اتصال به یک کشور خاص:
```bash
docker run -p 1080:1080 --rm psiphon-proxy --country US
```

تغییر پورت اتصال:
```bash
docker run -p 8080:8080 --rm psiphon-proxy --bind 0.0.0.0:8080 --country CA
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
docker pull ghcr.io/your-username/psiphon-proxy:latest
```

#### ساخت کانتینر (در صورت ساخت محلی)
```bash
docker build -t psiphon-proxy .
```

### ساخت از منبع

1. نصب گو نسخه 1.24 یا بالاتر
2. کلون کردن مخزن
3. اجرای دستورات زیر:

```bash
go mod tidy
go build -o psiphon-proxy
```