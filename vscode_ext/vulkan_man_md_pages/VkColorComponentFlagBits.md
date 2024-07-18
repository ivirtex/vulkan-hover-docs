# VkColorComponentFlagBits(3) Manual Page

## Name

VkColorComponentFlagBits - Bitmask controlling which components are
written to the framebuffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)::`colorWriteMask`,
determining whether the final color values R, G, B and A are written to
the framebuffer attachment, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkColorComponentFlagBits {
    VK_COLOR_COMPONENT_R_BIT = 0x00000001,
    VK_COLOR_COMPONENT_G_BIT = 0x00000002,
    VK_COLOR_COMPONENT_B_BIT = 0x00000004,
    VK_COLOR_COMPONENT_A_BIT = 0x00000008,
} VkColorComponentFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COLOR_COMPONENT_R_BIT` specifies that the R value is written to
  the color attachment for the appropriate sample. Otherwise, the value
  in memory is unmodified.

- `VK_COLOR_COMPONENT_G_BIT` specifies that the G value is written to
  the color attachment for the appropriate sample. Otherwise, the value
  in memory is unmodified.

- `VK_COLOR_COMPONENT_B_BIT` specifies that the B value is written to
  the color attachment for the appropriate sample. Otherwise, the value
  in memory is unmodified.

- `VK_COLOR_COMPONENT_A_BIT` specifies that the A value is written to
  the color attachment for the appropriate sample. Otherwise, the value
  in memory is unmodified.

The color write mask operation is applied regardless of whether blending
is enabled.

The color write mask operation is applied only if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-color-write-enable"
target="_blank" rel="noopener">Color Write Enable</a> is enabled for the
respective attachment. Otherwise the color write mask is ignored and
writes to all components of the attachment are disabled.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkColorComponentFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
