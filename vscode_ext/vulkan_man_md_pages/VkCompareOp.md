# VkCompareOp(3) Manual Page

## Name

VkCompareOp - Comparison operator for depth, stencil, and sampler
operations



## <a href="#_c_specification" class="anchor"></a>C Specification

*Comparison operators* compare a *reference* and a *test* value, and
return a true (“passed”) or false (“failed”) value depending on the
comparison operator chosen. The supported operators are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_COMPARE_OP_NEVER` specifies that the comparison always evaluates
  false.

- `VK_COMPARE_OP_LESS` specifies that the comparison evaluates
  *reference* \< *test*.

- `VK_COMPARE_OP_EQUAL` specifies that the comparison evaluates
  *reference* = *test*.

- `VK_COMPARE_OP_LESS_OR_EQUAL` specifies that the comparison evaluates
  *reference* ≤ *test*.

- `VK_COMPARE_OP_GREATER` specifies that the comparison evaluates
  *reference* \> *test*.

- `VK_COMPARE_OP_NOT_EQUAL` specifies that the comparison evaluates
  *reference* ≠ *test*.

- `VK_COMPARE_OP_GREATER_OR_EQUAL` specifies that the comparison
  evaluates *reference* ≥ *test*.

- `VK_COMPARE_OP_ALWAYS` specifies that the comparison always evaluates
  true.

Comparison operators are used for:

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-depth-compare-operation"
  target="_blank" rel="noopener">Depth Compare Operation</a> operator
  for a sampler, specified by
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`compareOp`.

- The stencil comparison operator for the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
  target="_blank" rel="noopener">stencil test</a>, specified by
  [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html)::`compareOp` or
  [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)::`compareOp`.

- The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth-comparison"
  target="_blank" rel="noopener">Depth Comparison</a> operator for the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
  target="_blank" rel="noopener">depth test</a>, specified by
  [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOp.html)::`depthCompareOp`
  or
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`depthCompareOp`.

Each such use describes how the *reference* and *test* values for that
comparison are determined.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html),
[VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html),
[vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOp.html),
[vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOpEXT.html),
[vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html),
[vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOpEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCompareOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
