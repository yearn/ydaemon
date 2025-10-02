#!/bin/bash

# CPU Monitor Service - Monitors CPU and restarts ydaemon.service if needed
# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Configuration
CPU_THRESHOLD=85
HIGH_CPU_DURATION=10
CHECK_INTERVAL=3

# Counter for high CPU duration
high_cpu_count=0

log_message() {
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] $1"
}

show_monitoring() {
    echo -e "${CYAN}"
    cat << "EOF"
    ╔═══════════════════════════════════════╗
    ║   🖥️  CPU MONITOR ACTIVE 🖥️          ║
    ║   Watching for CPU spikes...         ║
    ╚═══════════════════════════════════════╝
EOF
    echo -e "${NC}"
}

show_high_cpu_warning() {
    local cpu=$1
    local count=$2
    echo -e "${YELLOW}"
    cat << "EOF"
    ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️
         🔥 HOT HOT HOT! 🔥
    ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️
EOF
    echo -e "${NC}"
    log_message "CPU at ${cpu}% for ${count} seconds (threshold: ${HIGH_CPU_DURATION}s)"
}

show_restarting() {
    echo -e "${RED}"
    cat << "EOF"
    ╔═════════════════════════════════════════╗
    ║                                         ║
    ║    🚨 CRITICAL! RESTARTING SERVICE 🚨   ║
    ║                                         ║
    ║         ⟳  ⟳  ⟳  ⟳  ⟳  ⟳              ║
    ║                                         ║
    ║     Stopping ydaemon.service...        ║
    ║                                         ║
    ╚═════════════════════════════════════════╝
EOF
    echo -e "${NC}"
}

show_success() {
    echo -e "${GREEN}"
    cat << "EOF"
    ╔═════════════════════════════════════════╗
    ║                                         ║
    ║    ✨ SERVICE RESTARTED SUCCESSFULLY ✨  ║
    ║                                         ║
    ║          (•_•)                          ║
    ║          <)  )╯  Back to normal!        ║
    ║          /  \                           ║
    ║                                         ║
    ╚═════════════════════════════════════════╝
EOF
    echo -e "${NC}"
}

show_normal() {
    echo -e "${GREEN}"
    cat << "EOF"
    ✓ All systems normal - CPU cooling down
EOF
    echo -e "${NC}"
}

get_cpu_usage() {
    top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1}'
}

restart_service() {
    show_restarting

    systemctl stop ydaemon.service
    log_message "ydaemon.service stopped"

    sleep 5

    systemctl start ydaemon.service
    log_message "ydaemon.service started"

    show_success

    # Reset counter
    high_cpu_count=0
}

# Main loop
log_message "CPU Monitor starting..."
show_monitoring

while true; do
    cpu_usage=$(get_cpu_usage)
    cpu_int=${cpu_usage%.*}

    if [ "$cpu_int" -gt "$CPU_THRESHOLD" ]; then
        high_cpu_count=$((high_cpu_count + 1))
        show_high_cpu_warning "$cpu_usage" "$high_cpu_count"

        if [ "$high_cpu_count" -ge "$HIGH_CPU_DURATION" ]; then
            restart_service
        fi
    else
        if [ "$high_cpu_count" -gt 0 ]; then
            show_normal
            log_message "CPU back to normal: ${cpu_usage}%"
            high_cpu_count=0
        fi
    fi

    sleep "$CHECK_INTERVAL"
done