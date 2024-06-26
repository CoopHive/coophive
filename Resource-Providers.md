# Resource Provider Instructions
This page contains additional instructions for users that wish to run their own resource providers.  The process is fairly different based on machine, distro and personal preference.  However, one may find it usefull to see details of an example 2204 Ubuntu installation to help inform their own process

## Docker
https://docs.docker.com/engine/install/ubuntu/

https://docs.docker.com/engine/install/linux-postinstall/


## Nvidia Container Toolkit
https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html


```
curl -fsSL https://nvidia.github.io/libnvidia-container/gpgkey | sudo gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg \
  && curl -s -L https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list | \
    sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' | \
    sudo tee /etc/apt/sources.list.d/nvidia-container-toolkit.list
```

```
sudo apt-get update
```

```
sudo apt-get install -y nvidia-container-toolkit
```

```
sudo nvidia-ctk runtime configure --runtime=docker
```

```
sudo systemctl restart docker
```


## CUDA

https://docs.nvidia.com/cuda/cuda-installation-guide-linux/


```
gcc --version
```

```
sudo apt-get install linux-headers-$(uname -r)
```

```
sudo apt-key del 7fa2af80
```

```
# For example
sudo dpkg -i cuda-repo-<distro>_<version>_<architecture>.deb
```
as
```
sudo dpkg -i cuda-repo-ubuntu_2204__x86_64.deb
```

```
sudo cp /var/cuda-repo-ubuntu2204-X-Y-local/cuda-*-keyring.gpg /usr/share/keyrings/
```

```
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu2204/x86_64/cuda-ubuntu.pin
 sudo mv cuda-ubuntu.pin /etc/apt/preferences.d/cuda-repository-pin-600
```

```
sudo apt-get update

```

```
sudo apt-get install cuda-toolkit
```

```
sudo reboot
```


https://docs.nvidia.com/cuda/cuda-installation-guide-linux/index.html#post-installation-actions

export PATH=/usr/local/cuda-12.4/bin${PATH:+:${PATH}}


https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/sample-workload.html


```
sudo docker run --rm --runtime=nvidia --gpus all ubuntu nvidia-smi
```

## Coophive 

```
git clone https://github.com/CoopHive/coophive.git
```
`touch resource-provider-gpu.env`
add WEB3_PRIVATE_KEY=YOUR_PRIVATE_KEY

either `go build` and `sudo mv coophive /usr/bin/` or install from release and copy to /usr/bin

```
cd coophive/etc/systemd/
```


edit the configuration of `resource-provider-gpu.service`
```
[Unit]
Description=CoopHive Resource Provider GPU
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

[Service]
Environment="LOG_TYPE=json"
Environment="LOG_LEVEL=debug"
Environment="HOME=/app/coophive"
Environment="OFFER_GPU=1" # Your GPU amounts
Environment="OFFER_RAM=8192" # Your RAM spec (MiB)
Environment="OFFER_CPU=1000" # Your milli CPU points (1000 = 1 CPU)
EnvironmentFile=/app/coophive/resource-provider-gpu.env # path to where you cloned coophive and added the resource-provider-gpu.env
Restart=always
RestartSec=5s
ExecStart=/usr/bin/coophive resource-provider 

[Install]
WantedBy=multi-user.target

```

```sudo cp resource-provider-gpu.service /etc/systemd/system/```

```
sudo systemctl enable resource-provider-gpu.service
```

```
sudo systemctl start resource-provider-gpu.service
```


