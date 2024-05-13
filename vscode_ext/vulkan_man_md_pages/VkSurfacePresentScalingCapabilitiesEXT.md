# VkSurfacePresentScalingCapabilitiesEXT(3) Manual Page

## Name

VkSurfacePresentScalingCapabilitiesEXT - Structure describing the
presentation scaling capabilities of the surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfacePresentScalingCapabilitiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_surface_maintenance1
typedef struct VkSurfacePresentScalingCapabilitiesEXT {
    VkStructureType             sType;
    void*                       pNext;
    VkPresentScalingFlagsEXT    supportedPresentScaling;
    VkPresentGravityFlagsEXT    supportedPresentGravityX;
    VkPresentGravityFlagsEXT    supportedPresentGravityY;
    VkExtent2D                  minScaledImageExtent;
    VkExtent2D                  maxScaledImageExtent;
} VkSurfacePresentScalingCapabilitiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `supportedPresentScaling` is a bitmask of
  [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagBitsEXT.html)
  representing the scaling methods supported by the surface, or `0` if
  application-defined scaling is not supported.

- `supportedPresentGravityX` is a bitmask of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html)
  representing the X-axis pixel gravity supported by the surface, or `0`
  if Vulkan-defined pixel gravity is not supported for the X axis.

- `supportedPresentGravityY` is a bitmask of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html)
  representing the Y-axis pixel gravity supported by the surface, or `0`
  if Vulkan-defined pixel gravity is not supported for the Y axis.

- `minScaledImageExtent` contains the smallest valid swapchain extent
  for the surface on the specified device when one of the scaling
  methods specified in `supportedPresentScaling` is used, or the special
  value (0xFFFFFFFF, 0xFFFFFFFF) indicating that the surface size will
  be determined by the extent of a swapchain targeting the surface. The
  `width` and `height` of the extent will each be smaller than or equal
  to the corresponding `width` and `height` of
  [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageExtent`.

- `maxScaledImageExtent` contains the largest valid swapchain extent for
  the surface on the specified device when one of the scaling methods
  specified in `supportedPresentScaling` is used, or the special value
  described above for `minScaledImageExtent`. The `width` and `height`
  of the extent will each be greater than or equal to the corresponding
  `width` and `height` of
  [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`maxImageExtent`.

## <a href="#_description" class="anchor"></a>Description

Before creating a swapchain whose scaling mode **can** be specified
through the use of
[VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html),
obtain the set of supported scaling modes by including a
[VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) structure in the
`pNext` chain of
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
when calling
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).
The implementation **must** return the same values in
`VkSurfacePresentScalingCapabilitiesEXT` for any of the compatible
present modes as obtained through
[VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html).

Valid Usage (Implicit)

- <a href="#VUID-VkSurfacePresentScalingCapabilitiesEXT-sType-sType"
  id="VUID-VkSurfacePresentScalingCapabilitiesEXT-sType-sType"></a>
  VUID-VkSurfacePresentScalingCapabilitiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_PRESENT_SCALING_CAPABILITIES_EXT`

- <a
  href="#VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentScaling-parameter"
  id="VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentScaling-parameter"></a>
  VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentScaling-parameter  
  `supportedPresentScaling` **must** be a valid combination of
  [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagBitsEXT.html) values

- <a
  href="#VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityX-parameter"
  id="VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityX-parameter"></a>
  VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityX-parameter  
  `supportedPresentGravityX` **must** be a valid combination of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html) values

- <a
  href="#VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityY-parameter"
  id="VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityY-parameter"></a>
  VUID-VkSurfacePresentScalingCapabilitiesEXT-supportedPresentGravityY-parameter  
  `supportedPresentGravityY` **must** be a valid combination of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagsEXT.html),
[VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfacePresentScalingCapabilitiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
