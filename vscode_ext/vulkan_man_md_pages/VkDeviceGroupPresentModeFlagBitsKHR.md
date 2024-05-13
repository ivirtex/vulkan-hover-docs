# VkDeviceGroupPresentModeFlagBitsKHR(3) Manual Page

## Name

VkDeviceGroupPresentModeFlagBitsKHR - Bitmask specifying supported
device group present modes



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)::`modes`,
indicating which device group presentation modes are supported, are:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_surface
typedef enum VkDeviceGroupPresentModeFlagBitsKHR {
    VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR = 0x00000001,
    VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR = 0x00000002,
    VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR = 0x00000004,
    VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR = 0x00000008,
} VkDeviceGroupPresentModeFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR` specifies that any
  physical device with a presentation engine **can** present its own
  swapchain images.

- `VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR` specifies that any
  physical device with a presentation engine **can** present swapchain
  images from any physical device in its `presentMask`.

- `VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR` specifies that any physical
  device with a presentation engine **can** present the sum of swapchain
  images from any physical devices in its `presentMask`.

- `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR` specifies
  that multiple physical devices with a presentation engine **can** each
  present their own swapchain images.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentInfoKHR.html),
[VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupPresentModeFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
