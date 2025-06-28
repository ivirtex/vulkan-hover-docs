# VkSharedPresentSurfaceCapabilitiesKHR(3) Manual Page

## Name

VkSharedPresentSurfaceCapabilitiesKHR - Structure describing capabilities of a surface for shared presentation



## [](#_c_specification)C Specification

The `VkSharedPresentSurfaceCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_shared_presentable_image
typedef struct VkSharedPresentSurfaceCapabilitiesKHR {
    VkStructureType      sType;
    void*                pNext;
    VkImageUsageFlags    sharedPresentSupportedUsageFlags;
} VkSharedPresentSurfaceCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sharedPresentSupportedUsageFlags` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) representing the ways the application **can** use the shared presentable image from a swapchain created with [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) set to `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` for the surface on the specified device. `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` **must** be included in the set but implementations **may** support additional usages.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSharedPresentSurfaceCapabilitiesKHR-sType-sType)VUID-VkSharedPresentSurfaceCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_shared\_presentable\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shared_presentable_image.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSharedPresentSurfaceCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0