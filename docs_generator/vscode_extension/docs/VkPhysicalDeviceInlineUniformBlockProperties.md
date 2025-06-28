# VkPhysicalDeviceInlineUniformBlockProperties(3) Manual Page

## Name

VkPhysicalDeviceInlineUniformBlockProperties - Structure describing inline uniform block properties that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceInlineUniformBlockProperties` structure is defined as:

```c++
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

```c++
// Provided by VK_EXT_inline_uniform_block
typedef VkPhysicalDeviceInlineUniformBlockProperties VkPhysicalDeviceInlineUniformBlockPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`maxInlineUniformBlockSize` is the maximum size in bytes of an [inline uniform block](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-inlineuniformblock) binding.
- []()`maxPerStageDescriptorInlineUniformBlocks` is the maximum number of inline uniform block bindings that **can** be accessible to a single shader stage in a pipeline layout. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks` is similar to `maxPerStageDescriptorInlineUniformBlocks` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetInlineUniformBlocks` is the maximum number of inline uniform block bindings that **can** be included in descriptor bindings in a pipeline layout across all pipeline shader stages and descriptor set numbers. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxDescriptorSetUpdateAfterBindInlineUniformBlocks` is similar to `maxDescriptorSetInlineUniformBlocks` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

If the `VkPhysicalDeviceInlineUniformBlockProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceInlineUniformBlockProperties-sType-sType)VUID-VkPhysicalDeviceInlineUniformBlockProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_inline\_uniform\_block](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_inline_uniform_block.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceInlineUniformBlockProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0