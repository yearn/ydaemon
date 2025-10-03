# CPU Monitor Service

## Installation

1. Copy the service file to systemd:
   ```bash
   sudo cp scripts/cpu-monitor.service /etc/systemd/system/
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
