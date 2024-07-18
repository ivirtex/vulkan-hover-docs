# VkPresentScalingFlagBitsEXT(3) Manual Page

## Name

VkPresentScalingFlagBitsEXT - Bitmask specifying presentation scaling
methods



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`,
specifying scaling modes supported by the surface, are:

``` c
// Provided by VK_EXT_surface_maintenance1
typedef enum VkPresentScalingFlagBitsEXT {
    VK_PRESENT_SCALING_ONE_TO_ONE_BIT_EXT = 0x00000001,
    VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_EXT = 0x00000002,
    VK_PRESENT_SCALING_STRETCH_BIT_EXT = 0x00000004,
} VkPresentScalingFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PRESENT_SCALING_ONE_TO_ONE_BIT_EXT` specifies that no scaling
  occurs, and pixels in the swapchain image are mapped to one and only
  one pixel in the surface. The mapping between pixels is defined by the
  chosen presentation gravity.

- `VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_EXT` specifies that the
  swapchain image will be minified or magnified such that at least one
  of the resulting width or height is equal to the corresponding surface
  dimension, and the other resulting dimension is less than or equal to
  the corresponding surface dimension, with the aspect ratio of the
  resulting image being identical to that of the original swapchain
  image.

- `VK_PRESENT_SCALING_STRETCH_BIT_EXT` specifies that the swapchain
  image will be minified or magnified such that the resulting image
  dimensions are equal to those of the surface.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html),
[VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentScalingFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
