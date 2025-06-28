# VK\_EXT\_descriptor\_buffer(3) Manual Page

## Name

VK\_EXT\_descriptor\_buffer - device extension



## [](#_registered_extension_number)Registered Extension Number

317

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

                 [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
                 or  
                 [Vulkan Version 1.1](#versions-1.1)  
             and  
             [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html)  
             and  
             [VK\_EXT\_descriptor\_indexing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_indexing.html)  
         or  
         [Vulkan Version 1.2](#versions-1.2)  
     and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_acceleration\_structure
- Interacts with VK\_NV\_ray\_tracing

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_descriptor_buffer%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_descriptor_buffer%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_descriptor\_buffer](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_descriptor_buffer.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-07

**IP Status**

No known IP claims.

**Contributors**

- Tobias Hector, AMD
- Stu Smith, AMD
- Maciej Jesionowski, AMD
- Boris Zanin, AMD
- Hans-Kristian Arntzen, Valve
- Connor Abbott, Valve
- Baldur Karlsson, Valve
- Mike Blumenkrantz, Valve
- Graeme Leese, Broadcom
- Jan-Harald Fredriksen, Arm
- Rodrigo Locatti, NVIDIA
- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA
- Jeff Leger, QUALCOMM
- Lionel Landwerlin, Intel
- Slawomir Grajewski, Intel

## [](#_description)Description

This extension introduces new commands to put shader-accessible descriptors directly in memory, making the management of descriptor data more explicit.

## [](#_new_commands)New Commands

- [vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html)
- [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindDescriptorBuffersEXT.html)
- [vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html)
- [vkGetBufferOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureDescriptorDataEXT.html)
- [vkGetDescriptorEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorEXT.html)
- [vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html)
- [vkGetDescriptorSetLayoutSizeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSizeEXT.html)
- [vkGetImageOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageOpaqueCaptureDescriptorDataEXT.html)
- [vkGetImageViewOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewOpaqueCaptureDescriptorDataEXT.html)
- [vkGetSamplerOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSamplerOpaqueCaptureDescriptorDataEXT.html)

If [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html) or [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) is supported:

- [vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureOpaqueCaptureDescriptorDataEXT.html)

## [](#_new_structures)New Structures

- [VkBufferCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCaptureDescriptorDataInfoEXT.html)
- [VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorAddressInfoEXT.html)
- [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html)
- [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html)
- [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCaptureDescriptorDataInfoEXT.html)
- [VkImageViewCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCaptureDescriptorDataInfoEXT.html)
- [VkSamplerCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCaptureDescriptorDataInfoEXT.html)
- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html), [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html), [VkAccelerationStructureCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoNV.html), [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html), [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html):
  
  - [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
- Extending [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingInfoEXT.html):
  
  - [VkDescriptorBufferBindingPushDescriptorBufferHandleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferBindingPushDescriptorBufferHandleEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDescriptorBufferFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT.html)
  - [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)

If [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html) or [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) is supported:

- [VkAccelerationStructureCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCaptureDescriptorDataInfoEXT.html)

## [](#_new_unions)New Unions

- [VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorDataEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DESCRIPTOR_BUFFER_EXTENSION_NAME`
- `VK_EXT_DESCRIPTOR_BUFFER_SPEC_VERSION`
- Extending [VkAccelerationStructureCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateFlagBitsKHR.html):
  
  - `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT`
- Extending [VkBufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateFlagBits.html):
  
  - `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html):
  
  - `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkImageViewCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateFlagBits.html):
  
  - `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html):
  
  - `VK_SAMPLER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_ADDRESS_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_INFO_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_PUSH_DESCRIPTOR_BUFFER_HANDLE_EXT`
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_GET_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMAGE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`
  - `VK_STRUCTURE_TYPE_OPAQUE_CAPTURE_DESCRIPTOR_DATA_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_DENSITY_MAP_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_SAMPLER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

If [VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html) or [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-06-07 (Stu Smith)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_descriptor_buffer)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0