# VK\_MAX\_DRIVER\_INFO\_SIZE(3) Manual Page

## Name

VK\_MAX\_DRIVER\_INFO\_SIZE - Length of a physical device driver information string



## [](#_c_specification)C Specification

`VK_MAX_DRIVER_INFO_SIZE` is the length in `char` values of an array containing a driver information string, as returned in [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverProperties.html)::`driverInfo`.

```c++
#define VK_MAX_DRIVER_INFO_SIZE           256U
```

or the equivalent

```c++
#define VK_MAX_DRIVER_INFO_SIZE_KHR       VK_MAX_DRIVER_INFO_SIZE
```

## [](#_see_also)See Also

[VK\_KHR\_driver\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_driver_properties.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MAX_DRIVER_INFO_SIZE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0