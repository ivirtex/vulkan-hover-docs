# VkDisplayPlaneProperties2KHR(3) Manual Page

## Name

VkDisplayPlaneProperties2KHR - Structure describing an available display plane



## [](#_c_specification)C Specification

The `VkDisplayPlaneProperties2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayPlaneProperties2KHR {
    VkStructureType                sType;
    void*                          pNext;
    VkDisplayPlanePropertiesKHR    displayPlaneProperties;
} VkDisplayPlaneProperties2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `displayPlaneProperties` is a [VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlanePropertiesKHR.html) structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayPlaneProperties2KHR-sType-sType)VUID-VkDisplayPlaneProperties2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR`
- [](#VUID-VkDisplayPlaneProperties2KHR-pNext-pNext)VUID-VkDisplayPlaneProperties2KHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlanePropertiesKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayPlaneProperties2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0