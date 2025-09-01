# VkStencilOpState(3) Manual Page

## Name

VkStencilOpState - Structure specifying stencil operation state



## [](#_c_specification)C Specification

The `VkStencilOpState` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkStencilOpState {
    VkStencilOp    failOp;
    VkStencilOp    passOp;
    VkStencilOp    depthFailOp;
    VkCompareOp    compareOp;
    uint32_t       compareMask;
    uint32_t       writeMask;
    uint32_t       reference;
} VkStencilOpState;
```

## [](#_members)Members

- `failOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that fail the stencil test.
- `passOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that pass both the depth and stencil tests.
- `depthFailOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value specifying the action performed on samples that pass the stencil test and fail the depth test.
- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value specifying the comparison operator used in the stencil test.
- `compareMask` selects the bits of the unsigned integer stencil values participating in the stencil test.
- `writeMask` selects the bits of the unsigned integer stencil values updated by the stencil test in the stencil framebuffer attachment.
- `reference` is an integer stencil reference value that is used in the unsigned stencil comparison.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkStencilOpState-failOp-parameter)VUID-VkStencilOpState-failOp-parameter  
  `failOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-VkStencilOpState-passOp-parameter)VUID-VkStencilOpState-passOp-parameter  
  `passOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-VkStencilOpState-depthFailOp-parameter)VUID-VkStencilOpState-depthFailOp-parameter  
  `depthFailOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html) value
- [](#VUID-VkStencilOpState-compareOp-parameter)VUID-VkStencilOpState-compareOp-parameter  
  `compareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html), [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html), [VkStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOp.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkStencilOpState).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0