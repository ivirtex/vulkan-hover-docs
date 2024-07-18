# VkFramebufferCreateFlagBits(3) Manual Page

## Name

VkFramebufferCreateFlagBits - Bitmask specifying framebuffer properties



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags`,
specifying options for framebuffers, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkFramebufferCreateFlagBits {
  // Provided by VK_VERSION_1_2
    VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT = 0x00000001,
  // Provided by VK_KHR_imageless_framebuffer
    VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR = VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT,
} VkFramebufferCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` specifies that image views are
  not specified, and only attachment compatibility information will be
  provided via a
  [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html)
  structure.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFramebufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFramebufferCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
