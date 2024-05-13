# VK_EXT_image_view_min_lod(3) Manual Page

## Name

VK_EXT_image_view_min_lod - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

392

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Joshua Ashton <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_view_min_lod%5D%20@Joshua-Ashton%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_view_min_lod%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>Joshua-Ashton</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-11-09

**IP Status**  
No known IP claims.

**Contributors**  
- Joshua Ashton, Valve

- Hans-Kristian Arntzen, Valve

- Samuel Iglesias Gonsalvez, Igalia

- Tobias Hector, AMD

- Faith Ekstrand, Intel

- Tom Olson, ARM

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to clamp the minimum LOD value during
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-image-level-selection"
target="_blank" rel="noopener">Image Level(s) Selection</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-gather"
target="_blank" rel="noopener">Texel Gathering</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-integer-coordinate-operations"
target="_blank" rel="noopener">Integer Texel Coordinate Operations</a>
with a given [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) by
[VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod`.

This extension may be useful to restrict a
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) to only mips which have been uploaded,
and the use of fractional `minLod` can be useful for smoothly
introducing new mip levels when using linear mipmap filtering.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageViewMinLodFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageViewMinLodFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_VIEW_MIN_LOD_EXTENSION_NAME`

- `VK_EXT_IMAGE_VIEW_MIN_LOD_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_MIN_LOD_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_VIEW_MIN_LOD_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-07-06 (Joshua Ashton)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_view_min_lod"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
