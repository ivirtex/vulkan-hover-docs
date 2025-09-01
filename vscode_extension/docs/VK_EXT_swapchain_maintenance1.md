# VK\_EXT\_swapchain\_maintenance1(3) Manual Page

## Name

VK\_EXT\_swapchain\_maintenance1 - device extension



## [](#_registered_extension_number)Registered Extension Number

276

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html)  
and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain_maintenance1.html) extension

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_swapchain_maintenance1%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_swapchain_maintenance1%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_swapchain\_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_swapchain_maintenance1.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-12-16

**Interactions and External Dependencies**

- Promoted to `VK_KHR_swapchain_maintenance1`

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Shahbaz Youssefi, Google
- Chris Forbes, Google
- Ian Elliott, Google
- Yiwei Zhang, Google
- Charlie Lao, Google
- Lina Versace, Google
- Ralph Potter, Samsung
- Igor Nazarov, Samsung
- Hyunchang Kim, Samsung
- Suenghwan Lee, Samsung
- Munseong Kang, Samsung
- Joonyong Park, Samsung
- Hans-Kristian Arntzen, Valve
- Lisa Wu, Arm
- Daniel Stone, Collabora
- Pan Gao, Huawei

## [](#_description)Description

`VK_EXT_swapchain_maintenance1` adds a collection of window system integration features that were intentionally left out or overlooked in the original `VK_KHR_swapchain` extension.

The new features are as follows:

- Specify a fence that will be signaled when the resources associated with a present operation **can** be safely destroyed.
- Allow changing the present mode a swapchain is using at per-present granularity.
- Allow applications to define the behavior when presenting a swapchain image to a surface with different dimensions than the image. Using this feature **may** allow implementations to avoid returning `VK_ERROR_OUT_OF_DATE_KHR` in this situation.
- Allow applications to defer swapchain memory allocation for improved startup time and memory footprint.
- Allow applications to release previously acquired images without presenting them.

## [](#_promotion_to_vk_khr_swapchain_maintenance1)Promotion to `VK_KHR_swapchain_maintenance1`

All functionality in this extension is included in `VK_KHR_swapchain_maintenance1`, with the suffix changed to KHR. The original type, enum and command names are still available as aliases of the KHR functionality.

## [](#_new_commands)New Commands

- [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesEXT.html)

## [](#_new_structures)New Structures

- [VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoEXT.html)
  - [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoEXT.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html)
  - [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SWAPCHAIN_MAINTENANCE_1_EXTENSION_NAME`
- `VK_EXT_SWAPCHAIN_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SWAPCHAIN_MAINTENANCE_1_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_RELEASE_SWAPCHAIN_IMAGES_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_FENCE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_EXT`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT`

## [](#_version_history)Version History

- Revision 0, 2019-05-28 (James Jones)
  
  - Initial revisions
- Revision 1, 2022-08-21 (Shahbaz Youssefi)
  
  - Add functionality and complete spec

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_swapchain_maintenance1).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0