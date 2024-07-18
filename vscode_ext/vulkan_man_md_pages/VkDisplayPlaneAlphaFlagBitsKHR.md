# VkDisplayPlaneAlphaFlagBitsKHR(3) Manual Page

## Name

VkDisplayPlaneAlphaFlagBitsKHR - Alpha blending type



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html)::`alphaMode`,
specifying the type of alpha blending to use on a display, are:

``` c
// Provided by VK_KHR_display
typedef enum VkDisplayPlaneAlphaFlagBitsKHR {
    VK_DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR = 0x00000001,
    VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR = 0x00000002,
    VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR = 0x00000004,
    VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR = 0x00000008,
} VkDisplayPlaneAlphaFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DISPLAY_PLANE_ALPHA_OPAQUE_BIT_KHR` specifies that the source
  image will be treated as opaque.

- `VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR` specifies that a global alpha
  value **must** be specified that will be applied to all pixels in the
  source image.

- `VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR` specifies that the alpha
  value will be determined by the alpha component of the source image’s
  pixels. If the source format contains no alpha values, no blending
  will be applied. The source alpha values are not premultiplied into
  the source image’s other color components.

- `VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_PREMULTIPLIED_BIT_KHR` is equivalent
  to `VK_DISPLAY_PLANE_ALPHA_PER_PIXEL_BIT_KHR`, except the source alpha
  values are assumed to be premultiplied into the source image’s other
  color components.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayPlaneAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagsKHR.html),
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPlaneAlphaFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
