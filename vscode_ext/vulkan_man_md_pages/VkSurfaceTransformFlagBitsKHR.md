# VkSurfaceTransformFlagBitsKHR(3) Manual Page

## Name

VkSurfaceTransformFlagBitsKHR - Presentation transforms supported on a
device



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedTransforms`
indicating the presentation transforms supported for the surface on the
specified device, and possible values of
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html)::`currentTransform`
indicating the surface’s current transform relative to the presentation
engine’s natural orientation, are:

``` c
// Provided by VK_KHR_surface
typedef enum VkSurfaceTransformFlagBitsKHR {
    VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR = 0x00000001,
    VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR = 0x00000002,
    VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR = 0x00000004,
    VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR = 0x00000008,
    VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR = 0x00000010,
    VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR = 0x00000020,
    VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR = 0x00000040,
    VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR = 0x00000080,
    VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR = 0x00000100,
} VkSurfaceTransformFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` specifies that image content
  is presented without being transformed.

- `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR` specifies that image content
  is rotated 90 degrees clockwise.

- `VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR` specifies that image content
  is rotated 180 degrees clockwise.

- `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR` specifies that image content
  is rotated 270 degrees clockwise.

- `VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR` specifies that image
  content is mirrored horizontally.

- `VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR` specifies
  that image content is mirrored horizontally, then rotated 90 degrees
  clockwise.

- `VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR` specifies
  that image content is mirrored horizontally, then rotated 180 degrees
  clockwise.

- `VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR` specifies
  that image content is mirrored horizontally, then rotated 270 degrees
  clockwise.

- `VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR` specifies that the presentation
  transform is not specified, and is instead determined by
  platform-specific considerations and mechanisms outside Vulkan.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html),
[VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html),
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html),
[VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html),
[VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html),
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html),
[VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagsKHR.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceTransformFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
