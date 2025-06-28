# VkPhysicalDeviceLayeredDriverPropertiesMSFT(3) Manual Page

## Name

VkPhysicalDeviceLayeredDriverPropertiesMSFT - Structure containing information about driver layering for a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceLayeredDriverPropertiesMSFT` structure is defined as:

```c++
// Provided by VK_MSFT_layered_driver
typedef struct VkPhysicalDeviceLayeredDriverPropertiesMSFT {
    VkStructureType                     sType;
    void*                               pNext;
    VkLayeredDriverUnderlyingApiMSFT    underlyingAPI;
} VkPhysicalDeviceLayeredDriverPropertiesMSFT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `underlyingAPI` is a [VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayeredDriverUnderlyingApiMSFT.html) value indicating which underlying API is used to implement the layered driver, or `VK_LAYERED_DRIVER_UNDERLYING_API_NONE_MSFT` if the driver is not layered.

## [](#_description)Description

These are properties of the driver layering information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceLayeredDriverPropertiesMSFT-sType-sType)VUID-VkPhysicalDeviceLayeredDriverPropertiesMSFT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_DRIVER_PROPERTIES_MSFT`

## [](#_see_also)See Also

[VK\_MSFT\_layered\_driver](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_MSFT_layered_driver.html), [VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayeredDriverUnderlyingApiMSFT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLayeredDriverPropertiesMSFT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0