import requests
import time
import matplotlib.pyplot as plt
import numpy as np

def load_test(url, duration_seconds=30):
    response_times = []
    start_time = time.time()
    end_time = start_time + duration_seconds
    
    print(f"Starting load test for {duration_seconds} seconds...")
    print(f"Target URL: {url}")
    print("-" * 50)
    
    request_count = 0
    while time.time() < end_time:
        try:
            start_request = time.time()
            response = requests.get(url, timeout=10)
            end_request = time.time()
            
            response_time = (end_request - start_request) * 1000
            response_times.append(response_time)
            request_count += 1
            
            if response.status_code == 200:
                print(f"Request {request_count}: {response_time:.2f}ms - SUCCESS")
            else:
                print(f"Request {request_count}: {response_time:.2f}ms - FAILED")
                
        except requests.exceptions.RequestException as e:
            print(f"Request failed: {e}")
            request_count += 1
            
        time.sleep(0.1)
            
    return response_times

# Replace with your EC2 public IP
EC2_URL = "http://34.210.146.51:8080/albums"

# Test connectivity first
try:
    print("Testing connectivity...")
    test_response = requests.get(EC2_URL, timeout=5)
    if test_response.status_code == 200:
        print("✅ Server is responding!")
    else:
        print(f"⚠️ Server status: {test_response.status_code}")
except Exception as e:
    print(f"❌ Cannot connect: {e}")
    exit(1)

# Run the test
response_times = load_test(EC2_URL)

# Plot results
plt.figure(figsize=(12, 8))

plt.subplot(2, 1, 1)
plt.hist(response_times, bins=30, alpha=0.7, color='blue')
plt.xlabel('Response Time (ms)')
plt.ylabel('Frequency')
plt.title('Distribution of Response Times')

plt.subplot(2, 1, 2)
plt.scatter(range(len(response_times)), response_times, alpha=0.6)
plt.xlabel('Request Number')
plt.ylabel('Response Time (ms)')
plt.title('Response Times Over Time')

plt.tight_layout()
plt.show()

# Print statistics
print(f"\nStatistics:")
print(f"Total requests: {len(response_times)}")
print(f"Average: {np.mean(response_times):.2f}ms")
print(f"Median: {np.median(response_times):.2f}ms")
print(f"95th percentile: {np.percentile(response_times, 95):.2f}ms")
print(f"99th percentile: {np.percentile(response_times, 99):.2f}ms")
print(f"Max: {max(response_times):.2f}ms")
