import subprocess
import time

# Define the command to find the PID of the running Go application
pid_command = "pgrep -f 'go run ./cmd/chat'"
extra_com = "ps aux | grep -e '/tmp/go.*/exe/chat' | awk '{print $2}'"

pid_process = subprocess.Popen(pid_command, shell=True, stdout=subprocess.PIPE)
pid_output, _ = pid_process.communicate()
pid = pid_output.decode().strip()

extra_process = subprocess.Popen(extra_com, shell=True, stdout=subprocess.PIPE)
extra_output, _ = extra_process.communicate()
e_pid = extra_output.decode().strip().split('\n')[0]


if pid:
    # Send a SIGKILL signal to the Go application process
    kill_command = f"kill -9 {pid} {e_pid}"
    subprocess.call(kill_command, shell=True)
    print(f"Sent SIGKILL signal to processes with PID: {pid} {e_pid}")
    
    time_lag = 5
    for i in range(time_lag):
        print(f"Recovering after {time_lag - i} seconds...")
        time.sleep(1)
else:
    print("Go application process not found.")
