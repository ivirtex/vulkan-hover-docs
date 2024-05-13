# VkClearValue(3) Manual Page

## Name

VkClearValue - Structure specifying a clear value



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkClearValue` union is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef union VkClearValue {
    VkClearColorValue           color;
    VkClearDepthStencilValue    depthStencil;
} VkClearValue;
```

## <a href="#_members" class="anchor"></a>Members

- `color` specifies the color image clear values to use when clearing a
  color image or attachment.

- `depthStencil` specifies the depth and stencil clear values to use
  when clearing a depth/stencil image or attachment.

## <a href="#_description" class="anchor"></a>Description

This union is used where part of the API requires either color or
depth/stencil clear values, depending on the attachment, and defines the
initial clear values in the
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) structure.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearAttachment](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearAttachment.html),
[VkClearColorValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearColorValue.html),
[VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearDepthStencilValue.html),
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkClearValue"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
