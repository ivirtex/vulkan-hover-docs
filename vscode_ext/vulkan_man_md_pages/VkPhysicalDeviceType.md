# VkPhysicalDeviceType(3) Manual Page

## Name

VkPhysicalDeviceType - Supported physical device types



## <a href="#_c_specification" class="anchor"></a>C Specification

The physical device types which **may** be returned in
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`deviceType`
are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkPhysicalDeviceType {
    VK_PHYSICAL_DEVICE_TYPE_OTHER = 0,
    VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU = 1,
    VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU = 2,
    VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU = 3,
    VK_PHYSICAL_DEVICE_TYPE_CPU = 4,
} VkPhysicalDeviceType;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PHYSICAL_DEVICE_TYPE_OTHER` - the device does not match any other
  available types.

- `VK_PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU` - the device is typically one
  embedded in or tightly coupled with the host.

- `VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU` - the device is typically a
  separate processor connected to the host via an interlink.

- `VK_PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU` - the device is typically a
  virtual node in a virtualization environment.

- `VK_PHYSICAL_DEVICE_TYPE_CPU` - the device is typically running on the
  same processors as the host.

The physical device type is advertised for informational purposes only,
and does not directly affect the operation of the system. However, the
device type **may** correlate with other advertised properties or
capabilities of the system, such as how many memory heaps there are.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
