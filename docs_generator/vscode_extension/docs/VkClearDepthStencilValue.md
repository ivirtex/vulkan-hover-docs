# VkClearDepthStencilValue(3) Manual Page

## Name

VkClearDepthStencilValue - Structure specifying a clear depth stencil value



## [](#_c_specification)C Specification

The `VkClearDepthStencilValue` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkClearDepthStencilValue {
    float       depth;
    uint32_t    stencil;
} VkClearDepthStencilValue;
```

## [](#_members)Members

- `depth` is the clear value for the depth aspect of the depth/stencil attachment. It is a floating-point value which is automatically converted to the attachment’s format.
- `stencil` is the clear value for the stencil aspect of the depth/stencil attachment. It is a 32-bit integer value which is converted to the attachment’s format by taking the appropriate number of LSBs.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClearDepthStencilValue-depth-00022)VUID-VkClearDepthStencilValue-depth-00022  
  Unless the `VK_EXT_depth_range_unrestricted` extension is enabled `depth` **must** be between `0.0` and `1.0`, inclusive

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html), [vkCmdClearDepthStencilImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearDepthStencilImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClearDepthStencilValue)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0