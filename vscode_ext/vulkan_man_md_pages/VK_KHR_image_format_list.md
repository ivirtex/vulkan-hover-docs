# VK_KHR_image_format_list(3) Manual Page

## Name

VK_KHR_image_format_list - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

148

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_image_format_list%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_image_format_list%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-03-20

**IP Status**  
No known IP claims.

**Contributors**  
- Faith Ekstrand, Intel

- Jan-Harald Fredriksen, ARM

- Jeff Bolz, NVIDIA

- Jeff Leger, Qualcomm

- Neil Henning, Codeplay

## <a href="#_description" class="anchor"></a>Description

On some implementations, setting the
`VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` on image creation can cause access
to that image to perform worse than an equivalent image created without
`VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` because the implementation does not
know what view formats will be paired with the image.

This extension allows an application to provide the list of all formats
that **can** be used with an image when it is created. The
implementation may then be able to create a more efficient image that
supports the subset of formats required by the application without
having to support all formats in the format compatibility class of the
image format.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html):

  - [VkImageFormatListCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME`

- `VK_KHR_IMAGE_FORMAT_LIST_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_LIST_CREATE_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-03-20 (Faith Ekstrand)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_image_format_list"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
