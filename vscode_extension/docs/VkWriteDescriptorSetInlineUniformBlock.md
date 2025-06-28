# VkWriteDescriptorSetInlineUniformBlock(3) Manual Page

## Name

VkWriteDescriptorSetInlineUniformBlock - Structure specifying inline uniform block data



## [](#_c_specification)C Specification

If the `descriptorType` member of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then the data to write to the descriptor set is specified through a `VkWriteDescriptorSetInlineUniformBlock` structure included in the `pNext` chain of `VkWriteDescriptorSet`.

The `VkWriteDescriptorSetInlineUniformBlock` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkWriteDescriptorSetInlineUniformBlock {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           dataSize;
    const void*        pData;
} VkWriteDescriptorSetInlineUniformBlock;
```

or the equivalent

```c++
// Provided by VK_EXT_inline_uniform_block
typedef VkWriteDescriptorSetInlineUniformBlock VkWriteDescriptorSetInlineUniformBlockEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dataSize` is the number of bytes of inline uniform block data pointed to by `pData`.
- `pData` is a pointer to `dataSize` number of bytes of data to write to the inline uniform block.

## [](#_description)Description

Valid Usage

- [](#VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-02222)VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-02222  
  `dataSize` **must** be an integer multiple of `4`

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSetInlineUniformBlock-sType-sType)VUID-VkWriteDescriptorSetInlineUniformBlock-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK`
- [](#VUID-VkWriteDescriptorSetInlineUniformBlock-pData-parameter)VUID-VkWriteDescriptorSetInlineUniformBlock-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-arraylength)VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_inline\_uniform\_block](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_inline_uniform_block.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSetInlineUniformBlock)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0