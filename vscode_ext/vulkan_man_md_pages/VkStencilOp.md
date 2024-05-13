# VkStencilOp(3) Manual Page

## Name

VkStencilOp - Stencil comparison function



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `failOp`, `passOp`, and `depthFailOp` members of
[VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html), specifying what happens to
the stored stencil value if this or certain subsequent tests fail or
pass, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkStencilOp {
    VK_STENCIL_OP_KEEP = 0,
    VK_STENCIL_OP_ZERO = 1,
    VK_STENCIL_OP_REPLACE = 2,
    VK_STENCIL_OP_INCREMENT_AND_CLAMP = 3,
    VK_STENCIL_OP_DECREMENT_AND_CLAMP = 4,
    VK_STENCIL_OP_INVERT = 5,
    VK_STENCIL_OP_INCREMENT_AND_WRAP = 6,
    VK_STENCIL_OP_DECREMENT_AND_WRAP = 7,
} VkStencilOp;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_STENCIL_OP_KEEP` keeps the current value.

- `VK_STENCIL_OP_ZERO` sets the value to 0.

- `VK_STENCIL_OP_REPLACE` sets the value to `reference`.

- `VK_STENCIL_OP_INCREMENT_AND_CLAMP` increments the current value and
  clamps to the maximum representable unsigned value.

- `VK_STENCIL_OP_DECREMENT_AND_CLAMP` decrements the current value and
  clamps to 0.

- `VK_STENCIL_OP_INVERT` bitwise-inverts the current value.

- `VK_STENCIL_OP_INCREMENT_AND_WRAP` increments the current value and
  wraps to 0 when the maximum value would have been exceeded.

- `VK_STENCIL_OP_DECREMENT_AND_WRAP` decrements the current value and
  wraps to the maximum possible value when the value would go below 0.

For purposes of increment and decrement, the stencil bits are considered
as an unsigned integer.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html),
[vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html),
[vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOpEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkStencilOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
