# VkWriteDescriptorSetInlineUniformBlock(3) Manual Page

## Name

VkWriteDescriptorSetInlineUniformBlock - Structure specifying inline
uniform block data



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `descriptorType` member of
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) is
`VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then the data to write to the
descriptor set is specified through a
`VkWriteDescriptorSetInlineUniformBlock` structure included in the
`pNext` chain of `VkWriteDescriptorSet`.

The `VkWriteDescriptorSetInlineUniformBlock` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkWriteDescriptorSetInlineUniformBlock {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           dataSize;
    const void*        pData;
} VkWriteDescriptorSetInlineUniformBlock;
```

or the equivalent

``` c
// Provided by VK_EXT_inline_uniform_block
typedef VkWriteDescriptorSetInlineUniformBlock VkWriteDescriptorSetInlineUniformBlockEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `dataSize` is the number of bytes of inline uniform block data pointed
  to by `pData`.

- `pData` is a pointer to `dataSize` number of bytes of data to write to
  the inline uniform block.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-02222"
  id="VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-02222"></a>
  VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-02222  
  `dataSize` **must** be an integer multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-VkWriteDescriptorSetInlineUniformBlock-sType-sType"
  id="VUID-VkWriteDescriptorSetInlineUniformBlock-sType-sType"></a>
  VUID-VkWriteDescriptorSetInlineUniformBlock-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK`

- <a href="#VUID-VkWriteDescriptorSetInlineUniformBlock-pData-parameter"
  id="VUID-VkWriteDescriptorSetInlineUniformBlock-pData-parameter"></a>
  VUID-VkWriteDescriptorSetInlineUniformBlock-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a
  href="#VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-arraylength"
  id="VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-arraylength"></a>
  VUID-VkWriteDescriptorSetInlineUniformBlock-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_inline_uniform_block](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_inline_uniform_block.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWriteDescriptorSetInlineUniformBlock"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
