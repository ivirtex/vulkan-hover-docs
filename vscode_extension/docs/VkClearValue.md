# VkClearValue(3) Manual Page

## Name

VkClearValue - Structure specifying a clear value



## [](#_c_specification)C Specification

The `VkClearValue` union is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef union VkClearValue {
    VkClearColorValue           color;
    VkClearDepthStencilValue    depthStencil;
} VkClearValue;
```

## [](#_members)Members

- `color` specifies the color image clear values to use when clearing a color image or attachment.
- `depthStencil` specifies the depth and stencil clear values to use when clearing a depth/stencil image or attachment.

## [](#_description)Description

This union is used where part of the API requires either color or depth/stencil clear values, depending on the attachment, and defines the initial clear values in the [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) structure.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearAttachment](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearAttachment.html), [VkClearColorValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearColorValue.html), [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearDepthStencilValue.html), [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html), [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClearValue).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0