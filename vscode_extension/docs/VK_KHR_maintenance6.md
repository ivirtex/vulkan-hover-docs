# VK_KHR_maintenance6(3) Manual Page

## Name

VK_KHR_maintenance6 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

546

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_descriptor_buffer

- Interacts with VK_KHR_push_descriptor

## <a href="#_contact" class="anchor"></a>Contact

- Jon Leech <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance6%5D%20@oddhack%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance6%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>oddhack</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_maintenance6](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance6.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-08-03

**Interactions and External Dependencies**  
- Interacts with [`VK_EXT_robustness2`](VK_EXT_robustness2.html)

**Contributors**  
- Jon Leech, Khronos

- Stu Smith, AMD

- Mike Blumenkrantz, Valve

- Ralph Potter, Samsung

- James Fitzpatrick, Imagination Technologies

- Piers Daniell, NVIDIA

- Daniel Story, Nintendo

## <a href="#_description" class="anchor"></a>Description

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html) adds a collection of
minor features, none of which would warrant an entire extension of their
own.

The new features are as follows:

- [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindMemoryStatusKHR.html) may be included in
  the `pNext` chain of
  [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) and
  [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html), allowing
  applications to identify individual resources for which memory binding
  failed during calls to [vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html)
  and [vkBindImageMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html).

- A new property `fragmentShadingRateClampCombinerInputs` to indicate if
  an implementation clamps the inputs to fragment shading rate combiner
  operations.

- [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) is allowed to be used when
  binding an index buffer, instead of a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)
  handle. When the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is enabled, every index fetched results in a value of zero.

- A new property `maxCombinedImageSamplerDescriptorCount` to indicate
  the maximum number of descriptors needed for any of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">formats that require a sampler
  Y′C<sub>B</sub>C<sub>R</sub> conversion</a> supported by the
  implementation.

- A new property `blockTexelViewCompatibleMultipleLayers` indicating
  whether `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` is allowed
  to be used with `layerCount` \> 1

- `pNext` extensible \*2 versions of all descriptor binding commands.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBindDescriptorSets2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets2KHR.html)

- [vkCmdPushConstants2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushConstants2KHR.html)

If [VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html) is
supported:

- [vkCmdBindDescriptorBufferEmbeddedSamplers2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplers2EXT.html)

- [vkCmdSetDescriptorBufferOffsets2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsets2EXT.html)

If [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html) is supported:

- [vkCmdPushDescriptorSet2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSet2KHR.html)

- [vkCmdPushDescriptorSetWithTemplate2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplate2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBindDescriptorSetsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorSetsInfoKHR.html)

- [VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantsInfoKHR.html)

- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html),
  [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html):

  - [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindMemoryStatusKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceMaintenance6FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6FeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html)

If [VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html) is
supported:

- [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html)

- [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html)

If [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html) is supported:

- [VkPushDescriptorSetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetInfoKHR.html)

- [VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE_6_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_6_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_SETS_INFO_KHR`

  - `VK_STRUCTURE_TYPE_BIND_MEMORY_STATUS_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PUSH_CONSTANTS_INFO_KHR`

If [VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html) is
supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_BUFFER_EMBEDDED_SAMPLERS_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SET_DESCRIPTOR_BUFFER_OFFSETS_INFO_EXT`

If [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_WITH_TEMPLATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-08-01 (Jon Leech)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance6"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
