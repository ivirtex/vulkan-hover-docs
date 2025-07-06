# VkSurfaceCapabilities2KHR(3) Manual Page

## Name

VkSurfaceCapabilities2KHR - Structure describing capabilities of a surface



## [](#_c_specification)C Specification

The `VkSurfaceCapabilities2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkSurfaceCapabilities2KHR {
    VkStructureType             sType;
    void*                       pNext;
    VkSurfaceCapabilitiesKHR    surfaceCapabilities;
} VkSurfaceCapabilities2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `surfaceCapabilities` is a [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html) structure describing the capabilities of the specified surface.

## [](#_description)Description

If the `VK_GOOGLE_surfaceless_query` extension is enabled and [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface` in the [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) call is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the values returned in `minImageCount`, `maxImageCount`, `currentExtent`, and `currentTransform` will not reflect that of any surface and will instead be as such:

- `minImageCount` and `maxImageCount` will be 0xFFFFFFFF
- `currentExtent` will be (0xFFFFFFFF, 0xFFFFFFFF)
- `currentTransform` will be `VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-VkSurfaceCapabilities2KHR-sType-sType)VUID-VkSurfaceCapabilities2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR`
- [](#VUID-VkSurfaceCapabilities2KHR-pNext-pNext)VUID-VkSurfaceCapabilities2KHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html), [VkLatencySurfaceCapabilitiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySurfaceCapabilitiesNV.html), [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html), [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html), [VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html), [VkSurfaceCapabilitiesPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentId2KHR.html), [VkSurfaceCapabilitiesPresentWait2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentWait2KHR.html), [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html), [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html), or [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceProtectedCapabilitiesKHR.html)
- [](#VUID-VkSurfaceCapabilities2KHR-sType-unique)VUID-VkSurfaceCapabilities2KHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html), [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCapabilities2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0