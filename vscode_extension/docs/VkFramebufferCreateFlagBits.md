# VkFramebufferCreateFlagBits(3) Manual Page

## Name

VkFramebufferCreateFlagBits - Bitmask specifying framebuffer properties



## [](#_c_specification)C Specification

Bits which **can** be set in [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags`, specifying options for framebuffers, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkFramebufferCreateFlagBits {
  // Provided by VK_VERSION_1_2
    VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT = 0x00000001,
  // Provided by VK_KHR_imageless_framebuffer
    VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR = VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT,
} VkFramebufferCreateFlagBits;
```

## [](#_description)Description

- `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` specifies that image views are not specified, and only attachment compatibility information will be provided via a [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html) structure.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFramebufferCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebufferCreateFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0