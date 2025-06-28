# VK\_KHR\_device\_group(3) Manual Page

## Name

VK\_KHR\_device\_group - device extension



## [](#_registered_extension_number)Registered Extension Number

61

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_device\_group\_creation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group_creation.html)

## [](#_api_interactions)API Interactions

- Interacts with VK\_KHR\_bind\_memory2
- Interacts with VK\_KHR\_surface
- Interacts with VK\_KHR\_swapchain

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_device\_group](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_device_group.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_device_group%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_device_group%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-10-10

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Tobias Hector, Imagination Technologies

## [](#_description)Description

This extension provides functionality to use a logical device that consists of multiple physical devices, as created with the `VK_KHR_device_group_creation` extension. A device group can allocate memory across the subdevices, bind memory from one subdevice to a resource on another subdevice, record command buffers where some work executes on an arbitrary subset of the subdevices, and potentially present a swapchain image from one or more subdevices.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

The following enums, types and commands are included as interactions with `VK_KHR_swapchain`:

- `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR`
- `VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR`
- `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`
- `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`
- `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR`
- `VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR`
- `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`
- [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html)
- [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)
- [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSwapchainCreateInfoKHR.html)
- [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemorySwapchainInfoKHR.html)
- [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html)
- [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentInfoKHR.html)
- [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html)
- [vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html)
- [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html)
- [vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html)
- [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html)

If Vulkan 1.1 and `VK_KHR_swapchain` are supported, these are included by `VK_KHR_swapchain`.

The base functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkCmdDispatchBaseKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchBaseKHR.html)
- [vkCmdSetDeviceMaskKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDeviceMaskKHR.html)
- [vkGetDeviceGroupPeerMemoryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupPeerMemoryFeaturesKHR.html)

If [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html) is supported:

- [vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html)
- [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html)
- [vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html)

If [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html) is supported:

- [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html)

## [](#_new_structures)New Structures

- Extending [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html):
  
  - [VkDeviceGroupBindSparseInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupBindSparseInfoKHR.html)
- Extending [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html):
  
  - [VkDeviceGroupCommandBufferBeginInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupCommandBufferBeginInfoKHR.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkMemoryAllocateFlagsInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagsInfoKHR.html)
- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html):
  
  - [VkDeviceGroupRenderPassBeginInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfoKHR.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html):
  
  - [VkDeviceGroupSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSubmitInfoKHR.html)

If [VK\_KHR\_bind\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_bind_memory2.html) is supported:

- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html):
  
  - [VkBindBufferMemoryDeviceGroupInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryDeviceGroupInfoKHR.html)
- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html):
  
  - [VkBindImageMemoryDeviceGroupInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryDeviceGroupInfoKHR.html)

If [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)

If [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html) is supported:

- [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html)
- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html):
  
  - [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemorySwapchainInfoKHR.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSwapchainCreateInfoKHR.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentInfoKHR.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html)

## [](#_new_enums)New Enums

- [VkMemoryAllocateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBitsKHR.html)
- [VkPeerMemoryFeatureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlagBitsKHR.html)

If [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkMemoryAllocateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagsKHR.html)
- [VkPeerMemoryFeatureFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlagsKHR.html)

If [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DEVICE_GROUP_EXTENSION_NAME`
- `VK_KHR_DEVICE_GROUP_SPEC_VERSION`
- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlagBits.html):
  
  - `VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR`
- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html):
  
  - `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR`
- Extending [VkPeerMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPeerMemoryFeatureFlagBits.html):
  
  - `VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR`
  - `VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR`
  - `VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR`
  - `VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR`
- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits.html):
  
  - `VK_PIPELINE_CREATE_DISPATCH_BASE_BIT_KHR`
  - `VK_PIPELINE_CREATE_DISPATCH_BASE_KHR`
  - `VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR`

If [VK\_KHR\_bind\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_bind_memory2.html) is supported:

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR`
  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR`

If [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR`

If [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`

## [](#_new_built_in_variables)New Built-in Variables

- [`DeviceIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-deviceindex)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`DeviceGroup`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DeviceGroup)

## [](#_version_history)Version History

- Revision 1, 2016-10-19 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2017-05-19 (Tobias Hector)
  
  - Removed extended memory bind functions to VK\_KHR\_bind\_memory2, added dependency on that extension, and device-group-specific structs for those functions.
- Revision 3, 2017-10-06 (Ian Elliott)
  
  - Corrected Vulkan 1.1 interactions with the WSI extensions. All Vulkan 1.1 WSI interactions are with the VK\_KHR\_swapchain extension.
- Revision 4, 2017-10-10 (Jeff Bolz)
  
  - Rename “SFR” bits and structure members to use the phrase “split instance bind regions”.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_device_group)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0