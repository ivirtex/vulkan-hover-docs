# VkRenderPassCreateFlagBits(3) Manual Page

## Name

VkRenderPassCreateFlagBits - Bitmask specifying additional properties of
a render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)::`flags`,
describing additional properties of the render pass, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkRenderPassCreateFlagBits {
  // Provided by VK_QCOM_render_pass_transform
    VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM = 0x00000002,
} VkRenderPassCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM` specifies that the created
  render pass is compatible with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-renderpass-transform"
  target="_blank" rel="noopener">render pass transform</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
