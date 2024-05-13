# VkPresentGravityFlagBitsEXT(3) Manual Page

## Name

VkPresentGravityFlagBitsEXT - Bitmask specifying presentation pixel
gravity on either the x or y axis



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in the
[VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX`
or `supportedPresentGravityY` fields, specifying the gravity of
presented pixels supported by the surface, are:

``` c
// Provided by VK_EXT_surface_maintenance1
typedef enum VkPresentGravityFlagBitsEXT {
    VK_PRESENT_GRAVITY_MIN_BIT_EXT = 0x00000001,
    VK_PRESENT_GRAVITY_MAX_BIT_EXT = 0x00000002,
    VK_PRESENT_GRAVITY_CENTERED_BIT_EXT = 0x00000004,
} VkPresentGravityFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PRESENT_GRAVITY_MIN_BIT_EXT` means that the pixels will gravitate
  towards the top or left side of the surface.

- `VK_PRESENT_GRAVITY_MAX_BIT_EXT` means that the pixels will gravitate
  towards the bottom or right side of the surface.

- `VK_PRESENT_GRAVITY_CENTERED_BIT_EXT` means that the pixels will be
  centered in the surface.

If the value in
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`currentTransform`
is not `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, it is
implementation-defined whether the gravity configuration applies to the
presented image before or after transformation.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html),
[VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentGravityFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
