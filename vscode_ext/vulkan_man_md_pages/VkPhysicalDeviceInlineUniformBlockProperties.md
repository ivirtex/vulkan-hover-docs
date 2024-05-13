# VkPhysicalDeviceInlineUniformBlockProperties(3) Manual Page

## Name

VkPhysicalDeviceInlineUniformBlockProperties - Structure describing
inline uniform block properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceInlineUniformBlockProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceInlineUniformBlockProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxInlineUniformBlockSize;
    uint32_t           maxPerStageDescriptorInlineUniformBlocks;
    uint32_t           maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks;
    uint32_t           maxDescriptorSetInlineUniformBlocks;
    uint32_t           maxDescriptorSetUpdateAfterBindInlineUniformBlocks;
} VkPhysicalDeviceInlineUniformBlockProperties;
```

or the equivalent

``` c
// Provided by VK_EXT_inline_uniform_block
typedef VkPhysicalDeviceInlineUniformBlockProperties VkPhysicalDeviceInlineUniformBlockPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-maxInlineUniformBlockSize"></span>
  `maxInlineUniformBlockSize` is the maximum size in bytes of an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inlineuniformblock"
  target="_blank" rel="noopener">inline uniform block</a> binding.

- <span id="extension-limits-maxPerStageDescriptorInlineUniformBlocks"></span>
  `maxPerStageDescriptorInlineUniformBlocks` is the maximum number of
  inline uniform block bindings that **can** be accessible to a single
  shader stage in a pipeline layout. Descriptor bindings with a
  descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count
  against this limit. Only descriptor bindings in descriptor set layouts
  created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks"></span>
  `maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks` is similar
  to `maxPerStageDescriptorInlineUniformBlocks` but counts descriptor
  bindings from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetInlineUniformBlocks"></span>
  `maxDescriptorSetInlineUniformBlocks` is the maximum number of inline
  uniform block bindings that **can** be included in descriptor bindings
  in a pipeline layout across all pipeline shader stages and descriptor
  set numbers. Descriptor bindings with a descriptor type of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit.
  Only descriptor bindings in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindInlineUniformBlocks"></span>
  `maxDescriptorSetUpdateAfterBindInlineUniformBlocks` is similar to
  `maxDescriptorSetInlineUniformBlocks` but counts descriptor bindings
  from descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

If the `VkPhysicalDeviceInlineUniformBlockProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceInlineUniformBlockProperties-sType-sType"
  id="VUID-VkPhysicalDeviceInlineUniformBlockProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceInlineUniformBlockProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_inline_uniform_block](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_inline_uniform_block.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceInlineUniformBlockProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
