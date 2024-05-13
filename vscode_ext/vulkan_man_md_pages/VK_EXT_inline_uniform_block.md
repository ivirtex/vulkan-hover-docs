# VK_EXT_inline_uniform_block(3) Manual Page

## Name

VK_EXT_inline_uniform_block - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

139

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_inline_uniform_block%5D%20@aqnuep%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_inline_uniform_block%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>aqnuep</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-08-01

**IP Status**  
No known IP claims.

**Contributors**  
- Daniel Rakos, AMD

- Jeff Bolz, NVIDIA

- Slawomir Grajewski, Intel

- Neil Henning, Codeplay

## <a href="#_description" class="anchor"></a>Description

This extension introduces the ability to back uniform blocks directly
with descriptor sets by storing inline uniform data within descriptor
pool storage. Compared to push constants this new construct allows
uniform data to be reused across multiple disjoint sets of drawing or
dispatching commands and **may** enable uniform data to be accessed with
fewer indirections compared to uniforms backed by buffer memory.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html):

  - [VkDescriptorPoolInlineUniformBlockCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolInlineUniformBlockCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceInlineUniformBlockFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceInlineUniformBlockPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockPropertiesEXT.html)

- Extending [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html):

  - [VkWriteDescriptorSetInlineUniformBlockEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlockEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_INLINE_UNIFORM_BLOCK_EXTENSION_NAME`

- `VK_EXT_INLINE_UNIFORM_BLOCK_SPEC_VERSION`

- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html):

  - `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_INLINE_UNIFORM_BLOCK_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INLINE_UNIFORM_BLOCK_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET_INLINE_UNIFORM_BLOCK_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

Vulkan 1.3 adds <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-new-features"
target="_blank" rel="noopener">additional functionality related to this
extension</a> in the form of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxInlineUniformTotalSize"
target="_blank"
rel="noopener"><code>maxInlineUniformTotalSize</code></a> limit.

## <a href="#_issues" class="anchor"></a>Issues

1\) Do we need a new storage class for inline uniform blocks vs. uniform
blocks?

**RESOLVED**: No. The `Uniform` storage class is used to allow the same
syntax used for both uniform buffers and inline uniform blocks.

2\) Is the descriptor array index and array size expressed in terms of
bytes or dwords for inline uniform block descriptors?

**RESOLVED**: In bytes, but both **must** be a multiple of 4, similar to
how push constant ranges are specified. The `descriptorCount` of
`VkDescriptorSetLayoutBinding` thus provides the total number of bytes a
particular binding with an inline uniform block descriptor type can
hold, while the `srcArrayElement`, `dstArrayElement`, and
`descriptorCount` members of `VkWriteDescriptorSet`,
`VkCopyDescriptorSet`, and `VkDescriptorUpdateTemplateEntry` (where
applicable) specify the byte offset and number of bytes to write/copy to
the binding’s backing store. Additionally, the `stride` member of
`VkDescriptorUpdateTemplateEntry` is ignored for inline uniform blocks
and a default value of one is used, meaning that the data to update
inline uniform block bindings with must be contiguous in memory.

3\) What layout rules apply for uniform blocks corresponding to inline
constants?

**RESOLVED**: They use the same layout rules as uniform buffers.

4\) Do we need to add non-uniform indexing features/properties as
introduced by `VK_EXT_descriptor_indexing` for inline uniform blocks?

**RESOLVED**: No, because inline uniform blocks are not allowed to be
“arrayed”. A single binding with an inline uniform block descriptor type
corresponds to a single uniform block instance and the array indices
inside that binding refer to individual offsets within the uniform block
(see issue \#2). However, this extension does introduce new
features/properties about the level of support for update-after-bind
inline uniform blocks.

5\) Is the `descriptorBindingVariableDescriptorCount` feature introduced
by `VK_EXT_descriptor_indexing` supported for inline uniform blocks?

**RESOLVED**: Yes, as long as other inline uniform block specific limits
are respected.

6\) Do the robustness guarantees of `robustBufferAccess` apply to inline
uniform block accesses?

**RESOLVED**: No, similarly to push constants, as they are not backed by
buffer memory like uniform buffers.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-08-01 (Daniel Rakos)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_inline_uniform_block"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
