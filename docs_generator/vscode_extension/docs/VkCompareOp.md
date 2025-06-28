# VkCompareOp(3) Manual Page

## Name

VkCompareOp - Comparison operator for depth, stencil, and sampler operations



## [](#_c_specification)C Specification

*Comparison operators* compare a *reference* and a *test* value, and return a true (“passed”) or false (“failed”) value depending on the comparison operator chosen. The supported operators are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkCompareOp {
    VK_COMPARE_OP_NEVER = 0,
    VK_COMPARE_OP_LESS = 1,
    VK_COMPARE_OP_EQUAL = 2,
    VK_COMPARE_OP_LESS_OR_EQUAL = 3,
    VK_COMPARE_OP_GREATER = 4,
    VK_COMPARE_OP_NOT_EQUAL = 5,
    VK_COMPARE_OP_GREATER_OR_EQUAL = 6,
    VK_COMPARE_OP_ALWAYS = 7,
} VkCompareOp;
```

## [](#_description)Description

- `VK_COMPARE_OP_NEVER` specifies that the comparison always evaluates false.
- `VK_COMPARE_OP_LESS` specifies that the comparison evaluates *reference* &lt; *test*.
- `VK_COMPARE_OP_EQUAL` specifies that the comparison evaluates *reference* = *test*.
- `VK_COMPARE_OP_LESS_OR_EQUAL` specifies that the comparison evaluates *reference* ≤ *test*.
- `VK_COMPARE_OP_GREATER` specifies that the comparison evaluates *reference* &gt; *test*.
- `VK_COMPARE_OP_NOT_EQUAL` specifies that the comparison evaluates *reference* ≠ *test*.
- `VK_COMPARE_OP_GREATER_OR_EQUAL` specifies that the comparison evaluates *reference* ≥ *test*.
- `VK_COMPARE_OP_ALWAYS` specifies that the comparison always evaluates true.

Comparison operators are used for:

- The [Depth Compare Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-depth-compare-operation) operator for a sampler, specified by [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`compareOp`.
- The stencil comparison operator for the [stencil test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil), specified by [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOp.html)::`compareOp` or [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html)::`compareOp`.
- The [Depth Comparison](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth-comparison) operator for the [depth test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth), specified by [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOp.html)::`depthCompareOp` or [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`depthCompareOp`.

Each such use describes how the *reference* and *test* values for that comparison are determined.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html), [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html), [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html), [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOp.html), [vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOpEXT.html), [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOp.html), [vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOpEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCompareOp)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0