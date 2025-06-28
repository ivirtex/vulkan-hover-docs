# VK\_UUID\_SIZE(3) Manual Page

## Name

VK\_UUID\_SIZE - Length of a universally unique device or driver build identifier



## [](#_c_specification)C Specification

`VK_UUID_SIZE` is the length in `uint8_t` values of an array containing a universally unique device or driver build identifier, as returned in [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html)::`deviceUUID` and [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceIDProperties.html)::`driverUUID`.

```c++
#define VK_UUID_SIZE                      16U
```

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_UUID_SIZE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0