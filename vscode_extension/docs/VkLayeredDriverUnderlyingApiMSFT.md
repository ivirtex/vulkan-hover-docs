# VkLayeredDriverUnderlyingApiMSFT(3) Manual Page

## Name

VkLayeredDriverUnderlyingApiMSFT - Layered driver underlying APIs



## [](#_c_specification)C Specification

Underlying APIs which **may** be returned in [VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)::`underlyingAPI` are:

```c++
// Provided by VK_MSFT_layered_driver
typedef enum VkLayeredDriverUnderlyingApiMSFT {
    VK_LAYERED_DRIVER_UNDERLYING_API_NONE_MSFT = 0,
    VK_LAYERED_DRIVER_UNDERLYING_API_D3D12_MSFT = 1,
} VkLayeredDriverUnderlyingApiMSFT;
```

## [](#_see_also)See Also

[VK\_MSFT\_layered\_driver](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MSFT_layered_driver.html), [VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLayeredDriverUnderlyingApiMSFT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0