# VK\_NV\_device\_generated\_commands\_compute(3) Manual Page

## Name

VK\_NV\_device\_generated\_commands\_compute - device extension



## [](#_registered_extension_number)Registered Extension Number

429

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_generated_commands_compute%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_generated_commands_compute%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-07-21

**Contributors**

- Vikram Kushwaha, NVIDIA
- Jeff Bolz, NVIDIA
- Christoph Kubisch, NVIDIA
- Piers Daniell, NVIDIA
- Daniel Koch, NVIDIA
- Hans-Kristian Arntzen, Valve
- Mike Blumenkrantz, VALVE

## [](#_description)Description

This extension allows the device to generate commands for binding compute pipelines, setting push constants and launching compute dispatches.

## [](#_new_commands)New Commands

- [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)
- [vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectDeviceAddressNV.html)
- [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

## [](#_new_structures)New Structures

- [VkBindPipelineIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindPipelineIndirectCommandNV.html)
- [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)
- Extending [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html):
  
  - [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineIndirectBufferInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_DEVICE_GENERATED_COMMANDS_COMPUTE_EXTENSION_NAME`
- `VK_NV_DEVICE_GENERATED_COMMANDS_COMPUTE_SPEC_VERSION`
- Extending [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html):
  
  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_INDIRECT_BINDABLE_BIT_NV`
- Extending [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenTypeNV.html):
  
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV`
  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_INDIRECT_BUFFER_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_COMPUTE_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PIPELINE_INDIRECT_DEVICE_ADDRESS_INFO_NV`

## [](#_version_history)Version History

- Revision 2, 2023-07-21 (Vikram Kushwaha)
  
  - Rename vkCmdUpdatePipelineIndirectBuffer to vkCmdUpdatePipelineIndirectBufferNV
- Revision 1, 2023-06-09 (Vikram Kushwaha)
  
  - First Revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_device_generated_commands_compute).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0