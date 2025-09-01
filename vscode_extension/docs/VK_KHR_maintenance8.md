# VK\_KHR\_maintenance8(3) Manual Page

## Name

VK\_KHR\_maintenance8 - device extension



## [](#_registered_extension_number)Registered Extension Number

575

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance8%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance8%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_maintenance8](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance8.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-07

**Interactions and External Dependencies**

**Contributors**

- Jon Leech, Khronos
- Mike Blumenkrantz, Valve
- Spencer Fricke, LunarG
- Jan-Harald Fredriksen, ARM
- Piers Daniell, NVIDIA
- Matthew Netsch, Qualcomm
- Ricardo Garcia, Igalia
- Lionel Landwerlin, Intel
- Rick Hammerstone, Qualcomm
- Daniel Story, Nintendo
- Hans-Kristian Arntzen, Valve
- Caterina Shablia, Collabora
- Georg Lehmann, Valve
- Shahbaz Youssefi, Google
- Tobias Hector, AMD

## [](#_description)Description

[VK\_KHR\_maintenance8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance8.html) adds a collection of minor features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Allow copies between depth/stencil and “matching” color attachments
- Allow `dstCache` in `vkMergePipelineCaches` to be implicitly synchronized.
- Require src/dst sync scopes to work when doing queue family ownership transfers
- Support `Offset` (as an alternative to `ConstOffset`) image operand in texture sampling and fetch operations
- Use the SPIR-V definition of `OpSRem` and `OpSMod`, making these operations produce well-defined results for negative operands
- Loosen layer restrictions when blitting from 3D images to other image types
- Add space for an additional 64 access flags for use with VkMemoryBarrier2, VkBufferMemoryBarrier2, and VkImageMemoryBarrier2

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance8FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance8FeaturesKHR.html)
- Extending [VkSubpassDependency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency2.html), [VkBufferMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferMemoryBarrier2.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html):
  
  - [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html)

## [](#_new_enums)New Enums

- [VkAccessFlagBits3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits3KHR.html)
- [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlagBits.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_8_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_8_SPEC_VERSION`
- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html):
  
  - `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`
- Extending [VkPipelineCacheCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCacheCreateFlagBits.html):
  
  - `VK_PIPELINE_CACHE_CREATE_INTERNALLY_SYNCHRONIZED_MERGE_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MEMORY_BARRIER_ACCESS_FLAGS_3_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_8_FEATURES_KHR`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2024-06-20 (Jon Leech)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance8).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0