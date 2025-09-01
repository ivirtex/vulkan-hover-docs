# VkSurfaceCounterFlagBitsEXT(3) Manual Page

## Name

VkSurfaceCounterFlagBitsEXT - Surface-relative counter types



## [](#_c_specification)C Specification

Bits which **can** be set in [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html)::`supportedSurfaceCounters`, indicating supported surface counter types, are:

```c++
// Provided by VK_EXT_display_surface_counter
typedef enum VkSurfaceCounterFlagBitsEXT {
    VK_SURFACE_COUNTER_VBLANK_BIT_EXT = 0x00000001,
  // VK_SURFACE_COUNTER_VBLANK_EXT is a deprecated alias
    VK_SURFACE_COUNTER_VBLANK_EXT = VK_SURFACE_COUNTER_VBLANK_BIT_EXT,
} VkSurfaceCounterFlagBitsEXT;
```

## [](#_description)Description

- `VK_SURFACE_COUNTER_VBLANK_BIT_EXT` specifies a counter incrementing once every time a vertical blanking period occurs on the display associated with the surface.

## [](#_see_also)See Also

[VK\_EXT\_display\_surface\_counter](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_surface_counter.html), [VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagsEXT.html), [vkGetSwapchainCounterEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainCounterEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCounterFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0