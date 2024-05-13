# VK_KHR_device_group(3) Manual Page

## Name

VK_KHR_device_group - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

61

## <a href="#_revision" class="anchor"></a>Revision

4

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_device_group_creation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group_creation.html)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_KHR_bind_memory2

- Interacts with VK_KHR_surface

- Interacts with VK_KHR_swapchain

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_device_group](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_device_group.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_device_group%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_device_group%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-10-10

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Tobias Hector, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension provides functionality to use a logical device that
consists of multiple physical devices, as created with the
[`VK_KHR_device_group_creation`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group_creation.html)
extension. A device group can allocate memory across the subdevices,
bind memory from one subdevice to a resource on another subdevice,
record command buffers where some work executes on an arbitrary subset
of the subdevices, and potentially present a swapchain image from one or
more subdevices.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

The following enums, types and commands are included as interactions
with [`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html):

- `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR`

- `VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR`

- `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`

- `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`

- `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR`

- `VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR`

- `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`

- [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html)

- [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)

- [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html)

- [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html)

- [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html)

- [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentInfoKHR.html)

- [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html)

- [vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html)

- [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html)

- [vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html)

- [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html)

If Vulkan 1.1 and [`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) are
supported, these are included by `VK_KHR_swapchain`.

The base functionality in this extension is included in core Vulkan 1.1,
with the KHR suffix omitted. The original type, enum and command names
are still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdDispatchBaseKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchBaseKHR.html)

- [vkCmdSetDeviceMaskKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDeviceMaskKHR.html)

- [vkGetDeviceGroupPeerMemoryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPeerMemoryFeaturesKHR.html)

If [VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) is supported:

- [vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html)

- [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html)

- [vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html)

If [VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) is supported:

- [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindSparseInfo.html):

  - [VkDeviceGroupBindSparseInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupBindSparseInfoKHR.html)

- Extending [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html):

  - [VkDeviceGroupCommandBufferBeginInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfoKHR.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkMemoryAllocateFlagsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfoKHR.html)

- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html),
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html):

  - [VkDeviceGroupRenderPassBeginInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfoKHR.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html):

  - [VkDeviceGroupSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSubmitInfoKHR.html)

If [VK_KHR_bind_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html) is supported:

- Extending [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html):

  - [VkBindBufferMemoryDeviceGroupInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryDeviceGroupInfoKHR.html)

- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html):

  - [VkBindImageMemoryDeviceGroupInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfoKHR.html)

If [VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentCapabilitiesKHR.html)

If [VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) is supported:

- [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html)

- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html):

  - [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html)

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentInfoKHR.html)

- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkDeviceGroupSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupSwapchainCreateInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkMemoryAllocateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBitsKHR.html)

- [VkPeerMemoryFeatureFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBitsKHR.html)

If [VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkMemoryAllocateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsKHR.html)

- [VkPeerMemoryFeatureFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagsKHR.html)

If [VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) is supported:

- [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DEVICE_GROUP_EXTENSION_NAME`

- `VK_KHR_DEVICE_GROUP_SPEC_VERSION`

- Extending [VkDependencyFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlagBits.html):

  - `VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR`

- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagBits.html):

  - `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR`

- Extending
  [VkPeerMemoryFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlagBits.html):

  - `VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR`

  - `VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR`

  - `VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR`

  - `VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_DISPATCH_BASE_KHR`

  - `VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_BIND_SPARSE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_COMMAND_BUFFER_BEGIN_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_SUBMIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_FLAGS_INFO_KHR`

If [VK_KHR_bind_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html) is supported:

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_DEVICE_GROUP_INFO_KHR`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_DEVICE_GROUP_INFO_KHR`

If [VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_CAPABILITIES_KHR`

If [VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_PRESENT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR`

- Extending
  [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html):

  - `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-in Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-deviceindex"
  target="_blank" rel="noopener"><code>DeviceIndex</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-DeviceGroup"
  target="_blank" rel="noopener"><code>DeviceGroup</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-19 (Jeff Bolz)

  - Internal revisions

- Revision 2, 2017-05-19 (Tobias Hector)

  - Removed extended memory bind functions to VK_KHR_bind_memory2, added
    dependency on that extension, and device-group-specific structs for
    those functions.

- Revision 3, 2017-10-06 (Ian Elliott)

  - Corrected Vulkan 1.1 interactions with the WSI extensions. All
    Vulkan 1.1 WSI interactions are with the VK_KHR_swapchain extension.

- Revision 4, 2017-10-10 (Jeff Bolz)

  - Rename “SFR” bits and structure members to use the phrase “split
    instance bind regions”.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_device_group"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
