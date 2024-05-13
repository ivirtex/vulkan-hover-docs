# VkSurfaceCounterFlagBitsEXT(3) Manual Page

## Name

VkSurfaceCounterFlagBitsEXT - Surface-relative counter types



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html)::`supportedSurfaceCounters`,
indicating supported surface counter types, are:

``` c
// Provided by VK_EXT_display_surface_counter
typedef enum VkSurfaceCounterFlagBitsEXT {
    VK_SURFACE_COUNTER_VBLANK_BIT_EXT = 0x00000001,
    VK_SURFACE_COUNTER_VBLANK_EXT = VK_SURFACE_COUNTER_VBLANK_BIT_EXT,
} VkSurfaceCounterFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SURFACE_COUNTER_VBLANK_BIT_EXT` specifies a counter incrementing
  once every time a vertical blanking period occurs on the display
  associated with the surface.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_surface_counter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_surface_counter.html),
[VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagsEXT.html),
[vkGetSwapchainCounterEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainCounterEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCounterFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
