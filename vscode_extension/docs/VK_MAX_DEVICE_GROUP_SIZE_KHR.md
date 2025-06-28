# VK\_MAX\_DEVICE\_GROUP\_SIZE(3) Manual Page

## Name

VK\_MAX\_DEVICE\_GROUP\_SIZE - Length of a physical device handle array



## [](#_c_specification)C Specification

`VK_MAX_DEVICE_GROUP_SIZE` is the length of an array containing [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle values representing all physical devices in a group, as returned in [VkPhysicalDeviceGroupProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGroupProperties.html)::`physicalDevices`.

```c++
#define VK_MAX_DEVICE_GROUP_SIZE          32U
```

or the equivalent

```c++
#define VK_MAX_DEVICE_GROUP_SIZE_KHR      VK_MAX_DEVICE_GROUP_SIZE
```

## [](#_see_also)See Also

[VK\_KHR\_device\_group\_creation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group_creation.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MAX_DEVICE_GROUP_SIZE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0