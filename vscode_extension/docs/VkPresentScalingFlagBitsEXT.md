# VkPresentScalingFlagBitsKHR(3) Manual Page

## Name

VkPresentScalingFlagBitsKHR - Bitmask specifying presentation scaling methods



## [](#_c_specification)C Specification

Bits which **may** be set in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentScaling`, specifying scaling modes supported by the surface, are:

```c++
// Provided by VK_KHR_surface_maintenance1
typedef enum VkPresentScalingFlagBitsKHR {
    VK_PRESENT_SCALING_ONE_TO_ONE_BIT_KHR = 0x00000001,
    VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_KHR = 0x00000002,
    VK_PRESENT_SCALING_STRETCH_BIT_KHR = 0x00000004,
    VK_PRESENT_SCALING_ONE_TO_ONE_BIT_EXT = VK_PRESENT_SCALING_ONE_TO_ONE_BIT_KHR,
    VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_EXT = VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_KHR,
    VK_PRESENT_SCALING_STRETCH_BIT_EXT = VK_PRESENT_SCALING_STRETCH_BIT_KHR,
} VkPresentScalingFlagBitsKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_surface_maintenance1
typedef VkPresentScalingFlagBitsKHR VkPresentScalingFlagBitsEXT;
```

## [](#_description)Description

- `VK_PRESENT_SCALING_ONE_TO_ONE_BIT_KHR` specifies that no scaling occurs, and pixels in the swapchain image are mapped to one and only one pixel in the surface. The mapping between pixels is defined by the chosen presentation gravity.
- `VK_PRESENT_SCALING_ASPECT_RATIO_STRETCH_BIT_KHR` specifies that the swapchain image will be minified or magnified such that at least one of the resulting width or height is equal to the corresponding surface dimension, and the other resulting dimension is less than or equal to the corresponding surface dimension, with the aspect ratio of the resulting image being identical to that of the original swapchain image.
- `VK_PRESENT_SCALING_STRETCH_BIT_KHR` specifies that the swapchain image will be minified or magnified such that the resulting image dimensions are equal to those of the surface.

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html), [VkPresentScalingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentScalingFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0