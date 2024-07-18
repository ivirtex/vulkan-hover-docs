# VK_MAX_DEVICE_GROUP_SIZE(3) Manual Page

## Name

VK_MAX_DEVICE_GROUP_SIZE - Length of a physical device handle array



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAX_DEVICE_GROUP_SIZE` is the length of an array containing
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle values representing all
physical devices in a group, as returned in
[VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGroupProperties.html)::`physicalDevices`.

``` c
#define VK_MAX_DEVICE_GROUP_SIZE          32U
```

or the equivalent

``` c
#define VK_MAX_DEVICE_GROUP_SIZE_KHR      VK_MAX_DEVICE_GROUP_SIZE
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group_creation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group_creation.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAX_DEVICE_GROUP_SIZE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
