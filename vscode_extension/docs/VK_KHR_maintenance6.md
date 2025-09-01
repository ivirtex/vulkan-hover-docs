# VK\_KHR\_maintenance6(3) Manual Page

## Name

VK\_KHR\_maintenance6 - device extension



## [](#_registered_extension_number)Registered Extension Number

546

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_descriptor\_buffer
- Interacts with VK\_KHR\_push\_descriptor

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Jon Leech [\[GitHub\]oddhack](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance6%5D%20%40oddhack%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance6%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_maintenance6](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance6.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-08-03

**Interactions and External Dependencies**

- Interacts with `VK_EXT_robustness2`

**Contributors**

- Jon Leech, Khronos
- Stu Smith, AMD
- Mike Blumenkrantz, Valve
- Ralph Potter, Samsung
- James Fitzpatrick, Imagination Technologies
- Piers Daniell, NVIDIA
- Daniel Story, Nintendo

## [](#_description)Description

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html) adds a collection of minor features, none of which would warrant an entire extension of their own.

The new features are as follows:

- [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatusKHR.html) may be included in the `pNext` chain of [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html) and [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html), allowing applications to identify individual resources for which memory binding failed during calls to [vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory2.html) and [vkBindImageMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindImageMemory2.html).
- A new property `fragmentShadingRateClampCombinerInputs` to indicate if an implementation clamps the inputs to fragment shading rate combiner operations.
- [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) is allowed to be used when binding an index buffer, instead of a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle. When the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, every index fetched results in a value of zero.
- A new property `maxCombinedImageSamplerDescriptorCount` to indicate the maximum number of descriptors needed for any of the [formats that require a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) supported by the implementation.
- A new property `blockTexelViewCompatibleMultipleLayers` indicating whether `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` is allowed to be used with `layerCount` &gt; 1
- `pNext` extensible \*2 versions of all descriptor binding commands.

## [](#_new_commands)New Commands

- [vkCmdBindDescriptorSets2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorSets2KHR.html)
- [vkCmdPushConstants2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushConstants2KHR.html)

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- [vkCmdBindDescriptorBufferEmbeddedSamplers2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplers2EXT.html)
- [vkCmdSetDescriptorBufferOffsets2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsets2EXT.html)

If [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) is supported:

- [vkCmdPushDescriptorSet2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet2KHR.html)
- [vkCmdPushDescriptorSetWithTemplate2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplate2KHR.html)

## [](#_new_structures)New Structures

- [VkBindDescriptorSetsInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfoKHR.html)
- [VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfoKHR.html)
- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html), [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html):
  
  - [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatusKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance6FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6FeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html)

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html)
- [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html)

If [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) is supported:

- [VkPushDescriptorSetInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfoKHR.html)
- [VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_6_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_6_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_SETS_INFO_KHR`
  - `VK_STRUCTURE_TYPE_BIND_MEMORY_STATUS_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PUSH_CONSTANTS_INFO_KHR`

If [VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_BUFFER_EMBEDDED_SAMPLERS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SET_DESCRIPTOR_BUFFER_OFFSETS_INFO_EXT`

If [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_WITH_TEMPLATE_INFO_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2023-08-01 (Jon Leech)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance6).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0