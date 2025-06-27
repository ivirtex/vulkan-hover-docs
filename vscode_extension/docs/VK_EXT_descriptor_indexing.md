# VK_EXT_descriptor_indexing(3) Manual Page

## Name

VK_EXT_descriptor_indexing - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

162

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK_KHR_maintenance3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance3.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_descriptor_indexing](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_descriptor_indexing.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_descriptor_indexing%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_descriptor_indexing%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-10-02

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_nonuniform_qualifier`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_nonuniform_qualifier.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Daniel Rakos, AMD

- Slawomir Grajewski, Intel

- Tobias Hector, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension adds several small features which together enable
applications to create large descriptor sets containing substantially
all of their resources, and selecting amongst those resources with
dynamic (non-uniform) indexes in the shader. There are feature enables
and SPIR-V capabilities for non-uniform descriptor indexing in the
shader, and non-uniform indexing in the shader requires use of a new
`NonUniformEXT` decoration defined in the `SPV_EXT_descriptor_indexing`
SPIR-V extension. There are descriptor set layout binding creation flags
enabling several features:

- Descriptors can be updated after they are bound to a command buffer,
  such that the execution of the command buffer reflects the most recent
  update to the descriptors.

- Descriptors that are not used by any pending command buffers can be
  updated, which enables writing new descriptors for frame N+1 while
  frame N is executing.

- Relax the requirement that all descriptors in a binding that is
  “statically used” must be valid, such that descriptors that are not
  accessed by a submission need not be valid and can be updated while
  that submission is executing.

- The final binding in a descriptor set layout can have a variable size
  (and unsized arrays of resources are allowed in the
  `GL_EXT_nonuniform_qualifier` and `SPV_EXT_descriptor_indexing`
  extensions).

Note that it is valid for multiple descriptor arrays in a shader to use
the same set and binding number, as long as they are all compatible with
the descriptor type in the pipeline layout. This means a single array
binding in the descriptor set can serve multiple texture
dimensionalities, or an array of buffer descriptors can be used with
multiple different block layouts.

There are new descriptor set layout and descriptor pool creation flags
that are required to opt in to the update-after-bind functionality, and
there are separate `maxPerStage`\* and `maxDescriptorSet`\* limits that
apply to these descriptor set layouts which **may** be much higher than
the pre-existing limits. The old limits only count descriptors in
non-updateAfterBind descriptor set layouts, and the new limits count
descriptors in all descriptor set layouts in the pipeline layout.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetAllocateInfo.html):

  - [VkDescriptorSetVariableDescriptorCountAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfoEXT.html)

- Extending
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html):

  - [VkDescriptorSetLayoutBindingFlagsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfoEXT.html)

- Extending
  [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html):

  - [VkDescriptorSetVariableDescriptorCountLayoutSupportEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupportEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDescriptorIndexingFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDescriptorIndexingPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingPropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDescriptorBindingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDescriptorBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME`

- `VK_EXT_DESCRIPTOR_INDEXING_SPEC_VERSION`

- Extending
  [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html):

  - `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT`

  - `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT`

  - `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT`

  - `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT`

- Extending
  [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html):

  - `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT`

- Extending
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html):

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_FRAGMENTATION_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT`

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

Functionality in this extension is included in core Vulkan 1.2, with the
EXT suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, the `descriptorIndexing` capability is optional. The
original type, enum and command names are still available as aliases of
the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-07-26 (Jeff Bolz)

  - Internal revisions

- Revision 2, 2017-10-02 (Jeff Bolz)

  - ???

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_descriptor_indexing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
