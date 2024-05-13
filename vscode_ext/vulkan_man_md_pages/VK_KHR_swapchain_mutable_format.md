# VK_KHR_swapchain_mutable_format(3) Manual Page

## Name

VK_KHR_swapchain_mutable_format - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

201

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
     [VK_KHR_maintenance2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
     [VK_KHR_image_format_list](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_image_format_list.html)  
     or  
     [Version 1.2](#versions-1.2)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_swapchain_mutable_format%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_swapchain_mutable_format%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows processing of swapchain images as different
formats to that used by the window system, which is particularly useful
for switching between sRGB and linear RGB formats.

It adds a new swapchain creation flag that enables creating image views
from presentable images with a different format than the one used to
create the swapchain.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_EXTENSION_NAME`

- `VK_KHR_SWAPCHAIN_MUTABLE_FORMAT_SPEC_VERSION`

- Extending
  [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagBitsKHR.html):

  - `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Are there any new capabilities needed?

**RESOLVED**: No. It is expected that all implementations exposing this
extension support swapchain image format mutability.

2\) Do we need a separate `VK_SWAPCHAIN_CREATE_EXTENDED_USAGE_BIT_KHR`?

**RESOLVED**: No. This extension requires `VK_KHR_maintenance2` and
presentable images of swapchains created with
`VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` are created internally in a
way equivalent to specifying both `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`
and `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR`.

3\) Do we need a separate structure to allow specifying an image format
list for swapchains?

**RESOLVED**: No. We simply use the same
[VkImageFormatListCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfoKHR.html)
structure introduced by `VK_KHR_image_format_list`. The structure is
required to be included in the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) for swapchains
created with `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR`.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-03-28 (Daniel Rakos)

  - Internal revisions.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_swapchain_mutable_format"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
