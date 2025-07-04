# VkPresentScalingFlagBitsEXT(3) Manual Page

## Name

VkPresentScalingFlagBitsEXT - Bitmask specifying presentation scaling methods



## [](#_c_specification)C Specification

Bits which **may** be set in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`, specifying scaling modes supported by the surface, are:

```c++
// Provided by VK_EXT_surface_maintenance1
typedef enum VkPresentScalingFlagBitsEXT {
    VK_PRESENT_SCALING_ONE_TO_ONE_BIT_EXT = 0x00000001,
    VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_EXT = 0x00000002,
    VK_PRESENT_SCALING_STRETCH_BIT_EXT = 0x00000004,
} VkPresentScalingFlagBitsEXT;
```

## [](#_description)Description

- `VK_PRESENT_SCALING_ONE_TO_ONE_BIT_EXT` specifies that no scaling occurs, and pixels in the swapchain image are mapped to one and only one pixel in the surface. The mapping between pixels is defined by the chosen presentation gravity.
- `VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_EXT` specifies that the swapchain image will be minified or magnified such that at least one of the resulting width or height is equal to the corresponding surface dimension, and the other resulting dimension is less than or equal to the corresponding surface dimension, with the aspect ratio of the resulting image being identical to that of the original swapchain image.
- `VK_PRESENT_SCALING_STRETCH_BIT_EXT` specifies that the swapchain image will be minified or magnified such that the resulting image dimensions are equal to those of the surface.

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentScalingFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0