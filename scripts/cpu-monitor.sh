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

# Load Telegram configuration if available
SCRIPT_DIR="$(dirname "$0")"
if [ -f "$SCRIPT_DIR/cpu-monitor.env" ]; then
    export $(grep -v '^#' "$SCRIPT_DIR/cpu-monitor.env" | sed 's/ *= */=/g' | xargs)
fi

# Configuration
CPU_THRESHOLD=85
CHECK_INTERVAL=3
MAX_READINGS=20  # Number of readings to track (20 * 3s = 60 second window)
TRIGGER_RATIO=70  # Restart if this percentage of readings are high (70%)

# Sliding window for CPU readings
readings=()

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
    local high_count=$2
    local total_count=$3
    local percentage=$4
    echo -e "${YELLOW}"
    cat << "EOF"
    ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️
         🔥 HOT HOT HOT! 🔥
    ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️  ⚠️
EOF
    echo -e "${NC}"
    log_message "CPU at ${cpu}% - ${high_count}/${total_count} readings high (${percentage}%, trigger: ${TRIGGER_RATIO}%)"
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

send_telegram_alert() {
    local cpu=$1
    local high_count=$2
    local total_count=$3
    local percentage=$4
    local test_mode=$5

    # Check if Telegram is configured
    if [ -z "$TG_TOKEN" ] || [ -z "$TG_CHAT_ID" ]; then
        return 0
    fi

    local alert_text="⚠️ High CPU usage detected - restarting service"
    local action_text="<i>Service ydaemon.service is being restarted...</i>"

    if [ "$test_mode" == "true" ]; then
        alert_text="🧪 <b>THIS IS A TEST</b>"
        action_text="<i>No action taken - test mode only</i>"
    fi

    local message="🚨 <b>yDaemon CPU Alert</b>

${alert_text}

<b>Server:</b> $(hostname)
<b>CPU Usage:</b> ${cpu}%
<b>Threshold:</b> ${CPU_THRESHOLD}%
<b>High Readings:</b> ${high_count}/${total_count} (${percentage}%)
<b>Time:</b> $(date '+%Y-%m-%d %H:%M:%S')

${action_text}"

    local api_url="https://api.telegram.org/bot${TG_TOKEN}/sendMessage"

    # Build JSON payload
    local json_payload
    if [ -n "$TG_ALERT_MESSAGE_THREAD_ID" ]; then
        json_payload=$(cat <<EOF
{
    "chat_id": "${TG_CHAT_ID}",
    "message_thread_id": ${TG_ALERT_MESSAGE_THREAD_ID},
    "text": "${message}",
    "parse_mode": "HTML"
}
EOF
)
    else
        json_payload=$(cat <<EOF
{
    "chat_id": "${TG_CHAT_ID}",
    "text": "${message}",
    "parse_mode": "HTML"
}
EOF
)
    fi

    # Send message
    curl -s -X POST "$api_url" \
        -H "Content-Type: application/json" \
        -d "$json_payload" > /dev/null 2>&1
}

restart_service() {
    local cpu=$1
    local high_count=$2
    local total_count=$3
    local percentage=$4

    show_restarting

    # Send Telegram alert
    send_telegram_alert "$cpu" "$high_count" "$total_count" "$percentage"

    systemctl stop ydaemon.service
    log_message "ydaemon.service stopped"

    sleep 8

    systemctl start ydaemon.service
    log_message "ydaemon.service started"

    show_success

    # Clear readings array
    readings=()
}

# Handle command line arguments
if [ "$1" == "--test-telegram-alert" ]; then
    echo -e "${CYAN}Testing Telegram alert...${NC}"

    if [ -z "$TG_TOKEN" ] || [ -z "$TG_CHAT_ID" ]; then
        echo -e "${RED}Error: Telegram not configured in cpu-monitor.env${NC}"
        exit 1
    fi

    echo "Chat ID: $TG_CHAT_ID"
    echo "Thread ID: ${TG_ALERT_MESSAGE_THREAD_ID:-none}"
    echo ""

    # Send test alert with sample data
    send_telegram_alert "42.00" "69" "69" "69" "true"

    echo -e "${GREEN}Test alert sent!${NC}"
    exit 0
fi

# Main loop
log_message "CPU Monitor starting..."
show_monitoring

while true; do
    cpu_usage=$(get_cpu_usage)
    cpu_int=${cpu_usage%.*}

    # Add current reading to sliding window
    readings+=("$cpu_int")

    # Keep only last MAX_READINGS readings
    if [ ${#readings[@]} -gt $MAX_READINGS ]; then
        readings=("${readings[@]:1}")
    fi

    # Count how many readings exceed threshold
    high_count=0
    for reading in "${readings[@]}"; do
        if [ "$reading" -gt "$CPU_THRESHOLD" ]; then
            high_count=$((high_count + 1))
        fi
    done

    # Calculate percentage of high readings
    total_readings=${#readings[@]}
    if [ $total_readings -gt 0 ]; then
        high_percentage=$((high_count * 100 / total_readings))
    else
        high_percentage=0
    fi

    # Check if we should restart
    if [ $total_readings -eq $MAX_READINGS ] && [ $high_percentage -ge $TRIGGER_RATIO ]; then
        show_high_cpu_warning "$cpu_usage" "$high_count" "$total_readings" "$high_percentage"
        restart_service "$cpu_usage" "$high_count" "$total_readings" "$high_percentage"
    elif [ "$cpu_int" -gt "$CPU_THRESHOLD" ]; then
        show_high_cpu_warning "$cpu_usage" "$high_count" "$total_readings" "$high_percentage"
    else
        if [ $high_count -gt 0 ]; then
            log_message "CPU at ${cpu_usage}% - ${high_count}/${total_readings} readings high (${high_percentage}%)"
        fi
    fi

    sleep "$CHECK_INTERVAL"
done
