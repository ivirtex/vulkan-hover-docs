# VK_EXT_image_sliced_view_of_3d(3) Manual Page

## Name

VK_EXT_image_sliced_view_of_3d - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

419

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html)  
     and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Mike Blumenkrantz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_sliced_view_of_3d%5D%20@zmike%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_sliced_view_of_3d%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>zmike</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_image_sliced_view_of_3d](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_image_sliced_view_of_3d.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-01-24

**IP Status**  
No known IP claims.

**Contributors**  
- Mike Blumenkrantz, Valve

- Hans-Kristian Arntzen, Valve

- Ricardo Garcia, Igalia

- Shahbaz Youssefi, Google

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows creating 3D views of 3D images such that the views
contain a subset of the slices in the image, using a Z offset and range,
for the purpose of using the views as storage image descriptors. This
matches functionality in D3D12 and is primarily intended to support
D3D12 emulation.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewSlicedCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSlicedCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_SLICED_VIEW_OF_3D_EXTENSION_NAME`

- `VK_EXT_IMAGE_SLICED_VIEW_OF_3D_SPEC_VERSION`

- `VK_REMAINING_3D_SLICES_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_SLICED_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_SLICED_VIEW_OF_3D_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-10-21 (Mike Blumenkrantz)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_sliced_view_of_3d"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
