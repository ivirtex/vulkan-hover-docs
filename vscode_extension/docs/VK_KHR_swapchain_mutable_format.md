# VK\_KHR\_swapchain\_mutable\_format(3) Manual Page

## Name

VK\_KHR\_swapchain\_mutable\_format - device extension



## [](#_registered_extension_number)Registered Extension Number

201

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
     [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
     [VK\_KHR\_image\_format\_list](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_image_format_list.html)  
     or  
     [Vulkan Version 1.2](#versions-1.2)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_swapchain_mutable_format%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_swapchain_mutable_format%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-03-28

**IP Status**

No known IP claims.

**Contributors**

- Faith Ekstrand, Intel
- Jan-Harald Fredriksen, ARM
- Jesse Hall, Google
- Daniel Rakos, AMD
- Ray Smith, ARM

## [](#_description)Description

This extension allows processing of swapchain images as different formats to that used by the window system, which is particularly useful for switching between sRGB and linear RGB formats.

It adds a new swapchain creation flag that enables creating image views from presentable images with a different format than the one used to create the swapchain.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_EXTENSION_NAME`
- `VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_SPEC_VERSION`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR`

## [](#_issues)Issues

1\) Are there any new capabilities needed?

**RESOLVED**: No. It is expected that all implementations exposing this extension support swapchain image format mutability.

2\) Do we need a separate `VK_SWAPCHAIN_CREATE_EXTENDED_USAGE_BIT_KHR`?

**RESOLVED**: No. This extension requires `VK_KHR_maintenance2` and presentable images of swapchains created with `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` are created internally in a way equivalent to specifying both `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` and `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR`.

3\) Do we need a separate structure to allow specifying an image format list for swapchains?

**RESOLVED**: No. We simply use the same [VkImageFormatListCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfoKHR.html) structure introduced by `VK_KHR_image_format_list`. The structure is required to be included in the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) for swapchains created with `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR`.

## [](#_version_history)Version History

- Revision 1, 2018-03-28 (Daniel Rakos)
  
  - Internal revisions.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_swapchain_mutable_format).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0