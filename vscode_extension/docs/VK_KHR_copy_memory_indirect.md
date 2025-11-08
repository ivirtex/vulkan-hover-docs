# VK\_KHR\_copy\_memory\_indirect(3) Manual Page

## Name

VK\_KHR\_copy\_memory\_indirect - device extension



## [](#_registered_extension_number)Registered Extension Number

550

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     and  
     [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_copy_memory_indirect%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_copy_memory_indirect%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_copy\_memory\_indirect](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_copy_memory_indirect.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-25

**Contributors**

- Daniel Koch, NVIDIA
- Vikram Kushwaha, NVIDIA
- Jeff Bolz, NVIDIA
- Christoph Kubisch, NVIDIA
- Stuart Smith, AMD
- Faith Ekstrand, Collabora
- Caterina Shablia, Collabora
- Spencer Fricke, LunarG
- Matthew Netsch, Qualcomm Technologies, Inc
- Mike Blumenkrantz, Valve
- Alyssa Rosenzweig, Valve

## [](#_description)Description

This extension adds support for performing copies between memory and image regions using indirect parameters that are read by the device from a buffer during execution. This functionality may be useful for performing copies where the copy parameters are not known during the command buffer creation time.

## [](#_new_commands)New Commands

- [vkCmdCopyMemoryIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryIndirectKHR.html)
- [vkCmdCopyMemoryToImageIndirectKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectKHR.html)

## [](#_new_structures)New Structures

- [VkCopyMemoryIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectCommandKHR.html)
- [VkCopyMemoryIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryIndirectInfoKHR.html)
- [VkCopyMemoryToImageIndirectCommandKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectCommandKHR.html)
- [VkCopyMemoryToImageIndirectInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageIndirectInfoKHR.html)
- [VkStridedDeviceAddressRangeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressRangeKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkAddressCopyFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_COPY_MEMORY_INDIRECT_EXTENSION_NAME`
- `VK_KHR_COPY_MEMORY_INDIRECT_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_COPY_IMAGE_INDIRECT_DST_BIT_KHR`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_COPY_INDIRECT_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COPY_MEMORY_INDIRECT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_IMAGE_INDIRECT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_PROPERTIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-01-25 (Daniel Koch, Vikram Kushwaha)
  
  - Initial external release

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_copy_memory_indirect).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0