# VkLogicOp(3) Manual Page

## Name

VkLogicOp - Framebuffer logical operations



## <a href="#_c_specification" class="anchor"></a>C Specification

Logical operations are controlled by the `logicOpEnable` and `logicOp`
members of
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html).
The `logicOpEnable` state can also be controlled by
[vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html) if graphics
pipeline is created with `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
The `logicOp` state can also be controlled by
[vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html) if graphics pipeline is
created with `VK_DYNAMIC_STATE_LOGIC_OP_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
If `logicOpEnable` is `VK_TRUE`, then a logical operation selected by
`logicOp` is applied between each color attachment and the fragment’s
corresponding output value, and blending of all attachments is treated
as if it were disabled. Any attachments using color formats for which
logical operations are not supported simply pass through the color
values unmodified. The logical operation is applied independently for
each of the red, green, blue, and alpha components. The `logicOp` is
selected from the following operations:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkLogicOp {
    VK_LOGIC_OP_CLEAR = 0,
    VK_LOGIC_OP_AND = 1,
    VK_LOGIC_OP_AND_REVERSE = 2,
    VK_LOGIC_OP_COPY = 3,
    VK_LOGIC_OP_AND_INVERTED = 4,
    VK_LOGIC_OP_NO_OP = 5,
    VK_LOGIC_OP_XOR = 6,
    VK_LOGIC_OP_OR = 7,
    VK_LOGIC_OP_NOR = 8,
    VK_LOGIC_OP_EQUIVALENT = 9,
    VK_LOGIC_OP_INVERT = 10,
    VK_LOGIC_OP_OR_REVERSE = 11,
    VK_LOGIC_OP_COPY_INVERTED = 12,
    VK_LOGIC_OP_OR_INVERTED = 13,
    VK_LOGIC_OP_NAND = 14,
    VK_LOGIC_OP_SET = 15,
} VkLogicOp;
```

## <a href="#_description" class="anchor"></a>Description

The logical operations supported by Vulkan are summarized in the
following table in which

- ¬ is bitwise invert,

- ∧ is bitwise and,

- ∨ is bitwise or,

- ⊕ is bitwise exclusive or,

- s is the fragment’s R<sub>s0</sub>, G<sub>s0</sub>, B<sub>s0</sub> or
  A<sub>s0</sub> component value for the fragment output corresponding
  to the color attachment being updated, and

- d is the color attachment’s R, G, B or A component value:

| Mode                        | Operation |
|-----------------------------|-----------|
| `VK_LOGIC_OP_CLEAR`         | 0         |
| `VK_LOGIC_OP_AND`           | s ∧ d     |
| `VK_LOGIC_OP_AND_REVERSE`   | s ∧ ¬ d   |
| `VK_LOGIC_OP_COPY`          | s         |
| `VK_LOGIC_OP_AND_INVERTED`  | ¬ s ∧ d   |
| `VK_LOGIC_OP_NO_OP`         | d         |
| `VK_LOGIC_OP_XOR`           | s ⊕ d     |
| `VK_LOGIC_OP_OR`            | s ∨ d     |
| `VK_LOGIC_OP_NOR`           | ¬ (s ∨ d) |
| `VK_LOGIC_OP_EQUIVALENT`    | ¬ (s ⊕ d) |
| `VK_LOGIC_OP_INVERT`        | ¬ d       |
| `VK_LOGIC_OP_OR_REVERSE`    | s ∨ ¬ d   |
| `VK_LOGIC_OP_COPY_INVERTED` | ¬ s       |
| `VK_LOGIC_OP_OR_INVERTED`   | ¬ s ∨ d   |
| `VK_LOGIC_OP_NAND`          | ¬ (s ∧ d) |
| `VK_LOGIC_OP_SET`           | all 1s    |

Table 1. Logical Operations

The result of the logical operation is then written to the color
attachment as controlled by the component write mask, described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blendoperations"
target="_blank" rel="noopener">Blend Operations</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html),
[vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLogicOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
