# Deploying the Service Domain image on a bare metal server

## About the Service Domain Image File

The Service Domain bare metal image is available as a raw disk image provided by Nutanix for installation on approved hardware.

Download the Service Domain VM image file from the Nutanix Support portal [Downloads](https://portal.nutanix.com/page/downloads?product=karbonplatformservices) page.

There are two available images for bare metal installations, depending on whether the system boots with traditional BIOS or EFI:

  * Service Domain Raw compressed file (for bare metal installation)
  * Service Domain **EFI** Raw compressed file (for bare metal installation)

## Prepare the bare metal server

Before imaging the bare metal server, prepare it by updating any required firmware and performing hardware RAID virtual disk configuration.

## Download and boot a live image

The bare metal server can be imaged by "live booting" any linux operating system that includes lshw, tar, gzip, dd, and destination disk driver support. This may be completed via BMC or USB.
 
Example Usage with Ubuntu

1. Create an Ubuntu bootable USB drive using [macOS](https://tutorials.ubuntu.com/tutorial/tutorial-create-a-usb-stick-on-macos#0) or [Windows](https://tutorials.ubuntu.com/tutorial/tutorial-create-a-usb-stick-on-windows#0).
1. Download the appropriate image (RAW or EFI RAW compressed) onto the same USB drive or another.
1. Boot from the Ubuntu bootable USB drive with “Try Ubuntu” option.

## Identify the destination disk and image it

1. Open Terminal and identify disks with (minimum destination disk size: 200GB):
   
   `sudo lshw -c disk`
1. Change to directory where the Service Domain image is located, usually something like:
   
   `cd /media/ubuntu/<drive label>`
1. Execute the following command (replace <destination disk> with correct disk, e.g. /dev/sda or /dev/nvme0n1):
   
   `sudo tar -xOzvf <baseimage.raw.tgz> | sudo dd of=<destination disk> bs=1M status=progress`

   Example Usage

   `sudo tar -xOzvf base-image_175.raw.tgz | sudo dd of=/dev/sda bs=1M status=progress`

## Reboot the server and onboard it

1. When imaging is complete, connect the primary network interface to a network with DHCP, remove the USB drive, and reboot the system. Note the IP address received from DHCP on boot.

1. Continue with [Adding a Single Node Service Domain](https://portal.nutanix.com/#/page/docs/details?targetId=Karbon-Platform-Services-Admin-Guide:ks-service-domain-add-t.html).
 


