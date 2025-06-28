# VkPresentGravityFlagBitsEXT(3) Manual Page

## Name

VkPresentGravityFlagBitsEXT - Bitmask specifying presentation pixel gravity on either the x or y axis



## [](#_c_specification)C Specification

Bits which **may** be set in the [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX` or `supportedPresentGravityY` fields, specifying the gravity of presented pixels supported by the surface, are:

```c++
// Provided by VK_EXT_surface_maintenance1
typedef enum VkPresentGravityFlagBitsEXT {
    VK_PRESENT_GRAVITY_MIN_BIT_EXT = 0x00000001,
    VK_PRESENT_GRAVITY_MAX_BIT_EXT = 0x00000002,
    VK_PRESENT_GRAVITY_CENTERED_BIT_EXT = 0x00000004,
} VkPresentGravityFlagBitsEXT;
```

## [](#_description)Description

- `VK_PRESENT_GRAVITY_MIN_BIT_EXT` means that the pixels will gravitate towards the top or left side of the surface.
- `VK_PRESENT_GRAVITY_MAX_BIT_EXT` means that the pixels will gravitate towards the bottom or right side of the surface.
- `VK_PRESENT_GRAVITY_CENTERED_BIT_EXT` means that the pixels will be centered in the surface.

If the value in [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`currentTransform` is not `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, it is implementation-defined whether the gravity configuration applies to the presented image before or after transformation.

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentGravityFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0