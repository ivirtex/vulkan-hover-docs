# VK_NV_device_generated_commands_compute(3) Manual Page

## Name

VK_NV_device_generated_commands_compute - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

429

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_generated_commands_compute%5D%20@vkushwaha-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_generated_commands_compute%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows the device to generate commands for binding
compute pipelines, setting push constants and launching compute
dispatches.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)

- [vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectDeviceAddressNV.html)

- [vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBindPipelineIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindPipelineIndirectCommandNV.html)

- [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)

- Extending
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html):

  - [VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_DEVICE_GENERATED_COMMANDS_COMPUTE_EXTENSION_NAME`

- `VK_NV_DEVICE_GENERATED_COMMANDS_COMPUTE_SPEC_VERSION`

- Extending
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html):

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_INDIRECT_BINDABLE_BIT_NV`

- Extending
  [VkIndirectCommandsTokenTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsTokenTypeNV.html):

  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_DISPATCH_NV`

  - `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COMPUTE_PIPELINE_INDIRECT_BUFFER_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_COMPUTE_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PIPELINE_INDIRECT_DEVICE_ADDRESS_INFO_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-07-21 (Vikram Kushwaha)

  - Rename vkCmdUpdatePipelineIndirectBuffer to
    vkCmdUpdatePipelineIndirectBufferNV

- Revision 1, 2023-06-09 (Vikram Kushwaha)

  - First Revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_device_generated_commands_compute"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
