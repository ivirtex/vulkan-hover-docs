# VK\_KHR\_swapchain\_maintenance1(3) Manual Page

## Name

VK\_KHR\_swapchain\_maintenance1 - device extension



## [](#_registered_extension_number)Registered Extension Number

488

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
or  
[VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html)  
or  
[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_swapchain_maintenance1%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_swapchain_maintenance1%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_swapchain\_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_swapchain_maintenance1.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-31

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Shahbaz Youssefi, Google
- Chris Forbes, Google
- Ian Elliott, Google
- Yiwei Zhang, Google
- Charlie Lao, Google
- Chad Versace, Google
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

This extension is based off the `VK_EXT_swapchain_maintenance1` extension.

`VK_KHR_swapchain_maintenance1` adds a collection of window system integration features that were intentionally left out or overlooked in the original `VK_KHR_swapchain` extension.

The new features are as follows:

- Specify a fence that will be signaled when the resources associated with a present operation **can** be safely destroyed.
- Allow changing the present mode a swapchain is using at per-present granularity.
- Allow applications to define the behavior when presenting a swapchain image to a surface with different dimensions than the image. Using this feature **may** allow implementations to avoid returning `VK_ERROR_OUT_OF_DATE_KHR` in this situation.
- Allow applications to defer swapchain memory allocation for improved startup time and memory footprint.
- Allow applications to release previously acquired images without presenting them.

## [](#_new_commands)New Commands

- [vkReleaseSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesKHR.html)

## [](#_new_structures)New Structures

- [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSwapchainMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSwapchainMaintenance1FeaturesKHR.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkSwapchainPresentFenceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentFenceInfoKHR.html)
  - [VkSwapchainPresentModeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoKHR.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html)
  - [VkSwapchainPresentScalingCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SWAPCHAIN_MAINTENANCE_1_EXTENSION_NAME`
- `VK_KHR_SWAPCHAIN_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SWAPCHAIN_MAINTENANCE_1_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_RELEASE_SWAPCHAIN_IMAGES_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_FENCE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_KHR`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-03-31 (Shahbaz Youssefi)
  
  - Based on VK\_EXT\_swapchain\_maintenance1

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_swapchain_maintenance1)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0