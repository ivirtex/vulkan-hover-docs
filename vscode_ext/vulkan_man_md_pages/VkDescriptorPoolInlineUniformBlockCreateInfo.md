# VkDescriptorPoolInlineUniformBlockCreateInfo(3) Manual Page

## Name

VkDescriptorPoolInlineUniformBlockCreateInfo - Structure specifying the
maximum number of inline uniform block bindings of a newly created
descriptor pool



## <a href="#_c_specification" class="anchor"></a>C Specification

In order to be able to allocate descriptor sets having <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inlineuniformblock"
target="_blank" rel="noopener">inline uniform block</a> bindings the
descriptor pool **must** be created with specifying the inline uniform
block binding capacity of the descriptor pool, in addition to the total
inline uniform data capacity in bytes which is specified through a
[VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolSize.html) structure with a
`descriptorType` value of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`.
This **can** be done by adding a
`VkDescriptorPoolInlineUniformBlockCreateInfo` structure to the `pNext`
chain of [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html).

The `VkDescriptorPoolInlineUniformBlockCreateInfo` structure is defined
as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkDescriptorPoolInlineUniformBlockCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           maxInlineUniformBlockBindings;
} VkDescriptorPoolInlineUniformBlockCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_inline_uniform_block
typedef VkDescriptorPoolInlineUniformBlockCreateInfo VkDescriptorPoolInlineUniformBlockCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxInlineUniformBlockBindings` is the number of inline uniform block
  bindings to allocate.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorPoolInlineUniformBlockCreateInfo-sType-sType"
  id="VUID-VkDescriptorPoolInlineUniformBlockCreateInfo-sType-sType"></a>
  VUID-VkDescriptorPoolInlineUniformBlockCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_inline_uniform_block](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_inline_uniform_block.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorPoolInlineUniformBlockCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
