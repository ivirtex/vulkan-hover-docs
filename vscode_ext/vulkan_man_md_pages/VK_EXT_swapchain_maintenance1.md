# VK_EXT_swapchain_maintenance1(3) Manual Page

## Name

VK_EXT_swapchain_maintenance1 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

276

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html)  
and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_swapchain_maintenance1%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_swapchain_maintenance1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_swapchain_maintenance1](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_swapchain_maintenance1.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-10-28

**Contributors**  
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

## <a href="#_description" class="anchor"></a>Description

[`VK_EXT_swapchain_maintenance1`](VK_EXT_swapchain_maintenance1.html)
adds a collection of window system integration features that were
intentionally left out or overlooked in the original
[`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) extension.

The new features are as follows:

- Specify a fence that will be signaled when the resources associated
  with a present operation **can** be safely destroyed.

- Allow changing the present mode a swapchain is using at per-present
  granularity.

- Allow applications to define the behavior when presenting a swapchain
  image to a surface with different dimensions than the image. Using
  this feature **may** allow implementations to avoid returning
  `VK_ERROR_OUT_OF_DATE_KHR` in this situation.

- Allow applications to defer swapchain memory allocation for improved
  startup time and memory footprint.

- Allow applications to release previously acquired images without
  presenting them.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseSwapchainImagesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT.html)

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentFenceInfoEXT.html)

  - [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModeInfoEXT.html)

- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)

  - [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SWAPCHAIN_MAINTENANCE_1_EXTENSION_NAME`

- `VK_EXT_SWAPCHAIN_MAINTENANCE_1_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SWAPCHAIN_MAINTENANCE_1_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_RELEASE_SWAPCHAIN_IMAGES_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_FENCE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_EXT`

- Extending
  [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html):

  - `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 0, 2019-05-28

  - Initial revisions

- Revision 1, 2022-08-21 (Shahbaz Youssefi)

  - Add functionality and complete spec

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_swapchain_maintenance1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
