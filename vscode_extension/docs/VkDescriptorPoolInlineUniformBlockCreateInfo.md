# VkDescriptorPoolInlineUniformBlockCreateInfo(3) Manual Page

## Name

VkDescriptorPoolInlineUniformBlockCreateInfo - Structure specifying the maximum number of inline uniform block bindings of a newly created descriptor pool



## [](#_c_specification)C Specification

In order to be able to allocate descriptor sets having [inline uniform block](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-inlineuniformblock) bindings the descriptor pool **must** be created with specifying the inline uniform block binding capacity of the descriptor pool, in addition to the total inline uniform data capacity in bytes which is specified through a [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolSize.html) structure with a `descriptorType` value of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`. This **can** be done by adding a `VkDescriptorPoolInlineUniformBlockCreateInfo` structure to the `pNext` chain of [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html).

The `VkDescriptorPoolInlineUniformBlockCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkDescriptorPoolInlineUniformBlockCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maxInlineUniformBlockBindings;
} VkDescriptorPoolInlineUniformBlockCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_inline_uniform_block
typedef VkDescriptorPoolInlineUniformBlockCreateInfo VkDescriptorPoolInlineUniformBlockCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxInlineUniformBlockBindings` is the number of inline uniform block bindings to allocate.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDescriptorPoolInlineUniformBlockCreateInfo-sType-sType)VUID-VkDescriptorPoolInlineUniformBlockCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO`

## [](#_see_also)See Also

[VK\_EXT\_inline\_uniform\_block](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_inline_uniform_block.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorPoolInlineUniformBlockCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0