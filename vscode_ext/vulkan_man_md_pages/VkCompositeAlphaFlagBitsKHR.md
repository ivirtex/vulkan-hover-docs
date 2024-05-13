# VkCompositeAlphaFlagBitsKHR(3) Manual Page

## Name

VkCompositeAlphaFlagBitsKHR - Alpha compositing modes supported on a
device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `supportedCompositeAlpha` member is of type
[VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html),
containing the following values:

``` c
// Provided by VK_KHR_surface
typedef enum VkCompositeAlphaFlagBitsKHR {
    VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR = 0x00000001,
    VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR = 0x00000002,
    VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR = 0x00000004,
    VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR = 0x00000008,
} VkCompositeAlphaFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

These values are described as follows:

- `VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR`: The alpha component, if it
  exists, of the images is ignored in the compositing process. Instead,
  the image is treated as if it has a constant alpha of 1.0.

- `VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR`: The alpha component, if
  it exists, of the images is respected in the compositing process. The
  non-alpha components of the image are expected to already be
  multiplied by the alpha component by the application.

- `VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR`: The alpha component, if
  it exists, of the images is respected in the compositing process. The
  non-alpha components of the image are not expected to already be
  multiplied by the alpha component by the application; instead, the
  compositor will multiply the non-alpha components of the image by the
  alpha component during compositing.

- `VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR`: The way in which the
  presentation engine treats the alpha component in the images is
  unknown to the Vulkan API. Instead, the application is responsible for
  setting the composite alpha blending mode using native window system
  commands. If the application does not set the blending mode using
  native window system commands, then a platform-specific default will
  be used.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagsKHR.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCompositeAlphaFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
