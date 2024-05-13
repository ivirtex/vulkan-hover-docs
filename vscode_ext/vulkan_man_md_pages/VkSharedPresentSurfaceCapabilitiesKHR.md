# VkSharedPresentSurfaceCapabilitiesKHR(3) Manual Page

## Name

VkSharedPresentSurfaceCapabilitiesKHR - Structure describing
capabilities of a surface for shared presentation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSharedPresentSurfaceCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_shared_presentable_image
typedef struct VkSharedPresentSurfaceCapabilitiesKHR {
    VkStructureType      sType;
    void*                pNext;
    VkImageUsageFlags    sharedPresentSupportedUsageFlags;
} VkSharedPresentSurfaceCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `sharedPresentSupportedUsageFlags` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) representing the
  ways the application **can** use the shared presentable image from a
  swapchain created with [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) set
  to `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` or
  `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` for the surface on the
  specified device. `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` **must** be
  included in the set but implementations **may** support additional
  usages.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSharedPresentSurfaceCapabilitiesKHR-sType-sType"
  id="VUID-VkSharedPresentSurfaceCapabilitiesKHR-sType-sType"></a>
  VUID-VkSharedPresentSurfaceCapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SHARED_PRESENT_SURFACE_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shared_presentable_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shared_presentable_image.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSharedPresentSurfaceCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
