# VK\_EXT\_descriptor\_indexing(3) Manual Page

## Name

VK\_EXT\_descriptor\_indexing - device extension



## [](#_registered_extension_number)Registered Extension Number

162

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_maintenance3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance3.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_descriptor\_indexing](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_descriptor_indexing.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_descriptor_indexing%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_descriptor_indexing%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-10-02

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_nonuniform_qualifier`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_nonuniform_qualifier.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Daniel Rakos, AMD
- Slawomir Grajewski, Intel
- Tobias Hector, Imagination Technologies

## [](#_description)Description

This extension adds several small features which together enable applications to create large descriptor sets containing substantially all of their resources, and selecting amongst those resources with dynamic (non-uniform) indexes in the shader. There are feature enables and SPIR-V capabilities for non-uniform descriptor indexing in the shader, and non-uniform indexing in the shader requires use of a new `NonUniformEXT` decoration defined in the `SPV_EXT_descriptor_indexing` SPIR-V extension. There are descriptor set layout binding creation flags enabling several features:

- Descriptors can be updated after they are bound to a command buffer, such that the execution of the command buffer reflects the most recent update to the descriptors.
- Descriptors that are not used by any pending command buffers can be updated, which enables writing new descriptors for frame N+1 while frame N is executing.
- Relax the requirement that all descriptors in a binding that is “statically used” must be valid, such that descriptors that are not accessed by a submission need not be valid and can be updated while that submission is executing.
- The final binding in a descriptor set layout can have a variable size (and unsized arrays of resources are allowed in the `GL_EXT_nonuniform_qualifier` and `SPV_EXT_descriptor_indexing` extensions).

Note that it is valid for multiple descriptor arrays in a shader to use the same set and binding number, as long as they are all compatible with the descriptor type in the pipeline layout. This means a single array binding in the descriptor set can serve multiple texture dimensionalities, or an array of buffer descriptors can be used with multiple different block layouts.

There are new descriptor set layout and descriptor pool creation flags that are required to opt in to the update-after-bind functionality, and there are separate `maxPerStage`* and `maxDescriptorSet`* limits that apply to these descriptor set layouts which **may** be much higher than the pre-existing limits. The old limits only count descriptors in non-updateAfterBind descriptor set layouts, and the new limits count descriptors in all descriptor set layouts in the pipeline layout.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

Vulkan APIs in this extension are included in core Vulkan 1.2, with the EXT suffix omitted. However, if Vulkan 1.2 is supported and this extension is not, the `descriptorIndexing` capability is optional. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

If Vulkan 1.4 is supported, support for the `shaderUniformTexelBufferArrayDynamicIndexing` and `shaderStorageTexelBufferArrayDynamicIndexing` capabilities is required.

## [](#_new_structures)New Structures

- Extending [VkDescriptorSetAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetAllocateInfo.html):
  
  - [VkDescriptorSetVariableDescriptorCountAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfoEXT.html)
- Extending [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html):
  
  - [VkDescriptorSetLayoutBindingFlagsCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfoEXT.html)
- Extending [VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutSupport.html):
  
  - [VkDescriptorSetVariableDescriptorCountLayoutSupportEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupportEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDescriptorIndexingFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDescriptorIndexingPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingPropertiesEXT.html)

## [](#_new_enums)New Enums

- [VkDescriptorBindingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDescriptorBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME`
- `VK_EXT_DESCRIPTOR_INDEXING_SPEC_VERSION`
- Extending [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html):
  
  - `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT`
  - `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT`
  - `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT`
  - `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT`
- Extending [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_FRAGMENTATION_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT`

## [](#_version_history)Version History

- Revision 1, 2017-07-26 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2017-10-02 (Jeff Bolz)
  
  - ???

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_descriptor_indexing).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0