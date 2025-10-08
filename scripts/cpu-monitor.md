# CPU Monitor Service

The CPU monitor is a systemd service that continuously monitors the yDaemon process CPU usage and automatically restarts the service when sustained high CPU is detected. It uses a sliding window strategy (60 second window with 3 second intervals) and triggers a restart when 70% or more of the readings exceed the 85% CPU threshold. This prevents temporary CPU spikes from causing unnecessary restarts while ensuring recovery from sustained performance issues.

## Installation

1. Copy the service file to systemd:
   ```bash
   sudo cp cpu-monitor.service /etc/systemd/system/
   ```

2. Reload systemd:
   ```bash
   sudo systemctl daemon-reload
   ```

3. Enable the service to start on boot:
   ```bash
   sudo systemctl enable cpu-monitor.service
   ```

4. Start the service:
   ```bash
   sudo systemctl start cpu-monitor.service
   ```

5. Check the service status:
   ```bash
   sudo systemctl status cpu-monitor.service
   ```

## Viewing Logs

```bash
sudo journalctl -u cpu-monitor.service -f
```

## Stopping the Service

```bash
sudo systemctl stop cpu-monitor.service
```

## Disabling the Service

```bash
sudo systemctl disable cpu-monitor.service
```
