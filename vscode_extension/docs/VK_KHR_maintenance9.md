# VK\_KHR\_maintenance9(3) Manual Page

## Name

VK\_KHR\_maintenance9 - device extension



## [](#_registered_extension_number)Registered Extension Number

585

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance9%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance9%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_maintenance9](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance9.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-05-29

**Interactions and External Dependencies**

**Contributors**

- Mike Blumenkrantz, Valve
- Shahbaz Youssefi, Google
- Hans-Kristian Arntzen, Valve
- Piers Daniell, NVIDIA
- Daniel Story, Nintendo
- Jeff Bolz, NVIDIA

## [](#_description)Description

[VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html) adds a collection of minor features, none of which would warrant an entire extension of their own.

The new features are as follows:

- Support VkDevice with no queues. These can be used as effectively an offline compiler to prepopulate pipeline caches, without expensive queue creation or internal memory allocations.
- Allow `vkCmdSetEvent2` to not provide a dependency, providing `vkCmdSetEvent`-style usage using enums from `VK_KHR_synchronization2`
- Add a [VkQueryPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlagBits.html)::`VK_QUERY_POOL_CREATE_RESET_BIT_KHR` flag to create a query pool with all queries initialized to the reset state.
- Allow any integer bit width for specific bit-wise operations.
- Add a property to enable sparse support with `VK_EXT_image_2d_view_of_3d`.
- Add a property to indicate the implementation will return (0,0,0,0) or (0,0,0,1) to vertex shaders that read unassigned attributes.
- The effects of image memory barriers and image layout transitions on 3D images created with VK\_IMAGE\_CREATE\_2D\_ARRAY\_COMPATIBLE\_BIT are scoped to the slices specified by the user-provided VkImageSubresourceRange.
- Queue family ownership transfers are no longer required for buffers and linear images, and a new physical device queue family property is exposed to indicate whether queue family ownership transfers are required for optimal images.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance9FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance9FeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMaintenance9PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance9PropertiesKHR.html)
- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html):
  
  - [VkQueueFamilyOwnershipTransferPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyOwnershipTransferPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkDefaultVertexAttributeValueKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDefaultVertexAttributeValueKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_9_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_9_SPEC_VERSION`
- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html):
  
  - `VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR`
- Extending [VkQueryPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlagBits.html):
  
  - `VK_QUERY_POOL_CREATE_RESET_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_9_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_9_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_OWNERSHIP_TRANSFER_PROPERTIES_KHR`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2025-05-29 (Contributors)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance9).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0