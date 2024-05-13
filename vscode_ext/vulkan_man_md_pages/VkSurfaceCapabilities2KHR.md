# VkSurfaceCapabilities2KHR(3) Manual Page

## Name

VkSurfaceCapabilities2KHR - Structure describing capabilities of a
surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceCapabilities2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkSurfaceCapabilities2KHR {
    VkStructureType             sType;
    void*                       pNext;
    VkSurfaceCapabilitiesKHR    surfaceCapabilities;
} VkSurfaceCapabilities2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `surfaceCapabilities` is a
  [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html) structure
  describing the capabilities of the specified surface.

## <a href="#_description" class="anchor"></a>Description

If the [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
extension is enabled and
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface`
in the
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
call is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the values returned in
`minImageCount`, `maxImageCount`, `currentExtent`, and
`currentTransform` will not reflect that of any surface and will instead
be as such:

- `minImageCount` and `maxImageCount` will be 0xFFFFFFFF

- `currentExtent` will be (0xFFFFFFFF, 0xFFFFFFFF)

- `currentTransform` will be `VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceCapabilities2KHR-sType-sType"
  id="VUID-VkSurfaceCapabilities2KHR-sType-sType"></a>
  VUID-VkSurfaceCapabilities2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_KHR`

- <a href="#VUID-VkSurfaceCapabilities2KHR-pNext-pNext"
  id="VUID-VkSurfaceCapabilities2KHR-pNext-pNext"></a>
  VUID-VkSurfaceCapabilities2KHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html),
  [VkLatencySurfaceCapabilitiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySurfaceCapabilitiesNV.html),
  [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html),
  [VkSurfaceCapabilitiesFullScreenExclusiveEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesFullScreenExclusiveEXT.html),
  [VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html),
  [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html),
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html),
  or
  [VkSurfaceProtectedCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceProtectedCapabilitiesKHR.html)

- <a href="#VUID-VkSurfaceCapabilities2KHR-sType-unique"
  id="VUID-VkSurfaceCapabilities2KHR-sType-unique"></a>
  VUID-VkSurfaceCapabilities2KHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html),
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCapabilities2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
