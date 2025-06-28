# VK\_EXT\_inline\_uniform\_block(3) Manual Page

## Name

VK\_EXT\_inline\_uniform\_block - device extension



## [](#_registered_extension_number)Registered Extension Number

139

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_2
- Interacts with VK\_EXT\_descriptor\_indexing
- Interacts with VkPhysicalDeviceVulkan12Features::descriptorIndexing

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]aqnuep](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_inline_uniform_block%5D%20%40aqnuep%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_inline_uniform_block%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-08-01

**IP Status**

No known IP claims.

**Contributors**

- Daniel Rakos, AMD
- Jeff Bolz, NVIDIA
- Slawomir Grajewski, Intel
- Neil Henning, Codeplay

## [](#_description)Description

This extension introduces the ability to back uniform blocks directly with descriptor sets by storing inline uniform data within descriptor pool storage. Compared to push constants this new construct allows uniform data to be reused across multiple disjoint sets of drawing or dispatching commands and **may** enable uniform data to be accessed with fewer indirections compared to uniforms backed by buffer memory.

## [](#_new_structures)New Structures

- Extending [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html):
  
  - [VkDescriptorPoolInlineUniformBlockCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolInlineUniformBlockCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceInlineUniformBlockFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceInlineUniformBlockPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockPropertiesEXT.html)
- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html):
  
  - [VkWriteDescriptorSetInlineUniformBlockEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlockEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_INLINE_UNIFORM_BLOCK_EXTENSION_NAME`
- `VK_EXT_INLINE_UNIFORM_BLOCK_SPEC_VERSION`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

Vulkan 1.3 adds [additional functionality related to this extension](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-new-features) in the form of the [`maxInlineUniformTotalSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxInlineUniformTotalSize) limit.

## [](#_issues)Issues

1\) Do we need a new storage class for inline uniform blocks vs. uniform blocks?

**RESOLVED**: No. The `Uniform` storage class is used to allow the same syntax used for both uniform buffers and inline uniform blocks.

2\) Is the descriptor array index and array size expressed in terms of bytes or dwords for inline uniform block descriptors?

**RESOLVED**: In bytes, but both **must** be a multiple of 4, similar to how push constant ranges are specified. The `descriptorCount` of `VkDescriptorSetLayoutBinding` thus provides the total number of bytes a particular binding with an inline uniform block descriptor type can hold, while the `srcArrayElement`, `dstArrayElement`, and `descriptorCount` members of `VkWriteDescriptorSet`, `VkCopyDescriptorSet`, and `VkDescriptorUpdateTemplateEntry` (where applicable) specify the byte offset and number of bytes to write/copy to the binding’s backing store. Additionally, the `stride` member of `VkDescriptorUpdateTemplateEntry` is ignored for inline uniform blocks and a default value of one is used, meaning that the data to update inline uniform block bindings with must be contiguous in memory.

3\) What layout rules apply for uniform blocks corresponding to inline constants?

**RESOLVED**: They use the same layout rules as uniform buffers.

4\) Do we need to add non-uniform indexing features/properties as introduced by `VK_EXT_descriptor_indexing` for inline uniform blocks?

**RESOLVED**: No, because inline uniform blocks are not allowed to be “arrayed”. A single binding with an inline uniform block descriptor type corresponds to a single uniform block instance and the array indices inside that binding refer to individual offsets within the uniform block (see issue #2). However, this extension does introduce new features/properties about the level of support for update-after-bind inline uniform blocks.

5\) Is the [`descriptorBindingVariableDescriptorCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBindingVariableDescriptorCount) feature introduced by `VK_EXT_descriptor_indexing` supported for inline uniform blocks?

**RESOLVED**: Yes, as long as other inline uniform block specific limits are respected.

6\) Do the robustness guarantees of `robustBufferAccess` apply to inline uniform block accesses?

**RESOLVED**: No, similarly to push constants, as they are not backed by buffer memory like uniform buffers.

## [](#_version_history)Version History

- Revision 1, 2018-08-01 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_inline_uniform_block)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0